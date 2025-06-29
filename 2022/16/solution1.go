package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func main() {
	f, err := os.Open("C:\\Dateien\\Programieren\\Advent of Code\\2022\\16\\input.txt")
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

	var valves = make(map[string]Valve)

	for _, line := range lines {
		line = strings.ReplaceAll(line, "valves", "valve")
		var name string
		var flowRate int
		_, err := fmt.Sscanf(line, "Valve %s has flow rate=%d;", &name, &flowRate)
		if err != nil {
			fmt.Printf("Error parsing line: %q\n", line)
			panic(err)
		}
		connections := strings.Split(strings.Split(line, "valve ")[1], ", ")
		fmt.Println(connections)

		valve := Valve{
			name:        name,
			flowRate:    flowRate,
			connections: connections,
		}

		valves[name] = valve
	}
	var maximum = 0
	recursive(valves, &maximum, "AA", 0, 30, map[string]bool{})
	fmt.Println(maximum)
}

func estimateMaxFlow(valves map[string]Valve, currentValve string, timeLeft int, openedValves map[string]bool) int {
	estimatedMaxFlow := 0
	if openedValves[currentValve] {
		timeLeft -= 1
	}

	var flowRates []int
	for _, valve := range valves {
		if !openedValves[valve.name] && valve.flowRate > 0 {
			flowRates = append(flowRates, valve.flowRate)
		}
	}
	slices.Sort(flowRates)
	slices.Reverse(flowRates)

	for _, valve := range flowRates {
		estimatedMaxFlow += valve * timeLeft
		timeLeft -= 2
		if timeLeft <= 0 {
			break
		}
	}
	return estimatedMaxFlow
}

func recursive(valves map[string]Valve, maximum *int, currentValve string, currentFlow int, timeLeft int, openedValves map[string]bool) {
	// First decrement time for whatever action will be taken
	timeLeft -= 1

	// If time is up, end the branch
	if timeLeft <= 0 {
		if currentFlow > *maximum {
			*maximum = currentFlow
		}
		return
	}

	// If it can't beat the maximum, end the branch
	if currentFlow+estimateMaxFlow(valves, currentValve, timeLeft, openedValves) < *maximum {
		return
	}

	// Open the current valve, if it is not already and its flow rate is greater than 0
	if !openedValves[currentValve] && valves[currentValve].flowRate > 0 {
		openedValves[currentValve] = true
		recursive(valves, maximum, currentValve, currentFlow+valves[currentValve].flowRate*timeLeft, timeLeft, openedValves)
		openedValves[currentValve] = false
	}

	// Check if all valves with flow rate > 0 have been opened
	allValvesOpened := true
	for _, valve := range valves {
		if valve.flowRate > 0 && !openedValves[valve.name] {
			allValvesOpened = false
		}
	}
	if allValvesOpened {
		if currentFlow > *maximum {
			*maximum = currentFlow
		}
		return
	}

	// Move to an adjacent valve
	for _, nextValve := range valves[currentValve].connections {
		recursive(valves, maximum, nextValve, currentFlow, timeLeft, openedValves)
	}
}

type Valve struct {
	name        string
	flowRate    int
	connections []string
}
