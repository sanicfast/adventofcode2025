package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Region struct {
	r           int
	c           int
	shapeCounts [6]int
}
type Shape struct {
	vec         []int
	filledCells int
}

func readInput() ([]Shape, []Region) {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")
	shapes := make([]Shape, 6)
	for shapeIdx := range 6 {
		shapeVec := make([]int, 0, 9)
		filledCells := 0
		startRow := 1 + shapeIdx*5
		for row := startRow; row < startRow+3; row++ {
			for _, char := range lines[row] {
				element := 0
				if char == '#' {
					element = 1
					filledCells += 1
				}
				shapeVec = append(shapeVec, element)
			}
			shapes[shapeIdx].vec = shapeVec
			shapes[shapeIdx].filledCells = filledCells
		}
	}
	regions := make([]Region, 0, len(lines[30:]))
	for _, line := range lines[30:] {
		parts := strings.Split(line, ":")
		dim := strings.Split(parts[0], "x")
		shapesChar := strings.Fields(parts[1])
		r, eR := strconv.Atoi(dim[0])
		c, eC := strconv.Atoi(dim[1])
		check(eR)
		check(eC)
		i0, eI0 := strconv.Atoi(shapesChar[0])
		i1, eI1 := strconv.Atoi(shapesChar[1])
		i2, eI2 := strconv.Atoi(shapesChar[2])
		i3, eI3 := strconv.Atoi(shapesChar[3])
		i4, eI4 := strconv.Atoi(shapesChar[4])
		i5, eI5 := strconv.Atoi(shapesChar[5])
		check(eI0)
		check(eI1)
		check(eI2)
		check(eI3)
		check(eI4)
		check(eI5)
		newRegion := Region{r, c, [6]int{i0, i1, i2, i3, i4, i5}}
		regions = append(regions, newRegion)
	}

	return shapes, regions
}

func packPrezzies(shapes []Shape, regions []Region) int {

	successCount := 0
	for regIdx, reg := range regions {
		regionArea := reg.r * reg.c
		filledCells := 0
		shapeCount := 0
		for i := range 6 {
			filledCells += reg.shapeCounts[i] * shapes[i].filledCells
			shapeCount += reg.shapeCounts[i]
		}
		if filledCells >= regionArea {
			// fmt.Println("Space too small!", regIdx)
			continue
		}
		if regionArea >= shapeCount*9 {
			// fmt.Println("It fits with lazy Packing!", regIdx)
			successCount += 1
			continue
		}
		fmt.Println("We actually need to try and pack it manually...", regIdx)
	}
	return successCount
}

func main() {
	startTime := time.Now()
	shapes, regions := readInput()
	part1 := packPrezzies(shapes, regions)
	fmt.Println("Day 12 Part 1:", part1)
	fmt.Println(time.Since(startTime))

}

// thoughts:
// 	with optimal packing if r*c < sum(filledCells) it is impossible
//  even with laziest possible packing if r*c> numshapes*9 it is possible

// Every one of the input cases is one of the easy examples!!!! HAH!
