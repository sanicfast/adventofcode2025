package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() map[string][]string {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")
	edgeMap := make(map[string][]string)
	for _, line := range lines {
		split := strings.Split(line, ": ")
		key := split[0]
		vals := strings.Fields(split[1])
		edgeMap[key] = vals
	}
	return edgeMap
}

// func sortAndDedup(slice []string) []string {
// 	slices.Sort(slice)
// 	j := 0
// 	for i := 1; i < len(slice); i++ {
// 		if slice[i] != slice[j] {
// 			j++
// 			slice[j] = slice[i] // Move unique element to the next available position
// 		}
// 	}
// 	return slice[:j+1] // Return the slice up to the last unique element
// }

type Stack []string

func (stack *Stack) Push(c string) {
	*stack = append(*stack, c)
}

func (stack *Stack) Pop() string {
	n := len(*stack)
	out := (*stack)[n-1]
	*stack = (*stack)[:n-1]
	return out
}

// count all the ways from "you" to "out"
func traverse1(edgeMap map[string][]string, start, dest string) (int, error) {
	theStack := make(Stack, 0, 100)

	theStack.Push(start)

	waysOut := 0
	for len(theStack) > 0 {
		currentLocation := theStack.Pop()
		if currentLocation == dest {
			waysOut += 1
			// fmt.Println(waysOut)
			continue
		}
		if currentLocation == "out" {
			continue
		}
		newLocs, found := edgeMap[currentLocation]
		if !found {
			errorString := "NODE NOT FOUND:" + currentLocation
			return 0, errors.New(errorString)
		}
		for _, new := range newLocs {
			theStack.Push(new)
		}
	}
	return waysOut, nil
}

type infoTuple struct {
	currentLocation string
	fftPassed       bool
	dacPassed       bool
}

func recurse(edgeMap map[string][]string, cache map[infoTuple]int, current infoTuple, dest string) int {
	if current.currentLocation == dest {
		if current.dacPassed && current.fftPassed {
			return 1
		} else {
			return 0
		}
	}
	newLocs, found := edgeMap[current.currentLocation]
	if !found {
		errorString := "NODE NOT FOUND:" + current.currentLocation
		panic(errorString)
	}
	switch current.currentLocation {
	case "dac":
		current.dacPassed = true
	case "fft":
		current.fftPassed = true
	}
	sum := 0
	for _, loc := range newLocs {
		newTuple := infoTuple{loc, current.fftPassed, current.dacPassed}
		countPaths, found := cache[newTuple]
		if !found {
			countPaths = recurse(edgeMap, cache, newTuple, dest)
			cache[newTuple] = countPaths
		}
		sum += countPaths
	}
	return sum

}

func traverseWithCache(edgeMap map[string][]string, start, dest string) int {
	cache := make(map[infoTuple]int)
	startInfo := infoTuple{start, false, false}
	return recurse(edgeMap, cache, startInfo, dest)
}

func main() {
	startTime := time.Now()
	edgeMap := readInput()
	p1Time := time.Now()
	part1, e := traverse1(edgeMap, "you", "out")
	check(e)
	fmt.Println("Day 11, Part 1:", part1, time.Since(p1Time))

	p2Time := time.Now()
	part2 := traverseWithCache(edgeMap, "svr", "out")
	fmt.Println("Day 11, Part 2:", part2, time.Since(p2Time))
	fmt.Println(time.Since(startTime))
}
