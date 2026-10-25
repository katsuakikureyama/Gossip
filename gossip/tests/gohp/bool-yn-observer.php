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

$observations = [];
foreach ($cases as $case) {
    if (!is_array($case) || !is_string($case['name'] ?? null) || ($case['input_type'] ?? null) !== 'string' ||
        !is_string($case['input'] ?? null) || !array_key_exists('expected', $case) || !is_bool($case['expected'])) {
        throw new RuntimeException('malformed bool_from_yn fixture case');
    }
    $input = $case['input'];
    $value = bool_from_yn($input);
    if (get_debug_type($value) !== 'bool') {
        throw new RuntimeException('bool_from_yn did not return bool');
    }
    $observations[] = [
        'name' => $case['name'],
        'input_type' => 'string',
        'type' => 'bool',
        'value' => $value,
        'input_after' => [
            'type' => 'string',
            'value' => $input,
        ],
    ];
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
