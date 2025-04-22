package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Knot struct {
	posX int
	posY int
}

func main() {
	f, err := os.Open("C:\\Dateien\\Programieren\\Advent of Code\\2022\\09\\input.txt")
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
	var positions map[string]struct{} = map[string]struct{}{}
	var knots []Knot = make([]Knot, 10)

	for _, command := range lines {
		cmdParts := strings.Split(command, " ")
		quantifier, _ := strconv.Atoi(cmdParts[1])

		for i := 0; i < quantifier; i++ {
			switch cmdParts[0] {
			case "R":
				knots[0].posX++
				break
			case "L":
				knots[0].posX--
				break
			case "D":
				knots[0].posY++
				break
			case "U":
				knots[0].posY--
				break
			default:
				panic("Invalid command")
			}

			for j := 1; j < len(knots); j++ {
				move(knots, j)
			}

			var hash string = strconv.Itoa(knots[9].posX) + ":" + strconv.Itoa(knots[9].posY)
			positions[hash] = struct{}{}
		}
	}
	for position := range positions {
		fmt.Println(position)
	}
	fmt.Println(len(positions))
}

func move(knots []Knot, knotIndex int) {
	head := &knots[knotIndex-1]
	tail := &knots[knotIndex]
	headX := &head.posX
	headY := &head.posY
	tailX := &tail.posX
	tailY := &tail.posY

	if math.Abs(float64(*headX-*tailX))+math.Abs(float64(*headY-*tailY)) >= 3 {
		switch true {
		case *headX > *tailX && *headY > *tailY:
			*tailY++
			*tailX++
			break
		case *headX > *tailX && *headY < *tailY:
			*tailX++
			*tailY--
			break
		case *headX < *tailX && *headY < *tailY:
			*tailX--
			*tailY--
			break
		case *headX < *tailX && *headY > *tailY:
			*tailX--
			*tailY++
			break
		default:
			panic("Distance between head and tail too big")
		}
	} else if math.Abs(float64(*headX-*tailX)) > 1 || math.Abs(float64(*headY-*tailY)) > 1 {
		switch true {
		case *headX > *tailX+1:
			*tailX++
			break
		case *headX < *tailX-1:
			*tailX--
			break
		case *headY > *tailY+1:
			*tailY++
			break
		case *headY < *tailY-1:
			*tailY--
			break
		default:
			panic("Can't find direction to move")
		}
	}
}
