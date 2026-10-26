<?php

require_once '/input/class-wp-block-parser-block.php';
require_once '/input/class-wp-block-parser-frame.php';
require_once '/input/class-wp-block-parser.php';
require '/input/drivers.php';

if ( PHP_VERSION_ID < 80300 || PHP_VERSION_ID >= 80400 || setlocale( LC_CTYPE, 'C' ) !== 'C' || setlocale( LC_CTYPE, 0 ) !== 'C' ) {
	throw new RuntimeException( 'requires PHP8.3 and C locale' );
}

function parser_add_freeform_require_originals(): void {
	$parser = new ReflectionClass( 'WP_Block_Parser' );
	if ( ! $parser->isUserDefined() || $parser->isInternal() || null !== $parser->getConstructor() ) {
		throw new RuntimeException( 'requires original constructorless WP_Block_Parser' );
	}
	$property_names = [ 'document', 'offset', 'output', 'stack' ];
	$properties = $parser->getProperties();
	if ( count( $properties ) !== count( $property_names ) || array_map( static fn( ReflectionProperty $property ): string => $property->getName(), $properties ) !== $property_names ) {
		throw new RuntimeException( 'unexpected WP_Block_Parser property declaration order' );
	}
	foreach ( $properties as $property ) {
		if ( $property->getDeclaringClass()->getName() !== $parser->getName() || ! $property->isPublic() ) {
			throw new RuntimeException( 'requires original public WP_Block_Parser properties only' );
		}
	}
	$method_names = [ 'parse', 'proceed', 'next_token', 'freeform', 'add_freeform', 'add_inner_block', 'add_block_from_stack' ];
	$methods = $parser->getMethods();
	if ( count( $methods ) !== count( $method_names ) || array_map( static fn( ReflectionMethod $method ): string => $method->getName(), $methods ) !== $method_names ) {
		throw new RuntimeException( 'unexpected WP_Block_Parser method declaration order' );
	}
	$freeform = $parser->getMethod( 'freeform' );
	$add_freeform = $parser->getMethod( 'add_freeform' );
	if ( ! $freeform->isPublic() || $freeform->getDeclaringClass()->getName() !== $parser->getName() || 1 !== $freeform->getNumberOfParameters() || 1 !== $freeform->getNumberOfRequiredParameters() || 314 !== $freeform->getStartLine() || 316 !== $freeform->getEndLine() || ! $add_freeform->isPublic() || $add_freeform->getDeclaringClass()->getName() !== $parser->getName() || 1 !== $add_freeform->getNumberOfParameters() || 0 !== $add_freeform->getNumberOfRequiredParameters() || 327 !== $add_freeform->getStartLine() || 335 !== $add_freeform->getEndLine() ) {
		throw new RuntimeException( 'unexpected original freeform/add_freeform signature or lines' );
	}
	$parameter = $add_freeform->getParameters()[0];
	if ( $parameter->hasType() || ! $parameter->isDefaultValueAvailable() || null !== $parameter->getDefaultValue() ) {
		throw new RuntimeException( 'requires original untyped null default add_freeform argument' );
	}

	$block = new ReflectionClass( 'WP_Block_Parser_Block' );
	$constructor = $block->getConstructor();
	if ( ! $block->isUserDefined() || $block->isInternal() || null === $constructor || 5 !== $constructor->getNumberOfParameters() || 5 !== $constructor->getNumberOfRequiredParameters() ) {
		throw new RuntimeException( 'requires original WP_Block_Parser_Block' );
	}
	$block_names = [ 'blockName', 'attrs', 'innerBlocks', 'innerHTML', 'innerContent' ];
	$block_properties = $block->getProperties();
	if ( count( $block_properties ) !== count( $block_names ) || array_map( static fn( ReflectionProperty $property ): string => $property->getName(), $block_properties ) !== $block_names ) {
		throw new RuntimeException( 'unexpected WP_Block_Parser_Block property declaration order' );
	}
	foreach ( $block_properties as $property ) {
		if ( $property->getDeclaringClass()->getName() !== $block->getName() || ! $property->isPublic() ) {
			throw new RuntimeException( 'requires original public WP_Block_Parser_Block properties only' );
		}
	}
}

function parser_add_freeform_string( mixed $value, string $field ): string {
	if ( ! is_string( $value ) || ! preg_match( '//u', $value ) || strlen( $value ) > 4096 ) {
		throw new RuntimeException( 'invalid UTF-8 string field: ' . $field );
	}
	return $value;
}

function parser_add_freeform_expected_html( mixed $value, string $field ): ?string {
	if ( null === $value ) {
		return null;
	}
	$encoded = parser_add_freeform_string( $value, $field );
	$decoded = base64_decode( $encoded, true );
	if ( false === $decoded || base64_encode( $decoded ) !== $encoded ) {
		throw new RuntimeException( 'expected_html requires canonical standard base64: ' . $field );
	}
	return $decoded;
}

function parser_add_freeform_string_observation( string $value ): array {
	return [ 'type' => 'string', 'base64' => base64_encode( $value ) ];
}

function parser_add_freeform_empty_array( mixed $value, string $field ): array {
	if ( ! is_array( $value ) || [] !== $value ) {
		throw new RuntimeException( $field . ' requires empty array' );
	}
	return [];
}

function parser_add_freeform_block_observation( mixed $value, string $html, string $field ): array {
	if ( ! is_array( $value ) || array_keys( $value ) !== [ 'blockName', 'attrs', 'innerBlocks', 'innerHTML', 'innerContent' ] || null !== $value['blockName'] || ! is_string( $value['innerHTML'] ) || $value['innerHTML'] !== $html || ! is_array( $value['innerContent'] ) || [ $html ] !== $value['innerContent'] ) {
		throw new RuntimeException( 'unexpected actual freeform block array: ' . $field );
	}
	return [
		'type' => 'array',
		'properties' => [
			[ 'blockName', null ],
			[ 'attrs', parser_add_freeform_empty_array( $value['attrs'], $field . ' attrs' ) ],
			[ 'innerBlocks', parser_add_freeform_empty_array( $value['innerBlocks'], $field . ' innerBlocks' ) ],
			[ 'innerHTML', parser_add_freeform_string_observation( $value['innerHTML'] ) ],
			[ 'innerContent', [ parser_add_freeform_string_observation( $value['innerContent'][0] ) ] ],
		],
	];
}

parser_add_freeform_require_originals();
$cases = json_decode( file_get_contents( '/input/parser-add-freeform-cases.json' ), true, 512, JSON_THROW_ON_ERROR );
if ( ! is_array( $cases ) || ! array_is_list( $cases ) || 16 !== count( $cases ) ) {
	throw new RuntimeException( 'parser-add-freeform cases must be an ordered list of 16 cases' );
}

$observations = [];
$seen = [];
foreach ( $cases as $case ) {
	if ( ! is_array( $case ) || [ 'name', 'document', 'offset', 'length', 'mode', 'seed', 'expected_html' ] !== array_keys( $case ) ) {
		throw new RuntimeException( 'invalid parser-add-freeform case fields' );
	}
	$name = parser_add_freeform_string( $case['name'], 'name' );
	$document = parser_add_freeform_string( $case['document'], $name . ' document' );
	$offset = $case['offset'];
	$length = $case['length'];
	$mode = parser_add_freeform_string( $case['mode'], $name . ' mode' );
	$seed = $case['seed'];
	$expected_html = parser_add_freeform_expected_html( $case['expected_html'], $name . ' expected_html' );
	if ( '' === $name || isset( $seen[ $name ] ) || ! is_int( $offset ) || $offset < -1000000 || $offset > 1000000 || ( ! is_int( $length ) && null !== $length ) || ( is_int( $length ) && ( $length < -1000000 || $length > 1000000 ) ) || ( 'omitted' !== $mode && 'explicit' !== $mode ) || ! is_array( $seed ) || ! array_is_list( $seed ) || count( $seed ) > 4 ) {
		throw new RuntimeException( 'invalid parser-add-freeform case domain: ' . $name );
	}
	foreach ( $seed as $index => $value ) {
		$seed[ $index ] = parser_add_freeform_string( $value, $name . ' seed' . $index );
	}
	$seen[ $name ] = true;
	$seed_before = $seed;
	$data = gossip_parser_add_freeform( $document, $offset, $length, $mode, $seed );
	if ( ! is_array( $data ) || [ 'parser', 'return' ] !== array_keys( $data ) || ! ( $data['parser'] instanceof WP_Block_Parser ) || null !== $data['return'] || $seed !== $seed_before ) {
		throw new RuntimeException( 'source helper result/caller arguments mismatch: ' . $name );
	}
	$parser = $data['parser'];
	if ( ! is_string( $parser->document ) || $parser->document !== $document || ! is_int( $parser->offset ) || $parser->offset !== $offset || ! is_array( $parser->stack ) || [] !== $parser->stack || ! is_array( $parser->output ) || ! array_is_list( $parser->output ) ) {
		throw new RuntimeException( 'parser state mismatch: ' . $name );
	}
	$expected_count = count( $seed ) + ( null === $expected_html ? 0 : 1 );
	if ( count( $parser->output ) !== $expected_count ) {
		throw new RuntimeException( 'parser output count mismatch: ' . $name );
	}
	$output = [];
	foreach ( $parser->output as $index => $value ) {
		if ( $index < count( $seed ) ) {
			if ( ! is_string( $value ) || $value !== $seed[ $index ] ) {
				throw new RuntimeException( 'seed output mismatch: ' . $name );
			}
			$output[] = parser_add_freeform_string_observation( $value );
		} else {
			$output[] = parser_add_freeform_block_observation( $value, $expected_html, $name . ' freeform' );
		}
	}
	$seed_base64 = [];
	foreach ( $seed as $value ) {
		$seed_base64[] = base64_encode( $value );
	}
	$observations[] = [
		'name' => $name,
		'return_type' => 'null',
		'return_value' => null,
		'parser' => [ 'type' => 'object', 'class' => 'WP_Block_Parser', 'document_base64' => base64_encode( $parser->document ), 'offset' => $parser->offset, 'stack' => [], 'output' => $output ],
		'arguments_after' => [ 'document_base64' => base64_encode( $document ), 'offset' => $offset, 'length' => $length, 'mode' => $mode, 'seed_base64' => $seed_base64 ],
	];
}

echo json_encode( $observations, JSON_THROW_ON_ERROR ), "\n";
