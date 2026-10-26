<?php

// Owned observation driver, not a tokenizer implementation or replacement class.
// The runner checks PHP admission in this process before requiring this file.
require_once '/input/wp-includes/class-wp-block-parser.php';
if (setlocale(LC_CTYPE, 'C') !== 'C') {
    throw new RuntimeException('requires C LC_CTYPE');
}
foreach (['pcre.jit' => '0', 'pcre.recursion_limit' => '10000', 'memory_limit' => '128M'] as $key => $value) {
    if (false === ini_set($key, $value) || ini_get($key) !== $value) {
        throw new RuntimeException('cannot fix oracle setting: ' . $key);
    }
}
$method = new ReflectionMethod('WP_Block_Parser', 'next_token');
if (!$method->isPublic() || $method->isInternal() || $method->getDeclaringClass()->getName() !== 'WP_Block_Parser' || $method->getNumberOfParameters() !== 0 || $method->getStartLine() !== 239 || $method->getEndLine() !== 303 || $method->getFileName() !== '/input/wp-includes/class-wp-block-parser.php' || count((new ReflectionClass('WP_Block_Parser'))->getMethods()) !== 7) {
    throw new RuntimeException('requires complete original parser and exact next_token method');
}
foreach (['WP_Block_Parser_Block', 'WP_Block_Parser_Frame'] as $class) {
    if (!(new ReflectionClass($class))->isUserDefined()) {
        throw new RuntimeException('requires original parser dependencies');
    }
}

// Preserve types, key order and raw bytes; never stringify a missing/null value.
function token_reference_value(mixed $value, int $depth = 0): array {
    if ($depth > 32) {
        throw new RuntimeException('observation nesting exceeds fixture bound');
    }
    if (is_string($value)) {
        return ['type' => 'string', 'base64' => base64_encode($value)];
    }
    if (is_null($value) || is_bool($value) || is_int($value)) {
        return ['type' => get_debug_type($value), 'value' => $value];
    }
    if (is_array($value)) {
        $entries = [];
        foreach ($value as $key => $child) {
            $entries[] = ['key' => token_reference_value($key, $depth + 1), 'value' => token_reference_value($child, $depth + 1)];
        }
        return ['type' => 'array', 'entries' => $entries];
    }
    throw new RuntimeException('value outside reference observation domain');
}

$pattern = file_get_contents('/input/pattern.txt');
$cases = json_decode(file_get_contents('/input/cases.json'), true, 32, JSON_THROW_ON_ERROR);
if (!is_array($cases) || !array_is_list($cases) || count($cases) !== 24) {
    throw new RuntimeException('requires 24 fixed cases');
}
$observations = [];
$seen = [];
foreach ($cases as $case) {
    if (array_keys($case) !== ['name', 'document_base64', 'offset', 'backtrack_limit', 'expected_match', 'expected_pcre_error', 'expected_token']) {
        throw new RuntimeException('unexpected case fields');
    }
    $name = $case['name'];
    $document = base64_decode($case['document_base64'], true);
    $offset = $case['offset'];
    if (!is_string($name) || $name === '' || isset($seen[$name]) || $document === false || base64_encode($document) !== $case['document_base64'] || strlen($document) > 4096 || !is_int($offset) || $offset < -strlen($document) || $offset > strlen($document) || !in_array($case['backtrack_limit'], [1, 100000], true)) {
        throw new RuntimeException('case outside fixed domain');
    }
    $seen[$name] = true;
    ini_set('pcre.backtrack_limit', (string) $case['backtrack_limit']);
    if (ini_get('pcre.backtrack_limit') !== (string) $case['backtrack_limit']) {
        throw new RuntimeException('backtrack profile not applied');
    }

    $document_before = $document;
    $offset_before = $offset;
    $matches = null; // Same initial value as the original next_token.
    $match = preg_match($pattern, $document, $matches, PREG_OFFSET_CAPTURE, $offset);
    $match_error = preg_last_error();
    $match_error_message = preg_last_error_msg();
    if ($match !== $case['expected_match'] || $match_error !== $case['expected_pcre_error']) {
        throw new RuntimeException($name . ': direct preg_match differs from predeclared expectation: ' . json_encode([$match, $match_error]));
    }
    $match_observation = token_reference_value($matches);
    $capture_presence = [];
    foreach (['closer', 'namespace', 'name', 'attrs', 'void'] as $capture) {
        $capture_presence[$capture] = is_array($matches) && array_key_exists($capture, $matches);
    }

    $parser = new WP_Block_Parser();
    $parser->document = $document;
    $parser->offset = $offset;
    $parser->output = ['output-sentinel'];
    $parser->stack = ['stack-sentinel'];
    $state_before = get_object_vars($parser);
    json_decode('null', true); // Defined clean JSON-error baseline, not a parser substitute.
    $token = $parser->next_token();
    $parser_pcre_error = preg_last_error();
    $parser_pcre_message = preg_last_error_msg();
    $parser_json_error = json_last_error();
    if ($token !== $case['expected_token'] || $parser_pcre_error !== $case['expected_pcre_error']) {
        throw new RuntimeException($name . ': original next_token differs from predeclared expectation: ' . json_encode(token_reference_value($token)));
    }
    if (get_object_vars($parser) !== $state_before || $document !== $document_before || $offset !== $offset_before) {
        throw new RuntimeException($name . ': unexpected parser/caller mutation');
    }
    $observations[] = [
        'name' => $name,
        'input' => ['document_base64' => base64_encode($document), 'offset' => $offset],
        'profile' => ['pcre.jit' => ini_get('pcre.jit'), 'pcre.backtrack_limit' => ini_get('pcre.backtrack_limit'), 'pcre.recursion_limit' => ini_get('pcre.recursion_limit')],
        'preg_match' => ['return' => token_reference_value($match), 'error_code' => $match_error, 'error_message' => $match_error_message, 'matches_before' => token_reference_value(null), 'matches_after' => $match_observation, 'named_capture_presence' => $capture_presence],
        'next_token' => ['value' => token_reference_value($token), 'pcre_error_code' => $parser_pcre_error, 'pcre_error_message' => $parser_pcre_message, 'json_error_code' => $parser_json_error, 'parser_after' => token_reference_value(get_object_vars($parser)), 'state_unchanged' => true],
        'caller_unchanged' => true,
    ];
}
echo json_encode([
    'schema' => 'gossip.parser-token-native-observation/v1',
    'status' => 'native-reference-pass',
    'php_version' => PHP_VERSION, 'php_version_id' => PHP_VERSION_ID,
    'pcre_version' => PCRE_VERSION, 'pcre_jit_support' => PCRE_JIT_SUPPORT,
    'lc_ctype' => setlocale(LC_CTYPE, 0), 'int_size' => PHP_INT_SIZE,
    'memory_limit' => ini_get('memory_limit'), 'preg_offset_capture' => PREG_OFFSET_CAPTURE,
    'pattern_base64' => base64_encode($pattern),
    'pattern_sha256' => hash('sha256', $pattern),
    'method' => ['class' => 'WP_Block_Parser', 'name' => 'next_token', 'first_line' => $method->getStartLine(), 'last_line' => $method->getEndLine()],
    'cases' => $observations,
], JSON_THROW_ON_ERROR), "\n";
