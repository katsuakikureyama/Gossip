<?php

// Both retained observers load the same original definitions with require_once.
ob_start();
require '/input/yn-observer.php';
$legacy = json_decode(ob_get_clean(), true, 512, JSON_THROW_ON_ERROR);
ob_start();
require '/input/numeric-observer.php';
$numeric = json_decode(ob_get_clean(), true, 512, JSON_THROW_ON_ERROR);

$cases = json_decode(file_get_contents('/input/composition-cases.json'), true, 512, JSON_THROW_ON_ERROR);
$composition = [];
$seen = [];
foreach ($cases as $case) {
    if (!is_string($case['name']) || isset($seen[$case['name']])) {
        throw new RuntimeException('invalid composition name');
    }
    $seen[$case['name']] = true;
    $input = numeric_array_fixture_pairs($case['input'], 'composition input');
    foreach ($input as $value) {
        if (!is_string($value)) { throw new RuntimeException('composition values must be strings'); }
    }
    $keys = yn_bundle_fixture_selectors($case['keys'], $case['operation']);
    $value = match ($case['operation']) {
        'gossip_numeric_slice_shape' => gossip_numeric_slice_shape($input, $keys),
        'gossip_numeric_values' => gossip_numeric_values($input, $keys),
        'gossip_numeric_yn_all' => gossip_numeric_yn_all($input, $keys),
        default => throw new RuntimeException('unknown numeric composition'),
    };
    if (!is_array($value) || !array_is_list($value)) {
        throw new RuntimeException('composition must return an actual list');
    }
    foreach ($value as $flag) {
        if (!is_bool($flag)) { throw new RuntimeException('composition results must be actual bools'); }
    }
    $after = numeric_array_fixture_entries($input);
    if ($value !== $case['expected'] || $after !== $case['input_after'] || $keys !== $case['keys_after']) {
        throw new RuntimeException('composition expectation mismatch: ' . $case['name']);
    }
    $composition[] = ['name' => $case['name'], 'type' => 'array', 'value' => $value, 'args_after' => [$after, $keys]];
}
echo json_encode(['legacy' => $legacy, 'numeric' => $numeric, 'composition' => $composition], JSON_THROW_ON_ERROR), "\n";
