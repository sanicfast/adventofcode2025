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
func readInput() []int {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")
	combo := make([]int, 0, len(lines))
	for _, line := range lines {
		numChar := line[1:]
		num, e := strconv.Atoi(numChar)
		check(e)
		if line[0] == 'L' {
			num = -num
		}
		combo = append(combo, num)
	}
	return combo
}

func moduloPython(a, b int) int {
	residual := a % b
	if (residual < 0 && b > 0) || (residual > 0 && b < 0) {
		return residual + b
	}
	return residual
}

func spin(combo []int) (int, int) {
	part1 := 0
	part2 := 0
	// part2_2 := 0

	dial := 50

	for _, turn := range combo {
		increment := 0
		dial_init := dial
		stop := dial + turn
		dial = moduloPython(stop, 100)

		// part 1
		if dial == 0 {
			part1 += 1
		}

		// part 2
		if stop == 0 {
			increment += 1
		} else if stop >= 100 {
			increment += stop / 100 //integer division
		} else if stop < 0 {
			increment += -stop/100 + 1
			if dial_init == 0 { //don't double count if we started on a zero
				increment -= 1
			}
		}
		part2 += increment
		// fmt.Println("turn", turn, "finish", dial, "actual", stop, increment)

		// // slow way!
		// increment2 := 0
		// step := 1
		// if dial_init > stop {
		// 	step = -1
		// }
		// for i := dial_init + step; i != stop+step; i += step {
		// 	if i%100 == 0 {
		// 		increment2 += 1
		// 	}
		// }
		// part2_2 += increment2
		// if increment != increment2 {
		// 	fmt.Println("start", dial_init, "turn", turn, "finish", dial, "actual", stop, increment, increment2)
		// }

	}

	return part1, part2
}

func main() {
	startTime := time.Now()
	combo := readInput()
	part1, part2 := spin(combo)
	fmt.Println("Day 01, Part 1", part1)
	fmt.Println("Day 01, Part 2", part2)
	fmt.Println(time.Since(startTime))
}
