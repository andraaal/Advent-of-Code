<?php

$counter = 0;
$max = 9999;

// custom vars - START
$numbers = [];
// custom vars - END

$handle = fopen('input', 'rb');
while (($line = fgets($handle)) !== false) {
    $line = trim($line);

    // put code here - START
    $numbers[] = (int) $line;
    // put code here - END

    $counter++;
    if ($counter >= $max) {
        break;
    }
}
fclose($handle);

sort($numbers);

// $ways = recursion(0, $numbers, $ways);
// var_dump($numbers);
// var_dump($ways);
function recursion(int $number, array $numbers, int $ways): int {
    $last = true;

    foreach ($numbers as $currentNumber) {
        if ($currentNumber > $number && $currentNumber <= $number + 3) {
            $last = false;
            $ways = recursion($currentNumber, $numbers, $ways);
        }
    }

    if ($last) {
        $ways++;
    }

    return $ways;
}

$ways = 1;
$previous = 0;
foreach ($numbers as $number) {
    foreach ($numbers as $secondNumber) {
        if ($number - 3 <= $secondNumber && $secondNumber < $number) {
            $previous += 1;
            echo $previous;
        }
    }
    $ways *= $previous;
    $previous = 0;
}
var_dump($ways);


// output result
// echo 'gefundene Bäume: ' . "\r\n";
// echo 'gefundene Zeilen: ' . $counter . "\r\n";
