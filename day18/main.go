package main

import (
	"adventofcode2023/lib"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

const NORTH int = 1
const EAST int = 2
const SOUTH int = 4
const WEST int = 8

type trench struct {
	dir   int
	dist  int
	color string
}

func expandSite(site [][]int, addRows int, addCols int) [][]int {
	if addCols > 0 {
		cols := len(site[0])
		for r := range site {
			site[r] = append(site[r], make([]int, addCols)...)
			for c := cols; c < len(site[r]); c++ {
				site[r][c] = 0
			}
		}
	} else if addCols < 0 {
		for r := range site {
			site[r] = append(make([]int, -addCols), site[r]...)
			for c := 0; c < -addCols; c++ {
				site[r][c] = 0
			}
		}
	}
	if addRows > 0 {
		cols := len(site[0])
		for r := 0; r < addRows; r++ {
			row := make([]int, cols)
			for c := 0; c < cols; c++ {
				row[c] = 0
			}
			site = append(site, row)
		}
	} else if addRows < 0 {
		cols := len(site[0])
		for r := 0; r < -addRows; r++ {
			row := make([]int, cols)
			for c := 0; c < cols; c++ {
				row[c] = 0
			}
			site = append([][]int{row}, site...)
		}
	}
	return site
}

func part1(stdin *bufio.Scanner) string {
	trenches := []trench{}
	for stdin.Scan() {
		line := stdin.Text()
		parts := strings.Split(line, " ")

		direction := parts[0]
		distance, _ := strconv.Atoi(parts[1])
		color := parts[2][2 : len(parts[2])-1]

		dir := 0
		switch direction {
		case "U":
			dir = NORTH
		case "R":
			dir = EAST
		case "D":
			dir = SOUTH
		case "L":
			dir = WEST
		}

		trenches = append(trenches, trench{dir, distance, color})
	}

	site := lib.Create2dArray(5, 5, 0)
	site[0][0] = EAST

	row, col := 0, 0
	for i := range trenches {
		// fmt.Println(i+1, ":", row, col, trenches[i].dir, trenches[i].dist)

		site[row][col] |= trenches[i].dir
		switch trenches[i].dir {
		case NORTH:
			if row-trenches[i].dist < 0 {
				site = expandSite(site, -trenches[i].dist, 0)
				row += trenches[i].dist
			}
			for r := 0; r < trenches[i].dist; r++ {
				row--
				site[row][col] |= trenches[i].dir
			}
		case EAST:
			if col+trenches[i].dist >= len(site[row]) {
				site = expandSite(site, 0, trenches[i].dist)
			}
			for c := 0; c < trenches[i].dist; c++ {
				col++
				site[row][col] |= trenches[i].dir
			}
		case SOUTH:
			if row+trenches[i].dist >= len(site) {
				site = expandSite(site, trenches[i].dist, 0)
			}
			for r := 0; r < trenches[i].dist; r++ {
				row++
				site[row][col] |= trenches[i].dir
			}
		case WEST:
			if col-trenches[i].dist < 0 {
				site = expandSite(site, 0, -trenches[i].dist)
				col += trenches[i].dist
			}
			for c := 0; c < trenches[i].dist; c++ {
				col--
				site[row][col] |= trenches[i].dir
			}
		}
	}

	// for i := range site {
	// 	for j := range site[i] {
	// 		fmt.Print(site[i][j])
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	inside := 0
	for r := range site {
		in := false
		need := 0
		for c, _ := range site[r] {
			if site[r][c] == NORTH || site[r][c] == SOUTH {
				in = !in
				inside++
			} else if site[r][c] == EAST || site[r][c] == WEST {
				inside++
				// no op
			} else if site[r][c] > 0 {
				inside++
				if need == 0 {
					if site[r][c]&NORTH > 0 {
						need = NORTH
					} else {
						need = SOUTH
					}
				} else if need > 0 && site[r][c]&need > 0 {
					in = !in
					need = 0
				} else {
					need = 0
				}
			} else if in {
				inside++
				site[r][c] = -2
			}
		}
	}

	// for i := range site {
	// 	for j := range site[i] {
	// 		if site[i][j] == 0 {
	// 			fmt.Print(".")
	// 		} else if site[i][j] < 0 {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print("@")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	return fmt.Sprint(inside)
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	return fmt.Sprint(result)
}
