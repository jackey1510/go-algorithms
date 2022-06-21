package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	res := spiralOrder(matrix)
	fmt.Println(res)
}

type Direction string

const (
	Left  Direction = "l"
	Right           = "r"
	Up              = "u"
	Down            = "d"
)

type Node struct {
	x         int
	y         int
	direction Direction
}

type Visited [][]bool

func (v Visited) get(y, x int) bool {
	width := len(v[0])
	height := len(v)

	if x < 0 || x >= width {
		return true
	}
	if y < 0 || y >= height {
		return true
	}

	return v[y][x]
}

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	width := len(matrix[0])
	height := len(matrix)
	visited := make(Visited, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}
	queue := []Node{{x: 0, y: 0, direction: Right}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.x >= width {
			queue = append(queue, Node{x: node.x - 1, y: node.y + 1, direction: Down})
			continue
		}
		if node.x < 0 {
			queue = append(queue, Node{x: 0, y: node.y - 1, direction: Up})
			continue
		}

		if node.y >= height {
			queue = append(queue, Node{x: node.x - 1, y: node.y - 1, direction: Left})
			continue
		}

		if node.y < 0 {
			continue
		}

		if visited[node.y][node.x] {

			if node.direction == Left && !visited.get(node.y-1, node.x+1) {
				queue = append(queue, Node{x: node.x + 1, y: node.y - 1, direction: Up})
			}
			if node.direction == Right && !visited.get(node.y+1, node.x-1) {
				queue = append(queue, Node{x: node.x - 1, y: node.y + 1, direction: Down})
			}

			if node.direction == Up && !visited.get(node.y+1, node.x+1) {
				queue = append(queue, Node{x: node.x + 1, y: node.y + 1, direction: Right})
			}

			if node.direction == Down && !visited.get(node.y-1, node.x-1) {
				queue = append(queue, Node{x: node.x - 1, y: node.y - 1, direction: Left})
			}
			continue
		}

		visited[node.y][node.x] = true
		result = append(result, matrix[node.y][node.x])

		switch node.direction {
		case Left:
			{
				queue = append(queue, Node{x: node.x - 1, y: node.y, direction: Left})
				break
			}
		case Right:
			{
				queue = append(queue, Node{x: node.x + 1, y: node.y, direction: Right})
				break
			}
		case Up:
			{
				queue = append(queue, Node{x: node.x, y: node.y - 1, direction: Up})
				break
			}
		case Down:
			{
				queue = append(queue, Node{x: node.x, y: node.y + 1, direction: Down})
				break
			}
		}

	}

	return result
}
