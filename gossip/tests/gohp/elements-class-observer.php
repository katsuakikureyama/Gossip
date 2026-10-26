<?php

require '/input/helpers.php';

if ( PHP_VERSION_ID < 80300 || PHP_VERSION_ID >= 80400 || setlocale( LC_CTYPE, 'C' ) !== 'C' || setlocale( LC_CTYPE, 0 ) !== 'C' ) {
	throw new RuntimeException( 'requires PHP8.3 and C locale' );
}

ob_start();
require '/input/retained-settings-observer.php';
$retained = json_decode( ob_get_clean(), true, 512, JSON_THROW_ON_ERROR );

$function = new ReflectionFunction( 'wp_should_add_elements_class_name' );
if ( ! $function->isUserDefined() || $function->isInternal() ) {
	throw new RuntimeException( 'requires original user-defined wp_should_add_elements_class_name' );
}

$cases = json_decode( file_get_contents( '/input/elements-class-cases.json' ), true, 512, JSON_THROW_ON_ERROR );
if ( ! is_array( $cases ) || ! array_is_list( $cases ) || count( $cases ) !== 40 ) {
	throw new RuntimeException( 'elements-class cases must be an ordered list of 40 cases' );
}

$observations = [];
$seen         = [];
foreach ( $cases as $case ) {
	if ( ! is_array( $case ) || array_keys( $case ) !== [ 'name', 'block', 'options', 'expected' ] || ! is_string( $case['name'] ) || $case['name'] === '' || isset( $seen[ $case['name'] ] ) || ! is_bool( $case['expected'] ) ) {
		throw new RuntimeException( 'invalid elements-class case' );
	}
	$seen[ $case['name'] ] = true;

	$block_entries   = 0;
	$options_entries = 0;
	$block   = array_get_fixture_decode_node( $case['block'], $case['name'] . ' block', 0, $block_entries );
	$options = array_get_fixture_decode_node( $case['options'], $case['name'] . ' options', 0, $options_entries );
	if ( ! is_array( $block ) || ! is_array( $options ) ) {
		throw new RuntimeException( 'elements-class block and options roots must be arrays: ' . $case['name'] );
	}

	$block_before   = $block;
	$options_before = $options;
	$value          = wp_should_add_elements_class_name( $block, $options );
	$block_after_entries   = 0;
	$options_after_entries = 0;
	$block_after   = array_get_fixture_encode_node( $block, $case['name'] . ' block after', 0, $block_after_entries );
	$options_after = array_get_fixture_encode_node( $options, $case['name'] . ' options after', 0, $options_after_entries );
	if ( get_debug_type( $value ) !== 'bool' || $value !== $case['expected'] || $block !== $block_before || $options !== $options_before ) {
		throw new RuntimeException( 'elements-class expectation or mutation mismatch: ' . $case['name'] );
	}

	$observations[] = [
		'name'          => $case['name'],
		'type'          => 'bool',
		'value'         => $value,
		'block_after'   => $block_after,
		'options_after' => $options_after,
	];
}

echo json_encode( [ 'retained' => $retained, 'elements' => $observations ], JSON_THROW_ON_ERROR ), "\n";
