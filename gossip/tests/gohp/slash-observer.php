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
    $untrailing = untrailingslashit($case['input']);
    $trailing = trailingslashit($case['input']);
    $observations[] = [
        'name' => $case['name'],
        'input' => $case['input'],
        'untrailing' => [
            'type' => get_debug_type($untrailing),
            'value' => $untrailing,
        ],
        'trailing' => [
            'type' => get_debug_type($trailing),
            'value' => $trailing,
        ],
    ];
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
