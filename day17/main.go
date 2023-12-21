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

// Note: this algorithm for part 1 feels VERY overcomplicated...
func traverse(heatmap [][]int) (int, []int) {

	rows := len(heatmap)
	cols := len(heatmap[0])

	heatloss := make([][]lossmap, rows)
	for i := range heatloss {
		heatloss[i] = make([]lossmap, cols)
		for j := range heatloss[i] {
			heatloss[i][j] = lossmap{map[int][]int{
				NORTH: {math.MaxInt, math.MaxInt, math.MaxInt},
				EAST:  {math.MaxInt, math.MaxInt, math.MaxInt},
				SOUTH: {math.MaxInt, math.MaxInt, math.MaxInt},
				WEST:  {math.MaxInt, math.MaxInt, math.MaxInt},
			}}
		}
	}

	routes := []route{}
	routes = append(routes, route{0, 0, 0, 0, []int{}, []pos{{0, 0}}})

	// Build naive case
	naivePath := []int{}
	naiveLoss := 0
	for i, j := 0, 0; i < len(heatmap)-1 && j < len(heatmap[0])-1; {
		naivePath = append(naivePath, EAST)
		j++
		naiveLoss += heatmap[i][j]

		naivePath = append(naivePath, SOUTH)
		i++
		naiveLoss += heatmap[i][j]
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
		if curr.row < len(heatmap)-1 &&
			curr.dir != NORTH &&
			(curr.dir != SOUTH || distance(curr.path, SOUTH) < 3) {

			loss := curr.loss + heatmap[curr.row+1][curr.col]
			if loss < heatloss[curr.row+1][curr.col].state[SOUTH][distance(curr.path, SOUTH)] && !isInHistory(curr.row+1, curr.col, curr.history) {
				heatloss[curr.row+1][curr.col].state[SOUTH][distance(curr.path, SOUTH)] = loss
				r := route{curr.row + 1, curr.col, loss, SOUTH, addPath(curr.path, SOUTH), addHistory(curr.history, curr.row+1, curr.col)}
				routes = append(routes, r)
			}
		}
		if curr.col < len(heatmap[0])-1 &&
			curr.dir != WEST &&
			(curr.dir != EAST || distance(curr.path, EAST) < 3) {

			loss := curr.loss + heatmap[curr.row][curr.col+1]
			if loss < heatloss[curr.row][curr.col+1].state[EAST][distance(curr.path, EAST)] && !isInHistory(curr.row, curr.col+1, curr.history) {
				heatloss[curr.row][curr.col+1].state[EAST][distance(curr.path, EAST)] = loss
				r := route{curr.row, curr.col + 1, loss, EAST, addPath(curr.path, EAST), addHistory(curr.history, curr.row, curr.col+1)}
				routes = append(routes, r)
			}
		}
		if curr.col > 0 &&
			curr.dir != EAST &&
			(curr.dir != WEST || distance(curr.path, WEST) < 3) {

			loss := curr.loss + heatmap[curr.row][curr.col-1]
			if loss < heatloss[curr.row][curr.col-1].state[WEST][distance(curr.path, WEST)] && !isInHistory(curr.row, curr.col-1, curr.history) {
				heatloss[curr.row][curr.col-1].state[WEST][distance(curr.path, WEST)] = loss
				r := route{curr.row, curr.col - 1, loss, WEST, addPath(curr.path, WEST), addHistory(curr.history, curr.row, curr.col-1)}
				routes = append(routes, r)
			}
		}
		if curr.row > 0 &&
			curr.dir != SOUTH &&
			(curr.dir != NORTH || distance(curr.path, NORTH) < 3) {

			loss := curr.loss + heatmap[curr.row-1][curr.col]
			if loss < heatloss[curr.row-1][curr.col].state[NORTH][distance(curr.path, NORTH)] && !isInHistory(curr.row-1, curr.col, curr.history) {
				heatloss[curr.row-1][curr.col].state[NORTH][distance(curr.path, NORTH)] = loss
				r := route{curr.row - 1, curr.col, loss, NORTH, addPath(curr.path, NORTH), addHistory(curr.history, curr.row-1, curr.col)}
				routes = append(routes, r)
			}
		}
	}

	return minLoss, bestPath
}

func part1(stdin *bufio.Scanner) string {
	heatmap := lib.Read2dArray(stdin, false)
	loss, path := traverse(heatmap)

	printPath(path)
	return fmt.Sprint(loss)
}

func part2(stdin *bufio.Scanner) string {
	return "part2"
}
