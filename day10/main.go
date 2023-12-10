package main

import (
	"adventofcode2023/lib"
	"bufio"
	"fmt"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

const NORTH int = 1
const EAST int = 2
const WEST int = 4
const SOUTH int = 8

type node struct {
	row, col int
}

func walk(rstart int, cstart int, pipes [][]int, dist [][]int) {
	queue := []node{{rstart, cstart}}

	for len(queue) > 0 {

		n := queue[0]
		queue = queue[1:]
		row, col := n.row, n.col
		v := dist[row][col]

		if row > 0 && pipes[row][col]&NORTH > 0 && (dist[row-1][col] == -1) {
			dist[row-1][col] = v + 1
			queue = append(queue, node{row - 1, col})
		}
		if row < len(dist)-1 && pipes[row][col]&SOUTH > 0 && (dist[row+1][col] == -1) {
			dist[row+1][col] = v + 1
			queue = append(queue, node{row + 1, col})
		}
		if col > 0 && pipes[row][col]&WEST > 0 && (dist[row][col-1] == -1) {
			dist[row][col-1] = v + 1
			queue = append(queue, node{row, col - 1})
		}
		if col < len(dist[row])-1 && pipes[row][col]&EAST > 0 && (dist[row][col+1] == -1) {
			dist[row][col+1] = v + 1
			queue = append(queue, node{row, col + 1})
		}
	}
}

func part1(stdin *bufio.Scanner) string {
	pipes := [][]int{}

	rstart, cstart := 0, 0
	for stdin.Scan() {
		line := stdin.Text()

		row := []int{}
		for i, v := range line {
			switch v {
			case '|':
				row = append(row, NORTH|SOUTH)
				break
			case '-':
				row = append(row, EAST|WEST)
				break
			case 'L':
				row = append(row, NORTH|EAST)
				break
			case 'J':
				row = append(row, NORTH|WEST)
				break
			case '7':
				row = append(row, SOUTH|WEST)
				break
			case 'F':
				row = append(row, SOUTH|EAST)
				break
			case '.':
				row = append(row, 0)
				break
			case 'S':
				row = append(row, 0)
				rstart, cstart = len(pipes), i
				break
			default:
				panic("Unkonwn pipe")
			}
		}
		pipes = append(pipes, row)
	}

	dist := [][]int{}
	for i := 0; i < len(pipes); i++ {
		dist = append(dist, make([]int, len(pipes[i])))
		for j := 0; j < len(pipes[i]); j++ {
			dist[i][j] = -1
		}
	}
	dist[rstart][cstart] = 0

	// connect start
	if rstart > 0 && pipes[rstart-1][cstart]&SOUTH > 0 {
		pipes[rstart][cstart] |= NORTH
	}
	if rstart < len(pipes)-1 && pipes[rstart+1][cstart]&NORTH > 0 {
		pipes[rstart][cstart] |= SOUTH
	}
	if cstart > 0 && pipes[rstart][cstart-1]&EAST > 0 {
		pipes[rstart][cstart] |= WEST
	}
	if cstart < len(pipes[rstart])-1 && pipes[rstart][cstart+1]&WEST > 0 {
		pipes[rstart][cstart] |= EAST
	}

	walk(rstart, cstart, pipes, dist)

	// for _, r := range pipes {
	// 	fmt.Println(r)
	// }

	max := 0
	for _, r := range dist {
		max = lib.Max(max, lib.ArrayMax(r))
		// fmt.Println(r)
	}

	return fmt.Sprint(max)
}

func part2(stdin *bufio.Scanner) string {
	return "part2"
}
