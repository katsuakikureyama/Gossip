<?php

require '/input/helpers.php';

$cases = json_decode(
    file_get_contents('/input/cases.json'),
    true,
    512,
    JSON_THROW_ON_ERROR
);

function array_slice_assoc_fixture_key(mixed $value): int|string {
    if (!is_int($value) && !is_string($value)) {
        throw new RuntimeException('array-slice-assoc fixture keys must be int or string');
    }
    if (is_int($value) && ($value < -1_000_000 || $value > 1_000_000)) {
        throw new RuntimeException('array-slice-assoc fixture keys are bounded');
    }
    return $value;
}

function array_slice_assoc_fixture_type(mixed $value): string {
    return match (get_debug_type($value)) {
        'null', 'bool', 'int', 'string', 'array' => get_debug_type($value),
        default => throw new RuntimeException('fixture transports only null, bool, int, string, and sequential list values'),
    };
}

function validate_array_slice_assoc_fixture_value(mixed $value, int $depth = 0): void {
    if ($depth > 8) {
        throw new RuntimeException('array-slice-assoc fixture values are limited to depth eight');
    }
    $type = array_slice_assoc_fixture_type($value);
    if ($type === 'int' && ($value < -1_000_000 || $value > 1_000_000)) {
        throw new RuntimeException('array-slice-assoc fixture values are bounded modest integers');
    }
    if ($type !== 'array') {
        return;
    }
    if (!array_is_list($value)) {
        throw new RuntimeException('array-slice-assoc fixture values do not transport maps or non-list nested keys');
    }
    foreach ($value as $item) {
        validate_array_slice_assoc_fixture_value($item, $depth + 1);
    }
}

function array_slice_assoc_fixture_pairs(array $pairs, string $field): array {
    $array = [];
    foreach ($pairs as $index => $pair) {
        if (!is_array($pair) || !array_is_list($pair) || count($pair) !== 2) {
            throw new RuntimeException('array-slice-assoc fixture ' . $field . ' pair ' . $index . ' must contain exactly key and value');
        }
        $key = array_slice_assoc_fixture_key($pair[0]);
        validate_array_slice_assoc_fixture_value($pair[1]);
        $array[$key] = $pair[1];
    }
    return $array;
}

function array_slice_assoc_fixture_selectors(array $selectors, string $field): array {
    foreach ($selectors as $selector) {
        array_slice_assoc_fixture_key($selector);
    }
    return $selectors;
}

function array_slice_assoc_fixture_entries(array $value): array {
    $entries = [];
    foreach ($value as $key => $item) {
        array_slice_assoc_fixture_key($key);
        validate_array_slice_assoc_fixture_value($item);
        $entries[] = [$key, $item];
    }
    return $entries;
}

$observations = [];
foreach ($cases as $case) {
    if (!is_array($case) || !is_string($case['name'] ?? null) || !is_array($case['input'] ?? null) || !is_array($case['keys'] ?? null) ||
        !is_array($case['expected'] ?? null) || !is_array($case['input_after'] ?? null) || !is_array($case['keys_after'] ?? null)) {
        throw new RuntimeException('malformed array-slice-assoc fixture case');
    }
    $input = array_slice_assoc_fixture_pairs($case['input'], 'input');
    $keys = array_slice_assoc_fixture_selectors($case['keys'], 'keys');
    // Validate explicit host expectations without using them to form a result.
    array_slice_assoc_fixture_pairs($case['expected'], 'expected');
    array_slice_assoc_fixture_pairs($case['input_after'], 'input_after');
    array_slice_assoc_fixture_selectors($case['keys_after'], 'keys_after');
    $value = wp_array_slice_assoc($input, $keys);
    $observations[] = [
        'name' => $case['name'],
        'type' => array_slice_assoc_fixture_type($value),
        'entries' => array_slice_assoc_fixture_entries($value),
        'input_after' => [
            'type' => array_slice_assoc_fixture_type($input),
            'entries' => array_slice_assoc_fixture_entries($input),
        ],
        'keys_after' => [
            'type' => array_slice_assoc_fixture_type($keys),
            'value' => $keys,
        ],
    ];
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
