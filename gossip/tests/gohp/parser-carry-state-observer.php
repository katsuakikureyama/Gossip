<?php

// Owned observation driver. It executes the complete pinned WordPress parser;
// it neither substitutes parser code nor provides a generated-Go comparison.
require_once '/input/wp-includes/class-wp-block-parser.php';

if ( PHP_INT_SIZE !== 8 ) {
	throw new RuntimeException( 'requires 64-bit PHP integers' );
}
if ( setlocale( LC_CTYPE, 'C' ) !== 'C' || setlocale( LC_CTYPE, 0 ) !== 'C' ) {
	throw new RuntimeException( 'requires C LC_CTYPE' );
}
foreach ( [ 'pcre.jit' => '0', 'pcre.backtrack_limit' => '100000', 'pcre.recursion_limit' => '10000', 'memory_limit' => '128M' ] as $key => $value ) {
	if ( false === ini_set( $key, $value ) || ini_get( $key ) !== $value ) {
		throw new RuntimeException( 'cannot fix oracle setting: ' . $key );
	}
}

$parser_class     = new ReflectionClass( 'WP_Block_Parser' );
$expected_methods = [ 'parse', 'proceed', 'next_token', 'freeform', 'add_freeform', 'add_inner_block', 'add_block_from_stack' ];
$methods          = $parser_class->getMethods();
if ( ! $parser_class->isUserDefined() || $parser_class->isInternal() || $parser_class->getFileName() !== '/input/wp-includes/class-wp-block-parser.php' || count( $methods ) !== count( $expected_methods ) || array_map( static fn( ReflectionMethod $method ): string => $method->getName(), $methods ) !== $expected_methods ) {
	throw new RuntimeException( 'requires complete original WP_Block_Parser methods' );
}
foreach ( $methods as $method ) {
	if ( ! $method->isPublic() || $method->isInternal() || $method->getDeclaringClass()->getName() !== 'WP_Block_Parser' || $method->getFileName() !== '/input/wp-includes/class-wp-block-parser.php' ) {
		throw new RuntimeException( 'requires public original parser method' );
	}
}
$parse_method = new ReflectionMethod( 'WP_Block_Parser', 'parse' );
if ( $parse_method->getStartLine() !== 64 || $parse_method->getEndLine() !== 75 || $parse_method->getNumberOfParameters() !== 1 || $parse_method->getNumberOfRequiredParameters() !== 1 ) {
	throw new RuntimeException( 'requires original parse entrypoint lines 64-75' );
}
foreach ( [ 'WP_Block_Parser_Block' => '/input/wp-includes/class-wp-block-parser-block.php', 'WP_Block_Parser_Frame' => '/input/wp-includes/class-wp-block-parser-frame.php' ] as $class => $path ) {
	$dependency = new ReflectionClass( $class );
	if ( ! $dependency->isUserDefined() || $dependency->isInternal() || $dependency->getFileName() !== $path ) {
		throw new RuntimeException( 'requires original parser dependency: ' . $class );
	}
}

// Preserve scalar types, associative/list insertion order, and every string
// byte. This is intentionally an observation transport, not an expected result.
function parser_carry_state_value( mixed $value, int $depth = 0 ): array {
	if ( $depth > 64 ) {
		throw new RuntimeException( 'observation nesting exceeds fixture bound' );
	}
	if ( is_string( $value ) ) {
		return [ 'type' => 'string', 'base64' => base64_encode( $value ) ];
	}
	if ( is_null( $value ) || is_bool( $value ) || is_int( $value ) ) {
		return [ 'type' => get_debug_type( $value ), 'value' => $value ];
	}
	if ( is_array( $value ) ) {
		$entries = [];
		foreach ( $value as $key => $child ) {
			$entries[] = [ 'key' => parser_carry_state_value( $key, $depth + 1 ), 'value' => parser_carry_state_value( $child, $depth + 1 ) ];
		}
		return [ 'type' => 'array', 'entries' => $entries ];
	}
	if ( is_object( $value ) ) {
		$properties = [];
		foreach ( get_object_vars( $value ) as $key => $child ) {
			$properties[] = [ 'key' => parser_carry_state_value( $key, $depth + 1 ), 'value' => parser_carry_state_value( $child, $depth + 1 ) ];
		}
		return [ 'type' => 'object', 'class' => get_class( $value ), 'properties' => $properties ];
	}
	throw new RuntimeException( 'value outside carry-state observation domain: ' . get_debug_type( $value ) );
}

function parser_carry_state_errors(): array {
	return [
		'json_error_code'    => json_last_error(),
		'json_error_message' => json_last_error_msg(),
		'pcre_error_code'    => preg_last_error(),
	];
}

function parser_carry_state_seed( string $json_seed, string $pcre_seed ): array {
	if ( $json_seed === 'syntax' ) {
		$json_result = json_decode( '{', true );
		if ( $json_result !== null || json_last_error() !== JSON_ERROR_SYNTAX ) {
			throw new RuntimeException( 'syntax JSON seed did not establish error 4' );
		}
	} elseif ( $json_seed === 'null' ) {
		$json_result = json_decode( 'null', true );
		if ( $json_result !== null || json_last_error() !== JSON_ERROR_NONE ) {
			throw new RuntimeException( 'null JSON seed did not establish error 0' );
		}
	} else {
		throw new RuntimeException( 'unknown JSON seed' );
	}

	$matches = null;
	if ( $pcre_seed === 'clean' ) {
		$pcre_result = preg_match( '/a/s', '', $matches, PREG_OFFSET_CAPTURE, 0 );
		if ( ! is_int( $pcre_result ) || $pcre_result !== 0 || preg_last_error() !== PREG_NO_ERROR ) {
			throw new RuntimeException( 'clean PCRE seed did not establish error 0' );
		}
	} elseif ( $pcre_seed === 'limit' ) {
		$pcre_result = preg_match( '/^(a+)+$/s', 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa!', $matches, PREG_OFFSET_CAPTURE, 0 );
		if ( $pcre_result !== false || preg_last_error() !== PREG_BACKTRACK_LIMIT_ERROR ) {
			throw new RuntimeException( 'limit PCRE seed did not establish error 2' );
		}
	} else {
		throw new RuntimeException( 'unknown PCRE seed' );
	}

	return [
		'json_result' => parser_carry_state_value( $json_result ),
		'pcre_result' => parser_carry_state_value( $pcre_result ),
		'error_state' => parser_carry_state_errors(),
	];
}

function parser_carry_state_step( WP_Block_Parser $parser, string $document_name, string $document, string $document_base64 ): array {
	$document_before  = $document;
	$before           = parser_carry_state_errors();
	$result           = $parser->parse( $document );
	$after            = parser_carry_state_errors();
	$parser_after     = get_object_vars( $parser );
	$caller_unchanged = $document === $document_before;
	if ( ! $caller_unchanged || ! array_key_exists( 'document', $parser_after ) || ! array_key_exists( 'output', $parser_after ) || ! array_key_exists( 'stack', $parser_after ) || $parser_after['document'] !== $document || $parser_after['output'] !== $result || $parser_after['stack'] !== [] ) {
		throw new RuntimeException( 'parser postcondition failed: ' . $document_name );
	}

	return [
		'document_name'    => $document_name,
		'input'            => [ 'document_base64' => $document_base64 ],
		'before'           => $before,
		'result'           => parser_carry_state_value( $result ),
		'after'            => $after,
		'parser_after'     => parser_carry_state_value( $parser_after ),
		'caller_unchanged' => true,
	];
}

function parser_carry_state_pass( WP_Block_Parser $parser, string $instance, array $scenario, array $documents ): array {
	$steps = [];
	$seed  = parser_carry_state_seed( $scenario['json_seed'], $scenario['pcre_seed'] );
	foreach ( $scenario['documents'] as $document_name ) {
		$document = $documents[ $document_name ];
		$steps[]  = parser_carry_state_step( $parser, $document_name, $document['text'], $document['base64'] );
	}
	return [ 'instance' => $instance, 'seed' => $seed, 'steps' => $steps ];
}

// Both input files are decoded before a seed can establish the carry-in state.
$document_cases = json_decode( file_get_contents( '/input/documents.json' ), true, 32, JSON_THROW_ON_ERROR );
$sequences      = json_decode( file_get_contents( '/input/sequences.json' ), true, 32, JSON_THROW_ON_ERROR );
if ( ! is_array( $document_cases ) || ! array_is_list( $document_cases ) || count( $document_cases ) !== 12 ) {
	throw new RuntimeException( 'requires 12 fixed parser documents' );
}
if ( ! is_array( $sequences ) || ! array_is_list( $sequences ) || count( $sequences ) !== 4 ) {
	throw new RuntimeException( 'requires 4 fixed carry-state scenarios' );
}

$documents      = [];
$seen_documents = [];
foreach ( $document_cases as $case ) {
	if ( ! is_array( $case ) || array_keys( $case ) !== [ 'name', 'document_base64' ] || ! is_string( $case['name'] ) || $case['name'] === '' || isset( $seen_documents[ $case['name'] ] ) || ! is_string( $case['document_base64'] ) ) {
		throw new RuntimeException( 'invalid parser document' );
	}
	$document = base64_decode( $case['document_base64'], true );
	if ( $document === false || base64_encode( $document ) !== $case['document_base64'] || strlen( $document ) > 4096 ) {
		throw new RuntimeException( 'document outside fixed parser domain: ' . $case['name'] );
	}
	$seen_documents[ $case['name'] ] = true;
	$documents[ $case['name'] ]      = [ 'text' => $document, 'base64' => $case['document_base64'] ];
}

$seen_scenarios = [];
$total_steps    = 0;
foreach ( $sequences as $scenario ) {
	if ( ! is_array( $scenario ) || array_keys( $scenario ) !== [ 'name', 'json_seed', 'pcre_seed', 'documents' ] || ! is_string( $scenario['name'] ) || $scenario['name'] === '' || isset( $seen_scenarios[ $scenario['name'] ] ) || ! is_string( $scenario['json_seed'] ) || ! in_array( $scenario['json_seed'], [ 'syntax', 'null' ], true ) || ! is_string( $scenario['pcre_seed'] ) || ! in_array( $scenario['pcre_seed'], [ 'clean', 'limit' ], true ) || ! is_array( $scenario['documents'] ) || ! array_is_list( $scenario['documents'] ) || count( $scenario['documents'] ) < 1 || count( $scenario['documents'] ) > 3 ) {
		throw new RuntimeException( 'invalid carry-state scenario' );
	}
	foreach ( $scenario['documents'] as $document_name ) {
		if ( ! is_string( $document_name ) || ! array_key_exists( $document_name, $documents ) ) {
			throw new RuntimeException( 'scenario references unknown parser document' );
		}
	}
	$seen_scenarios[ $scenario['name'] ] = true;
	$total_steps                        += count( $scenario['documents'] );
}
if ( $total_steps !== 10 ) {
	throw new RuntimeException( 'requires 10 carry-state document steps' );
}

$shared_parser = new WP_Block_Parser();
$observations  = [];
foreach ( $sequences as $scenario ) {
	$first = parser_carry_state_pass( new WP_Block_Parser(), 'fresh-sequence', $scenario, $documents );
	$reuse = parser_carry_state_pass( $shared_parser, 'reused-across-sequences', $scenario, $documents );

	$first_comparable = $first;
	$reuse_comparable = $reuse;
	unset( $first_comparable['instance'], $reuse_comparable['instance'] );
	$observations[] = [
		'name'        => $scenario['name'],
		'json_seed'   => $scenario['json_seed'],
		'pcre_seed'   => $scenario['pcre_seed'],
		'first'       => $first,
		'reuse'       => $reuse,
		'reuse_reset' => $first_comparable === $reuse_comparable,
	];
}

echo json_encode( [
	'schema'           => 'gossip.parser-carry-state-observation/v1',
	'parse_call_count' => 20,
	'scenarios'        => $observations,
	'native_environment' => [
		'php_version'    => PHP_VERSION,
		'php_version_id' => PHP_VERSION_ID,
		'pcre_version'   => PCRE_VERSION,
		'int_size'       => PHP_INT_SIZE,
		'lc_ctype'       => setlocale( LC_CTYPE, 0 ),
		'php_os'         => PHP_OS,
		'php_arch'       => php_uname( 'm' ),
		'profile'        => [ 'pcre.jit' => ini_get( 'pcre.jit' ), 'pcre.backtrack_limit' => ini_get( 'pcre.backtrack_limit' ), 'pcre.recursion_limit' => ini_get( 'pcre.recursion_limit' ), 'memory_limit' => ini_get( 'memory_limit' ) ],
	],
], JSON_THROW_ON_ERROR ), "\n";
