package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("C:\\Dateien\\Programieren\\Advent of Code\\2022\\13\\input.txt")
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

	lines := append(strings.Split(strings.ReplaceAll(string(data), "\r\n\r\n", "\r\n"), "\r\n"), "[[2]]", "[[6]]")
	slices.SortFunc(lines, sortPair)
	index1 := slices.Index(lines, "[[2]]") + 1
	index2 := slices.Index(lines, "[[6]]") + 1

	fmt.Println(index1 * index2)
	fmt.Println(lines)
}

func sortPair(leftS string, rightS string) int {
	left, right := []byte(leftS), []byte(rightS)
	leftI, rightI := 0, 0
	leftDepth, rightDepth := 0, 0

	for {
		log.Print(leftI, rightI, leftDepth, rightDepth, left, right, string(left[leftI]), string(right[rightI]))
		for left[leftI] == ',' {
			leftI++
		}
		for right[rightI] == ',' {
			rightI++
		}

		log.Print(leftI, rightI, leftDepth, rightDepth, bytes.Runes(left), right, string(left[leftI]), string(right[rightI]))
		if left[leftI] == '[' {
			leftDepth++
			leftI++
			if right[rightI] == '[' {
				rightDepth++
				rightI++
			} else if right[rightI] == ']' {
				// wrong order
				return 1
			} else {
				rightDepth++
				numbers := right[rightI : rightI+2]
				if numbers[1] == '0' {
					right = append(right[:rightI+2], append([]byte{']'}, right[rightI+2:]...)...)
				} else {
					right = append(right[:rightI+1], append([]byte{']'}, right[rightI+1:]...)...)
				}
			}
		} else if left[leftI] == ']' {
			leftDepth--
			leftI++
			if right[rightI] == ']' {
				rightDepth--
				rightI++
			} else {
				// correct order
				return -1
			}
		} else {
			if right[rightI] == ']' {
				// wrong order
				return 1
			} else if right[rightI] == '[' {
				numbers := left[leftI : leftI+2]
				if numbers[1] == '0' {
					left = append(left[:leftI+2], append([]byte{']'}, left[leftI+2:]...)...)
				} else {
					left = append(left[:leftI+1], append([]byte{']'}, left[leftI+1:]...)...)
				}
				leftDepth++
				rightDepth++
				rightI++
			} else {
				leftN, rightN := 0, 0
				var erro error
				if byte(left[leftI+1]) == ',' || byte(left[leftI+1]) == ']' || byte(left[leftI+1]) == '[' {
					leftN, erro = strconv.Atoi(string(left[leftI]))
				} else {
					leftN, erro = strconv.Atoi(string(left[leftI : leftI+2]))
				}
				if byte(right[rightI+1]) == ',' || byte(right[rightI+1]) == ']' || byte(right[rightI+1]) == '[' {
					rightN, erro = strconv.Atoi(string(right[rightI]))
				} else {
					rightN, erro = strconv.Atoi(string(right[rightI : rightI+2]))
				}
				log.Print(leftN, rightN, erro, left[leftI:leftI+2], "   ==> in comparison")
				if leftN < rightN {
					// correct order
					return -1
				} else if leftN > rightN {
					// wrong order
					return 1
				}
				rightI++
				leftI++
				if right[rightI] == '0' {
					rightI++
				}
				if left[leftI] == '0' {
					leftI++
				}
			}
		}
	}
}
