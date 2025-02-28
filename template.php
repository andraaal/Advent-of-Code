<?php

$counter = 0;
$max = 9999;

// custom vars - START
// custom vars - END

$handle = fopen('input', 'rb');
while (($line = fgets($handle)) !== false) {
    $line = trim($line);

    // put code here - START
    // put code here - END

    $counter++;
    if ($counter >= $max) {
        break;
    }
}
fclose($handle);

// output result
echo 'gefundene Bäume: ' . $trees . "\r\n";
echo 'gefundene Zeilen: ' . $counter . "\r\n";
