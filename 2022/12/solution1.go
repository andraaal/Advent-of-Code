package main

import (
	"io"
	"os"
	"strings"
)

type square struct {
	row, line int
	f, g, h   int
	parent    *square
}

func main() {
	f, err := os.Open("C:\\Dateien\\Programieren\\Advent of Code\\2022\\12\\input.txt")
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

	grid := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	closedList := make(map[square]struct{})
	openList := make(map[square]struct{})

	row, line := find(grid, 'E')
	start := square{row, line, 0, 0, 0, nil}
	openList[start] = struct{}{}

	for len(openList) > 0 {
		closedList[start] = struct{}{}
		delete(openList, start)

	}

}

func findLowest(arr map[square]struct{}) *square {
	var maximum *square
	for _, elem := range arr {
		if elem.f > maximum.f {
			maximum = elem
		}
	}
}

func find(grid []string, target rune) (int, int) {
	for i, row := range grid {
		return i, strings.IndexRune(row, target)
	}
	return -1, -1
}
