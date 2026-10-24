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
    if (!is_string($case['name']) || !is_array($case['input']) || !array_is_list($case['input']) || count($case['input']) === 0) {
        throw new RuntimeException('case must contain a non-empty list input: ' . ($case['name'] ?? '<unknown>'));
    }
    foreach ($case['input'] as $entry) {
        if (!is_string($entry)) {
            throw new RuntimeException('case input must contain only strings: ' . $case['name']);
        }
    }
    $value = _autop_newline_preservation_helper($case['input']);
    if (get_debug_type($value) !== 'string') {
        throw new RuntimeException('_autop_newline_preservation_helper did not return string');
    }
    $observations[] = [
        'name' => $case['name'],
        'type' => 'string',
        'value' => $value,
    ];
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
