<?php

$counter = 0;
$max = 9999;

// custom vars - START
$highestID = 0;
$array = [];
for ($i = 0; $i < 128; $i++) {
    $array[$i] = '........';
}
// custom vars - END

$handle = fopen('input', 'rb');
while (($line = fgets($handle)) !== false) {
    $line = trim($line);

    // put code here - START
    $row = substr($line, 0, 7);
    $column = substr($line, 7, 3);
    $row = str_replace('F', '0', $row);
    $row = str_replace('B', '1', $row);
    $column = str_replace('R', '1', $column);
    $column = str_replace('L', '0', $column);
    $row = bindec($row);
    $column = bindec($column);
    $array[$row] = substr_replace($array[$row], 'X', $column, 1);
    // var_dump($array[$row]);
    if ($highestID < ($row * 8) + $column) {
        $highestID = ($row * 8) + $column;
    }
    // put code here - END

    $counter++;
    if ($counter >= $max) {
        break;
    }
}
fclose($handle);

// output result
echo 'höchste Sitz ID: ' . $highestID . "\r\n";
echo 'gefundene Zeilen: ' . $counter . "\r\n";
var_dump($array);
// 83 7
// echo (83 * 8) + 7;
