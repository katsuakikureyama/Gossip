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

$parser_class = new ReflectionClass( 'WP_Block_Parser' );
$expected_methods = [ 'parse', 'proceed', 'next_token', 'freeform', 'add_freeform', 'add_inner_block', 'add_block_from_stack' ];
$methods = $parser_class->getMethods();
if ( ! $parser_class->isUserDefined() || $parser_class->isInternal() || $parser_class->getFileName() !== '/input/wp-includes/class-wp-block-parser.php' || count( $methods ) !== count( $expected_methods ) || array_map( static fn( ReflectionMethod $method ): string => $method->getName(), $methods ) !== $expected_methods ) {
	throw new RuntimeException( 'requires complete original WP_Block_Parser methods' );
}
$method_observation = [];
foreach ( $methods as $method ) {
	if ( ! $method->isPublic() || $method->isInternal() || $method->getDeclaringClass()->getName() !== 'WP_Block_Parser' || $method->getFileName() !== '/input/wp-includes/class-wp-block-parser.php' ) {
		throw new RuntimeException( 'requires public original parser method' );
	}
	$method_observation[] = [ 'name' => $method->getName(), 'first_line' => $method->getStartLine(), 'last_line' => $method->getEndLine() ];
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
function full_parser_reference_value( mixed $value, int $depth = 0 ): array {
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
			$entries[] = [ 'key' => full_parser_reference_value( $key, $depth + 1 ), 'value' => full_parser_reference_value( $child, $depth + 1 ) ];
		}
		return [ 'type' => 'array', 'entries' => $entries ];
	}
	if ( is_object( $value ) ) {
		$properties = [];
		foreach ( get_object_vars( $value ) as $key => $child ) {
			$properties[] = [ 'key' => full_parser_reference_value( $key, $depth + 1 ), 'value' => full_parser_reference_value( $child, $depth + 1 ) ];
		}
		return [ 'type' => 'object', 'class' => get_class( $value ), 'properties' => $properties ];
	}
	throw new RuntimeException( 'value outside full parser observation domain: ' . get_debug_type( $value ) );
}

$cases = json_decode( file_get_contents( '/input/cases.json' ), true, 32, JSON_THROW_ON_ERROR );
if ( ! is_array( $cases ) || ! array_is_list( $cases ) || count( $cases ) !== 12 ) {
	throw new RuntimeException( 'requires 12 fixed full-parser cases' );
}
$parser = new WP_Block_Parser();
$observations = [];
$seen = [];
foreach ( $cases as $case ) {
	if ( ! is_array( $case ) || array_keys( $case ) !== [ 'name', 'document_base64' ] || ! is_string( $case['name'] ) || $case['name'] === '' || isset( $seen[ $case['name'] ] ) || ! is_string( $case['document_base64'] ) ) {
		throw new RuntimeException( 'invalid full-parser case' );
	}
	$document = base64_decode( $case['document_base64'], true );
	if ( $document === false || base64_encode( $document ) !== $case['document_base64'] || strlen( $document ) > 4096 ) {
		throw new RuntimeException( 'case outside fixed full-parser domain: ' . $case['name'] );
	}
	$seen[ $case['name'] ] = true;
	$document_before = $document;
	json_decode( 'null', true ); // Defined JSON-error baseline before parser-owned json_decode.
	$fresh_parser = new WP_Block_Parser();
	$first = $fresh_parser->parse( $document );
	$first_pcre_error = preg_last_error();
	$first_json_error = json_last_error();
	$first_json_message = json_last_error_msg();
	$first_state = get_object_vars( $fresh_parser );
	json_decode( 'null', true ); // Same error baseline, without resetting reused parser state.
	$reuse = $parser->parse( $document );
	$reuse_pcre_error = preg_last_error();
	$reuse_json_error = json_last_error();
	$reuse_json_message = json_last_error_msg();
	$reuse_state = get_object_vars( $parser );
	$caller_unchanged = $document === $document_before;
	$reuse_reset = $first === $reuse && $first_state === $reuse_state && $first_pcre_error === $reuse_pcre_error && $first_json_error === $reuse_json_error && $reuse_state['document'] === $document && $reuse_state['output'] === $reuse && $reuse_state['stack'] === [];
	if ( ! $caller_unchanged || ! $reuse_reset ) {
		throw new RuntimeException( 'caller mutation or parser reuse/reset mismatch: ' . $case['name'] );
	}
	$observations[] = [
		'name' => $case['name'],
		'input' => [ 'document_base64' => base64_encode( $document ) ],
		'first_parse' => [ 'instance' => 'fresh', 'result' => full_parser_reference_value( $first ), 'pcre_error_code' => $first_pcre_error, 'json_error_code' => $first_json_error, 'json_error_message' => $first_json_message, 'parser_after' => full_parser_reference_value( $first_state ) ],
		'reuse_parse' => [ 'instance' => 'reused-across-cases', 'result' => full_parser_reference_value( $reuse ), 'pcre_error_code' => $reuse_pcre_error, 'json_error_code' => $reuse_json_error, 'json_error_message' => $reuse_json_message, 'parser_after' => full_parser_reference_value( $reuse_state ) ],
		'caller_unchanged' => true,
		'reuse_reset' => true,
	];
}

echo json_encode( [
	'schema' => 'gossip.parser-full-native-observation/v1',
	'status' => 'native-reference-recorded',
	'php_version' => PHP_VERSION,
	'php_version_id' => PHP_VERSION_ID,
	'pcre_version' => PCRE_VERSION,
	'pcre_jit_support' => PCRE_JIT_SUPPORT,
	'lc_ctype' => setlocale( LC_CTYPE, 0 ),
	'int_size' => PHP_INT_SIZE,
	'php_os' => PHP_OS,
	'php_arch' => php_uname( 'm' ),
	'profile' => [ 'pcre.jit' => ini_get( 'pcre.jit' ), 'pcre.backtrack_limit' => ini_get( 'pcre.backtrack_limit' ), 'pcre.recursion_limit' => ini_get( 'pcre.recursion_limit' ), 'memory_limit' => ini_get( 'memory_limit' ) ],
	'entrypoint' => [ 'class' => 'WP_Block_Parser', 'name' => 'parse', 'first_line' => $parse_method->getStartLine(), 'last_line' => $parse_method->getEndLine() ],
	'methods' => $method_observation,
	'cases' => $observations,
], JSON_THROW_ON_ERROR ), "\n";
