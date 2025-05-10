package main

/*
import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("C:\\Dateien\\Programieren\\Advent of Code\\2022\\15\\input.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	data, readErr := io.ReadAll(f)
	if readErr != nil {
		panic(readErr)
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	targetY := 2000000
	var occupiedSpaces map[int]struct{} = make(map[int]struct{})
	var beacons map[int]struct{} = make(map[int]struct{})

	for _, line := range lines {
		positions := strings.Split(line, ": closest beacon is at x=")
		senderPosString := strings.Split(positions[0], ", y=")
		beaconPosString := strings.Split(positions[1], ", y=")
		var senderX, senderY, beaconX, beaconY int
		senderX, err = strconv.Atoi(senderPosString[0][12:])
		beaconX, _ = strconv.Atoi(beaconPosString[0])
		senderY, _ = strconv.Atoi(senderPosString[1])
		beaconY, _ = strconv.Atoi(beaconPosString[1])

		if beaconY == targetY {
			beacons[beaconX] = struct{}{}
		}

		distBetween := abs(senderX-beaconX) + abs(senderY-beaconY)
		distToTarget := abs(senderY - targetY)
		filledRange := distBetween - distToTarget

		// fmt.Println(distBetween, distToTarget, filledRange)

		for i := senderX - filledRange; i <= senderX+filledRange; i++ {
			occupiedSpaces[i] = struct{}{}
		}

	}
	fmt.Println(len(occupiedSpaces) - len(beacons))

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
*/
