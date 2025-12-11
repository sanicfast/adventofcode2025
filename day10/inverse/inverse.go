package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// --- 1. Define the Input Matrix (A) ---
	// Define a 2x2 matrix data (flat, column-major order is common,
	// but here we use row-major for simplicity: 1, 2, 3, 4)
	data := []float64{
		4, 7, // Row 1
		2, 6, // Row 2
	}
	A := mat.NewDense(2, 2, data)
	fmt.Println("Original Matrix (A):")
	// Print the matrix nicely
	fa := mat.Formatted(A, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("%v\n\n", fa)

	// --- 2. Create the Destination Matrix (A_inv) ---
	// The result will be stored in this new matrix.
	var A_inv mat.Dense

	// --- 3. Compute the Inverse ---
	// Inverse(A) computes the inverse of A and stores it in A_inv.
	err := A_inv.Inverse(A)
	if err != nil {
		// Handle the case where the matrix is singular (non-invertible)
		log.Fatalf("Error computing inverse: %v", err)
	}

	fmt.Println("Inverse Matrix (A⁻¹):")
	// Print the inverse matrix
	fainv := mat.Formatted(&A_inv, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("%v\n\n", fainv)

	// --- 4. Verification (Optional but Recommended) ---
	// Multiply the original matrix A by its inverse A_inv.
	// The result should be the Identity Matrix (I).
	var I mat.Dense
	I.Mul(A, &A_inv) // I = A * A_inv

	fmt.Println("Verification (A * A⁻¹): Should be Identity Matrix (I)")
	fI := mat.Formatted(&I, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("%v\n", fI)
}

// intMatrix := makeJoltButtonMatrix(buttons, jolts)
// flatMat, numRows, numCols := convertMatrixToFlatFloat(intMatrix)

// joltsFloat := make([]float64, len(jolts))
// for i, j := range jolts {
// 	joltsFloat[i] = float64(j)
// }
// A := mat.NewDense(numRows, numCols, flatMat)
// B := mat.NewDense(len(jolts), 1, joltsFloat)
// var A_inv mat.Dense
// e := A_inv.Inverse(A)
// check(e)

// var X mat.Dense
// X.Mul(&A_inv, B) // I = A * A_inv
// fI := mat.Formatted(&X, mat.Prefix("    "), mat.Squeeze())
// fmt.Printf("%v\n", fI)
