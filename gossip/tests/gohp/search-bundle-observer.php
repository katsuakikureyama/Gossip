<?php

ob_start();
require '/input/retained-observer.php';
$retainedBundle = json_decode(ob_get_clean(),true,512,JSON_THROW_ON_ERROR);
ob_start();
require '/input/search-observer.php';
$searchUnit = json_decode(ob_get_clean(),true,512,JSON_THROW_ON_ERROR);
$cases = json_decode(file_get_contents('/input/search-composition-cases.json'),true,512,JSON_THROW_ON_ERROR);
$composed = [];
$seen = [];
foreach ($cases as $case) {
    if (!is_string($case['name']) || isset($seen[$case['name']])) { throw new RuntimeException('invalid composition name'); }
    $seen[$case['name']] = true;
    $input = numeric_array_fixture_pairs($case['input'],'search composition');
    foreach ($input as $item) {
        if (!is_string($item) && $item !== null) { throw new RuntimeException('expected string or null input'); }
    }
    $keys = yn_bundle_fixture_selectors($case['keys'],$case['operation']);
    $value = match ($case['operation']) {
        'gossip_slice_search_yn' => gossip_slice_search_yn($input,$keys),
        'gossip_slice_search_falsy' => gossip_slice_search_falsy($input,$keys),
        default => throw new RuntimeException('unknown search composition'),
    };
    if (!is_array($value) || !array_is_list($value)) { throw new RuntimeException('expected actual list'); }
    foreach ($value as $item) { numeric_array_fixture_scalar($item,'observed result element'); }
    $after = numeric_array_fixture_entries($input);
    if ($value !== $case['expected'] || $after !== $case['input_after'] || $keys !== $case['keys_after']) {
        throw new RuntimeException('composition expectation mismatch: ' . $case['name']);
    }
    $composed[] = ['name'=>$case['name'],'type'=>'array','value'=>$value,'args_after'=>[$after,$keys]];
}
echo json_encode(['retained'=>$retainedBundle,'search'=>$searchUnit,'composition'=>$composed],JSON_THROW_ON_ERROR), "\n";
