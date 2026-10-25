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

function integrated_fixture_type(mixed $value): string {
    return match (get_debug_type($value)) {
        'null', 'bool', 'int', 'string', 'array' => get_debug_type($value),
        default => throw new RuntimeException('fixture transports only null, bool, int, string, and list array values'),
    };
}

function integrated_validate_fixture_value(mixed $value): void {
    if (integrated_fixture_type($value) !== 'array') {
        return;
    }
    if (!array_is_list($value)) {
        throw new RuntimeException('fixture does not transport PHP maps or non-list array keys');
    }
    foreach ($value as $entry) {
        integrated_validate_fixture_value($entry);
    }
}

function integrated_no_args(array $args, string $operation): void {
    if ($args !== []) {
        throw new RuntimeException($operation . ' requires no arguments');
    }
}

function integrated_string_arg(array $args, string $operation): string {
    if (count($args) !== 1 || get_debug_type($args[0]) !== 'string') {
        throw new RuntimeException($operation . ' requires exactly one string argument');
    }
    return $args[0];
}

function integrated_width(mixed $value, string $operation): int {
    if (get_debug_type($value) !== 'int' || $value < 0 || $value > 4096) {
        throw new RuntimeException($operation . ' width must be an int in 0..4096');
    }
    return $value;
}

function integrated_value_and_width(array $args, string $operation): array {
    if (count($args) !== 2 || !in_array(get_debug_type($args[0]), ['string', 'int'], true)) {
        throw new RuntimeException($operation . ' requires a string or int value and int width');
    }
    return [$args[0], integrated_width($args[1], $operation)];
}

function integrated_captures(array $args, string $operation): array {
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

function integrated_array_arg(array $args, string $operation): array {
    if (count($args) !== 1 || !is_array($args[0]) || !array_is_list($args[0])) {
        throw new RuntimeException($operation . ' requires exactly one list array');
    }
    return $args[0];
}

function integrated_observation(string $name, mixed $value): array {
    integrated_validate_fixture_value($value);
    return [
        'name' => $name,
        'type' => integrated_fixture_type($value),
        'value' => $value,
    ];
}

$observations = [];
foreach ($cases as $case) {
    if (!is_array($case) || !is_string($case['name'] ?? null) || !is_string($case['operation'] ?? null) || !is_array($case['args'] ?? null) ||
        !is_array($case['expected'] ?? null) || !is_string($case['expected']['type'] ?? null) || !array_key_exists('value', $case['expected'])) {
        throw new RuntimeException('malformed integrated helper bundle case');
    }
    integrated_validate_fixture_value($case['args']);
    integrated_validate_fixture_value($case['expected']['value']);
    if (integrated_fixture_type($case['expected']['value']) !== $case['expected']['type']) {
        throw new RuntimeException('case expected type is outside the fixture contract: ' . $case['name']);
    }

    $operation = $case['operation'];
    $args = $case['args'];
    switch ($operation) {
        case '__return_true':
            integrated_no_args($args, $operation);
            $value = __return_true();
            break;
        case '__return_false':
            integrated_no_args($args, $operation);
            $value = __return_false();
            break;
        case '__return_zero':
            integrated_no_args($args, $operation);
            $value = __return_zero();
            break;
        case '__return_empty_array':
            integrated_no_args($args, $operation);
            $value = __return_empty_array();
            break;
        case '__return_null':
            integrated_no_args($args, $operation);
            $value = __return_null();
            break;
        case '__return_empty_string':
            integrated_no_args($args, $operation);
            $value = __return_empty_string();
            break;
        case 'wp_validate_boolean':
            $value = wp_validate_boolean(integrated_string_arg($args, $operation));
            break;
        case '_autop_newline_preservation_helper':
            $value = _autop_newline_preservation_helper(integrated_captures($args, $operation));
            break;
        case 'zeroise':
            [$input, $width] = integrated_value_and_width($args, $operation);
            $value = zeroise($input, $width);
            break;
        case 'trailingslashit':
            $value = trailingslashit(integrated_string_arg($args, $operation));
            break;
        case 'untrailingslashit':
            $value = untrailingslashit(integrated_string_arg($args, $operation));
            break;
        case 'stripslashes_from_strings_only':
            if (count($args) !== 1) {
                throw new RuntimeException($operation . ' requires exactly one fixture value');
            }
            $value = stripslashes_from_strings_only($args[0]);
            break;
        case 'gossip_capture_newline_unslash':
            $value = gossip_capture_newline_unslash(integrated_captures($args, $operation));
            break;
        case 'gossip_unslash_array_keys':
            $value = gossip_unslash_array_keys(integrated_array_arg($args, $operation));
            break;
        case 'gossip_unslash_array_shape':
            $value = gossip_unslash_array_shape(integrated_array_arg($args, $operation));
            break;
        case 'gossip_path_pad':
            [$path, $width] = integrated_value_and_width($args, $operation);
            if (get_debug_type($path) !== 'string') {
                throw new RuntimeException($operation . ' requires a string path');
            }
            $value = gossip_path_pad($path, $width);
            break;
        case 'gossip_zero_pad':
            if (count($args) !== 1) {
                throw new RuntimeException($operation . ' requires one width');
            }
            $value = gossip_zero_pad(integrated_width($args[0], $operation));
            break;
        case 'gossip_false_boolean':
            integrated_no_args($args, $operation);
            $value = gossip_false_boolean();
            break;
        case 'gossip_empty_string_path':
            integrated_no_args($args, $operation);
            $value = gossip_empty_string_path();
            break;
        default:
            throw new RuntimeException('unknown operation: ' . $operation);
    }
    $observations[] = integrated_observation($case['name'], $value);
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
