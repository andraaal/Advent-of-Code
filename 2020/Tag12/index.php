<?php

/* ---------------------- PART 1 ----------------------
  $counter = 0;
  $max = 99999;

  // custom vars - START
  $direction = 'E';
  $y = 0;
  $x = 0;
  $left = ['E' => 'N', 'N' => 'W', 'W' => 'S', 'S' => 'E'];
  $right = ['E' => 'S', 'S' => 'W', 'W' => 'N', 'N' => 'E'];
  // custom vars - END

  $handle = fopen('input', 'rb');
  while (($line = fgets($handle)) !== false) {
  $line = trim($line);
  $command = substr($line, 0, 1);
  $number = (int) substr($line, 1);

  // put code here - START
  // var_dump($command);
  switch ($command) {
  case 'E':
  $y += $number;
  break;
  case 'W':
  $y -= $number;
  break;
  case 'S':
  $x -= $number;
  break;
  case 'N':
  $x += $number;
  break;
  case 'R':
  $deegrees = $number / 90;
  for ($i = 0; $i < $deegrees; $i++) {
  $direction = $right[$direction];
  }
  break;
  case 'L':
  $deegrees = $number / 90;
  for ($i = 0; $i < $deegrees; $i++) {
  $direction = $left[$direction];
  }
  break;
  case 'F':
  switch ($direction) {
  case 'E':
  $y += $number;
  break;
  case 'W':
  $y -= $number;
  break;
  case 'S':
  $x -= $number;
  break;
  case 'N':
  $x += $number;
  break;
  }
  break;
  }
  // put code here - END

  $counter++;
  if ($counter >= $max) {
  break;
  }
  }
  fclose($handle);

  // var_dump($e, $w, $s, $n);

  $distance = abs($y) + abs($x);


  // output result
  echo 'Entfernung zum Schiff: ' . $distance . "\r\n";
  echo 'gefundene Zeilen: ' . $counter . "\r\n"; */

// ---------------------- PART 2 ----------------------

$counter = 0;
$max = 99999;

// custom vars - START
$ship = new ship();
// custom vars - END

$handle = fopen('input', 'rb');
while (($line = fgets($handle)) !== false) {
    $line = trim($line);
    $command = substr($line, 0, 1);
    $number = (int) substr($line, 1);
    $ship->executeCommand($command, $number);
}
$ship->getManhattenDistance();
var_dump($ship->getManhattenDistance());

class ship {

    const LEFT = ['E' => 'N', 'N' => 'W', 'W' => 'S', 'S' => 'E'];
    const RIGHT = ['E' => 'S', 'S' => 'W', 'W' => 'N', 'N' => 'E'];
    const MOVE_WAYPOINT_RIGHT = ['E' => ['x' => '', 'y' => '-'], 'N' => ['x' => '', 'y' => '-'], 'W' => ['x' => '-', 'y' => ''], 'S' => ['x' => '-', 'y' => '']];
    const MOVE_WAYPOINT_LEFT = ['E' => ['x' => '-', 'y' => ''], 'N' => ['x' => '-', 'y' => ''], 'W' => ['x' => '', 'y' => '-'], 'S' => ['x' => '', 'y' => '-']];

    private $waypointY;
    private $waypointX;
    private $waypointDirection;
    private $x;
    private $y;

    public function __construct() {
        $this->waypointY = 10;
        $this->waypointX = 1;
        $this->waypointDirection = 'E';
        $this->x = 0;
        $this->y = 0;
    }

    public function executeCommand(string $command, int $number) {
        switch ($command) {
            case 'E':
                $this->waypointY += $number;
                break;
            case 'W':
                $this->waypointY -= $number;
                break;
            case 'S':
                $this->waypointX -= $number;
                break;
            case 'N':
                $this->waypointX += $number;
                break;
            case 'F':
                for ($i = 0; $i < $number; $i++) {
                    $this->y += $this->waypointY;
                    var_dump($this->waypointX);
                    while (substr($this->waypointX, 0, 1) === '0') {
                        $this->waypointX = substr($this->waypointX, 1);
                    }
                    $this->x += $this->waypointX;
                }
                break;
            case 'R':
                for ($i = 0; $i < $number / 90; $i++) {
                    $some = $this->waypointX;
                    $this->waypointX = $this->waypointY;
                    $this->waypointY = $some;
                    $this->waypointX = (int) self::MOVE_WAYPOINT_RIGHT[$this->waypointDirection]['x'] . $this->waypointX;
                    $this->waypointY = (int) self::MOVE_WAYPOINT_RIGHT[$this->waypointDirection]['y'] . $this->waypointY;
                    $this->waypointDirection = self::RIGHT[$this->waypointDirection];
                }
                break;
            case 'L':
                for ($i = 0; $i < $number / 90; $i++) {
                    $some = $this->waypointX;
                    $this->waypointX = $this->waypointY;
                    $this->waypointY = $some;
                    $this->waypointX = (int) self::MOVE_WAYPOINT_LEFT[$this->waypointDirection]['x'] . $this->waypointX;
                    $this->waypointY = (int) self::MOVE_WAYPOINT_LEFT[$this->waypointDirection]['y'] . $this->waypointY;
                    $this->waypointDirection = self::LEFT[$this->waypointDirection];
                }
                break;
            default:
        }
    }

    public function getManhattenDistance() {
        return abs($this->x) + abs($this->y);
    }

}

// es kommen manchmal 90 Nuller vorne dazu