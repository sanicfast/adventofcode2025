package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	x int
	y int
	z int
}

func readInput() ([]Coord, int) {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")
	outCoords := []Coord{}
	for _, line := range lines {
		f := strings.Split(line, ",")
		x, e1 := strconv.Atoi(f[0])
		y, e2 := strconv.Atoi(f[1])
		z, e3 := strconv.Atoi(f[2])
		check(e1)
		check(e2)
		check(e3)
		newCoord := Coord{x, y, z}
		outCoords = append(outCoords, newCoord)
	}

	n := 1000
	if len(os.Args) > 1 {
		n = 10
	}

	return outCoords, n
}
func square(n int) int {
	return n * n
}

func distSq(a, b Coord) int {
	hyp := square(a.x-b.x) + square(a.y-b.y) + square(a.z-b.z)
	return hyp
}

// combines two sets and merges their contents
func setMerge(setA, setB map[Coord]struct{}) {
	for k := range setB {
		setA[k] = struct{}{}
	}
}

func connectN(coords []Coord, n int) int {
	nCoords := len(coords)
	// key=coord id -> value=set containing coord
	coordMap := make(map[Coord]map[Coord]struct{}, nCoords)
	for _, c := range coords {
		circuitSet := make(map[Coord]struct{}, 1)
		circuitSet[c] = struct{}{}
		coordMap[c] = circuitSet
	}
	connections := []Coord{}
	for i := range nCoords {
		for j := i + 1; j < nCoords; j++ {
			// fmt.Println(i, j)
			dist := distSq(coords[i], coords[j])
			// a connection is stored as the indexes of the points and the dist between
			connections = append(connections, Coord{i, j, dist})
		}
	}
	lessFunc := func(a, b int) bool { return connections[a].z < connections[b].z }
	sort.Slice(connections, lessFunc)

	for i := range n {
		d := connections[i]
		a, b := coords[d.x], coords[d.y]
		// is point b in the same circuit set as point a?
		// if so do nothing. if not add it and then
		_, found := coordMap[a][b]
		if !found {
			setMerge(coordMap[a], coordMap[b])
			for k := range coordMap[b] {
				coordMap[k] = coordMap[a]
			}
		}
	}

	finalSets := map[string]int{}
	for _, v := range coordMap {
		addr := fmt.Sprintf("%p", v)
		finalSets[addr] = len(v)
	}
	circuitSizes := []int{}
	for _, v := range finalSets {
		// fmt.Println(k, v)
		circuitSizes = append(circuitSizes, v)
	}
	slices.SortFunc(circuitSizes, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	output := 1
	for _, size := range circuitSizes[:3] {
		output *= size
	}
	return output
}

// check whether all of the values point to the same value in memory
func isOneSet(coordMap map[Coord]map[Coord]struct{}) bool {
	i := -1
	var addrPrev string
	for _, v := range coordMap {
		addr := fmt.Sprintf("%p", v)
		i += 1
		if i == 0 {
			addrPrev = addr
			continue
		}
		if addr != addrPrev {
			return false
		}
	}
	return true
}

func connectUntilOneCircuit(coords []Coord) int {
	nCoords := len(coords)
	// key=coord id -> value=set containing coord
	coordMap := make(map[Coord]map[Coord]struct{}, nCoords)
	for _, c := range coords {
		circuitSet := make(map[Coord]struct{}, 1)
		circuitSet[c] = struct{}{}
		coordMap[c] = circuitSet
	}
	connections := []Coord{}
	for i := range nCoords {
		for j := i + 1; j < nCoords; j++ {
			// fmt.Println(i, j)
			dist := distSq(coords[i], coords[j])
			// a connection is stored as the indexes of the points and the dist between
			connections = append(connections, Coord{i, j, dist})
		}
	}
	lessFunc := func(a, b int) bool { return connections[a].z < connections[b].z }
	sort.Slice(connections, lessFunc)

	// just loop until they're all in the same circuit and return the x coords of the last connection
	x1, x2 := -1, 999
	for i := 0; !isOneSet(coordMap); i++ {
		d := connections[i]
		a, b := coords[d.x], coords[d.y]
		x1, x2 = a.x, b.x
		// is point b in the same circuit set as point a?
		// if so do nothing. if not add it and then
		_, found := coordMap[a][b]
		if !found {
			setMerge(coordMap[a], coordMap[b])
			for k := range coordMap[b] {
				coordMap[k] = coordMap[a]
			}
		}
	}
	return x1 * x2
}

func main() {
	startTime := time.Now()

	coords, n := readInput()
	part1 := connectN(coords, n)
	fmt.Println("Day 8, Part 1:", part1)
	part2 := connectUntilOneCircuit(coords)
	fmt.Println("Day 8, Part 2:", part2)

	fmt.Println(time.Since(startTime))

}
