<?php

require '/input/helpers.php';

$cases = json_decode(
    file_get_contents('/input/cases.json'),
    true,
    512,
    JSON_THROW_ON_ERROR
);

function stripslashes_fixture_type(mixed $value): string {
    return match (get_debug_type($value)) {
        'null', 'bool', 'int', 'string', 'array' => get_debug_type($value),
        default => throw new RuntimeException('fixture transports only null, bool, int, string, and list array values'),
    };
}

function validate_stripslashes_fixture_value(mixed $value): void {
    if (stripslashes_fixture_type($value) !== 'array') {
        return;
    }
    if (!array_is_list($value)) {
        throw new RuntimeException('fixture does not transport PHP maps or non-list array keys');
    }
    foreach ($value as $entry) {
        validate_stripslashes_fixture_value($entry);
    }
}

$observations = [];
foreach ($cases as $case) {
    if (!is_array($case) || !is_string($case['name'] ?? null) || !is_string($case['input_type'] ?? null) || !is_string($case['expected_type'] ?? null) ||
        !array_key_exists('input', $case) || !array_key_exists('expected', $case)) {
        throw new RuntimeException('malformed stripslashes fixture case');
    }
    $input = $case['input'];
    validate_stripslashes_fixture_value($input);
    if (stripslashes_fixture_type($input) !== $case['input_type']) {
        throw new RuntimeException('case input type did not survive JSON transport: ' . $case['name']);
    }
    validate_stripslashes_fixture_value($case['expected']);
    if (stripslashes_fixture_type($case['expected']) !== $case['expected_type']) {
        throw new RuntimeException('case expected type is outside the fixture contract: ' . $case['name']);
    }
    $value = stripslashes_from_strings_only($input);
    validate_stripslashes_fixture_value($value);
    $observations[] = [
        'name' => $case['name'],
        'input_type' => $case['input_type'],
        'type' => stripslashes_fixture_type($value),
        'value' => $value,
    ];
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
