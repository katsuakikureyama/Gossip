<?php

require '/input/helpers.php';
require '/input/transport.php';

if ( PHP_VERSION_ID < 80300 || PHP_VERSION_ID >= 80400 || setlocale( LC_CTYPE, 'C' ) !== 'C' || setlocale( LC_CTYPE, 0 ) !== 'C' ) {
	throw new RuntimeException( 'requires PHP8.3 and C locale' );
}

$class = new ReflectionClass( 'WP_Block_Parser_Block' );
if ( ! $class->isUserDefined() || $class->isInternal() ) {
	throw new RuntimeException( 'requires original user-defined WP_Block_Parser_Block' );
}
$constructor = $class->getConstructor();
if ( null === $constructor || 5 !== $constructor->getNumberOfParameters() || 5 !== $constructor->getNumberOfRequiredParameters() ) {
	throw new RuntimeException( 'requires five-required-parameter original constructor' );
}
$property_names = [ 'blockName', 'attrs', 'innerBlocks', 'innerHTML', 'innerContent' ];
$properties     = $class->getProperties();
if ( count( $properties ) !== 5 || array_map( static fn( ReflectionProperty $property ): string => $property->getName(), $properties ) !== $property_names ) {
	throw new RuntimeException( 'requires five original properties in declaration order' );
}
foreach ( $properties as $property ) {
	if ( $property->getDeclaringClass()->getName() !== $class->getName() || ! $property->isPublic() ) {
		throw new RuntimeException( 'requires public original properties only' );
	}
}
$methods = $class->getMethods();
if ( count( $methods ) !== 1 || '__construct' !== $methods[0]->getName() ) {
	throw new RuntimeException( 'requires no additional parser-block methods' );
}
foreach ( [ '_wp_array_get', 'wp_should_add_elements_class_name' ] as $name ) {
	$function = new ReflectionFunction( $name );
	if ( ! $function->isUserDefined() || $function->isInternal() ) {
		throw new RuntimeException( 'requires original user-defined ' . $name );
	}
}

$cases = json_decode( file_get_contents( '/input/parser-block-cases.json' ), true, 512, JSON_THROW_ON_ERROR );
if ( ! is_array( $cases ) || ! array_is_list( $cases ) || count( $cases ) !== 16 ) {
	throw new RuntimeException( 'parser-block cases must be an ordered list of 16 cases' );
}

$observations = [];
$seen         = [];
foreach ( $cases as $case ) {
	if ( ! is_array( $case ) || array_keys( $case ) !== [ 'name', 'block_name', 'attrs', 'inner_blocks', 'inner_html', 'inner_content', 'options', 'expected' ] || ! is_string( $case['name'] ) || '' === $case['name'] || isset( $seen[ $case['name'] ] ) || ! is_bool( $case['expected'] ) ) {
		throw new RuntimeException( 'invalid parser-block case' );
	}
	$seen[ $case['name'] ] = true;

	$block_name_entries    = 0;
	$attrs_entries         = 0;
	$inner_blocks_entries  = 0;
	$inner_html_entries    = 0;
	$inner_content_entries = 0;
	$options_entries       = 0;
	$block_name    = array_get_fixture_decode_node( $case['block_name'], $case['name'] . ' block_name', 0, $block_name_entries );
	$attrs         = array_get_fixture_decode_node( $case['attrs'], $case['name'] . ' attrs', 0, $attrs_entries );
	$inner_blocks  = array_get_fixture_decode_node( $case['inner_blocks'], $case['name'] . ' inner_blocks', 0, $inner_blocks_entries );
	$inner_html    = array_get_fixture_decode_node( $case['inner_html'], $case['name'] . ' inner_html', 0, $inner_html_entries );
	$inner_content = array_get_fixture_decode_node( $case['inner_content'], $case['name'] . ' inner_content', 0, $inner_content_entries );
	$options       = array_get_fixture_decode_node( $case['options'], $case['name'] . ' options', 0, $options_entries );
	if ( ( ! is_string( $block_name ) && ! is_null( $block_name ) ) || ( ! is_array( $attrs ) && ! is_null( $attrs ) ) || ! is_array( $inner_blocks ) || ! is_string( $inner_html ) || ! is_array( $inner_content ) || ! is_array( $options ) ) {
		throw new RuntimeException( 'invalid parser-block argument domain: ' . $case['name'] );
	}
	foreach ( $inner_blocks as $index => $inner_block ) {
		if ( ! is_array( $inner_block ) ) {
			throw new RuntimeException( 'inner_blocks must contain data arrays only: ' . $case['name'] . ' index ' . $index );
		}
	}
	foreach ( $inner_content as $index => $fragment ) {
		if ( ! is_string( $fragment ) && ! is_null( $fragment ) ) {
			throw new RuntimeException( 'inner_content must contain string/null fragments: ' . $case['name'] . ' index ' . $index );
		}
	}

	$block_name_before    = $block_name;
	$attrs_before         = $attrs;
	$inner_blocks_before  = $inner_blocks;
	$inner_html_before    = $inner_html;
	$inner_content_before = $inner_content;
	$options_before       = $options;
	$data = gossip_parser_block_data( $block_name, $attrs, $inner_blocks, $inner_html, $inner_content );
	if ( ! is_array( $data ) || array_keys( $data ) !== [ 'object', 'snapshot' ] || ! ( $data['object'] instanceof WP_Block_Parser_Block ) || ! is_array( $data['snapshot'] ) ) {
		throw new RuntimeException( 'parser-block source helper returned invalid data' );
	}
	$object   = $data['object'];
	$snapshot = $data['snapshot'];
	$expected_properties = [
		'blockName'    => $block_name,
		'attrs'        => $attrs,
		'innerBlocks'  => $inner_blocks,
		'innerHTML'    => $inner_html,
		'innerContent' => $inner_content,
	];
	$property_values = [];
	foreach ( $property_names as $property_name ) {
		$property_values[ $property_name ] = $object->{$property_name};
	}
	$snapshot_before = $snapshot;
	$value           = wp_should_add_elements_class_name( $snapshot, $options );
	$properties_entries = 0;
	$snapshot_entries = 0;
	$block_name_after_entries = 0;
	$attrs_after_entries = 0;
	$inner_blocks_after_entries = 0;
	$inner_html_after_entries = 0;
	$inner_content_after_entries = 0;
	$options_after_entries = 0;
	$properties_node = array_get_fixture_encode_node( $property_values, $case['name'] . ' properties', 0, $properties_entries );
	$snapshot_node = array_get_fixture_encode_node( $snapshot, $case['name'] . ' snapshot', 0, $snapshot_entries );
	$arguments_after = [
		'block_name'    => array_get_fixture_encode_node( $block_name, $case['name'] . ' block_name after', 0, $block_name_after_entries ),
		'attrs'         => array_get_fixture_encode_node( $attrs, $case['name'] . ' attrs after', 0, $attrs_after_entries ),
		'inner_blocks'  => array_get_fixture_encode_node( $inner_blocks, $case['name'] . ' inner_blocks after', 0, $inner_blocks_after_entries ),
		'inner_html'    => array_get_fixture_encode_node( $inner_html, $case['name'] . ' inner_html after', 0, $inner_html_after_entries ),
		'inner_content' => array_get_fixture_encode_node( $inner_content, $case['name'] . ' inner_content after', 0, $inner_content_after_entries ),
		'options'       => array_get_fixture_encode_node( $options, $case['name'] . ' options after', 0, $options_after_entries ),
	];
	if ( get_debug_type( $value ) !== 'bool' || $value !== $case['expected'] || $property_values !== $expected_properties || $snapshot !== $expected_properties || $snapshot !== $snapshot_before || $block_name !== $block_name_before || $attrs !== $attrs_before || $inner_blocks !== $inner_blocks_before || $inner_html !== $inner_html_before || $inner_content !== $inner_content_before || $options !== $options_before ) {
		throw new RuntimeException( 'parser-block expectation or mutation mismatch: ' . $case['name'] );
	}
	$observations[] = [
		'name'            => $case['name'],
		'type'            => 'object',
		'class'           => 'WP_Block_Parser_Block',
		'properties'      => $properties_node,
		'snapshot'        => $snapshot_node,
		'decision'        => $value,
		'arguments_after' => $arguments_after,
	];
}

echo json_encode( $observations, JSON_THROW_ON_ERROR ), "\n";
