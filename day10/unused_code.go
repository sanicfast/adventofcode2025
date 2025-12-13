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

func convertMatrixToFlatFloat(matrix [][]int) ([]float64, int, int) {
	numRows := len(matrix)
	numCols := len(matrix[0])

	flatFloatMatrix := make([]float64, 0, numRows*numCols)

	for r := range numRows {
		for c := range numCols {
			flatFloatMatrix = append(flatFloatMatrix, float64(matrix[r][c]))
		}
	}
	return flatFloatMatrix, numRows, numCols
}
func convertFlatFloatToIntSliceMatrix(flatFloatMatrix []float64, numRows, numCols int) [][]int {
	matrix := make([][]int, numRows)
	i := 0
	for r := range numRows {
		matrix[r] = make([]int, numCols)
		for c := range numCols {
			matrix[r][c] = int(flatFloatMatrix[i])
			i += 1
		}
	}
	return matrix
}

// p1Time := time.Now()
// part1 := getMinButtonTotalOld(lightsInt, buttonsIntList)
// fmt.Println("Day 10, Part 1:", part1, time.Since(p1Time))
