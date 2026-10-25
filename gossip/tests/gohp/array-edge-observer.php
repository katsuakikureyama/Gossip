<?php

require '/input/helpers.php';
require '/input/transport.php';

if ( PHP_VERSION_ID >= 80400 || setlocale( LC_CTYPE, 'C' ) !== 'C' ) {
	throw new RuntimeException( 'requires PHP8.3 and C locale' );
}

foreach ( [ 'array_first', 'array_last' ] as $name ) {
	$function = new ReflectionFunction( $name );
	if ( ! $function->isUserDefined() || $function->isInternal() ) {
		throw new RuntimeException( 'requires original user-defined polyfill' );
	}
}

$cases        = json_decode( file_get_contents( '/input/cases.json' ), true, 512, JSON_THROW_ON_ERROR );
$observations = [];
$seen         = [];

foreach ( $cases as $case ) {
	if ( ! is_string( $case['name'] ) || isset( $seen[ $case['name'] ] ) ) {
		throw new RuntimeException( 'invalid edge case name' );
	}
	$seen[ $case['name'] ] = true;
	$input                 = numeric_array_fixture_pairs( $case['input'], 'edge input' );
	$value                 = match ( $case['operation'] ) {
		'array_first' => array_first( $input ),
		'array_last'  => array_last( $input ),
		default       => throw new RuntimeException( 'unknown edge operation' ),
	};

	numeric_array_fixture_scalar( $value, 'observed edge result' );
	$type  = get_debug_type( $value );
	$after = numeric_array_fixture_entries( $input );
	if ( $type !== $case['expected']['type'] || $value !== $case['expected']['value'] || $after !== $case['input_after'] ) {
		throw new RuntimeException( 'edge expectation mismatch: ' . $case['name'] );
	}
	$observations[] = [
		'name'        => $case['name'],
		'type'        => $type,
		'value'       => $value,
		'input_after' => $after,
	];
}

echo json_encode( $observations, JSON_THROW_ON_ERROR ), "\n";
