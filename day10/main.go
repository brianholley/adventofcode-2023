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
const SOUTH int = 4
const WEST int = 8

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

func parsePipes(stdin *bufio.Scanner) ([][]int, int, int) {
	pipes := [][]int{}

	rstart, cstart := 0, 0
	for stdin.Scan() {
		line := stdin.Text()

		row := []int{}
		for i, v := range line {
			switch v {
			case '|':
				row = append(row, NORTH|SOUTH)
			case '-':
				row = append(row, EAST|WEST)
			case 'L':
				row = append(row, NORTH|EAST)
			case 'J':
				row = append(row, NORTH|WEST)
			case '7':
				row = append(row, SOUTH|WEST)
			case 'F':
				row = append(row, SOUTH|EAST)
			case '.':
				row = append(row, 0)
			case 'S':
				row = append(row, 0)
				rstart, cstart = len(pipes), i
			default:
				fmt.Println(string(v))
				panic("Unknown pipe")
			}
		}
		pipes = append(pipes, row)
	}

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
	return pipes, rstart, cstart
}

func printPipes(pipes [][]int) {
	for r := range pipes {
		for c := range pipes[r] {
			switch pipes[r][c] {
			case NORTH | SOUTH:
				fmt.Print("|")
			case EAST | WEST:
				fmt.Print("-")
			case NORTH | EAST:
				fmt.Print("L")
			case NORTH | WEST:
				fmt.Print("J")
			case SOUTH | WEST:
				fmt.Print("7")
			case SOUTH | EAST:
				fmt.Print("F")
			case 0:
				fmt.Print(".")
			case -1:
				fmt.Print(".")
			case -2:
				fmt.Print("I")
			default:
				fmt.Print("?")
			}
		}
		fmt.Println()
	}
}

func part1(stdin *bufio.Scanner) string {
	pipes, rstart, cstart := parsePipes(stdin)

	dist := [][]int{}
	for i := 0; i < len(pipes); i++ {
		dist = append(dist, make([]int, len(pipes[i])))
		for j := 0; j < len(pipes[i]); j++ {
			dist[i][j] = -1
		}
	}
	dist[rstart][cstart] = 0

	walk(rstart, cstart, pipes, dist)

	max := 0
	for _, r := range dist {
		max = lib.Max(max, lib.ArrayMax(r))
		// fmt.Println(r)
	}

	return fmt.Sprint(max)
}

func markMainLoop(pipes [][]int, dist [][]int) [][]int {

	loop := [][]int{}
	for i := 0; i < len(pipes); i++ {
		loop = append(loop, make([]int, len(pipes[i])))
		for j := 0; j < len(pipes[i]); j++ {
			loop[i][j] = 0
		}
	}

	// rstart, cstart := 0, 0
	rend, cend := 0, 0
	max := 0
	for r, _ := range dist {
		for c, _ := range dist[r] {
			// if dist[r][c] == 0 {
			// 	rstart, cstart = r, c
			// }
			if dist[r][c] > max {
				rend, cend = r, c
				max = dist[r][c]
			}
		}
	}

	adj := []node{{rend, cend}}

	for len(adj) > 0 {
		n := adj[0]
		adj = adj[1:]
		loop[n.row][n.col] = pipes[n.row][n.col]
		if n.row > 0 && pipes[n.row][n.col]&NORTH > 0 && dist[n.row-1][n.col] < dist[n.row][n.col] && dist[n.row-1][n.col] >= 0 {
			adj = append(adj, node{n.row - 1, n.col})
		}
		if n.row < len(pipes)-1 && pipes[n.row][n.col]&SOUTH > 0 && dist[n.row+1][n.col] < dist[n.row][n.col] && dist[n.row+1][n.col] >= 0 {
			adj = append(adj, node{n.row + 1, n.col})
		}
		if n.col > 0 && pipes[n.row][n.col]&WEST > 0 && dist[n.row][n.col-1] < dist[n.row][n.col] && dist[n.row][n.col-1] >= 0 {
			adj = append(adj, node{n.row, n.col - 1})
		}
		if n.col < len(pipes[n.row])-1 && pipes[n.row][n.col]&EAST > 0 && dist[n.row][n.col+1] < dist[n.row][n.col] && dist[n.row][n.col+1] >= 0 {
			adj = append(adj, node{n.row, n.col + 1})
		}
	}

	return loop
}

func part2(stdin *bufio.Scanner) string {
	pipes, rstart, cstart := parsePipes(stdin)

	dist := [][]int{}
	for i := 0; i < len(pipes); i++ {
		dist = append(dist, make([]int, len(pipes[i])))
		for j := 0; j < len(pipes[i]); j++ {
			dist[i][j] = -1
		}
	}
	dist[rstart][cstart] = 0

	walk(rstart, cstart, pipes, dist)

	loop := markMainLoop(pipes, dist)

	inside := 0
	for r, _ := range loop {
		in := false
		need := 0
		for c, _ := range loop[r] {
			if loop[r][c]&(NORTH|SOUTH) == (NORTH | SOUTH) {
				in = !in
			} else if loop[r][c]&(EAST|WEST) == (EAST | WEST) {
				// no op
			} else if loop[r][c] > 0 {
				if need == 0 {
					if loop[r][c]&NORTH > 0 {
						need = SOUTH
					} else {
						need = NORTH
					}
				} else if need > 0 && loop[r][c]&need > 0 {
					in = !in
					need = 0
				} else {
					need = 0
				}
			} else if in {
				inside++
				loop[r][c] = -2
			}
		}
	}

	// printPipes(loop)

	return fmt.Sprint(inside)
}
