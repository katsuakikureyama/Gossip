<?php

require '/input/helpers.php';

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
    if (get_debug_type($case['threshold']) !== 'int' || $case['threshold'] < 0 || $case['threshold'] > 4096) {
        throw new RuntimeException('case threshold must be an int in 0..4096: ' . $case['name']);
    }
    $value = zeroise($input, $case['threshold']);
    if (get_debug_type($value) !== 'string') {
        throw new RuntimeException('zeroise did not return string');
    }
    $observations[] = [
        'name' => $case['name'],
        'input_type' => $case['input_type'],
        'type' => 'string',
        'value' => $value,
    ];
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
