<?php

$counter = 0;
$max = 99999;

// custom vars - START
$necessaryKeys = ['ecl', 'byr', 'iyr', 'hcl', 'pid', 'hgt', 'eyr'];
$validECL = ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth',];

$validPasswords = 0;
$keys = [];
// custom vars - END

$handle = fopen('input', 'rb');
while (($line = fgets($handle)) !== false) { // hgt:181cm pid:192793676 hcl:8f6ae6 iyr:2021 eyr:1978
    $line = trim($line);
    // var_dump($line);
// put code here - START
    if ($line === '') {
// var_dump($line);

        if (array_diff($necessaryKeys, array_keys($keys)) === []) {
            $valid = true;
            if (!in_array($keys['ecl'], $validECL)) {
                $valid = false;
            }
            if (strlen($keys['byr']) !== 4 || (int) $keys['byr'] < 1920 || (int) $keys['byr'] > 2002) {
                $valid = false;
            }
            if (strlen($keys['iyr']) !== 4 || (int) $keys['iyr'] < 2010 || (int) $keys['iyr'] > 2020) {
                $valid = false;
            }
            $pattern = '/^#[a-f0-9]{6}$/i';
            // var_dump($keys['hcl']);
            if (preg_match($pattern, $keys['hcl']) === 0) {
                $valid = false;
            }
            $pattern = '/^[0-9]{9}$/';
            if (preg_match($pattern, $keys['pid']) === 0) {
                $valid = false;
            }
            $in = (int) str_replace('/in$/', '', $keys['hgt']);
            $cm = (int) str_replace('/cm$/', '', $keys['hgt']);
            if (preg_match('/in$/', $keys['hgt']) === 0 && preg_match('/cm$/', $keys['hgt']) === 0) {
                $valid = false;
            } else
            if (preg_match('/in$/', $keys['hgt']) === 1 && ( $in < 59 || $in > 76)) {
                $valid = false;
            } else
            if (preg_match('/cm$/', $keys['hgt']) === 1 && ( $cm < 150 || $cm > 193)) {
                $valid = false;
            }
            if (strlen($keys['eyr']) !== 4 || (int) $keys['eyr'] < 2020 || (int) $keys['eyr'] > 2030) {
                $valid = false;
            }

            // var_dump(array_keys($keys), $counter);
            if ($valid) {
                var_dump($keys);
                $validPasswords++;
            }
// }
        }
        $keys = [];
        $counter++;
        if ($counter >= $max) {
            break;
        }
    } else {
        $pairs = explode(' ', $line);
        foreach ($pairs as $pair) {
            list($key, $value) = explode(':', $pair);
            $keys[$key] = $value;
        }
    }

// put code here - END
}
if (array_diff($necessaryKeys, array_keys($keys)) === []) {
    $valid = true;
    if (!in_array($keys['ecl'], $validECL)) {
        $valid = false;
    }
    if (strlen($keys['byr']) !== 4 || (int) $keys['byr'] < 1920 || (int) $keys['byr'] > 2002) {
        $valid = false;
    }
    if (strlen($keys['iyr']) !== 4 || (int) $keys['iyr'] < 2010 || (int) $keys['iyr'] > 2020) {
        $valid = false;
    }
    $pattern = '/^#[a-f0-9]{6}$/i';
    // var_dump($keys['hcl']);
    if (preg_match($pattern, $keys['hcl']) === 0) {
        $valid = false;
    }
    $pattern = '/^[0-9]{9}$/';
    if (preg_match($pattern, $keys['pid']) === 0) {
        $valid = false;
    }
    $in = (int) str_replace('/in$/', '', $keys['hgt']);
    $cm = (int) str_replace('/cm$/', '', $keys['hgt']);
    if (preg_match('/in$/', $keys['hgt']) === 0 && preg_match('/cm$/', $keys['hgt']) === 0) {
        $valid = false;
    } else
    if (preg_match('/in$/', $keys['hgt']) === 1 && ( $in < 59 || $in > 76)) {
        $valid = false;
    } else
    if (preg_match('/cm$/', $keys['hgt']) === 1 && ( $cm < 150 || $cm > 193)) {
        $valid = false;
    }
    if (strlen($keys['eyr']) !== 4 || (int) $keys['eyr'] < 2020 || (int) $keys['eyr'] > 2030) {
        $valid = false;
    }

    // var_dump(array_keys($keys), $counter);
    if ($valid) {
        var_dump($keys);
        $validPasswords++;
    }
}
fclose($handle);

// output result
echo 'gefundene korrekte Passwörter: ' . $validPasswords . "\r\n";
echo 'gefundene Pässe: ' . $counter . "\r\n";




/*$allValid = true;
            foreach ($keys as $key => $values) {
                foreach ($values as $value) {
                    switch ($key) {
                        case 'byr':
                        case 'iyr':
                        case 'eyr':
                            if (!is_numeric($value) || strlen($value) !== 4) {
                                $allValid = false;
                            }
                            break;
                        case 'pid':
                            if (!is_numeric($value) || strlen($value) !== 9) {
                                $allValid = false;
                            }
                            break;
                        case 'hgt':
                            if () {
                                $allValid = false;
                            }
                            break;
                    }
                    if (!$allValid) {
                        break 2;
                    }
                }
            }

            if ($allValid) {*/