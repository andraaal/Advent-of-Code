package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("C:\\Dateien\\Programieren\\Advent of Code\\2022\\08\\input.txt")
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

	var forest [99][99]int
	var valid [99][99]bool

	for i, line := range strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n") {
		for j, num := range strings.Split(line, "") {
			log.Print(i, j)
			forest[i][j], _ = strconv.Atoi(num)
			valid[i][j] = true
		}
	}

	maximum := 0
	for i, r := range valid {
		for j := range r {
			tree := checkTree(i, j, forest)
			if tree > maximum {
				maximum = tree
			}
		}
	}
	fmt.Println(maximum)

}

func checkTree(line int, col int, forest [99][99]int) int {
	var height int = forest[line][col]
	rightSum, leftSum, upSum, downSum := 0, 0, 0, 0
	for i := line + 1; i < len(forest); i++ {
		rightSum++
		if forest[i][col] >= height {
			break
		}
	}
	for i := line - 1; i >= 0; i-- {
		leftSum++
		if forest[i][col] >= height {
			break
		}
	}

	for i := col + 1; i < len(forest[0]); i++ {
		downSum++
		if forest[line][i] >= height {
			break
		}
	}
	for i := col - 1; i >= 0; i-- {
		upSum++
		if forest[line][i] >= height {
			break
		}
	}

	return downSum * upSum * leftSum * rightSum
}
