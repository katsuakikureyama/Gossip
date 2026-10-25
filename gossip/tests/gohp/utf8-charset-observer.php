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

function utf8_charset_fixture_type(mixed $value): string {
    return match (get_debug_type($value)) {
        'null', 'bool', 'int', 'string', 'array' => get_debug_type($value),
        default => throw new RuntimeException('fixture transports only null, bool, int, string, and sequential list values'),
    };
}

function validate_utf8_charset_fixture_input(mixed $value, int $depth = 0): void {
    if ($depth > 8) {
        throw new RuntimeException('UTF-8 charset fixture inputs are limited to depth eight');
    }
    $type = utf8_charset_fixture_type($value);
    if ($type === 'int' && ($value < -1_000_000 || $value > 1_000_000)) {
        throw new RuntimeException('UTF-8 charset fixture inputs are bounded modest integers');
    }
    if ($type !== 'array') {
        return;
    }
    if (!array_is_list($value)) {
        throw new RuntimeException('UTF-8 charset fixture does not transport maps or non-list nested keys');
    }
    foreach ($value as $item) {
        validate_utf8_charset_fixture_input($item, $depth + 1);
    }
}

$observations = [];
foreach ($cases as $case) {
    if (!is_array($case) || !is_string($case['name'] ?? null) || !is_string($case['input_type'] ?? null) ||
        !array_key_exists('input', $case) || !array_key_exists('expected', $case) || !is_bool($case['expected'])) {
        throw new RuntimeException('malformed UTF-8 charset fixture case');
    }
    $input = $case['input'];
    validate_utf8_charset_fixture_input($input);
    if (utf8_charset_fixture_type($input) !== $case['input_type']) {
        throw new RuntimeException('case input type did not survive JSON transport: ' . $case['name']);
    }
    $value = _is_utf8_charset($input);
    if (get_debug_type($value) !== 'bool') {
        throw new RuntimeException('_is_utf8_charset did not return bool');
    }
    $observations[] = [
        'name' => $case['name'],
        'input_type' => $case['input_type'],
        'type' => 'bool',
        'value' => $value,
        'input_after' => [
            'type' => utf8_charset_fixture_type($input),
            'value' => $input,
        ],
    ];
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
