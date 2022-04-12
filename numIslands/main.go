package main

func main() {
}

type visitParams struct {
	grid *[][]byte
	rows int
	cols int
	i    int
	j    int
}

func numIslands(grid [][]byte) int {

	rows := len(grid)
	cols := len(grid[0])

	num := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == []byte("1")[0] {
				visit(visitParams{i: i, j: j, rows: rows, cols: cols, grid: &grid})
				num++
			}
		}
	}

	return num
}

func visit(param visitParams) {
	i, j, rows, cols, grid := param.i, param.j, param.rows, param.cols, *(param.grid)
	if i < 0 || j < 0 || i >= rows || i >= cols || grid[i][j] != []byte("1")[0] {
		return
	}
	if grid[i][j] == []byte("1")[0] {
		grid[i][j] = 2
	}

	visit(visitParams{i: i + 1, j: j, rows: rows, cols: cols, grid: &grid})
	visit(visitParams{i: i, j: j + 1, rows: rows, cols: cols, grid: &grid})
	visit(visitParams{i: i - 1, j: j, rows: rows, cols: cols, grid: &grid})
	visit(visitParams{i: i, j: j - 1, rows: rows, cols: cols, grid: &grid})
}
