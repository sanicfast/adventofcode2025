package main

import (
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

func readInput() [][]int {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")
	manifold := make([][]int, len(lines))
	for r, line := range lines {
		manifold[r] = make([]int, len(line))
		for c, char := range line {
			switch char {
			case 'S':
				manifold[r][c] = 1
			case '^':
				manifold[r][c] = 2
			}
		}
	}
	return manifold
}

func printManifold(manifold [][]int) {
	for _, row := range manifold {
		fmt.Println(row)
	}
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func intSum(s []int) int {
	sum := 0
	for _, num := range s {
		sum += num
	}
	return sum
}

func inBounds(r, c, dimr, dimc int) bool {
	if r >= 0 && r < dimr && c >= 0 && c < dimc {
		return true
	}
	return false
}

func simulateTachyons(manifold [][]int) int {
	dimr, dimc := len(manifold), len(manifold[0])
	splitCnt := 0
	for r := range manifold {
		if r == len(manifold)-1 {
			break
		}
		for c := range manifold[0] {
			if manifold[r][c] != 1 {
				continue
			}
			manifold[r+1][c] = maxInt(manifold[r+1][c], 1)
			if manifold[r+1][c] == 2 {
				splitCnt += 1
				if inBounds(r+1, c+1, dimr, dimc) {
					manifold[r+1][c+1] = maxInt(manifold[r+1][c+1], 1)
				}
				if inBounds(r+1, c-1, dimr, dimc) {
					manifold[r+1][c-1] = maxInt(manifold[r+1][c-1], 1)
				}
			}
		}
	}
	// printManifold(manifold)
	return splitCnt
}

// type Coord struct {
// 	r int
// 	c int
// }
// type Stack []Coord

// func (stack *Stack) Push(c Coord) {
// 	*stack = append(*stack, c)
// }

// func (stack *Stack) Pop() Coord {
// 	n := len(*stack)
// 	out := (*stack)[n-1]
// 	*stack = (*stack)[:n-1]
// 	return out
// }

// // this works on example data but wayyyyy too slow for the input.txt
// func simQuantumStack(manifold [][]int) int {
// 	theStack := make(Stack, 0, 100)
// 	for i := range manifold[0] {
// 		if manifold[0][i] == 1 {
// 			theStack.Push(Coord{0, i})
// 			break
// 		}
// 	}
// 	rows, cols := len(manifold), len(manifold[0])
// 	timelineCnt := 0
// 	for len(theStack) > 0 {
// 		position := theStack.Pop()
// 		fmt.Println(position)
// 		if position.r == len(manifold)-1 { // we reached the bottom!
// 			timelineCnt += 1
// 			continue
// 		}
// 		if manifold[position.r+1][position.c] != 2 {
// 			theStack.Push(Coord{position.r + 1, position.c})
// 		} else {
// 			left, right := Coord{position.r + 1, position.c - 1}, Coord{position.r + 1, position.c + 1}
// 			if inBounds(left.r, left.c, rows, cols) {
// 				theStack.Push(left)
// 			}
// 			if inBounds(right.r, right.c, rows, cols) {
// 				theStack.Push(right)
// 			}
// 		}
// 	}
// 	return timelineCnt
// }

func countTimelines(manifold [][]int) int {

	rows, cols := len(manifold), len(manifold[0])

	scratch := make([][]int, rows)
	for r := range manifold {
		scratch[r] = make([]int, cols)
	}
	scratch[0] = manifold[0]

	for r := 1; r < rows; r++ {
		for c := range cols {
			above := scratch[r-1][c]
			if manifold[r][c] == 2 {
				scratch[r][c+1] += above
				scratch[r][c-1] += above
			} else {
				scratch[r][c] += above
			}
		}
	}
	timelineCnt := intSum(scratch[rows-1])
	return timelineCnt
}

func main() {
	startTime := time.Now()

	manifold := readInput()
	part1 := simulateTachyons(manifold)
	fmt.Println("Day 7, Part 1:", part1)
	// printManifold(manifold)
	// part2 := quantumSimulateTachyons(manifold)
	// fmt.Println("Day 7, Part 2:", part2)
	part2 := countTimelines(manifold)
	fmt.Println("Day 7, Part 2:", part2)

	fmt.Println(time.Since(startTime))

}
