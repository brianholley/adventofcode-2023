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

type beam struct {
	row, col int
	dir      int
}

func energizedCount(energized [][]int) int {
	sum := 0
	for i := range energized {
		for j := range energized[i] {
			if energized[i][j] > 0 {
				sum++
			}
		}
	}
	return sum
}

func runLasers(field [][]rune, startRow int, startCol int, startDir int) int {
	energized := make([][]int, len(field))
	for i := range energized {
		energized[i] = make([]int, len(field[i]))
		for j := range energized[i] {
			energized[i][j] = 0
		}
	}

	beams := []beam{{startRow, startCol, startDir}}
	energized[startRow][startCol] = startDir

	newBeams := []beam{}
	for len(beams) > 0 {
		for i := range beams {
			switch field[beams[i].row][beams[i].col] {
			case '/':
				if beams[i].dir == NORTH {
					beams[i].dir = EAST
					beams[i].col++
				} else if beams[i].dir == EAST {
					beams[i].dir = NORTH
					beams[i].row--
				} else if beams[i].dir == SOUTH {
					beams[i].dir = WEST
					beams[i].col--
				} else if beams[i].dir == WEST {
					beams[i].dir = SOUTH
					beams[i].row++
				}
			case '\\':
				if beams[i].dir == NORTH {
					beams[i].dir = WEST
					beams[i].col--
				} else if beams[i].dir == EAST {
					beams[i].dir = SOUTH
					beams[i].row++
				} else if beams[i].dir == SOUTH {
					beams[i].dir = EAST
					beams[i].col++
				} else if beams[i].dir == WEST {
					beams[i].dir = NORTH
					beams[i].row--
				}
			case '|':
				if beams[i].dir == NORTH {
					beams[i].row--
				} else if beams[i].dir == EAST {
					newBeams = append(newBeams, beam{beams[i].row + 1, beams[i].col, SOUTH})
					beams[i].dir = NORTH
					beams[i].row--
				} else if beams[i].dir == SOUTH {
					beams[i].row++
				} else if beams[i].dir == WEST {
					newBeams = append(newBeams, beam{beams[i].row + 1, beams[i].col, SOUTH})
					beams[i].dir = NORTH
					beams[i].row--
				}
			case '-':
				if beams[i].dir == NORTH {
					newBeams = append(newBeams, beam{beams[i].row, beams[i].col - 1, WEST})
					beams[i].dir = EAST
					beams[i].col++
				} else if beams[i].dir == EAST {
					beams[i].col++
				} else if beams[i].dir == SOUTH {
					newBeams = append(newBeams, beam{beams[i].row, beams[i].col - 1, WEST})
					beams[i].dir = EAST
					beams[i].col++
				} else if beams[i].dir == WEST {
					beams[i].col--
				}
			case '.':
				if beams[i].dir == NORTH {
					beams[i].row--
				} else if beams[i].dir == EAST {
					beams[i].col++
				} else if beams[i].dir == SOUTH {
					beams[i].row++
				} else if beams[i].dir == WEST {
					beams[i].col--
				}
			}
		}

		// Add any any newly created beams
		beams = append(beams, newBeams...)
		newBeams = []beam{}

		// Remove beams off the map and energize covered tiles
		for i := 0; i < len(beams); {
			if beams[i].row < 0 || beams[i].row >= len(field) || beams[i].col < 0 || beams[i].col >= len(field[0]) || energized[beams[i].row][beams[i].col]&beams[i].dir > 0 {
				if i < len(beams)-1 {
					beams = append(beams[:i], beams[i+1:]...)
				} else {
					beams = beams[:i]
				}
			} else {
				energized[beams[i].row][beams[i].col] |= beams[i].dir
				i++
			}
		}

		// fmt.Println(len(beams), energizedCount(energized))
	}

	sum := energizedCount(energized)
	return sum
}

func part1(stdin *bufio.Scanner) string {
	field := [][]rune{}

	for stdin.Scan() {
		line := stdin.Text()
		field = append(field, []rune(line))
	}

	sum := runLasers(field, 0, 0, EAST)

	return fmt.Sprint(sum)
}

func part2(stdin *bufio.Scanner) string {
	field := [][]rune{}

	for stdin.Scan() {
		line := stdin.Text()
		field = append(field, []rune(line))
	}

	max := 0
	for r := range field {
		energized := runLasers(field, r, 0, EAST)
		max = lib.Max(max, energized)

		energized = runLasers(field, r, len(field[0])-1, WEST)
		max = lib.Max(max, energized)
	}

	for c := range field[0] {
		energized := runLasers(field, 0, c, SOUTH)
		max = lib.Max(max, energized)

		energized = runLasers(field, len(field)-1, c, NORTH)
		max = lib.Max(max, energized)
	}

	return fmt.Sprint(max)
}
