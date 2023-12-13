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

func isPatternMirroredAroundRow(p pattern, l int) bool {
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
			return true
		}
	}
	return false
}

func part1(stdin *bufio.Scanner) string {

	patterns := readPatterns(stdin)

	sum := 0
	for _, p := range patterns {
		score := 0

		// Horizontal
		for l := 0; l < len(p.rows)-1; l++ {
			if isPatternMirroredAroundRow(p, l) {
				score += 100 * (l + 1)
				// fmt.Println("Horizontal match at line", (l + 1))
			}
		}

		// Vertical
		if score == 0 {
			rotated := rotate(p)
			for l := 0; l < len(rotated.rows)-1; l++ {
				if isPatternMirroredAroundRow(rotated, l) {
					score += (l + 1)
					// fmt.Println("Vertical match at column", (l + 1))
				}
			}
		}

		// fmt.Println("Score for", (i + 1), ":", score)

		sum += score
	}

	return fmt.Sprint(sum)
}

func almostEqual(l1 string, l2 string) bool {
	allowedErrors := 1
	for i := range l1 {
		if l1[i] != l2[i] {
			if allowedErrors > 0 {
				allowedErrors--
			} else {
				return false
			}
		}
	}
	return allowedErrors == 0
}

func findSmudgeCandidates(p pattern) []int {
	smudges := []int{}
	for l := range p.rows {
		for check := l + 1; check < len(p.rows); check += 2 {
			if almostEqual(p.rows[l], p.rows[check]) {
				// fmt.Println("Pattern", i, ", Rows almost equal", l, "and", check)
				smudges = append(smudges, l, check)
			}
		}
	}
	return smudges
}

func isPatternMirroredAroundRowWithSmudge(p pattern, l int) bool {
	usedSmudge := false

	initialRowMatches := p.rows[l] == p.rows[l+1]
	if !initialRowMatches && almostEqual(p.rows[l], p.rows[l+1]) {
		initialRowMatches = true
		usedSmudge = true
	}

	if initialRowMatches {
		up, down := l-1, l+2
		for up >= 0 && down < len(p.rows) {
			rowsMatch := p.rows[up] == p.rows[down]
			if !rowsMatch && !usedSmudge && almostEqual(p.rows[up], p.rows[down]) {
				rowsMatch = true
				usedSmudge = true
			}

			if !rowsMatch {
				break
			}
			up--
			down++
		}
		if up < 0 || down == len(p.rows) {
			return usedSmudge
		}
	}
	return false
}

func part2(stdin *bufio.Scanner) string {

	patterns := readPatterns(stdin)

	sum := 0
	for _, p := range patterns {
		score := 0

		// Horizontal
		for l := 0; l < len(p.rows)-1 && score == 0; l++ {
			if isPatternMirroredAroundRowWithSmudge(p, l) {
				score += 100 * (l + 1)
				// fmt.Println("Horizontal match at line", (l + 1))
			}
		}

		// Vertical
		if score == 0 {
			rotated := rotate(p)
			for l := 0; l < len(rotated.rows)-1 && score == 0; l++ {
				if isPatternMirroredAroundRowWithSmudge(rotated, l) {
					score += (l + 1)
					// fmt.Println("Vertical match at column", (l + 1))
					break
				}
			}
		}

		// fmt.Println("Score for", (i + 1), ":", score)

		sum += score
	}

	return fmt.Sprint(sum)
}
