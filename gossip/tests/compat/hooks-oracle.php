<?php
// Owned test adapter. Execute the verified ORIGINAL source from this fork.
// The immutable container supplies PHP only; its bundled WordPress is unused.
error_reporting(E_ALL);
set_error_handler(static function ($severity, $message, $file, $line) {
    throw new ErrorException($message, 0, $severity, $file, $line);
});
define('ABSPATH', '/wordpress/');
define('WPINC', 'wp-includes');
require ABSPATH . WPINC . '/version.php';
require ABSPATH . WPINC . '/plugin.php';
$fixture = json_decode(file_get_contents('/fixtures/hooks-v1.json'), true, 512, JSON_THROW_ON_ERROR);
if ($fixture['schema'] !== 'gossip.hook-fixture/v1') {
    throw new RuntimeException('Unsupported fixture schema');
}
$results = [];
foreach ($fixture['cases'] as $case) {
    $hook = 'gossip_fixture_' . $case['id'];
    $trace = [];
    foreach ($case['handlers'] as $index => $handler) {
        add_filter($hook, static function ($value) use ($index, $handler, &$trace) {
            $trace[] = $index;
            return $handler['prefix'] . $value . $handler['suffix'];
        }, $handler['priority'], 1);
    }
    $runs = [];
    for ($run = 0; $run < 2; $run++) {
        $trace = [];
        $value = apply_filters($hook, $case['input']);
        $runs[] = ['value' => $value, 'trace' => $trace];
    }
    remove_all_filters($hook);
    $results[$case['id']] = $runs;
}
$source_files = [];
foreach (['license.txt', 'wp-includes/version.php', 'wp-includes/plugin.php', 'wp-includes/class-wp-hook.php', 'wp-includes/class-wp-filter-sentinel.php'] as $path) {
    $source_files[$path] = hash_file('sha256', ABSPATH . $path);
}
echo json_encode([
    'schema' => 'gossip.hook-observation/v1',
    'wordpress' => $wp_version,
    'php' => PHP_VERSION,
    'source_files' => $source_files,
    'results' => $results,
], JSON_THROW_ON_ERROR | JSON_UNESCAPED_UNICODE | JSON_UNESCAPED_SLASHES);
