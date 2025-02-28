<?php

$counter = 0;
$skippedCommands = 0;
$max = 9999;

// custom vars - START
$programm = [];
$accumulator = 0;
// custom vars - END

$handle = fopen('input', 'rb');
while (($line = fgets($handle)) !== false) {
    $line = trim($line);

    // put code here - START
    list($command, $position) = explode(' ', $line);
    $programLine = [
        'cmd' => $command,
        'pos' => (int) $position,
        'visited' => false,
    ];

    $programm[] = $programLine;
    // put code here - END
}
fclose($handle);

// -----------------------------------------------------------------------------



function checkPath(array $programm, int $changeLine) {

    $currentLine = 0;
    $accumulator = 0;
    $commandCounter = 0;

    while (isset($programm[$currentLine])) {

        $currentCommand = $programm[$currentLine]['cmd'];
        $currentNumber = $programm[$currentLine]['pos'];
        $visited = $programm[$currentLine]['visited'];

        if ($visited) {
            return false;
        }

        // change currentCommand
        if ($commandCounter === $changeLine && $currentCommand !== 'acc') {
            switch ($currentCommand) {
                case 'nop':
                    $currentCommand = 'jmp';
                    break;
                case 'jmp':
                    $currentCommand = 'nop';
                    break;
            }
        }

        $programm[$currentLine]['visited'] = true;
        switch ($currentCommand) {
            case 'acc':
                $accumulator += $currentNumber;
                $currentLine++;
                break;

            case 'jmp':
                $currentLine += $currentNumber;
                break;

            default:
                $currentLine++;
                break;
        }

        $commandCounter++;

        // var_dump($commandCounter);
        // var_dump($changeLine);
        // echo sprintf('%s: %s %s --> %s%s', $currentLine, $currentCommand, $currentNumber, $accumulator, "\r\n");
    }

    return $accumulator;
}

$isPathValid = false;
$changeLine = 0;
while ($isPathValid === false) {
    $isPathValid = checkPath($programm, $changeLine);
    $changeLine++;
}

// output result
echo 'Accumulator: ' . $isPathValid . "\r\n";
echo 'gefundene Zeilen: ' . $counter . "\r\n";
