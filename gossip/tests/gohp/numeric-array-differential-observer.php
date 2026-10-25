<?php

// Reuse the exact previous reference observer and its complete guarded source.
// Its assertions include PHP 8.3, C locale and user-defined array_all provenance.
ob_start();
require '/input/reference.php';
$predicate = json_decode(ob_get_clean(), true, 512, JSON_THROW_ON_ERROR);

$cases = json_decode(file_get_contents('/input/polyfill-cases.json'), true, 512, JSON_THROW_ON_ERROR);
$polyfill = [];
$seen = [];
foreach ($cases as $case) {
    if (!is_string($case['name']) || isset($seen[$case['name']]) || !is_bool($case['expected'])) {
        throw new RuntimeException('invalid polyfill case');
    }
    $seen[$case['name']] = true;
    $input = numeric_array_fixture_input($case['input'], 'array', 'polyfill input');
    $after = numeric_array_fixture_input($case['input_after'], 'array', 'polyfill input_after');
    $value = match ($case['operation']) {
        'gossip_array_all_key_value' => gossip_array_all_key_value($input),
        'gossip_array_all_short_circuit' => gossip_array_all_short_circuit($input),
        'gossip_array_all_empty' => gossip_array_all_empty($input),
        'gossip_array_all_mutation' => gossip_array_all_mutation($input),
        default => throw new RuntimeException('unknown polyfill operation'),
    };
    if (!is_bool($value) || $value !== $case['expected'] || $input !== $after) {
        throw new RuntimeException('polyfill expectation mismatch: ' . $case['name']);
    }
    $polyfill[] = [
        'name' => $case['name'], 'input_type' => 'array', 'type' => 'bool', 'value' => $value,
        'is_list' => array_is_list($input),
        'input_after' => ['type' => 'array', 'value' => numeric_array_fixture_entries($input)],
    ];
}
echo json_encode(['predicate' => $predicate, 'polyfill' => $polyfill], JSON_THROW_ON_ERROR), "\n";
