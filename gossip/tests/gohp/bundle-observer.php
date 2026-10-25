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

function bundle_require_no_args(array $args, string $operation): void {
    if ($args !== []) {
        throw new RuntimeException($operation . ' requires no arguments');
    }
}

function bundle_require_string(array $args, string $operation): string {
    if (count($args) !== 1 || get_debug_type($args[0]) !== 'string') {
        throw new RuntimeException($operation . ' requires exactly one string argument');
    }
    return $args[0];
}

function bundle_require_width($value, string $operation): int {
    if (get_debug_type($value) !== 'int' || $value < 0 || $value > 4096) {
        throw new RuntimeException($operation . ' width must be an int in 0..4096');
    }
    return $value;
}

function bundle_require_width_arg(array $args, string $operation): int {
    if (count($args) !== 1) {
        throw new RuntimeException($operation . ' requires exactly one int width');
    }
    return bundle_require_width($args[0], $operation);
}

function bundle_require_value_and_width(array $args, string $operation) {
    if (count($args) !== 2 || !in_array(get_debug_type($args[0]), ['string', 'int'], true)) {
        throw new RuntimeException($operation . ' requires a string or int value and int width');
    }
    bundle_require_width($args[1], $operation);
    return [$args[0], $args[1]];
}

function bundle_require_path_and_width(array $args, string $operation): array {
    if (count($args) !== 2 || get_debug_type($args[0]) !== 'string') {
        throw new RuntimeException($operation . ' requires a string path and int width');
    }
    return [$args[0], bundle_require_width($args[1], $operation)];
}

function bundle_observation(string $name, $value): array {
    $type = get_debug_type($value);
    return match ($type) {
        'bool', 'int', 'string', 'null' => [
            'name' => $name,
            'type' => $type,
            'value' => $value,
        ],
        'array' => $value === [] ? [
            'name' => $name,
            'type' => 'array',
            'value' => [],
        ] : throw new RuntimeException('only empty array results are in scope'),
        default => throw new RuntimeException('unsupported result type: ' . $type),
    };
}

$observations = [];
foreach ($cases as $case) {
    if (!is_string($case['name'] ?? null) || !is_string($case['operation'] ?? null) || !is_array($case['args'] ?? null)) {
        throw new RuntimeException('case must have string name, string operation, and array args');
    }

    $operation = $case['operation'];
    $args      = $case['args'];
    switch ($operation) {
        case '__return_true':
            bundle_require_no_args($args, $operation);
            $value = __return_true();
            break;
        case '__return_false':
            bundle_require_no_args($args, $operation);
            $value = __return_false();
            break;
        case '__return_zero':
            bundle_require_no_args($args, $operation);
            $value = __return_zero();
            break;
        case '__return_empty_array':
            bundle_require_no_args($args, $operation);
            $value = __return_empty_array();
            break;
        case '__return_null':
            bundle_require_no_args($args, $operation);
            $value = __return_null();
            break;
        case '__return_empty_string':
            bundle_require_no_args($args, $operation);
            $value = __return_empty_string();
            break;
        case 'trailingslashit':
            $value = trailingslashit(bundle_require_string($args, $operation));
            break;
        case 'untrailingslashit':
            $value = untrailingslashit(bundle_require_string($args, $operation));
            break;
        case 'wp_validate_boolean':
            $value = wp_validate_boolean(bundle_require_string($args, $operation));
            break;
        case 'zeroise':
            [$input, $width] = bundle_require_value_and_width($args, $operation);
            $value = zeroise($input, $width);
            break;
        case 'gossip_path_pad':
            [$path, $width] = bundle_require_path_and_width($args, $operation);
            $value = gossip_path_pad($path, $width);
            break;
        case 'gossip_empty_array_boolean':
            bundle_require_no_args($args, $operation);
            $value = gossip_empty_array_boolean();
            break;
        case 'gossip_zero_pad':
            $value = gossip_zero_pad(bundle_require_width_arg($args, $operation));
            break;
        case 'gossip_false_boolean':
            bundle_require_no_args($args, $operation);
            $value = gossip_false_boolean();
            break;
        case 'gossip_empty_string_path':
            bundle_require_no_args($args, $operation);
            $value = gossip_empty_string_path();
            break;
        default:
            throw new RuntimeException('unknown operation: ' . $operation);
    }
    $observations[] = bundle_observation($case['name'], $value);
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
