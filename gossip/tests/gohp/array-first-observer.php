<?php

require '/input/helpers.php';
require '/input/transport.php';

if ( PHP_VERSION_ID >= 80400 || setlocale( LC_CTYPE, 'C' ) !== 'C' ) {
	throw new RuntimeException( 'requires PHP8.3 and C locale' );
}

$function = new ReflectionFunction( 'array_first' );
if ( ! $function->isUserDefined() || $function->isInternal() ) {
	throw new RuntimeException( 'requires original user-defined polyfill' );
}

$cases        = json_decode( file_get_contents( '/input/cases.json' ), true, 512, JSON_THROW_ON_ERROR );
$observations = [];
$seen         = [];

foreach ( $cases as $case ) {
	if ( ! is_string( $case['name'] ) || isset( $seen[ $case['name'] ] ) ) {
		throw new RuntimeException( 'invalid array_first case name' );
	}
	$seen[ $case['name'] ] = true;
	$input                 = numeric_array_fixture_pairs( $case['input'], 'array_first input' );
	if ( $case['operation'] !== 'array_first' ) {
		throw new RuntimeException( 'unknown array_first operation' );
	}
	$value = array_first( $input );

	numeric_array_fixture_scalar( $value, 'observed array_first result' );
	$type  = get_debug_type( $value );
	$after = numeric_array_fixture_entries( $input );
	if ( $type !== $case['expected']['type'] || $value !== $case['expected']['value'] || $after !== $case['input_after'] ) {
		throw new RuntimeException( 'array_first expectation mismatch: ' . $case['name'] );
	}
	$observations[] = [
		'name'        => $case['name'],
		'type'        => $type,
		'value'       => $value,
		'input_after' => $after,
	];
}

echo json_encode( $observations, JSON_THROW_ON_ERROR ), "\n";
