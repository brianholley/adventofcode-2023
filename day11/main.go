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

type galaxy struct {
	row, col int
}

func rowIsEmpty(row int, space [][]byte) bool {
	for i := range space[row] {
		if space[row][i] != '.' {
			return false
		}
	}
	return true
}

func colIsEmpty(col int, space [][]byte) bool {
	for i := range space {
		if space[i][col] != '.' {
			return false
		}
	}
	return true
}

func part1(stdin *bufio.Scanner) string {
	space := [][]byte{}
	for stdin.Scan() {
		row := stdin.Text()
		space = append(space, []byte(row))
	}

	for r := 0; r < len(space); r++ {
		if rowIsEmpty(r, space) {
			space = append(append(space[:r], space[r]), space[r:]...)
			r++
		}
	}

	for c := 0; c < len(space[0]); c++ {
		if colIsEmpty(c, space) {
			for r := range space {
				space[r] = append(append(space[r][:c], '.'), space[r][c:]...)
			}
			c++
		}
	}

	galaxies := []galaxy{}
	for r := range space {
		for c := range space[r] {
			if space[r][c] == '#' {
				galaxies = append(galaxies, galaxy{r, c})
			}
		}
	}

	sum := 0
	for g1 := range galaxies {
		for g2 := g1 + 1; g2 < len(galaxies); g2++ {
			d := lib.Abs(galaxies[g2].row-galaxies[g1].row) + lib.Abs(galaxies[g2].col-galaxies[g1].col)
			sum += d
		}

	}

	return fmt.Sprint(sum)
}

func part2(stdin *bufio.Scanner) string {
	expansionFactor := 1000000

	space := [][]byte{}
	for stdin.Scan() {
		row := stdin.Text()
		space = append(space, []byte(row))
	}

	galaxies := []galaxy{}
	for r := range space {
		for c := range space[r] {
			if space[r][c] == '#' {
				galaxies = append(galaxies, galaxy{r, c})
			}
		}
	}

	for r := len(space) - 1; r >= 0; r-- {
		if rowIsEmpty(r, space) {
			for g := range galaxies {
				if galaxies[g].row > r {
					// fmt.Println("Moving row", galaxies[g].row, "to", galaxies[g].row+expansionFactor-1)
					galaxies[g].row += expansionFactor - 1
				}
			}
		}
	}

	for c := len(space[0]) - 1; c >= 0; c-- {
		if colIsEmpty(c, space) {
			for g := range galaxies {
				if galaxies[g].col > c {
					// fmt.Println("Moving col", galaxies[g].col, "to", galaxies[g].col+expansionFactor-1)
					galaxies[g].col += expansionFactor - 1
				}
			}
		}
	}

	sum := 0
	for g1 := range galaxies {
		// fmt.Println(galaxies[g1])
		for g2 := g1 + 1; g2 < len(galaxies); g2++ {
			d := lib.Abs(galaxies[g2].row-galaxies[g1].row) + lib.Abs(galaxies[g2].col-galaxies[g1].col)
			sum += d
		}

	}

	return fmt.Sprint(sum)
}
