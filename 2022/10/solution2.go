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
	f, err := os.Open("C:\\Dateien\\Programieren\\Advent of Code\\2022\\10\\input.txt")
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
	sum := 0
	crt := [6][40]rune{}
	registerVal := 1
	executed := 0
	impending := 0

	for cycles := 1; executed < len(lines); cycles++ {
		log.Printf("Cycle %d: Register %d\n", cycles, registerVal)

		if cycles <= 220 && cycles%40-20 == 0 {
			sum += registerVal * cycles
			log.Printf("Logged in cycles: %d with register value of %d: %d", cycles, registerVal, registerVal*cycles)
		}

		log.Printf("Comparing: %d with %d", cycles%40, registerVal)
		if registerVal <= cycles%40 && cycles%40 <= registerVal+2 {
			crt[(cycles-1)/40][(cycles-1)%40] = '#'
		} else {
			crt[(cycles-1)/40][(cycles-1)%40] = '.'
		}

		if impending != 0 {
			registerVal += impending
			impending = 0
			executed++
		} else if lines[executed] == "noop" {
			executed++
		} else {
			amount, _ := strconv.Atoi(strings.Split(lines[executed], " ")[1])
			impending += amount
		}
	}

	fmt.Println(sum)
	for _, line := range crt {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}
