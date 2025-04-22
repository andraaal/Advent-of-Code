package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	id           int
	items        []int
	test         int
	ifTrue       int
	ifFalse      int
	operator     rune
	operand      int
	inspectCount int
}

func main() {
	f, err := os.Open("C:\\Dateien\\Programieren\\Advent of Code\\2022\\11\\input.txt")
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

	monkeyStrings := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n\n")
	monkeys := make([]monkey, len(monkeyStrings))

	for j, monkeyDef := range monkeyStrings {
		lines := strings.Split(monkeyDef, "\n")
		num, _ := strconv.Atoi(lines[0][7 : len(lines[0])-1])
		strItems := strings.Split(lines[1][18:], ", ")
		items := make([]int, len(strItems))
		for i := 0; i < len(strItems); i++ {
			items[i], _ = strconv.Atoi(strItems[i])
		}
		test, _ := strconv.Atoi(lines[3][21:])
		ifTrue, _ := strconv.Atoi(lines[4][29:])
		ifFalse, _ := strconv.Atoi(lines[5][30:])

		operation := lines[2][23:]
		operator := rune(operation[0])
		operand, err := strconv.Atoi(operation[2:])
		if err != nil {
			operand = 2
			operator = '^'
		}

		monk := monkey{
			id:           num,
			items:        items,
			test:         test,
			ifTrue:       ifTrue,
			ifFalse:      ifFalse,
			operator:     operator,
			operand:      operand,
			inspectCount: 0,
		}

		monkeys[j] = monk
	}

	for k := 0; k < 10000; k++ {
		for _, monk := range monkeys {
			simulateMonkey(monk.id, monkeys)
		}
	}
	highest := 0
	secondHighest := 0
	for _, monk := range monkeys {
		if monk.inspectCount > highest {
			secondHighest = highest
			highest = monk.inspectCount
		} else if monk.inspectCount > secondHighest {
			secondHighest = monk.inspectCount
		}
	}
	fmt.Printf("The solution is %d", secondHighest*highest)
}

func simulateMonkey(id int, monkeys []monkey) {
	monk := &monkeys[id]
	forTrue := make([]int, 0, len(monk.items))
	forFalse := make([]int, 0, len(monk.items))
	for i := range monk.items {
		log.Printf("Monkey %d: %d %s %d", monk.id, monk.items[i], string(monk.operator), monk.operand)
		switch monk.operator {
		case '+':
			monk.items[i] += monk.operand
			break
		case '*':
			monk.items[i] *= monk.operand
			break
		case '^':
			monk.items[i] *= monk.items[i]
		}
		monk.items[i] %= 9699690 // This is the least common multiple of all the test values
		log.Printf("After: %d", monk.items[i])
		if monk.items[i]%monk.test == 0 {
			forTrue = append(forTrue, monk.items[i])
		} else {
			forFalse = append(forFalse, monk.items[i])
		}
		monk.inspectCount++
	}

	monk.items = nil
	monkeys[monk.ifTrue].items = append(monkeys[monk.ifTrue].items, forTrue...)
	monkeys[monk.ifFalse].items = append(monkeys[monk.ifFalse].items, forFalse...)
}
