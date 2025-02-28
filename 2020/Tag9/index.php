<?php

$counter = 0;
$max = 99999;
$falseNumber = 530627549;

// custom vars - START
$beforeNumbers = [];
$firstNotValidNumber = 0;
$currentNumber = 0;
$addedNumbers = [];
$result = 0;
// custom vars - END

$handle = fopen('input', 'rb');
while (($line = fgets($handle)) !== false) {
    $line = (int) trim($line);
    var_dump($line, $currentNumber, '------------------------------------');
    if ($currentNumber === $falseNumber) {
        sort($addedNumbers);
        $result = $addedNumbers[0] + $addedNumbers[count($addedNumbers) - 1];
        break;
    }
    if ($currentNumber < $falseNumber) {
        $currentNumber += $line;
        $addedNumbers[] = $line;
    } else {
        while ($currentNumber > $falseNumber && isset($addedNumbers[0])) {
            var_dump($addedNumbers);
            $currentNumber -= $addedNumbers[0];
            unset($addedNumbers[0]);
            $addedNumbers = array_values($addedNumbers);
        }
        if ($currentNumber !== $falseNumber) {
            $currentNumber += $line;
            $addedNumbers[] = $line;
        }
    }

    // put code here - START
    /* if ($counter > 25) {
      while (true) {

      foreach ($beforeNumbers as $number) {
      foreach ($beforeNumbers as $secondNumber) {
      if ($number !== $secondNumber && $number + $secondNumber === $line) {
      break 3;
      }
      }
      }
      $firstNotValidNumber = $line;
      break;
      }
      }

      $beforeNumbers[$counter % 25] = $line; */
    // put code here - END
    $counter++;
    if ($counter >= $max) {
        break;
    }
}
fclose($handle);

function check($line, $beforeNumbers) {

}

// output result
var_dump($addedNumbers);
echo 'ergebnis: ' . $result . "\r\n";
echo 'gefundene Zeilen: ' . $counter . "\r\n";
