<?php

require '/input/helpers.php';
require '/input/transport.php';

if ( PHP_VERSION_ID < 80300 || PHP_VERSION_ID >= 80400 || setlocale( LC_CTYPE, 'C' ) !== 'C' || setlocale( LC_CTYPE, 0 ) !== 'C' ) {
	throw new RuntimeException( 'requires PHP8.3 and C locale' );
}

function parser_frame_require_class( string $name, array $property_names, int $total_parameters, int $required_parameters ): void {
	$class = new ReflectionClass( $name );
	if ( ! $class->isUserDefined() || $class->isInternal() ) {
		throw new RuntimeException( 'requires original user-defined ' . $name );
	}
	$constructor = $class->getConstructor();
	if ( null === $constructor || $total_parameters !== $constructor->getNumberOfParameters() || $required_parameters !== $constructor->getNumberOfRequiredParameters() ) {
		throw new RuntimeException( 'unexpected constructor signature for ' . $name );
	}
	foreach ( $constructor->getParameters() as $index => $parameter ) {
		if ( $parameter->hasType() || ( $index >= $required_parameters && ( ! $parameter->isDefaultValueAvailable() || null !== $parameter->getDefaultValue() ) ) ) {
			throw new RuntimeException( 'requires original untyped/null-default constructor for ' . $name );
		}
	}
	$properties = $class->getProperties();
	if ( count( $properties ) !== count( $property_names ) || array_map( static fn( ReflectionProperty $property ): string => $property->getName(), $properties ) !== $property_names ) {
		throw new RuntimeException( 'unexpected public property declaration order for ' . $name );
	}
	foreach ( $properties as $property ) {
		if ( $property->getDeclaringClass()->getName() !== $name || ! $property->isPublic() ) {
			throw new RuntimeException( 'requires original public properties only for ' . $name );
		}
	}
	$methods = $class->getMethods();
	if ( count( $methods ) !== 1 || '__construct' !== $methods[0]->getName() ) {
		throw new RuntimeException( 'requires constructor only for ' . $name );
	}
}

parser_frame_require_class( 'WP_Block_Parser_Block', [ 'blockName', 'attrs', 'innerBlocks', 'innerHTML', 'innerContent' ], 5, 5 );
parser_frame_require_class( 'WP_Block_Parser_Frame', [ 'block', 'token_start', 'token_length', 'prev_offset', 'leading_html_start' ], 5, 3 );

$cases = json_decode( file_get_contents( '/input/parser-frame-cases.json' ), true, 512, JSON_THROW_ON_ERROR );
if ( ! is_array( $cases ) || ! array_is_list( $cases ) || count( $cases ) !== 12 ) {
	throw new RuntimeException( 'parser-frame cases must be an ordered list of 12 cases' );
}

$observations = [];
$seen = [];
foreach ( $cases as $case ) {
	if ( ! is_array( $case ) || array_keys( $case ) !== [ 'name', 'arity', 'block_name', 'attrs', 'inner_html', 'inner_content', 'start', 'length', 'prev', 'leading', 'after_html', 'expected_prev' ] || ! is_string( $case['name'] ) || '' === $case['name'] || isset( $seen[ $case['name'] ] ) || ! is_int( $case['arity'] ) || $case['arity'] < 3 || $case['arity'] > 5 ) {
		throw new RuntimeException( 'invalid parser-frame case' );
	}
	$seen[ $case['name'] ] = true;
	$entries = array_fill( 0, 9, 0 );
	$block_name = array_get_fixture_decode_node( $case['block_name'], $case['name'] . ' block_name', 0, $entries[0] );
	$attrs = array_get_fixture_decode_node( $case['attrs'], $case['name'] . ' attrs', 0, $entries[1] );
	$inner_html = array_get_fixture_decode_node( $case['inner_html'], $case['name'] . ' inner_html', 0, $entries[2] );
	$inner_content = array_get_fixture_decode_node( $case['inner_content'], $case['name'] . ' inner_content', 0, $entries[3] );
	$start = array_get_fixture_decode_node( $case['start'], $case['name'] . ' start', 0, $entries[4] );
	$length = array_get_fixture_decode_node( $case['length'], $case['name'] . ' length', 0, $entries[5] );
	$prev = array_get_fixture_decode_node( $case['prev'], $case['name'] . ' prev', 0, $entries[6] );
	$leading = array_get_fixture_decode_node( $case['leading'], $case['name'] . ' leading', 0, $entries[7] );
	$after_html = array_get_fixture_decode_node( $case['after_html'], $case['name'] . ' after_html', 0, $entries[8] );
	$expected_prev_entries = 0;
	$expected_prev = array_get_fixture_decode_node( $case['expected_prev'], $case['name'] . ' expected_prev', 0, $expected_prev_entries );
	if ( ( ! is_string( $block_name ) && ! is_null( $block_name ) ) || ( ! is_array( $attrs ) && ! is_null( $attrs ) ) || ! is_string( $inner_html ) || ! is_array( $inner_content ) || ! is_int( $start ) || ! is_int( $length ) || ( ! is_int( $prev ) && ! is_null( $prev ) ) || ( ! is_int( $leading ) && ! is_null( $leading ) ) || ! is_string( $after_html ) || ! is_int( $expected_prev ) ) {
		throw new RuntimeException( 'invalid parser-frame argument domain: ' . $case['name'] );
	}
	foreach ( $inner_content as $fragment ) {
		if ( ! is_string( $fragment ) && ! is_null( $fragment ) ) {
			throw new RuntimeException( 'inner_content must contain string/null fragments: ' . $case['name'] );
		}
	}
	$arguments = [ $block_name, $attrs, $inner_html, $inner_content, $start, $length, $prev, $leading, $after_html ];
	$arguments_before = $arguments;
	$data = match ( $case['arity'] ) {
		3 => gossip_parser_frame_3( ...$arguments ),
		4 => gossip_parser_frame_4( ...$arguments ),
		5 => gossip_parser_frame_5( ...$arguments ),
	};
	if ( ! is_array( $data ) || array_keys( $data ) !== [ 'frame', 'block', 'snapshot' ] || ! ( $data['frame'] instanceof WP_Block_Parser_Frame ) || ! ( $data['block'] instanceof WP_Block_Parser_Block ) || ! is_array( $data['snapshot'] ) ) {
		throw new RuntimeException( 'parser-frame source helper returned invalid data' );
	}
	$frame = $data['frame'];
	$block = $data['block'];
	$snapshot = $data['snapshot'];
	if ( array_keys( $snapshot ) !== [ 'block', 'token_start', 'token_length', 'prev_offset', 'leading_html_start' ] || $frame->block !== $block || $snapshot['block'] !== $block || $frame->block !== $snapshot['block'] || $frame->token_start !== $start || $frame->token_length !== $length || $frame->prev_offset !== $expected_prev || $frame->leading_html_start !== ( 5 === $case['arity'] ? $leading : null ) || $block->innerHTML !== $after_html || $arguments !== $arguments_before ) {
		throw new RuntimeException( 'parser-frame constructor/default/identity mismatch: ' . $case['name'] );
	}
	$frame_entries = 0;
	$snapshot_entries = 0;
	$block_entries = 0;
	$frame_scalar = array_get_fixture_encode_node( [ 'token_start' => $frame->token_start, 'token_length' => $frame->token_length, 'prev_offset' => $frame->prev_offset, 'leading_html_start' => $frame->leading_html_start ], $case['name'] . ' frame scalar properties', 0, $frame_entries );
	$snapshot_scalar = array_get_fixture_encode_node( [ 'token_start' => $snapshot['token_start'], 'token_length' => $snapshot['token_length'], 'prev_offset' => $snapshot['prev_offset'], 'leading_html_start' => $snapshot['leading_html_start'] ], $case['name'] . ' snapshot scalar properties', 0, $snapshot_entries );
	$block_properties = array_get_fixture_encode_node( [ 'blockName' => $block->blockName, 'attrs' => $block->attrs, 'innerBlocks' => $block->innerBlocks, 'innerHTML' => $block->innerHTML, 'innerContent' => $block->innerContent ], $case['name'] . ' block properties', 0, $block_entries );
	$arguments_after = [];
	foreach ( [ 'block_name', 'attrs', 'inner_html', 'inner_content', 'start', 'length', 'prev', 'leading', 'after_html' ] as $index => $name ) {
		$argument_entries = 0;
		$arguments_after[ $name ] = array_get_fixture_encode_node( $arguments[ $index ], $case['name'] . ' ' . $name . ' after', 0, $argument_entries );
	}
	$observations[] = [
		'name' => $case['name'], 'arity' => $case['arity'],
		'frame' => [ 'type' => 'object', 'class' => 'WP_Block_Parser_Frame', 'scalar_properties' => $frame_scalar ],
		'snapshot' => [ 'type' => 'array', 'source_class' => 'WP_Block_Parser_Frame', 'property_order' => array_keys( $snapshot ), 'scalar_properties' => $snapshot_scalar ],
		'block' => [ 'type' => 'object', 'class' => 'WP_Block_Parser_Block', 'properties' => $block_properties ],
		'object_identity' => true, 'arguments_after' => $arguments_after,
	];
}

echo json_encode( $observations, JSON_THROW_ON_ERROR ), "\n";
