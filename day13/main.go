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

type pattern struct {
	rows    []string
	rotated bool
}

func readPatterns(stdin *bufio.Scanner) []pattern {
	patterns := []pattern{}
	rows := []string{}
	for stdin.Scan() {
		line := stdin.Text()
		if len(line) == 0 {
			patterns = append(patterns, pattern{rows, false})
			rows = []string{}
		} else {
			rows = append(rows, line)
		}
	}
	patterns = append(patterns, pattern{rows, false})
	return patterns
}

func rotate(p pattern) pattern {
	rows := []string{}
	for i := range p.rows[0] {
		row := make([]rune, len(p.rows))
		for r := range p.rows {
			row[r] = rune(p.rows[r][i])
		}
		rows = append(rows, string(row))
	}
	return pattern{rows, true}
}

func part1(stdin *bufio.Scanner) string {

	patterns := readPatterns(stdin)

	sum := 0
	for i, p := range patterns {
		// Horizontal
		score := 0
		for l := 0; l < len(p.rows)-1; l++ {
			if p.rows[l] == p.rows[l+1] {
				up, down := l-1, l+2
				for up >= 0 && down < len(p.rows) {
					if p.rows[up] != p.rows[down] {
						break
					}
					up--
					down++
				}
				if up < 0 || down == len(p.rows) {
					score += 100 * (l + 1)
					fmt.Println("Horizontal match at line", (l + 1))
				}
			}
		}

		// Vertical
		if score == 0 {
			rotated := rotate(p)
			for l := 0; l < len(rotated.rows)-1; l++ {
				if rotated.rows[l] == rotated.rows[l+1] {
					up, down := l-1, l+2
					for up >= 0 && down < len(rotated.rows) {
						if rotated.rows[up] != rotated.rows[down] {
							break
						}
						up--
						down++
					}
					if up < 0 || down == len(rotated.rows) {
						score += (l + 1)
						fmt.Println("Vertical match at column", (l + 1))
					}
				}
			}
		}

		fmt.Println("Score for", (i + 1), ":", score)

		sum += score
	}

	return fmt.Sprint(sum)
}

func part2(stdin *bufio.Scanner) string {
	return "part2"
}
