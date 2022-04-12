package main

func main() {

}

func setZeroes(matrix [][]int) [][]int {

	rows := make(map[int]bool)
	cols := make(map[int]bool)

	for col := range matrix {
		for row := range matrix[col] {
			if matrix[col][row] == 0 {
				// setRowToZeroes(&matrix, row)
				// setColToZeroes(&matrix, col)

				rows[row] = true
				cols[col] = true
			}
		}
	}

	// for col := range matrix {
	// 	for row := range matrix[col] {
	// 		if rows[row] {

	// 			setRowToZeroes(&matrix, row)
	// 		}
	// 		if cols[col] {
	// 			setColToZeroes(&matrix, col)
	// 		}
	// 	}
	// }

	for k, v := range rows {
		if v {
			setRowToZeroes(&matrix, k)
		}
	}

	for k, v := range cols {
		if v {
			setColToZeroes(&matrix, k)
		}
	}

	return matrix
}

func setRowToZeroes(matrix *[][]int, row int) {

	for i := range *matrix {
		(*matrix)[i][row] = 0
	}
}

func setColToZeroes(matrix *[][]int, col int) {
	for i := range (*matrix)[col] {
		(*matrix)[col][i] = 0
	}
}
