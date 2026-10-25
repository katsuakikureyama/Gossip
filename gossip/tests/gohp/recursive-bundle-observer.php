<?php

require '/input/helpers.php';

if (setlocale(LC_CTYPE, 'C') !== 'C' || setlocale(LC_CTYPE, 0) !== 'C') {
    throw new RuntimeException('LC_CTYPE must be C for this bounded observation');
}

$cases = json_decode(
    file_get_contents('/input/cases.json'),
    true,
    512,
    JSON_THROW_ON_ERROR
);

function recursive_fixture_type(mixed $value): string {
    return match (get_debug_type($value)) {
        'null', 'bool', 'int', 'string', 'array' => get_debug_type($value),
        default => throw new RuntimeException('fixture transports only null, bool, int, string, and list array values'),
    };
}

function recursive_validate_fixture_value(mixed $value): void {
    if (recursive_fixture_type($value) !== 'array') {
        return;
    }
    if (!array_is_list($value)) {
        throw new RuntimeException('fixture does not transport PHP maps or non-list array keys');
    }
    foreach ($value as $entry) {
        recursive_validate_fixture_value($entry);
    }
}

function recursive_no_args(array $args, string $operation): void {
    if ($args !== []) {
        throw new RuntimeException($operation . ' requires no arguments');
    }
}

function recursive_string_arg(array $args, string $operation): string {
    if (count($args) !== 1 || get_debug_type($args[0]) !== 'string') {
        throw new RuntimeException($operation . ' requires exactly one string argument');
    }
    return $args[0];
}

function recursive_any_arg(array $args, string $operation): mixed {
    if (count($args) !== 1) {
        throw new RuntimeException($operation . ' requires exactly one fixture value');
    }
    recursive_validate_fixture_value($args[0]);
    return $args[0];
}

function recursive_width(mixed $value, string $operation): int {
    if (get_debug_type($value) !== 'int' || $value < 0 || $value > 4096) {
        throw new RuntimeException($operation . ' width must be an int in 0..4096');
    }
    return $value;
}

function recursive_value_and_width(array $args, string $operation): array {
    if (count($args) !== 2 || !in_array(get_debug_type($args[0]), ['string', 'int'], true)) {
        throw new RuntimeException($operation . ' requires a string or int value and int width');
    }
    return [$args[0], recursive_width($args[1], $operation)];
}

function recursive_captures(array $args, string $operation): array {
    if (count($args) !== 1 || !is_array($args[0]) || !array_is_list($args[0]) || count($args[0]) === 0) {
        throw new RuntimeException($operation . ' requires one non-empty list of string captures');
    }
    foreach ($args[0] as $capture) {
        if (get_debug_type($capture) !== 'string') {
            throw new RuntimeException($operation . ' captures must all be strings');
        }
    }
    return $args[0];
}

function recursive_array_arg(array $args, string $operation): array {
    if (count($args) !== 1 || !is_array($args[0]) || !array_is_list($args[0])) {
        throw new RuntimeException($operation . ' requires exactly one list array');
    }
    return $args[0];
}

function recursive_observation(string $name, mixed $value, array $args): array {
    recursive_validate_fixture_value($value);
    recursive_validate_fixture_value($args);
    return [
        'name' => $name,
        'type' => recursive_fixture_type($value),
        'value' => $value,
        'args_after' => $args,
    ];
}

$observations = [];
foreach ($cases as $case) {
    if (!is_array($case) || !is_string($case['name'] ?? null) || !is_string($case['operation'] ?? null) || !is_array($case['args'] ?? null) ||
        !is_array($case['expected'] ?? null) || !is_string($case['expected']['type'] ?? null) || !array_key_exists('value', $case['expected'])) {
        throw new RuntimeException('malformed recursive helper bundle case');
    }
    recursive_validate_fixture_value($case['args']);
    recursive_validate_fixture_value($case['expected']['value']);
    if (recursive_fixture_type($case['expected']['value']) !== $case['expected']['type']) {
        throw new RuntimeException('case expected type is outside the fixture contract: ' . $case['name']);
    }

    $operation = $case['operation'];
    $args = $case['args'];
    switch ($operation) {
        case '__return_true':
            recursive_no_args($args, $operation);
            $value = __return_true();
            break;
        case '__return_false':
            recursive_no_args($args, $operation);
            $value = __return_false();
            break;
        case '__return_zero':
            recursive_no_args($args, $operation);
            $value = __return_zero();
            break;
        case '__return_empty_array':
            recursive_no_args($args, $operation);
            $value = __return_empty_array();
            break;
        case '__return_null':
            recursive_no_args($args, $operation);
            $value = __return_null();
            break;
        case '__return_empty_string':
            recursive_no_args($args, $operation);
            $value = __return_empty_string();
            break;
        case 'wp_validate_boolean':
            $value = wp_validate_boolean(recursive_string_arg($args, $operation));
            break;
        case '_autop_newline_preservation_helper':
            $value = _autop_newline_preservation_helper(recursive_captures($args, $operation));
            break;
        case 'zeroise':
            [$input, $width] = recursive_value_and_width($args, $operation);
            $value = zeroise($input, $width);
            break;
        case 'trailingslashit':
            $value = trailingslashit(recursive_string_arg($args, $operation));
            break;
        case 'untrailingslashit':
            $value = untrailingslashit(recursive_string_arg($args, $operation));
            break;
        case 'stripslashes_from_strings_only':
            $value = stripslashes_from_strings_only(recursive_any_arg($args, $operation));
            break;
        case 'wp_slash':
            $value = wp_slash(recursive_any_arg($args, $operation));
            break;
        case 'gossip_capture_newline_unslash':
            $value = gossip_capture_newline_unslash(recursive_captures($args, $operation));
            break;
        case 'gossip_unslash_array_keys':
            $value = gossip_unslash_array_keys(recursive_array_arg($args, $operation));
            break;
        case 'gossip_unslash_array_shape':
            $value = gossip_unslash_array_shape(recursive_array_arg($args, $operation));
            break;
        case 'gossip_path_pad':
            [$path, $width] = recursive_value_and_width($args, $operation);
            if (get_debug_type($path) !== 'string') {
                throw new RuntimeException('gossip_path_pad requires a string path');
            }
            $value = gossip_path_pad($path, $width);
            break;
        case 'gossip_zero_pad':
            if (count($args) !== 1) {
                throw new RuntimeException('gossip_zero_pad requires one width');
            }
            $value = gossip_zero_pad(recursive_width($args[0], $operation));
            break;
        case 'gossip_false_boolean':
            recursive_no_args($args, $operation);
            $value = gossip_false_boolean();
            break;
        case 'gossip_empty_string_path':
            recursive_no_args($args, $operation);
            $value = gossip_empty_string_path();
            break;
        case 'gossip_slash_scalar_roundtrip':
            $value = gossip_slash_scalar_roundtrip(recursive_string_arg($args, $operation));
            break;
        case 'gossip_slash_list_shallow_roundtrip':
            $value = gossip_slash_list_shallow_roundtrip(recursive_array_arg($args, $operation));
            break;
        case 'gossip_slash_capture_newline_unslash':
            $value = gossip_slash_capture_newline_unslash(recursive_captures($args, $operation));
            break;
        case 'gossip_slash_array_shape':
            $value = gossip_slash_array_shape(recursive_array_arg($args, $operation));
            break;
        default:
            throw new RuntimeException('unknown operation: ' . $operation);
    }
    $observations[] = recursive_observation($case['name'], $value, $args);
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
