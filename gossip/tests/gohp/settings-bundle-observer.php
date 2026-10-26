<?php

ob_start();
require '/input/retained-edge-observer.php';
$retained = json_decode( ob_get_clean(), true, 512, JSON_THROW_ON_ERROR );

ob_start();
require '/input/lookup-observer.php';
$lookup = json_decode( ob_get_clean(), true, 512, JSON_THROW_ON_ERROR );

$cases = json_decode( file_get_contents( '/input/settings-composition-cases.json' ), true, 512, JSON_THROW_ON_ERROR );
if ( ! is_array( $cases ) || ! array_is_list( $cases ) ) {
	throw new RuntimeException( 'settings composition cases must be an ordered list' );
}

$composition = [];
$seen        = [];
foreach ( $cases as $case ) {
	if ( ! is_array( $case ) || array_keys( $case ) !== [ 'name', 'operation', 'input', 'path', 'default', 'selector', 'expected' ] || ! is_string( $case['name'] ) || $case['name'] === '' || isset( $seen[ $case['name'] ] ) || ! is_string( $case['operation'] ) ) {
		throw new RuntimeException( 'invalid settings composition case' );
	}
	$seen[ $case['name'] ] = true;

	$input_entries    = 0;
	$path_entries     = 0;
	$default_entries  = 0;
	$selector_entries = 0;
	$expected_entries = 0;
	$input    = array_get_fixture_decode_node( $case['input'], $case['name'] . ' input', 0, $input_entries );
	$path     = array_get_fixture_decode_node( $case['path'], $case['name'] . ' path', 0, $path_entries );
	$default  = array_get_fixture_decode_node( $case['default'], $case['name'] . ' default', 0, $default_entries );
	$selector = array_get_fixture_decode_node( $case['selector'], $case['name'] . ' selector', 0, $selector_entries );
	$expected = array_get_fixture_decode_node( $case['expected'], $case['name'] . ' expected', 0, $expected_entries );
	if ( ! is_array( $selector ) ) {
		throw new RuntimeException( 'settings selector must be an array: ' . $case['name'] );
	}
	foreach ( $selector as $index => $value ) {
		if ( ! is_int( $value ) && ! is_string( $value ) ) {
			throw new RuntimeException( 'settings selector value must be int or string: ' . $case['name'] . ' index ' . $index );
		}
	}
	$expected_canonical_entries = 0;
	if ( array_get_fixture_encode_node( $expected, $case['name'] . ' expected', 0, $expected_canonical_entries ) !== $case['expected'] ) {
		throw new RuntimeException( 'settings composition expected node is not canonical: ' . $case['name'] );
	}

	$input_before    = $input;
	$path_before     = $path;
	$default_before  = $default;
	$selector_before = $selector;
	$value = match ( $case['operation'] ) {
		'gossip_settings_bool'  => gossip_settings_bool( $input, $path, $default ),
		'gossip_settings_edges' => gossip_settings_edges( $input, $path, $selector ),
		default                 => throw new RuntimeException( 'unknown settings composition operation: ' . $case['operation'] ),
	};

	$value_entries         = 0;
	$input_after_entries   = 0;
	$path_after_entries    = 0;
	$default_after_entries = 0;
	$selector_after_entries = 0;
	$value_node     = array_get_fixture_encode_node( $value, $case['name'] . ' observed value', 0, $value_entries );
	$input_after    = array_get_fixture_encode_node( $input, $case['name'] . ' input after', 0, $input_after_entries );
	$path_after     = array_get_fixture_encode_node( $path, $case['name'] . ' path after', 0, $path_after_entries );
	$default_after  = array_get_fixture_encode_node( $default, $case['name'] . ' default after', 0, $default_after_entries );
	$selector_after = array_get_fixture_encode_node( $selector, $case['name'] . ' selector after', 0, $selector_after_entries );
	if ( get_debug_type( $value ) !== get_debug_type( $expected ) || $value !== $expected || $input !== $input_before || $path !== $path_before || $default !== $default_before || $selector !== $selector_before ) {
		throw new RuntimeException( 'settings composition expectation or mutation mismatch: ' . $case['name'] );
	}
	$composition[] = [
		'name'       => $case['name'],
		'type'       => get_debug_type( $value ),
		'value'      => $value_node,
		'args_after' => [ $input_after, $path_after, $default_after, $selector_after ],
	];
}

echo json_encode( [ 'retained' => $retained, 'lookup' => $lookup, 'composition' => $composition ], JSON_THROW_ON_ERROR ), "\n";
