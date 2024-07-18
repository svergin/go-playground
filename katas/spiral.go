package katas

import (
	"strings"
)

// Kata: Make a spiral <3 kyu>
// https://www.codewars.com/kata/534e01fbbb17187c7e0000c6

func MakeSpiral(size int) string {
	matrix := make([][]string, size)

	for idx := range matrix {
		matrix[idx] = make([]string, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = "."
		}
	}

	// return first(matrix, size)
	return second(matrix, size)
}

func second(m [][]string, size int) string {

	return printMatrix(m)
}

func first(m [][]string, size int) string {
	maxIter := size - 3

	for curIter := 0; curIter < maxIter; curIter++ {

		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if curIter == 0 {
					if i == 0 || i == size-1 {
						m[i][j] = "0"
					}
					if j == size-1 && (i >= 1 && i <= size-2) {
						m[i][j] = "0"
					}
					if j == 0 && i == size-2 {
						m[i][j] = "0"
					}
				}
				// if curIter == 1 {
				// 	if matrix[i][j] == "." && matrix[i-1][j] == "." && matrix[i+1][j] == "0" {
				// 		matrix[i][j] = "0"
				// 	}
				// }
				//matrix[i][j] = "."
			}
		}
		// fmt.Printf("Iteration %v\n", curIter)
	}
	return printMatrix(m)
}

func printMatrix(m [][]string) string {
	sb := strings.Builder{}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			sb.WriteString(m[i][j])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
