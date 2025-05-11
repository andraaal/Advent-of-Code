package main

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

	var table [4000000]int
	var instructions [][]int = make([][]int, len(lines))

	for i, line := range lines {
		positions := strings.Split(line, ": closest beacon is at x=")
		senderPosString := strings.Split(positions[0], ", y=")
		beaconPosString := strings.Split(positions[1], ", y=")
		var senderX, senderY, beaconX, beaconY int
		senderX, _ = strconv.Atoi(senderPosString[0][12:])
		beaconX, _ = strconv.Atoi(beaconPosString[0])
		senderY, _ = strconv.Atoi(senderPosString[1])
		beaconY, _ = strconv.Atoi(beaconPosString[1])

		distBetween := abs(senderX-beaconX) + abs(senderY-beaconY)

		instructions[i] = []int{senderX, senderY, distBetween}
	}

	for {
		for i, val := range table {
			if val >= 4000000 {
				continue
			}
			changed := false
			for _, instr := range instructions {
				posX := instr[0]
				posY := instr[1]
				radius := instr[2]
				distToTarget := abs(i - posY)
				filledRange := radius - distToTarget
				if posX-filledRange <= val+1 && posX+filledRange > val {
					table[i] = posX + filledRange
					changed = true
				}
			}
			// fmt.Println(table)
			if !changed {
				fmt.Println("X: " + strconv.Itoa(val+1) + " | Y: " + strconv.Itoa(i))
				fmt.Println("Tuning value: " + strconv.Itoa((val+1)*4000000+i))
				return
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
