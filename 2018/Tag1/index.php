<?php

$counter = 0;
$max = 9999;

// custom vars - START
$sum = 0;
// custom vars - END

$handle = fopen('input', 'rb');
while (($line = fgets($handle)) !== false) {
    $line = (int) trim($line);

    // put code here - START
    $sum += $line;
    // put code here - END

    $counter++;
    if ($counter >= $max) {
        break;
    }
}
fclose($handle);

// output result
echo 'Summe: ' . $sum . "\r\n";
echo 'gefundene Zeilen: ' . $counter . "\r\n";
