<?php

require '/input/helpers.php';

if (setlocale(LC_CTYPE, 'C') !== 'C' || setlocale(LC_CTYPE, 0) !== 'C') {
    throw new RuntimeException('LC_CTYPE must be C for this bounded observation');
}

if (PHP_VERSION_ID >= 80400) {
    throw new RuntimeException('this observation requires the WordPress array_all polyfill, not PHP 8.4 array_all');
}

$array_all = new ReflectionFunction('array_all');
if (! $array_all->isUserDefined() || $array_all->isInternal()) {
    throw new RuntimeException('array_all must be the user-defined WordPress compat polyfill');
}

$cases = json_decode(
    file_get_contents('/input/cases.json'),
    true,
    512,
    JSON_THROW_ON_ERROR
);

function numeric_array_fixture_type(mixed $value): string {
    return match (get_debug_type($value)) {
        'null', 'bool', 'int', 'string', 'array' => get_debug_type($value),
        default => throw new RuntimeException('numeric-array fixture transport is limited to null, bool, int, string, and arrays'),
    };
}

function numeric_array_fixture_scalar(mixed $value, string $field): mixed {
    $type = numeric_array_fixture_type($value);
    if ($type === 'array') {
        throw new RuntimeException('numeric-array fixture ' . $field . ' must be scalar');
    }
    if ($type === 'int' && ($value < -1_000_000 || $value > 1_000_000)) {
        throw new RuntimeException('numeric-array fixture ' . $field . ' integer is outside the bounded range');
    }
    return $value;
}

function numeric_array_fixture_key(mixed $key, string $field): int|string {
    if (! is_int($key) && ! is_string($key)) {
        throw new RuntimeException('numeric-array fixture ' . $field . ' keys must be int or string');
    }
    if (is_int($key) && ($key < -1_000_000 || $key > 1_000_000)) {
        throw new RuntimeException('numeric-array fixture ' . $field . ' integer keys are outside the bounded range');
    }
    return $key;
}

function numeric_array_fixture_normalized_key(int|string $key, string $field): int|string {
    $probe = [];
    $probe[$key] = true;
    $normalized = array_key_first($probe);
    if (is_int($normalized) && ($normalized < -1_000_000 || $normalized > 1_000_000)) {
        throw new RuntimeException('numeric-array fixture ' . $field . ' normalized integer keys are outside the bounded range');
    }
    return $normalized;
}

function numeric_array_fixture_pairs(array $pairs, string $field): array {
    if (! array_is_list($pairs)) {
        throw new RuntimeException('numeric-array fixture ' . $field . ' must be an ordered pair list');
    }

    $array = [];
    foreach ($pairs as $index => $pair) {
        if (! is_array($pair) || ! array_is_list($pair) || count($pair) !== 2) {
            throw new RuntimeException('numeric-array fixture ' . $field . ' pair ' . $index . ' must contain exactly key and value');
        }
        $key = numeric_array_fixture_key($pair[0], $field . ' pair ' . $index);
        $normalized_key = numeric_array_fixture_normalized_key($key, $field . ' pair ' . $index);
        $value = numeric_array_fixture_scalar($pair[1], $field . ' pair ' . $index . ' value');
        if (array_key_exists($normalized_key, $array)) {
            throw new RuntimeException('numeric-array fixture ' . $field . ' has duplicate normalized key at pair ' . $index);
        }
        $array[$normalized_key] = $value;
    }

    return $array;
}

function numeric_array_fixture_input(mixed $value, string $input_type, string $field): mixed {
    if ($input_type === 'array') {
        if (! is_array($value)) {
            throw new RuntimeException('numeric-array fixture ' . $field . ' must be an ordered pair list');
        }
        return numeric_array_fixture_pairs($value, $field);
    }

    $scalar = numeric_array_fixture_scalar($value, $field);
    if (numeric_array_fixture_type($scalar) !== $input_type) {
        throw new RuntimeException('numeric-array fixture ' . $field . ' does not match input_type ' . $input_type);
    }
    return $scalar;
}

function numeric_array_fixture_entries(array $value): array {
    $entries = [];
    foreach ($value as $key => $item) {
        numeric_array_fixture_key($key, 'observed array');
        numeric_array_fixture_scalar($item, 'observed array value');
        $entries[] = [$key, $item];
    }
    return $entries;
}

if (! is_array($cases) || ! array_is_list($cases)) {
    throw new RuntimeException('numeric-array cases must be an ordered list');
}

$seen_names = [];
$observations = [];
foreach ($cases as $case) {
    if (! is_array($case) || ! is_string($case['name'] ?? null) || ! array_key_exists('input_type', $case) ||
        ! is_string($case['input_type']) || ! array_key_exists('input', $case) || ! is_bool($case['expected'] ?? null) ||
        ! array_key_exists('expected_is_list', $case) || ! array_key_exists('input_after', $case)) {
        throw new RuntimeException('malformed numeric-array fixture case');
    }
    if (array_key_exists($case['name'], $seen_names)) {
        throw new RuntimeException('numeric-array fixture case names must be unique');
    }
    $seen_names[$case['name']] = true;

    $input_type = $case['input_type'];
    if (! in_array($input_type, ['null', 'bool', 'int', 'string', 'array'], true)) {
        throw new RuntimeException('numeric-array fixture has unsupported input_type');
    }
    if (($input_type === 'array') !== is_bool($case['expected_is_list'])) {
        throw new RuntimeException('numeric-array fixture expected_is_list must be bool for arrays and null otherwise');
    }
    if ($input_type !== 'array' && $case['expected_is_list'] !== null) {
        throw new RuntimeException('numeric-array fixture expected_is_list must be null for non-arrays');
    }

    $input = numeric_array_fixture_input($case['input'], $input_type, 'input');
    $expected_input_after = numeric_array_fixture_input($case['input_after'], $input_type, 'input_after');
    $value = wp_is_numeric_array($input);
    if (get_debug_type($value) !== 'bool') {
        throw new RuntimeException('wp_is_numeric_array did not return bool');
    }
    $is_list = is_array($input) ? array_is_list($input) : null;

    // Explicit fixture expectations assert the PHP host contract; observations remain computed values.
    if ($value !== $case['expected'] || $is_list !== $case['expected_is_list'] || $input !== $expected_input_after) {
        throw new RuntimeException('numeric-array PHP observation differs from its explicit fixture expectation: ' . $case['name']);
    }

    $observations[] = [
        'name' => $case['name'],
        'input_type' => $input_type,
        'type' => 'bool',
        'value' => $value,
        'is_list' => $is_list,
        'input_after' => [
            'type' => numeric_array_fixture_type($input),
            'value' => is_array($input) ? numeric_array_fixture_entries($input) : $input,
        ],
    ];
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
