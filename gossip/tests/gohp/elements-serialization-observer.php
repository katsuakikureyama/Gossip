<?php

require '/input/helpers.php';
require '/input/transport.php';

if ( PHP_VERSION_ID < 80300 || PHP_VERSION_ID >= 80400 || setlocale( LC_CTYPE, 'C' ) !== 'C' || setlocale( LC_CTYPE, 0 ) !== 'C' ) {
	throw new RuntimeException( 'requires PHP8.3 and C locale' );
}

foreach ( array( '_wp_array_get', 'wp_should_skip_block_supports_serialization', 'wp_should_add_elements_class_name' ) as $name ) {
	$function = new ReflectionFunction( $name );
	if ( ! $function->isUserDefined() || $function->isInternal() ) {
		throw new RuntimeException( 'requires original user-defined ' . $name );
	}
}
if ( ! function_exists( 'gossip_elements_serialization_decision' ) ) {
	throw new RuntimeException( 'requires owned elements serialization decision driver' );
}

function elements_serialization_scalar( mixed $node, string $field ): mixed {
	$entries = 0;
	$value   = array_get_fixture_decode_node( $node, $field, 0, $entries );
	if ( is_array( $value ) ) {
		throw new RuntimeException( $field . ' must be a scalar node' );
	}
	$canonical_entries = 0;
	if ( array_get_fixture_encode_node( $value, $field, 0, $canonical_entries ) !== $node ) {
		throw new RuntimeException( $field . ' is not a canonical scalar node' );
	}
	return $value;
}

function elements_serialization_node( mixed $value, string $field ): mixed {
	$entries = 0;
	return array_get_fixture_encode_node( $value, $field, 0, $entries );
}

$cases = json_decode( file_get_contents( '/input/cases.json' ), true, 512, JSON_THROW_ON_ERROR );
if ( ! is_array( $cases ) || ! array_is_list( $cases ) || count( $cases ) !== 24 ) {
	throw new RuntimeException( 'elements serialization cases must be an ordered list of 24 cases' );
}

$observations = array();
$seen         = array();
$direct       = 0;
$composition  = 0;
foreach ( $cases as $case ) {
	if ( ! is_array( $case ) || array_keys( $case ) !== array( 'name', 'mode', 'object', 'input', 'feature_set', 'feature', 'block', 'expected' ) || ! is_string( $case['name'] ) || $case['name'] === '' || isset( $seen[ $case['name'] ] ) || ! is_string( $case['mode'] ) || ! in_array( $case['mode'], array( 'direct', 'composition' ), true ) || ! is_bool( $case['object'] ) ) {
		throw new RuntimeException( 'invalid elements serialization case' );
	}
	$seen[ $case['name'] ] = true;
	if ( $case['mode'] === 'direct' ) {
		$direct++;
	} else {
		$composition++;
	}

	$input_entries = 0;
	$block_entries = 0;
	$input = array_get_fixture_decode_node( $case['input'], $case['name'] . ' input', 0, $input_entries );
	$block = array_get_fixture_decode_node( $case['block'], $case['name'] . ' block', 0, $block_entries );
	$feature_set = elements_serialization_scalar( $case['feature_set'], $case['name'] . ' feature_set' );
	$feature = elements_serialization_scalar( $case['feature'], $case['name'] . ' feature' );
	$expected = elements_serialization_scalar( $case['expected'], $case['name'] . ' expected' );
	if ( ! is_array( $block ) || ( $case['mode'] === 'direct' && $block !== array() ) ) {
		throw new RuntimeException( 'elements serialization block has an invalid root: ' . $case['name'] );
	}

	if ( $case['object'] ) {
		// This data fixture only provides the owned supports holder, not a real
		// WP_Block_Type service, registry, or other WordPress runtime behavior.
		$block_type = new GossipBlockSupportsInput();
		$block_type->supports = $input;
		$input_before = $block_type->supports;
	} else {
		$block_type = $input;
		$input_before = $block_type;
	}
	$feature_set_before = $feature_set;
	$feature_before     = $feature;
	$block_before       = $block;
	$value = $case['mode'] === 'direct'
		? wp_should_skip_block_supports_serialization( $block_type, $feature_set, $feature )
		: gossip_elements_serialization_decision( $block, $block_type );
	$input_after_value = $case['object'] ? $block_type->supports : $block_type;
	if ( get_debug_type( $value ) !== get_debug_type( $expected ) || $value !== $expected || $input_after_value !== $input_before || $feature_set !== $feature_set_before || $feature !== $feature_before || $block !== $block_before ) {
		throw new RuntimeException( 'elements serialization expectation or mutation mismatch: ' . $case['name'] );
	}

	$observations[] = array(
		'name'              => $case['name'],
		'type'              => get_debug_type( $value ),
		'value'             => elements_serialization_node( $value, $case['name'] . ' observed value' ),
		'input_after'       => elements_serialization_node( $input_after_value, $case['name'] . ' input after' ),
		'feature_set_after' => elements_serialization_node( $feature_set, $case['name'] . ' feature_set after' ),
		'feature_after'     => elements_serialization_node( $feature, $case['name'] . ' feature after' ),
		'block_after'       => elements_serialization_node( $block, $case['name'] . ' block after' ),
	);
}
if ( $direct !== 16 || $composition !== 8 || count( $seen ) !== 24 ) {
	throw new RuntimeException( 'elements serialization case mode/count mismatch' );
}

echo json_encode( $observations, JSON_THROW_ON_ERROR ), "\n";
