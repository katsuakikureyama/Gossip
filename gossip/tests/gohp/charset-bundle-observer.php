<?php

require '/input/helpers.php';

if (setlocale(LC_CTYPE, 'C') !== 'C' || setlocale(LC_CTYPE, 0) !== 'C') {
    throw new RuntimeException('LC_CTYPE must be C for this bounded observation');
}

$cases = json_decode(file_get_contents('/input/cases.json'), true, 512, JSON_THROW_ON_ERROR);

function charset_bundle_fixture_key(mixed $value): int|string {
    if (!is_int($value) && !is_string($value)) {
        throw new RuntimeException('charset bundle fixture keys must be int or string');
    }
    if (is_int($value) && ($value < -1_000_000 || $value > 1_000_000)) {
        throw new RuntimeException('charset bundle fixture keys are bounded');
    }
    return $value;
}

function charset_bundle_fixture_type(mixed $value): string {
    return match (get_debug_type($value)) {
        'null', 'bool', 'int', 'string', 'array' => get_debug_type($value),
        default => throw new RuntimeException('fixture transports only null, bool, int, string, and list array values'),
    };
}

function charset_bundle_validate_fixture_value(mixed $value, int $depth = 0): void {
    if ($depth > 8) {
        throw new RuntimeException('charset bundle fixture values are limited to depth eight');
    }
    $type = charset_bundle_fixture_type($value);
    if ($type === 'int' && ($value < -1_000_000 || $value > 1_000_000)) {
        throw new RuntimeException('charset bundle fixture values are bounded modest integers');
    }
    if ($type !== 'array') {
        return;
    }
    if (!array_is_list($value)) {
        throw new RuntimeException('fixture does not transport PHP maps or non-list array keys');
    }
    foreach ($value as $entry) {
        charset_bundle_validate_fixture_value($entry, $depth + 1);
    }
}

function charset_bundle_fixture_pairs(array $pairs, string $field): array {
    $array = [];
    foreach ($pairs as $index => $pair) {
        if (!is_array($pair) || !array_is_list($pair) || count($pair) !== 2) {
            throw new RuntimeException('charset bundle fixture ' . $field . ' pair ' . $index . ' must contain exactly key and value');
        }
        $key = charset_bundle_fixture_key($pair[0]);
        charset_bundle_validate_fixture_value($pair[1]);
        $array[$key] = $pair[1];
    }
    return $array;
}

function charset_bundle_fixture_selectors(array $selectors, string $operation): array {
    foreach ($selectors as $selector) {
        charset_bundle_fixture_key($selector);
    }
    return $selectors;
}

function charset_bundle_fixture_entries(array $value): array {
    $entries = [];
    foreach ($value as $key => $item) {
        charset_bundle_fixture_key($key);
        charset_bundle_validate_fixture_value($item);
        $entries[] = [$key, $item];
    }
    return $entries;
}

function charset_bundle_pair_operation(string $operation): bool {
    return in_array($operation, [
        'wp_array_slice_assoc',
        'gossip_slice_slash',
        'gossip_slash_slice',
        'gossip_slice_shallow_roundtrip',
        'gossip_slice_shape',
        'gossip_slice_boolean',
        'gossip_slice_charset',
        'gossip_slice_unescape_charset',
    ], true);
}

function charset_bundle_pairs_result(string $operation): bool {
    return in_array($operation, [
        'wp_array_slice_assoc',
        'gossip_slice_slash',
        'gossip_slash_slice',
        'gossip_slice_shallow_roundtrip',
        'gossip_slice_charset',
        'gossip_slice_unescape_charset',
    ], true);
}

function charset_bundle_no_args(array $args, string $operation): void {
    if ($args !== []) {
        throw new RuntimeException($operation . ' requires no arguments');
    }
}

function charset_bundle_string_arg(array $args, string $operation): string {
    if (count($args) !== 1 || get_debug_type($args[0]) !== 'string') {
        throw new RuntimeException($operation . ' requires exactly one string argument');
    }
    return $args[0];
}

function charset_bundle_any_arg(array $args, string $operation): mixed {
    if (count($args) !== 1) {
        throw new RuntimeException($operation . ' requires exactly one fixture value');
    }
    charset_bundle_validate_fixture_value($args[0]);
    return $args[0];
}

function charset_bundle_width(mixed $value, string $operation): int {
    if (get_debug_type($value) !== 'int' || $value < 0 || $value > 4096) {
        throw new RuntimeException($operation . ' width must be an int in 0..4096');
    }
    return $value;
}

function charset_bundle_value_and_width(array $args, string $operation): array {
    if (count($args) !== 2 || !in_array(get_debug_type($args[0]), ['string', 'int'], true)) {
        throw new RuntimeException($operation . ' requires a string or int value and int width');
    }
    return [$args[0], charset_bundle_width($args[1], $operation)];
}

function charset_bundle_captures(array $args, string $operation): array {
    if (count($args) !== 1 || !is_array($args[0]) || !array_is_list($args[0]) || count($args[0]) === 0) {
        throw new RuntimeException($operation . ' requires one non-empty list of string captures');
    }
    foreach ($args[0] as $capture) {
        if (get_debug_type($capture) !== 'string') {
            throw new RuntimeException($operation . ' captures must all be strings');
        }
    }
    return $args[0];
}

function charset_bundle_array_arg(array $args, string $operation): array {
    if (count($args) !== 1 || !is_array($args[0]) || !array_is_list($args[0])) {
        throw new RuntimeException($operation . ' requires exactly one list array');
    }
    return $args[0];
}

function charset_bundle_observation(string $name, string $operation, mixed $value, array $args): array {
    $observedValue = $value;
    if (charset_bundle_pairs_result($operation)) {
        if (!is_array($value)) {
            throw new RuntimeException($operation . ' must return an array');
        }
        $observedValue = charset_bundle_fixture_entries($value);
    } else {
        charset_bundle_validate_fixture_value($value);
    }
    $argsAfter = $args;
    if (charset_bundle_pair_operation($operation)) {
        $argsAfter[0] = charset_bundle_fixture_entries($args[0]);
        charset_bundle_validate_fixture_value($argsAfter[1]);
    } else {
        charset_bundle_validate_fixture_value($argsAfter);
    }
    return [
        'name' => $name,
        'type' => charset_bundle_fixture_type($value),
        'value' => $observedValue,
        'args_after' => $argsAfter,
    ];
}

$observations = [];
foreach ($cases as $case) {
    if (!is_array($case) || !is_string($case['name'] ?? null) || !is_string($case['operation'] ?? null) || !is_array($case['args'] ?? null) ||
        !is_array($case['expected'] ?? null) || !is_string($case['expected']['type'] ?? null) || !array_key_exists('value', $case['expected']) ||
        !is_array($case['args_after'] ?? null)) {
        throw new RuntimeException('malformed charset helper bundle case');
    }
    charset_bundle_validate_fixture_value($case['args']);
    charset_bundle_validate_fixture_value($case['expected']['value']);
    charset_bundle_validate_fixture_value($case['args_after']);
    if (charset_bundle_fixture_type($case['expected']['value']) !== $case['expected']['type']) {
        throw new RuntimeException('case expected type is outside the fixture contract: ' . $case['name']);
    }

    $operation = $case['operation'];
    $args = $case['args'];
    if (charset_bundle_pair_operation($operation)) {
        if (count($args) !== 2 || !is_array($args[0]) || !is_array($args[1])) {
            throw new RuntimeException($operation . ' requires exactly an ordered-pairs array and a selector list');
        }
        $args = [
            charset_bundle_fixture_pairs($args[0], $operation . ' input'),
            charset_bundle_fixture_selectors($args[1], $operation),
        ];
    }

    switch ($operation) {
        case '__return_true': charset_bundle_no_args($args, $operation); $value = __return_true(); break;
        case '__return_false': charset_bundle_no_args($args, $operation); $value = __return_false(); break;
        case '__return_zero': charset_bundle_no_args($args, $operation); $value = __return_zero(); break;
        case '__return_empty_array': charset_bundle_no_args($args, $operation); $value = __return_empty_array(); break;
        case '__return_null': charset_bundle_no_args($args, $operation); $value = __return_null(); break;
        case '__return_empty_string': charset_bundle_no_args($args, $operation); $value = __return_empty_string(); break;
        case 'wp_validate_boolean': $value = wp_validate_boolean(charset_bundle_string_arg($args, $operation)); break;
        case '_autop_newline_preservation_helper': $value = _autop_newline_preservation_helper(charset_bundle_captures($args, $operation)); break;
        case 'zeroise': [$input, $width] = charset_bundle_value_and_width($args, $operation); $value = zeroise($input, $width); break;
        case 'trailingslashit': $value = trailingslashit(charset_bundle_string_arg($args, $operation)); break;
        case 'untrailingslashit': $value = untrailingslashit(charset_bundle_string_arg($args, $operation)); break;
        case 'stripslashes_from_strings_only': $value = stripslashes_from_strings_only(charset_bundle_any_arg($args, $operation)); break;
        case 'wp_slash': $value = wp_slash(charset_bundle_any_arg($args, $operation)); break;
        case 'gossip_capture_newline_unslash': $value = gossip_capture_newline_unslash(charset_bundle_captures($args, $operation)); break;
        case 'gossip_unslash_array_keys': $value = gossip_unslash_array_keys(charset_bundle_array_arg($args, $operation)); break;
        case 'gossip_unslash_array_shape': $value = gossip_unslash_array_shape(charset_bundle_array_arg($args, $operation)); break;
        case 'gossip_path_pad': [$path, $width] = charset_bundle_value_and_width($args, $operation); if (get_debug_type($path) !== 'string') { throw new RuntimeException('gossip_path_pad requires a string path'); } $value = gossip_path_pad($path, $width); break;
        case 'gossip_zero_pad': if (count($args) !== 1) { throw new RuntimeException('gossip_zero_pad requires one width'); } $value = gossip_zero_pad(charset_bundle_width($args[0], $operation)); break;
        case 'gossip_false_boolean': charset_bundle_no_args($args, $operation); $value = gossip_false_boolean(); break;
        case 'gossip_empty_string_path': charset_bundle_no_args($args, $operation); $value = gossip_empty_string_path(); break;
        case 'gossip_slash_scalar_roundtrip': $value = gossip_slash_scalar_roundtrip(charset_bundle_string_arg($args, $operation)); break;
        case 'gossip_slash_list_shallow_roundtrip': $value = gossip_slash_list_shallow_roundtrip(charset_bundle_array_arg($args, $operation)); break;
        case 'gossip_slash_capture_newline_unslash': $value = gossip_slash_capture_newline_unslash(charset_bundle_captures($args, $operation)); break;
        case 'gossip_slash_array_shape': $value = gossip_slash_array_shape(charset_bundle_array_arg($args, $operation)); break;
        case 'wp_array_slice_assoc': $value = wp_array_slice_assoc($args[0], $args[1]); break;
        case 'gossip_slice_slash': $value = gossip_slice_slash($args[0], $args[1]); break;
        case 'gossip_slash_slice': $value = gossip_slash_slice($args[0], $args[1]); break;
        case 'gossip_slice_shallow_roundtrip': $value = gossip_slice_shallow_roundtrip($args[0], $args[1]); break;
        case 'gossip_slice_shape': $value = gossip_slice_shape($args[0], $args[1]); break;
        case 'gossip_slice_boolean': $value = gossip_slice_boolean($args[0], $args[1]); break;
        case '_is_utf8_charset': $value = _is_utf8_charset(charset_bundle_any_arg($args, $operation)); break;
        case 'gossip_slice_charset': $value = gossip_slice_charset($args[0], $args[1]); break;
        case 'gossip_slice_unescape_charset': $value = gossip_slice_unescape_charset($args[0], $args[1]); break;
        case 'gossip_charset_boolean': $value = gossip_charset_boolean(charset_bundle_string_arg($args, $operation)); break;
        default: throw new RuntimeException('unknown operation: ' . $operation);
    }
    $observations[] = charset_bundle_observation($case['name'], $operation, $value, $args);
}

echo json_encode($observations, JSON_THROW_ON_ERROR), "\n";
