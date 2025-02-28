<?php

/* $counter = 0;
  $max = 9999;

  // custom vars - START
  const LENGTH = 1;
  $distance = 0;
  $trees = 0;
  // custom vars - END

  $handle = fopen('input', 'rb');
  while (($line = fgets($handle)) !== false) {
  $line = trim($line);

  // put code here - START
  if (substr($line, $distance, LENGTH) === '#') {
  $trees++;
  }
  $distance = ($distance + 3) % 31;
  // put code here - END

  $counter++;
  if ($counter >= $max) {
  break;
  }
  }
  fclose($handle);

  // output result
  echo 'gefundene Bäume: ' . $trees . "\r\n";
  echo 'gefundene Zeilen: ' . $counter . "\r\n"; */

// Part 2

function getTrees($x, $y) {
    $counter = 0;
    $max = 9999;

// custom vars - START
    $length = 1;
    $distance = 0;
    $trees = 0;
// custom vars - END

    $handle = fopen('input', 'rb');
    while (($line = fgets($handle)) !== false) {
        $continue = $counter % $y !== 0;
        $counter++;
        if ($continue) {
            continue;
        }
        $line = trim($line);

        // put code here - START
        if (substr($line, $distance, $length) === '#') {
            $trees++;
        }
        $distance = ($distance + $x) % 31;
        // put code here - END
        if ($counter >= $max) {
            break;
        }
    }
    fclose($handle);

// output result
    echo 'gefundene Bäume: ' . $trees . "\r\n";
    echo 'gefundene Zeilen: ' . $counter . "\r\n";
    return $trees;
}

$one = getTrees(1, 1);
$two = getTrees(3, 1);
$three = getTrees(5, 1);
$four = getTrees(7, 1);
$five = getTrees(1, 2);
echo $one * $two * $three * $four * $five;
