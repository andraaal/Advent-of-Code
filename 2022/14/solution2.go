package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("C:\\Dateien\\Programieren\\Advent of Code\\2022\\14\\input.txt")
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

	var cave [200][900]rune
	deepestRow := 0
	for i, row := range cave {
		cave[i] = [900]rune{}
		for j := range row {
			cave[i][j] = '.'
		}
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	for _, line := range lines {
		commands := strings.Split(line, " -> ")
		lastCol, lastRow := 0, 0
		for _, command := range commands {
			corrds := strings.Split(command, ",")
			col, _ := strconv.Atoi(corrds[0])
			row, _ := strconv.Atoi(corrds[1])
			if row > deepestRow {
				deepestRow = row
			}
			if lastCol < col && lastCol != 0 {
				for i := lastCol; i != col; i++ {
					cave[row][i] = '#'
				}
			} else if lastCol > col {
				for i := lastCol; i != col; i-- {
					cave[row][i] = '#'
				}
			}
			if lastRow < row && lastRow != 0 {
				for i := lastRow; i != row; i++ {
					cave[i][col] = '#'
				}
			} else if lastRow > row {
				for i := lastRow; i != row; i-- {
					cave[i][col] = '#'
				}
			}
			cave[row][col] = '#'
			lastRow = row
			lastCol = col
		}
	}

	for i := range 900 {
		cave[deepestRow+2][i] = '#'
	}

	count := 0
	for simulateSand(&cave) {
		count++
	}

	fmt.Println(count)
}

func simulateSand(cave *[200][900]rune) bool {
	col := 500
	row := 0

	for {
		if cave[0][500] == 'S' {
			return false
		}
		if cave[row+1][col] == '.' {
			row++
		} else if cave[row+1][col-1] == '.' {
			row++
			col--
		} else if cave[row+1][col+1] == '.' {
			row++
			col++
		} else {
			cave[row][col] = 'S'
			return true
		}
	}
}
