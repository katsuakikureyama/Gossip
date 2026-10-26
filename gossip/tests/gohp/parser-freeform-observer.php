<?php

require '/input/class-wp-block-parser.php';
require '/input/drivers.php';

if ( PHP_VERSION_ID < 80300 || PHP_VERSION_ID >= 80400 || setlocale( LC_CTYPE, 'C' ) !== 'C' || setlocale( LC_CTYPE, 0 ) !== 'C' ) {
	throw new RuntimeException( 'requires PHP8.3 and C locale' );
}

function parser_freeform_require_originals(): void {
	$parser = new ReflectionClass( 'WP_Block_Parser' );
	if ( ! $parser->isUserDefined() || $parser->isInternal() || null !== $parser->getConstructor() ) {
		throw new RuntimeException( 'requires original user-defined constructorless WP_Block_Parser' );
	}
	$parser_properties = [ 'document', 'offset', 'output', 'stack' ];
	$properties        = $parser->getProperties();
	if ( count( $properties ) !== count( $parser_properties ) || array_map( static fn( ReflectionProperty $property ): string => $property->getName(), $properties ) !== $parser_properties ) {
		throw new RuntimeException( 'unexpected WP_Block_Parser property declaration order' );
	}
	foreach ( $properties as $property ) {
		if ( $property->getDeclaringClass()->getName() !== $parser->getName() || ! $property->isPublic() ) {
			throw new RuntimeException( 'requires original public WP_Block_Parser properties only' );
		}
	}
	$method_names = [ 'parse', 'proceed', 'next_token', 'freeform', 'add_freeform', 'add_inner_block', 'add_block_from_stack' ];
	$methods      = $parser->getMethods();
	if ( count( $methods ) !== count( $method_names ) || array_map( static fn( ReflectionMethod $method ): string => $method->getName(), $methods ) !== $method_names ) {
		throw new RuntimeException( 'unexpected WP_Block_Parser method declaration order' );
	}
	$freeform = $parser->getMethod( 'freeform' );
	if ( ! $freeform->isPublic() || $freeform->getDeclaringClass()->getName() !== $parser->getName() || 1 !== $freeform->getNumberOfParameters() || 1 !== $freeform->getNumberOfRequiredParameters() || 314 !== $freeform->getStartLine() || 316 !== $freeform->getEndLine() ) {
		throw new RuntimeException( 'unexpected original WP_Block_Parser::freeform signature or lines' );
	}

	$block = new ReflectionClass( 'WP_Block_Parser_Block' );
	if ( ! $block->isUserDefined() || $block->isInternal() ) {
		throw new RuntimeException( 'requires original user-defined WP_Block_Parser_Block' );
	}
	$constructor = $block->getConstructor();
	if ( null === $constructor || 5 !== $constructor->getNumberOfParameters() || 5 !== $constructor->getNumberOfRequiredParameters() ) {
		throw new RuntimeException( 'unexpected WP_Block_Parser_Block constructor signature' );
	}
	$block_properties = [ 'blockName', 'attrs', 'innerBlocks', 'innerHTML', 'innerContent' ];
	$properties       = $block->getProperties();
	if ( count( $properties ) !== count( $block_properties ) || array_map( static fn( ReflectionProperty $property ): string => $property->getName(), $properties ) !== $block_properties ) {
		throw new RuntimeException( 'unexpected WP_Block_Parser_Block property declaration order' );
	}
	foreach ( $properties as $property ) {
		if ( $property->getDeclaringClass()->getName() !== $block->getName() || ! $property->isPublic() ) {
			throw new RuntimeException( 'requires original public WP_Block_Parser_Block properties only' );
		}
	}
}

function parser_freeform_string( mixed $value, string $field ): string {
	if ( ! is_string( $value ) || ! preg_match( '//u', $value ) || strlen( $value ) > 4096 ) {
		throw new RuntimeException( 'invalid UTF-8 string field: ' . $field );
	}
	return $value;
}

function parser_freeform_block_observation( WP_Block_Parser_Block $block, string $html, string $field ): array {
	$block_name    = $block->blockName;
	$attrs         = $block->attrs;
	$inner_blocks  = $block->innerBlocks;
	$inner_html    = $block->innerHTML;
	$inner_content = $block->innerContent;
	if ( null !== $block_name || ! is_array( $attrs ) || [] !== $attrs || ! is_array( $inner_blocks ) || [] !== $inner_blocks || ! is_string( $inner_html ) || $inner_html !== $html || ! is_array( $inner_content ) || [ $html ] !== $inner_content ) {
		throw new RuntimeException( 'unexpected actual freeform block properties: ' . $field );
	}
	return [
		'type'       => 'object',
		'class'      => 'WP_Block_Parser_Block',
		'properties' => [
			[ 'blockName', $block_name ],
			[ 'attrs', $attrs ],
			[ 'innerBlocks', $inner_blocks ],
			[ 'innerHTML', $inner_html ],
			[ 'innerContent', $inner_content ],
		],
	];
}

parser_freeform_require_originals();

$cases = json_decode( file_get_contents( '/input/parser-freeform-cases.json' ), true, 512, JSON_THROW_ON_ERROR );
if ( ! is_array( $cases ) || ! array_is_list( $cases ) || 12 !== count( $cases ) ) {
	throw new RuntimeException( 'parser-freeform cases must be an ordered list of 12 cases' );
}

$observations = [];
$seen         = [];
foreach ( $cases as $case ) {
	if ( ! is_array( $case ) || [ 'name', 'first', 'second' ] !== array_keys( $case ) ) {
		throw new RuntimeException( 'invalid parser-freeform case fields' );
	}
	$name   = parser_freeform_string( $case['name'], 'name' );
	$first  = parser_freeform_string( $case['first'], $name . ' first' );
	$second = parser_freeform_string( $case['second'], $name . ' second' );
	if ( '' === $name || isset( $seen[ $name ] ) ) {
		throw new RuntimeException( 'invalid parser-freeform case name' );
	}
	$seen[ $name ] = true;

	$first_before  = $first;
	$second_before = $second;
	$data          = gossip_parser_freeform( $first, $second );
	if ( ! is_array( $data ) || [ 'parser', 'first', 'second' ] !== array_keys( $data ) || ! ( $data['parser'] instanceof WP_Block_Parser ) || ! ( $data['first'] instanceof WP_Block_Parser_Block ) || ! ( $data['second'] instanceof WP_Block_Parser_Block ) ) {
		throw new RuntimeException( 'parser-freeform source helper returned invalid data: ' . $name );
	}
	$parser = $data['parser'];
	if ( null !== $parser->document || null !== $parser->offset || null !== $parser->output || null !== $parser->stack || $data['first'] === $data['second'] || $first !== $first_before || $second !== $second_before ) {
		throw new RuntimeException( 'parser/freeform identity or caller argument mismatch: ' . $name );
	}
	$observations[] = [
		'name' => $name,
		'parser' => [
			'type'       => 'object',
			'class'      => 'WP_Block_Parser',
			'properties' => [
				[ 'document', $parser->document ],
				[ 'offset', $parser->offset ],
				[ 'output', $parser->output ],
				[ 'stack', $parser->stack ],
			],
		],
		'first'           => parser_freeform_block_observation( $data['first'], $first, $name . ' first' ),
		'second'          => parser_freeform_block_observation( $data['second'], $second, $name . ' second' ),
		'distinct_blocks' => true,
		'arguments_after' => [ 'first' => $first, 'second' => $second ],
	];
}

echo json_encode( $observations, JSON_THROW_ON_ERROR ), "\n";

