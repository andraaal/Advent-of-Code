package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type square struct {
	row, col int
	f, g, h  int
	height   int
	parent   *square
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

	tempGrid := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	grid := make([][]byte, len(tempGrid))
	for i, row := range tempGrid {
		grid[i] = []byte(row)
	}
	goalRow, goalCol := findByte(grid, 'E')
	rowStart, colStart := findByte(grid, 'S')
	grid[rowStart][colStart] = 'a'

	lowest := 1000000000

	for row, gridRow := range grid {
		for col, char := range gridRow {
			if char == 'a' {
				distance := findPath(grid, goalRow, goalCol, row, col)
				if distance > 0 && distance < lowest {
					lowest = distance
				}
				fmt.Println(distance)
			}
		}
	}
	fmt.Println(lowest)
}

func findPath(grid [][]byte, goalRow int, goalCol int, rowStart int, colStart int) int {
	closedList := make(map[*square]struct{})
	openList := make(map[*square]struct{})

	// row, col := findByte(grid, 'S')
	start := square{rowStart, colStart, 0, 0, 0, 97, nil}
	openList[&start] = struct{}{}

	for len(openList) > 0 {
		current := findLowest(&openList)
		closedList[current] = struct{}{}
		delete(openList, current)

		if current.row == goalRow && current.col == goalCol {
			depth := 0
			level := current
			for level.parent != nil {
				level = level.parent
				depth++
				// grid[level.row][level.col] = byte('#')
				// fmt.Println(*level)
			}
			/*for _, row := range grid {
				fmt.Println(string(row))
			}*/
			// fmt.Println(depth)
			return depth
		}

		for i := -2; i < 2; i++ {
			currentRow := current.row + i%2
			currentCol := current.col + (i+1)%2

			if currentRow < 0 || currentCol < 0 || currentRow > len(grid)-1 || currentCol > len(grid[0])-1 {
				continue
			}
			if int(grid[currentRow][currentCol])-1 > current.height || findSquare(&closedList, currentRow, currentCol) != nil {
				continue
			}

			squareOpen := findSquare(&openList, currentRow, currentCol)
			currentG := current.g + 1
			currentH := intAbs(currentRow-goalRow) + intAbs(currentCol-goalCol) - 1
			if squareOpen == nil || currentG < squareOpen.g {
				openList[&square{currentRow, currentCol, currentG + currentH, currentG, currentH, int(grid[currentRow][currentCol]), current}] = struct{}{}
			}
		}
	}
	return -1
}

func findLowest(arr *map[*square]struct{}) *square {
	var minimum *square = &square{0, 0, 100000, 0, 0, 0, nil}
	for key, _ := range *arr {
		if key.f < minimum.f {
			minimum = key
		}
	}
	return minimum
}

func findByte(grid [][]byte, target byte) (int, int) {
	for i, row := range grid {
		for j, char := range row {
			if char == target {
				return i, j
			}
		}
	}
	return -1, -1
}

func findSquare(arr *map[*square]struct{}, row int, col int) *square {
	for key, _ := range *arr {
		if key.row == row && key.col == col {
			return key
		}
	}
	return nil
}

func intAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
