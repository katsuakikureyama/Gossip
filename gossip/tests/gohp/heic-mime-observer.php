<?php

require '/input/helpers.php';

$cases = json_decode(
    file_get_contents('/input/cases.json'),
    true,
    512,
    JSON_THROW_ON_ERROR
);

function heic_mime_fixture_type(mixed $value): string {
    return match (get_debug_type($value)) {
        'null', 'bool', 'int', 'string' => get_debug_type($value),
        default => throw new RuntimeException('HEIC MIME fixture transports scalar labels only, not arrays or objects'),
    };
}

function validate_heic_mime_fixture_input(mixed $value): void {
    $type = heic_mime_fixture_type($value);
    if ($type === 'int' && ($value < -1_000_000 || $value > 1_000_000)) {
        throw new RuntimeException('HEIC MIME fixture inputs are bounded modest integers');
    }
}

$observations = [];
foreach ($cases as $case) {
    if (!is_array($case) || !is_string($case['name'] ?? null) || !is_string($case['input_type'] ?? null) ||
        !array_key_exists('input', $case) || !array_key_exists('expected', $case) || !is_bool($case['expected'])) {
        throw new RuntimeException('malformed HEIC MIME fixture case');
    }
    $input = $case['input'];
    validate_heic_mime_fixture_input($input);
    if (heic_mime_fixture_type($input) !== $case['input_type']) {
        throw new RuntimeException('case input type did not survive JSON transport: ' . $case['name']);
    }
    $value = wp_is_heic_image_mime_type($input);
    if (get_debug_type($value) !== 'bool') {
        throw new RuntimeException('wp_is_heic_image_mime_type did not return bool');
    }
    $observations[] = [
        'name' => $case['name'],
        'input_type' => $case['input_type'],
        'type' => 'bool',
        'value' => $value,
        'input_after' => [
            'type' => heic_mime_fixture_type($input),
            'value' => $input,
        ],
    ];
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
