package main

import (
	"container/list"
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

func readInput() ([]int, [][]int, [][][]int, [][]int) {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")
	lights := make([]int, len(lines))
	buttons := make([][][]int, len(lines))
	jolts := make([][]int, len(lines))
	for i, line := range lines {
		myFields := strings.Fields(line)
		lightChar := myFields[0]
		buttonsCharList := myFields[1 : len(myFields)-1]
		joltageChar := myFields[len(myFields)-1]
		lightNum := 0
		// get binary representation of lights i.e. .# = 10 = 2
		for j := 1; j < len(lightChar)-1; j++ {
			digit := 0
			index := j - 1
			if lightChar[j] == '#' {
				digit = 1
			}
			// fmt.Println(digit, index, power(2, index), string(lightChar[j]))
			lightNum += digit * power(2, index)
		}
		lights[i] = lightNum
		// fmt.Println(buttonsCharList)
		buttonLine := make([][]int, len(buttonsCharList))
		// looping over buttons
		for j, b := range buttonsCharList {
			b1 := b[1 : len(b)-1]
			buttonCharNums := strings.Split(b1, ",")
			buttonNums := []int{}
			// looping within each button
			for _, n := range buttonCharNums {
				nInt, e := strconv.Atoi(n)
				check(e)
				buttonNums = append(buttonNums, nInt)
			}
			buttonLine[j] = buttonNums
			// fmt.Println(buttonNums)
		}
		buttons[i] = buttonLine

		j := joltageChar[1 : len(joltageChar)-1]
		splitJolts := strings.Split(j, ",")
		joltList := make([]int, len(splitJolts))
		for j, jolt := range splitJolts {
			n, e := strconv.Atoi(jolt)
			check(e)
			joltList[j] = n
		}
		jolts[i] = joltList
	}
	buttonsInt := make([][]int, len(lines))
	for i := range buttons {
		buttonOptions := make([]int, len(buttons[i]))
		for j := range buttons[i] {
			buttonIntRep := 0
			for k := range buttons[i][j] {
				buttonIntRep += 1 * power(2, buttons[i][j][k])
			}
			buttonOptions[j] = buttonIntRep
		}
		buttonsInt[i] = buttonOptions
	}
	// fmt.Println(lights)
	// fmt.Println(buttons)
	// fmt.Println(buttonsInt)
	// fmt.Println(jolts)
	return lights, buttonsInt, buttons, jolts

}

func power(n, p int) int {
	out := 1
	for range p {
		out *= n
	}
	return out
}

func printBinary(n int) {
	fmt.Printf("%b\n", n)
}
func printBinaryList(s []int) {
	for _, n := range s {
		printBinary(n)
	}
}

type qEntry struct {
	n       int
	presses int
}

func mashButtons(target int, buttonOptions []int) int {
	myQ := list.New()
	myQ.PushBack(qEntry{0, 0})
	visited := make(map[int]struct{})

	for myQ.Len() > 0 {
		front := myQ.Front()
		entry := front.Value.(qEntry)
		myQ.Remove(front)
		for _, button := range buttonOptions {
			newNum := entry.n ^ button
			if newNum == target {
				return entry.presses + 1
			}
			_, found := visited[newNum]
			if !found {
				visited[newNum] = struct{}{}
				myQ.PushBack(qEntry{newNum, entry.presses + 1})
			}
		}
	}
	return 0 //not found!
}

func test() {
	n := 0
	target := 0b110
	n ^= 0b101
	n ^= 0b11
	fmt.Println("basic xor op")
	fmt.Println(n == target)
	fmt.Println("mashButtons")
	fmt.Println(mashButtons(target, []int{8, 10, 4, 12, 5, 3}))
}

func getMinButtonTotal(lightsInt []int, buttonsListList [][]int) int {
	sum := 0
	for i := range lightsInt {
		minButtons := mashButtons(lightsInt[i], buttonsListList[i])
		sum += minButtons
	}
	return sum
}

func transpose(matrix [][]int) [][]int {
	numRows, numCols := len(matrix), len(matrix[0])
	matrixT := make([][]int, numCols)
	for c := range numCols {
		matrixT[c] = make([]int, numRows)
		for r := range numRows {
			matrixT[c][r] = matrix[r][c]
		}
	}
	return matrixT

}

func makeJoltButtonMatrix(buttons [][]int, jolts []int) [][]int {
	numRows := len(buttons)
	numCols := len(jolts)

	matrix := make([][]int, numRows)
	for r := range numRows {
		matrix[r] = make([]int, numCols)
	}
	for r, b := range buttons {
		for _, c := range b {
			matrix[r][c] = 1
		}
	}
	matrixT := transpose(matrix)
	return matrixT
}

// func convertMatrixToFlatFloat(matrix [][]int) ([]float64, int, int) {
// 	numRows := len(matrix)
// 	numCols := len(matrix[0])

// 	flatFloatMatrix := make([]float64, 0, numRows*numCols)

// 	for r := range numRows {
// 		for c := range numCols {
// 			flatFloatMatrix = append(flatFloatMatrix, float64(matrix[r][c]))
// 		}
// 	}
// 	return flatFloatMatrix, numRows, numCols
// }
// func convertFlatFloatToIntSliceMatrix(flatFloatMatrix []float64, numRows, numCols int) [][]int {
// 	matrix := make([][]int, numRows)
// 	i := 0
// 	for r := range numRows {
// 		matrix[r] = make([]int, numCols)
// 		for c := range numCols {
// 			matrix[r][c] = int(flatFloatMatrix[i])
// 			i += 1
// 		}
// 	}
// 	return matrix
// }

func main() {
	startTime := time.Now()
	// lightsInt, buttonsIntList, _, _ := readInput()
	lightsInt, buttonsIntList, buttons, jolts := readInput()
	p1Time := time.Now()
	part1 := getMinButtonTotal(lightsInt, buttonsIntList)
	fmt.Println("Day 10, Part 1:", part1, time.Since(p1Time))
	fmt.Println(time.Since(startTime))

	// couldn't get z3 working with go so i did it in python
	// I think a solution to this is possible without it but i have to do gaussian elimination
	// on the matrix from makeJoltButtonMatrix, solve as many parameters as possible,
	// then bfs over the free parameters to find the min sum. TODO!
}
