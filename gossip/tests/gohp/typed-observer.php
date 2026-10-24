<?php

require '/input/helpers.php';

$names = [
    'true' => '__return_true',
    'false' => '__return_false',
    'zero' => '__return_zero',
    'null' => '__return_null',
    'empty_string' => '__return_empty_string',
];

$observations = [];
foreach ($names as $name => $function) {
    $value = $function();
    $observations[] = [
        'name' => $name,
        'type' => get_debug_type($value),
        'value' => $value,
    ];
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
