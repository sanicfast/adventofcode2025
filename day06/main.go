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
func readInput1() ([][]int, []string, [][]int) {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")
	outNums := make([][]int, 0, len(lines)-1)
	outOps := make([]string, 0)
	for i, line := range lines {
		elementsStr := strings.Fields(line)
		if i == len(lines)-1 {
			outOps = elementsStr
			continue
		}
		tmp := make([]int, 0, len(elementsStr))
		for _, numChar := range elementsStr {
			num, e := strconv.Atoi(numChar)
			check(e)
			tmp = append(tmp, num)
		}
		outNums = append(outNums, tmp)
	}
	rows := len(lines) - 1
	cols := len(lines[0])
	transposedBytes := make([][]byte, cols)

	for c := range cols {
		transposedBytes[c] = make([]byte, rows)
	}
	for r := range rows {
		for c := range cols {
			transposedBytes[c][r] = lines[r][c]
		}
	}

	transposedStrings := make([][]string, cols)
	for i := range transposedBytes {
		transposedStrings[i] = strings.Fields(string(transposedBytes[i]))
	}

	part2 := [][]int{}
	j := -1
	for i := range transposedStrings {
		if i == 0 {
			j += 1
			part2 = append(part2, []int{})
		} else if len(transposedStrings[i]) == 0 || i == 0 {
			j += 1
			part2 = append(part2, []int{})
			continue
		}
		val, e := strconv.Atoi(transposedStrings[i][0])
		check(e)
		part2[j] = append(part2[j], val)
	}

	return outNums, outOps, part2
}

func doHw(nums [][]int, ops []string) int {
	sum := 0
	for i, op := range ops {
		answer := 0
		if op == "*" {
			answer = 1
		}
		for j := range nums {
			// fmt.Println(i, j, nums[j][i], answer, op)
			switch op {
			case "*":
				answer *= nums[j][i]
			case "+":
				answer += nums[j][i]

			}
		}
		// fmt.Println(answer)
		sum += answer
	}
	return sum
}

func readInput2() [][]int {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")
	rows := len(lines) - 1
	cols := len(lines[0])
	transposedBytes := make([][]byte, cols)

	for c := range cols {
		transposedBytes[c] = make([]byte, rows)
	}
	for r := range rows {
		for c := range cols {
			transposedBytes[c][r] = lines[r][c]
		}
	}

	transposedStrings := make([][]string, cols)
	for i := range transposedBytes {
		transposedStrings[i] = strings.Fields(string(transposedBytes[i]))
	}

	outIntSlice := [][]int{}
	j := -1
	for i := range transposedStrings {
		if i == 0 {
			j += 1
			outIntSlice = append(outIntSlice, []int{})
		} else if len(transposedStrings[i]) == 0 || i == 0 {
			j += 1
			outIntSlice = append(outIntSlice, []int{})
			continue
		}
		val, e := strconv.Atoi(transposedStrings[i][0])
		check(e)
		outIntSlice[j] = append(outIntSlice[j], val)
	}

	return outIntSlice
}

func doHw2(nums [][]int, ops []string) int {
	sum := 0
	for i, op := range ops {
		answer := 0
		if op == "*" {
			answer = 1
		}
		for j := range nums[i] {
			switch op {
			case "*":
				answer *= nums[i][j]
			case "+":
				answer += nums[i][j]

			}
		}
		// fmt.Println(answer)
		sum += answer
	}
	return sum
}

func main() {
	startTime := time.Now()
	nums, ops, nums2 := readInput1()
	part1 := doHw(nums, ops)
	fmt.Println("Day 6, Part 1:", part1)
	// fmt.Println(nums2, ops)
	part2 := doHw2(nums2, ops)
	fmt.Println("Day 6, Part 2:", part2)
	fmt.Println(time.Since(startTime))
}
