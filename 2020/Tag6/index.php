<?php

$counter = 0;
$max = 99999;

// custom vars - START
$sum = 0;
$commonChars = false;
// custom vars - END

$handle = fopen('input', 'rb');
while (($line = fgets($handle)) !== false) {
    $line = trim($line);

    // put code here - START
    if ($line === '') {
        $sum += count($commonChars);
        $commonChars = false;
        $counter++;
        var_dump('----------------------------------');
        if ($counter >= $max) {
            break;
        }
    } else {
        $currentChars = array_keys(count_chars($line, 1));
        var_dump($line);
        if ($commonChars === false) {
            $commonChars = $currentChars;
        } else {
            $commonChars = array_intersect($commonChars, $currentChars);
        }
        var_dump($commonChars);
    }
    // put code here - END
}
$sum += count($commonChars);
fclose($handle);

// output result
echo 'gefundene Antworten: ' . $sum . "\r\n";
echo 'gefundene Guppen: ' . $counter . "\r\n";
