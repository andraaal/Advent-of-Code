<?php

$counter = 0;
$max = 9999;

// custom vars - START
$validBags = 0;
$bags = [];
// custom vars - END

$handle = fopen('input', 'rb');
while (($line = fgets($handle)) !== false) {
    $line = trim($line);

    // put code here - START
    $currentInfos = explode(' bags contain ', $line);

    $currentBagColor = str_replace(' ', '_', $currentInfos[0]);
    // var_dump($currentBagColor);

    $innerBags = explode(', ', substr($currentInfos[1], 0, -1));
    $currentInnerBags = [];
    foreach ($innerBags as $i => $currentInnerBag) {
        if ($currentInnerBag !== 'no other bags') {
            $currentInnerBags[$i][0] = str_replace(' ', '_', trim(substr($currentInnerBag, 2, -4)));
            var_dump($currentInnerBag);
            $currentInnerBags[$i][1] = trim(substr($currentInnerBag, 0, 1));
        }
    }
    // var_dump($currentInnerBags);


    $bags[$currentBagColor] = $currentInnerBags;

    // put code here - END

    $counter++;
    if ($counter >= $max) {
        break;
    }
}
fclose($handle);

// var_dump($bags['shiny gold']);
// var_dump($bags['muted fuchsia']);

function getInnerBags(string $color, array $bags): int {

    if (!empty($bags[$color])) {
        $sum = 0;
        foreach ($bags[$color] as $currentInnerBag) {
            for ($i = 0; $i < (int) $currentInnerBag[1]; $i++) {
                $sum += getInnerBags($currentInnerBag[0], $bags) + 1;
            }
        }
        return $sum;
    }
    return 0;
    /* var_dump($bags[$color], $color);
      if (array_key_exists($color, $bags)) {
      $sum = [];
      foreach ($bags[$color] as $innerColor) {

      // var_dump($innerColor, getColors($innerColor, $bags));
      $sum = array_merge($sum, getColors($innerColor, $bags), [$innerColor]);
      $sum = array_unique($sum);
      }
      return $sum;
      }
      return []; */
}

// var_dump($bags);
var_dump(getInnerBags('shiny_gold', $bags));
// output result
// var_dump($bags);
// echo 'gefundene richtige Koffer: ' . $trees . "\r\n";
echo 'gefundene Regeln: ' . $counter . "\r\n";
