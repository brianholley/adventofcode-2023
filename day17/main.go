package main

import (
	"adventofcode2023/lib"
	"bufio"
	"fmt"
	"math"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

const NORTH int = 1
const EAST int = 2
const SOUTH int = 4
const WEST int = 8

type pos struct {
	row int
	col int
}
type route struct {
	row     int
	col     int
	loss    int
	dir     int
	path    []int
	history []pos
}

type lossmap struct {
	state map[int][]int
}

func distance(path []int, dir int) int {
	i := 0
	for p := len(path) - 1; p >= 0 && path[p] == dir; p-- {
		i++
	}
	return i
}

func addPath(path []int, dir int) []int {
	return append(append([]int{}, path...), dir)
}

func addHistory(history []pos, row int, col int) []pos {
	return append(append([]pos{}, history...), pos{row, col})
}

func isInHistory(row int, col int, history []pos) bool {
	for i := range history {
		if row == history[i].row && col == history[i].col {
			return true
		}
	}
	return false
}

func printPath(path []int) {
	for _, dir := range path {
		switch dir {
		case NORTH:
			fmt.Print("^")
		case SOUTH:
			fmt.Print("V")
		case WEST:
			fmt.Print("<")
		case EAST:
			fmt.Print(">")
		}
	}
	fmt.Println()
}

func arrayNew(length int, value int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = value
	}
	return arr
}

// Note: this algorithm for part 1 feels VERY overcomplicated...
func traverse(heatmap [][]int, minDist int, maxDist int) (int, []int) {

	rows := len(heatmap)
	cols := len(heatmap[0])

	heatloss := make([][]lossmap, rows)
	for i := range heatloss {
		heatloss[i] = make([]lossmap, cols)
		for j := range heatloss[i] {
			heatloss[i][j] = lossmap{map[int][]int{
				NORTH: arrayNew(maxDist+1, math.MaxInt),
				EAST:  arrayNew(maxDist+1, math.MaxInt),
				SOUTH: arrayNew(maxDist+1, math.MaxInt),
				WEST:  arrayNew(maxDist+1, math.MaxInt),
			}}
		}
	}

	routes := []route{}
	routes = append(routes, route{0, 0, 0, 0, []int{}, []pos{{0, 0}}})

	// Build naive case
	naivePath := []int{}
	naiveLoss := 0
	for i, j := 0, 0; i < len(heatmap)-1 || j < len(heatmap[0])-1; {
		colStep := lib.Max(lib.Min(maxDist, len(heatmap[0])-minDist-1), minDist)
		for r := 0; r < colStep && j < len(heatmap[0])-1; r++ {
			j++
			naivePath = append(naivePath, EAST)
			naiveLoss += heatmap[i][j]
		}

		rowStep := lib.Max(lib.Min(maxDist, len(heatmap)-minDist-1), minDist)
		for r := 0; r < rowStep && i < len(heatmap)-1; r++ {
			i++
			naivePath = append(naivePath, SOUTH)
			naiveLoss += heatmap[i][j]
		}
	}
	bestPath := naivePath
	minLoss := naiveLoss

	for len(routes) > 0 {
		curr := routes[0]
		routes = routes[1:]

		if curr.row == rows-1 && curr.col == cols-1 {
			if curr.loss <= minLoss {
				bestPath = curr.path
				minLoss = curr.loss
			}
			continue
		}

		// Check for whether there's already a better path known
		if curr.loss > minLoss {
			continue
		}
		if curr.dir != 0 && curr.loss > heatloss[curr.row][curr.col].state[curr.dir][distance(curr.path, curr.dir)-1] {
			continue
		}

		// Fan out
		if curr.row < len(heatmap)-minDist && curr.dir != NORTH && curr.dir != SOUTH {
			loss, path, history := curr.loss, curr.path, curr.history
			for i := 1; i <= maxDist; i++ {
				row := curr.row + i
				if row >= len(heatmap) {
					break
				}
				loss += heatmap[row][curr.col]

				if loss < heatloss[row][curr.col].state[SOUTH][i] && !isInHistory(row, curr.col, curr.history) {
					heatloss[row][curr.col].state[SOUTH][i] = loss
					if i >= minDist {
						path = addPath(path, SOUTH)
						history = addHistory(history, row, curr.col)
						r := route{row, curr.col, loss, SOUTH, path, history}
						routes = append(routes, r)
					}
				}
			}
		}
		if curr.col < len(heatmap[0])-minDist && curr.dir != WEST && curr.dir != EAST {
			loss, path, history := curr.loss, curr.path, curr.history
			for i := 1; i <= maxDist; i++ {
				col := curr.col + i
				if col >= len(heatmap[0]) {
					break
				}
				loss += heatmap[curr.row][col]

				if loss < heatloss[curr.row][col].state[EAST][i] && !isInHistory(curr.row, col, curr.history) {
					heatloss[curr.row][col].state[EAST][i] = loss
					if i >= minDist {
						path = addPath(path, EAST)
						history = addHistory(history, curr.row, col)
						r := route{curr.row, col, loss, EAST, path, history}
						routes = append(routes, r)
					}
				}
			}
		}
		if curr.row >= minDist && curr.dir != NORTH && curr.dir != SOUTH {
			loss, path, history := curr.loss, curr.path, curr.history
			for i := 1; i <= maxDist; i++ {
				row := curr.row - i
				if row < 0 {
					break
				}
				loss += heatmap[row][curr.col]

				if loss < heatloss[row][curr.col].state[NORTH][i] && !isInHistory(row, curr.col, curr.history) {
					heatloss[row][curr.col].state[NORTH][i] = loss
					if i >= minDist {
						path = addPath(path, NORTH)
						history = addHistory(history, row, curr.col)
						r := route{row, curr.col, loss, NORTH, path, history}
						routes = append(routes, r)
					}
				}
			}
		}
		if curr.col >= minDist && curr.dir != WEST && curr.dir != EAST {
			loss, path, history := curr.loss, curr.path, curr.history
			for i := 1; i <= maxDist; i++ {
				col := curr.col - i
				if col < 0 {
					break
				}
				loss += heatmap[curr.row][col]

				if loss < heatloss[curr.row][col].state[WEST][i] && !isInHistory(curr.row, col, curr.history) {
					heatloss[curr.row][col].state[WEST][i] = loss
					if i >= minDist {
						path = addPath(path, WEST)
						history = addHistory(history, curr.row, col)
						r := route{curr.row, col, loss, WEST, path, history}
						routes = append(routes, r)
					}
				}
			}
		}
	}

	return minLoss, bestPath
}

func part1(stdin *bufio.Scanner) string {
	heatmap := lib.Read2dArray(stdin, false)
	loss, path := traverse(heatmap, 1, 3)

	printPath(path)
	return fmt.Sprint(loss)
}

func part2(stdin *bufio.Scanner) string {
	heatmap := lib.Read2dArray(stdin, false)
	loss, path := traverse(heatmap, 4, 10)

	printPath(path)
	return fmt.Sprint(loss)
}
