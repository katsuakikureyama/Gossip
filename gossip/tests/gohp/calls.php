// This echo oracle deliberately covers only scalar returns. Array output
// coercion is not a type-preserving PHP observation; typed-observer.php
// compares __return_empty_array() separately with the other generated calls.
echo 'true=' . __return_true() . "\n";
echo 'false=' . __return_false() . "\n";
echo 'zero=' . __return_zero() . "\n";
echo 'null=' . __return_null() . "\n";
echo 'empty-string=' . __return_empty_string() . "\n";
