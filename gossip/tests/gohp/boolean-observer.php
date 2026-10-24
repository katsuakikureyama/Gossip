<?php

require '/input/helpers.php';

if (setlocale(LC_CTYPE, '0') !== 'C') {
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
    $input = $case['input'];
    if (get_debug_type($input) !== $case['input_type']) {
        throw new RuntimeException('case input type did not survive JSON transport: ' . $case['name']);
    }
    $value = wp_validate_boolean($input);
    if (get_debug_type($value) !== 'bool') {
        throw new RuntimeException('wp_validate_boolean did not return bool');
    }
    $observations[] = [
        'name' => $case['name'],
        'input_type' => $case['input_type'],
        'type' => 'bool',
        'value' => $value,
    ];
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
