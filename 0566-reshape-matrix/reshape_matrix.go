package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	matrixRows := len(mat)
	matrixCols := len(mat[0])

	if matrixRows*matrixCols != r*c {
		return mat
	}

	col := 0
	row := 0
	resultMatrix := make([][]int, r)
	for i := 0; i < r; i++ {
		resultMatrix[i] = make([]int, c)
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if col == c {
				row++
				col = 0
			}

			resultMatrix[row][col] = mat[i][j]
			col++
		}

	}

	return resultMatrix
}

func main() {
	matrix_1 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	result := matrixReshape(matrix_1, 2, 6)
	fmt.Println(result)
}
