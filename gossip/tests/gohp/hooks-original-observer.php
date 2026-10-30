<?php

// Native observation driver for the pinned, complete WordPress hook API. It
// only calls the original public API; the callbacks below are fixture inputs,
// not a substitute hook implementation.
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

$expected_workflows = [
	'no_callbacks',
	'priority_order',
	'accepted_arguments',
	'callback_identity_removal',
	'actions',
	'reentrant_filter',
	'mutation_during_iteration',
	'sentinel_identity',
];
$cases = json_decode( file_get_contents( '/input/cases.json' ), true, 16, JSON_THROW_ON_ERROR );
if ( ! is_array( $cases ) || ! array_is_list( $cases ) || $cases !== $expected_workflows ) {
	throw new RuntimeException( 'requires the fixed hook workflow sequence' );
}

// plugin.php inspects this caller global when normalizing preinitialized hooks.
// This explicit empty precondition does not exercise that migration branch.
$wp_filter = array();
require_once '/input/wp-includes/plugin.php';

foreach ( [
	'WP_Hook'            => '/input/wp-includes/class-wp-hook.php',
	'WP_Filter_Sentinel' => '/input/wp-includes/class-wp-filter-sentinel.php',
] as $class => $path ) {
	$reflection = new ReflectionClass( $class );
	if ( ! $reflection->isUserDefined() || $reflection->isInternal() || $reflection->getFileName() !== $path ) {
		throw new RuntimeException( 'requires original class: ' . $class );
	}
}
foreach ( [ 'add_filter', 'apply_filters', 'remove_filter', 'add_action', 'do_action', 'current_filter', 'did_filter', 'did_action' ] as $function ) {
	$reflection = new ReflectionFunction( $function );
	if ( ! $reflection->isUserDefined() || $reflection->isInternal() || $reflection->getFileName() !== '/input/wp-includes/plugin.php' ) {
		throw new RuntimeException( 'requires original plugin API function: ' . $function );
	}
}

/** @return array{type:string,value?:null|bool|int,base64?:string,entries?:array} */
function hooks_original_typed( mixed $value, int $depth = 0 ): array {
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
			$entries[] = [ 'key' => hooks_original_typed( $key, $depth + 1 ), 'value' => hooks_original_typed( $child, $depth + 1 ) ];
		}
		return [ 'type' => 'array', 'entries' => $entries ];
	}
	throw new RuntimeException( 'value outside hook observation transport: ' . get_debug_type( $value ) );
}

function hooks_original_reset_globals(): void {
	global $wp_filter, $wp_actions, $wp_filters, $wp_current_filter;
	$wp_filter         = array();
	$wp_actions        = array();
	$wp_filters        = array();
	$wp_current_filter = array();
}

/** @return array<string,mixed> */
function hooks_original_private_state( WP_Hook $hook ): array {
	$reflection = new ReflectionClass( $hook );
	$state      = [];
	foreach ( [ 'priorities', 'iterations', 'current_priority', 'nesting_level', 'doing_action' ] as $property_name ) {
		$property = $reflection->getProperty( $property_name );
		$state[ $property_name ] = $property->getValue( $hook );
	}
	return $state;
}

/** @return array<string,mixed> */
function hooks_original_state(): array {
	global $wp_filter, $wp_actions, $wp_filters, $wp_current_filter;
	$hooks = [];
	foreach ( $wp_filter as $name => $hook ) {
		if ( ! $hook instanceof WP_Hook ) {
			throw new RuntimeException( 'unexpected non-WP_Hook global entry' );
		}
		$callback_counts = [];
		foreach ( $hook->callbacks as $priority => $callbacks ) {
			$callback_counts[ $priority ] = count( $callbacks );
		}
		$hooks[ $name ] = [
			'callback_counts' => $callback_counts,
			'private'         => hooks_original_private_state( $hook ),
		];
	}
	return [
		'wp_filter'         => $hooks,
		'wp_actions'        => $wp_actions,
		'wp_filters'        => $wp_filters,
		'wp_current_filter' => $wp_current_filter,
	];
}

/** @param list<mixed> $args @return array<string,mixed> */
function hooks_original_event( string $label, array $args ): array {
	global $wp_filter;
	$hook     = current_filter();
	$priority = false;
	if ( false !== $hook && isset( $wp_filter[ $hook ] ) ) {
		$priority = $wp_filter[ $hook ]->current_priority();
	}
	return [
		'label'            => $label,
		'args'             => $args,
		'current_filter'   => $hook,
		'current_priority' => $priority,
		'doing_filter'     => doing_filter(),
		'doing_action'     => doing_action(),
		'did_filter'       => false === $hook ? 0 : did_filter( $hook ),
		'did_action'       => false === $hook ? 0 : did_action( $hook ),
	];
}

/** @return array<string,bool> */
function hooks_original_sentinel_summary( mixed $value, WP_Filter_Sentinel $original ): array {
	return [
		'same_as_input' => $value === $original,
		'is_null'       => null === $value,
		'is_false'      => false === $value,
		'is_sentinel'   => $value instanceof WP_Filter_Sentinel,
	];
}

/** @return array{result:array<string,mixed>,events:list<array<string,mixed>>} */
function hooks_original_run( string $name ): array {
	$events = [];
	$result = [];

	switch ( $name ) {
		case 'no_callbacks':
			$filtered = apply_filters( 'empty', 'seed' );
			$action_return = do_action( 'empty_action' );
			$result = [
				'filter_result'       => $filtered,
				'action_return'        => $action_return,
				'did_filter_empty'    => did_filter( 'empty' ),
				'did_action_empty'    => did_action( 'empty_action' ),
				'current_filter'      => current_filter(),
				'doing_filter_empty'  => doing_filter( 'empty' ),
				'doing_action_empty'  => doing_action( 'empty_action' ),
				'doing_filter_any'    => doing_filter(),
				'doing_action_any'    => doing_action(),
			];
			break;

		case 'priority_order':
			$twenty = static function ( string $value ) use ( &$events ): string {
				$events[] = hooks_original_event( 'priority-20', [ $value ] );
				return $value . '|twenty';
			};
			$first = static function ( string $value ) use ( &$events ): string {
				$events[] = hooks_original_event( 'priority-10-first', [ $value ] );
				return $value . '|first';
			};
			$second = static function ( string $value ) use ( &$events ): string {
				$events[] = hooks_original_event( 'priority-10-second', [ $value ] );
				return $value . '|second';
			};
			$registration = [ add_filter( 'order', $twenty, 20 ), add_filter( 'order', $first, 10 ), add_filter( 'order', $second, 10 ) ];
			$result = [
				'registration' => $registration,
				'filter_result' => apply_filters( 'order', 'seed' ),
				'did_filter' => did_filter( 'order' ),
				'current_filter' => current_filter(),
				'doing_filter' => doing_filter( 'order' ),
				'doing_action' => doing_action( 'order' ),
			];
			break;

		case 'accepted_arguments':
			$zero = static function () use ( &$events ): string {
				$events[] = hooks_original_event( 'accepted-0', [ func_num_args() ] );
				return 'zero';
			};
			$one = static function ( string $value ) use ( &$events ): string {
				$events[] = hooks_original_event( 'accepted-1', [ $value, func_num_args() ] );
				return $value . '|one';
			};
			$two = static function ( string $value, string $extra ) use ( &$events ): string {
				$events[] = hooks_original_event( 'accepted-2', [ $value, $extra, func_num_args() ] );
				return $value . '|' . $extra;
			};
			$result = [
				'registration' => [ add_filter( 'args', $zero, 5, 0 ), add_filter( 'args', $one, 10, 1 ), add_filter( 'args', $two, 15, 2 ) ],
				'filter_result' => apply_filters( 'args', 'seed', 'extra', 'ignored' ),
				'did_filter' => did_filter( 'args' ),
			];
			break;

		case 'callback_identity_removal':
			$same = static function ( string $value ) use ( &$events ): string {
				$events[] = hooks_original_event( 'same', [ $value ] );
				return $value . '|same';
			};
			$other = static function ( string $value ) use ( &$events ): string {
				$events[] = hooks_original_event( 'other', [ $value ] );
				return $value . '|other';
			};
			$registrations = [ add_filter( 'identity', $same, 10 ), add_filter( 'identity', $same, 10 ), add_filter( 'identity', $same, 20 ), add_filter( 'identity', $other, 10 ) ];
			$removed = remove_filter( 'identity', $same, 10 );
			$removed_again = remove_filter( 'identity', $same, 10 );
			$has_after_remove = has_filter( 'identity', $same );
			$applied = apply_filters( 'identity', 'start' );
			$remove_all = remove_all_filters( 'identity' );
			$result = [
				'registrations' => $registrations,
				'removed' => $removed,
				'removed_again' => $removed_again,
				'has_after_remove' => $has_after_remove,
				'filter_result' => $applied,
				'remove_all' => $remove_all,
				'has_after_all' => has_filter( 'identity' ),
			];
			break;

		case 'actions':
			$action = static function ( string $first, string $second ) use ( &$events ): string {
				$events[] = hooks_original_event( 'action', [ $first, $second, func_num_args(), current_action() ] );
				return 'ignored:' . $first . ':' . $second;
			};
			$registration = add_action( 'event', $action, 10, 2 );
			$return = do_action( 'event', 'x', 'y' );
			$result = [
				'registration' => $registration,
				'action_return' => $return,
				'did_action' => did_action( 'event' ),
				'current_action' => current_action(),
				'doing_action' => doing_action( 'event' ),
				'doing_filter' => doing_filter( 'event' ),
			];
			break;

		case 'reentrant_filter':
			$depth = 0;
			$nested_result = null;
			$head = static function ( string $value ) use ( &$events, &$depth, &$nested_result ): string {
				$events[] = hooks_original_event( 'reentrant-head', [ $value, $depth ] );
				if ( 0 === $depth ) {
					++$depth;
					$nested_result = apply_filters( 'recursive', 'inner' );
					--$depth;
				}
				return $value . '|head';
			};
			$tail = static function ( string $value ) use ( &$events, &$depth ): string {
				$events[] = hooks_original_event( 'reentrant-tail', [ $value, $depth ] );
				return $value . '|tail';
			};
			$result = [
				'registration' => [ add_filter( 'recursive', $head, 10 ), add_filter( 'recursive', $tail, 20 ) ],
				'filter_result' => apply_filters( 'recursive', 'outer' ),
				'nested_result' => $nested_result,
				'did_filter' => did_filter( 'recursive' ),
				'current_filter' => current_filter(),
				'doing_filter' => doing_filter( 'recursive' ),
			];
			break;

		case 'mutation_during_iteration':
			$mutated = false;
			$scheduled = static function ( string $value ) use ( &$events ): string {
				$events[] = hooks_original_event( 'scheduled-20', [ $value ] );
				return $value . '|scheduled';
			};
			$added = static function ( string $value ) use ( &$events ): string {
				$events[] = hooks_original_event( 'added-15', [ $value ] );
				return $value . '|added';
			};
			$mutator = static function ( string $value ) use ( &$events, &$mutated, &$scheduled, &$added ): string {
				$events[] = hooks_original_event( 'mutator-10', [ $value, $mutated ] );
				if ( ! $mutated ) {
					$mutated = true;
					remove_filter( 'mutation', $scheduled, 20 );
					add_filter( 'mutation', $added, 15 );
				}
				return $value . '|mutator';
			};
			$tail = static function ( string $value ) use ( &$events ): string {
				$events[] = hooks_original_event( 'tail-30', [ $value ] );
				return $value . '|tail';
			};
			$result = [
				'registration' => [ add_filter( 'mutation', $mutator, 10 ), add_filter( 'mutation', $scheduled, 20 ), add_filter( 'mutation', $tail, 30 ) ],
				'filter_result' => apply_filters( 'mutation', 'start' ),
				'did_filter' => did_filter( 'mutation' ),
				'mutated' => $mutated,
			];
			break;

		case 'sentinel_identity':
			$sentinel = new WP_Filter_Sentinel();
			$none = apply_filters( 'sentinel-none', $sentinel );
			$pass = static function ( WP_Filter_Sentinel $value ) use ( &$events ): WP_Filter_Sentinel {
				$events[] = hooks_original_event( 'sentinel-pass-through', [] );
				return $value;
			};
			$new = static function ( WP_Filter_Sentinel $value ) use ( &$events ): WP_Filter_Sentinel {
				$events[] = hooks_original_event( 'sentinel-new', [] );
				return new WP_Filter_Sentinel();
			};
			$to_null = static function ( WP_Filter_Sentinel $value ) use ( &$events ): mixed {
				$events[] = hooks_original_event( 'sentinel-null', [] );
				return null;
			};
			$to_false = static function ( WP_Filter_Sentinel $value ) use ( &$events ): mixed {
				$events[] = hooks_original_event( 'sentinel-false', [] );
				return false;
			};
			add_filter( 'sentinel-pass', $pass );
			add_filter( 'sentinel-new', $new );
			add_filter( 'sentinel-null', $to_null );
			add_filter( 'sentinel-false', $to_false );
			$result = [
				'no_callback' => hooks_original_sentinel_summary( $none, $sentinel ),
				'pass_through' => hooks_original_sentinel_summary( apply_filters( 'sentinel-pass', $sentinel ), $sentinel ),
				'new_sentinel' => hooks_original_sentinel_summary( apply_filters( 'sentinel-new', $sentinel ), $sentinel ),
				'null_callback' => hooks_original_sentinel_summary( apply_filters( 'sentinel-null', $sentinel ), $sentinel ),
				'false_callback' => hooks_original_sentinel_summary( apply_filters( 'sentinel-false', $sentinel ), $sentinel ),
			];
			break;

		default:
			throw new RuntimeException( 'unknown hook workflow: ' . $name );
	}

	return [ 'result' => $result, 'events' => $events ];
}

$workflows = [];
foreach ( $cases as $name ) {
	hooks_original_reset_globals();
	$initial = hooks_original_state();
	$observation = hooks_original_run( $name );
	$workflows[] = [
		'name' => $name,
		'initial' => hooks_original_typed( $initial ),
		'result' => hooks_original_typed( $observation['result'] ),
		'events' => hooks_original_typed( $observation['events'] ),
		'state_after' => hooks_original_typed( hooks_original_state() ),
	];
}

echo json_encode( [
	'schema' => 'gossip.hooks-original-native-observation/v1',
	'workflow_count' => count( $workflows ),
	'native_environment' => [
		'php_version' => PHP_VERSION,
		'php_version_id' => PHP_VERSION_ID,
		'pcre_version' => PCRE_VERSION,
		'int_size' => PHP_INT_SIZE,
		'lc_ctype' => setlocale( LC_CTYPE, 0 ),
		'php_os' => PHP_OS,
		'php_arch' => php_uname( 'm' ),
		'profile' => [
			'pcre.jit' => ini_get( 'pcre.jit' ),
			'pcre.backtrack_limit' => ini_get( 'pcre.backtrack_limit' ),
			'pcre.recursion_limit' => ini_get( 'pcre.recursion_limit' ),
			'memory_limit' => ini_get( 'memory_limit' ),
		],
	],
	'workflows' => $workflows,
], JSON_THROW_ON_ERROR ), "\n";
