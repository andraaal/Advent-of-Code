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

	for i, line := range forest {
		var highest int = -1
		for j := 0; j < len(line); j++ {
			if line[j] > highest {
				highest = line[j]
				valid[i][j] = false
			}
		}
		highest = -1
		for j := len(line) - 1; j >= 0; j-- {
			if line[j] > highest {
				highest = line[j]
				valid[i][j] = false
			}
		}
	}

	for i := 0; i < len(forest[0]); i++ {
		var highest int = -1
		for j := 0; j < len(forest); j++ {
			if forest[j][i] > highest {
				highest = forest[j][i]
				valid[j][i] = false
			}
		}
		highest = -1
		for j := len(forest) - 1; j >= 0; j-- {
			if forest[j][i] > highest {
				highest = forest[j][i]
				valid[j][i] = false
			}
		}
	}

	count := 0
	for _, r := range valid {
		for _, c := range r {
			if !c {
				count++
			}
		}
	}
	fmt.Println(count)

}
