<?php

require '/input/helpers.php';

if ( PHP_VERSION_ID < 80300 || PHP_VERSION_ID >= 80400 || setlocale( LC_CTYPE, 'C' ) !== 'C' || setlocale( LC_CTYPE, 0 ) !== 'C' ) {
	throw new RuntimeException( 'requires PHP8.3 and C locale' );
}

$function = new ReflectionFunction( '_wp_array_get' );
if ( ! $function->isUserDefined() || $function->isInternal() ) {
	throw new RuntimeException( 'requires original user-defined _wp_array_get' );
}

const ARRAY_GET_MAX_DEPTH   = 12;
const ARRAY_GET_MAX_ENTRIES = 256;
const ARRAY_GET_MAX_STRING  = 4096;

function array_get_fixture_key( mixed $key, string $field ): int|string {
	if ( ! is_int( $key ) && ! is_string( $key ) ) {
		throw new RuntimeException( $field . ' key must be int or string' );
	}
	if ( is_int( $key ) && ( $key < -1000000 || $key > 1000000 ) ) {
		throw new RuntimeException( $field . ' integer key is outside the bounded range' );
	}
	if ( is_string( $key ) && strlen( $key ) > ARRAY_GET_MAX_STRING ) {
		throw new RuntimeException( $field . ' string key exceeds the bounded size' );
	}
	return $key;
}

function array_get_fixture_normalized_key( int|string $key, string $field ): int|string {
	$probe       = [];
	$probe[ $key ] = true;
	$normalized  = array_key_first( $probe );
	if ( is_int( $normalized ) && ( $normalized < -1000000 || $normalized > 1000000 ) ) {
		throw new RuntimeException( $field . ' normalized integer key is outside the bounded range' );
	}
	return $normalized;
}

function array_get_fixture_decode_node( mixed $node, string $field, int $depth = 0, int &$entries = 0 ): mixed {
	if ( $depth > ARRAY_GET_MAX_DEPTH ) {
		throw new RuntimeException( $field . ' exceeds the bounded depth' );
	}
	if ( is_null( $node ) || is_bool( $node ) ) {
		return $node;
	}
	if ( is_int( $node ) ) {
		if ( $node < -1000000 || $node > 1000000 ) {
			throw new RuntimeException( $field . ' integer is outside the bounded range' );
		}
		return $node;
	}
	if ( is_string( $node ) ) {
		if ( strlen( $node ) > ARRAY_GET_MAX_STRING ) {
			throw new RuntimeException( $field . ' string exceeds the bounded size' );
		}
		return $node;
	}
	if ( ! is_array( $node ) || array_is_list( $node ) || array_keys( $node ) !== [ 'array' ] || ! is_array( $node['array'] ) || ! array_is_list( $node['array'] ) ) {
		throw new RuntimeException( $field . ' must be a scalar node or {array: ordered pairs}' );
	}

	$array = [];
	foreach ( $node['array'] as $index => $pair ) {
		if ( ! is_array( $pair ) || ! array_is_list( $pair ) || count( $pair ) !== 2 ) {
			throw new RuntimeException( $field . ' pair ' . $index . ' must contain exactly key and node' );
		}
		$key            = array_get_fixture_key( $pair[0], $field . ' pair ' . $index );
		$normalized_key = array_get_fixture_normalized_key( $key, $field . ' pair ' . $index );
		if ( array_key_exists( $normalized_key, $array ) ) {
			throw new RuntimeException( $field . ' has duplicate normalized key at pair ' . $index );
		}
		$entries++;
		if ( $entries > ARRAY_GET_MAX_ENTRIES ) {
			throw new RuntimeException( $field . ' exceeds the bounded entry count' );
		}
		$array[ $key ] = array_get_fixture_decode_node( $pair[1], $field . ' pair ' . $index . ' value', $depth + 1, $entries );
	}

	return $array;
}

function array_get_fixture_encode_node( mixed $value, string $field, int $depth = 0, int &$entries = 0 ): mixed {
	if ( $depth > ARRAY_GET_MAX_DEPTH ) {
		throw new RuntimeException( $field . ' exceeds the bounded depth' );
	}
	if ( is_null( $value ) || is_bool( $value ) ) {
		return $value;
	}
	if ( is_int( $value ) ) {
		if ( $value < -1000000 || $value > 1000000 ) {
			throw new RuntimeException( $field . ' integer is outside the bounded range' );
		}
		return $value;
	}
	if ( is_string( $value ) ) {
		if ( strlen( $value ) > ARRAY_GET_MAX_STRING ) {
			throw new RuntimeException( $field . ' string exceeds the bounded size' );
		}
		return $value;
	}
	if ( ! is_array( $value ) ) {
		throw new RuntimeException( $field . ' has unsupported node type ' . get_debug_type( $value ) );
	}

	$pairs = [];
	foreach ( $value as $key => $item ) {
		$key = array_get_fixture_key( $key, $field );
		$entries++;
		if ( $entries > ARRAY_GET_MAX_ENTRIES ) {
			throw new RuntimeException( $field . ' exceeds the bounded entry count' );
		}
		$pairs[] = [ $key, array_get_fixture_encode_node( $item, $field . ' value', $depth + 1, $entries ) ];
	}
	return [ 'array' => $pairs ];
}

$cases = json_decode( file_get_contents( '/input/cases.json' ), true, 512, JSON_THROW_ON_ERROR );
if ( ! is_array( $cases ) || ! array_is_list( $cases ) ) {
	throw new RuntimeException( 'array-get cases must be an ordered list' );
}

$observations = [];
$seen         = [];
foreach ( $cases as $case ) {
	if ( ! is_array( $case ) || array_keys( $case ) !== [ 'name', 'input', 'path', 'default', 'expected', 'omit_default' ] || ! is_string( $case['name'] ) || $case['name'] === '' || isset( $seen[ $case['name'] ] ) || ! is_bool( $case['omit_default'] ) ) {
		throw new RuntimeException( 'invalid array-get case' );
	}
	$seen[ $case['name'] ] = true;

	$input_entries = 0;
	$path_entries = 0;
	$default_entries = 0;
	$expected_entries = 0;
	$input    = array_get_fixture_decode_node( $case['input'], $case['name'] . ' input', 0, $input_entries );
	$path     = array_get_fixture_decode_node( $case['path'], $case['name'] . ' path', 0, $path_entries );
	$default  = array_get_fixture_decode_node( $case['default'], $case['name'] . ' default', 0, $default_entries );
	$expected = array_get_fixture_decode_node( $case['expected'], $case['name'] . ' expected', 0, $expected_entries );
	$expected_canonical_entries = 0;
	if ( array_get_fixture_encode_node( $expected, $case['name'] . ' expected', 0, $expected_canonical_entries ) !== $case['expected'] ) {
		throw new RuntimeException( 'array-get expected node is not canonical: ' . $case['name'] );
	}

	$input_before   = $input;
	$path_before    = $path;
	$default_before = $default;
	$value          = $case['omit_default'] ? _wp_array_get( $input, $path ) : _wp_array_get( $input, $path, $default );
	$value_entries  = 0;
	$value_node     = array_get_fixture_encode_node( $value, $case['name'] . ' observed value', 0, $value_entries );
	$input_after_entries = 0;
	$path_after_entries = 0;
	$default_after_entries = 0;
	$input_after   = array_get_fixture_encode_node( $input, $case['name'] . ' input after', 0, $input_after_entries );
	$path_after    = array_get_fixture_encode_node( $path, $case['name'] . ' path after', 0, $path_after_entries );
	$default_after = array_get_fixture_encode_node( $default, $case['name'] . ' default after', 0, $default_after_entries );
	if ( get_debug_type( $value ) !== get_debug_type( $expected ) || $value !== $expected || $input !== $input_before || $path !== $path_before || $default !== $default_before ) {
		throw new RuntimeException( 'array-get expectation or mutation mismatch: ' . $case['name'] );
	}

	$observations[] = [
		'name'          => $case['name'],
		'type'          => get_debug_type( $value ),
		'value'         => $value_node,
		'input_after'   => $input_after,
		'path_after'    => $path_after,
		'default_after' => $default_after,
	];
}

echo json_encode( $observations, JSON_THROW_ON_ERROR ), "\n";
