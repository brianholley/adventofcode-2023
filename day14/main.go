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

func part1(stdin *bufio.Scanner) string {
	field := [][]rune{}
	for stdin.Scan() {
		field = append(field, []rune(stdin.Text()))
	}

	// for r := range field {
	// 	fmt.Println(field[r])
	// }

	for r := range field {
		for c := range field[r] {
			if field[r][c] == 'O' {
				for i := r - 1; i >= 0; i-- {
					if field[i][c] == '.' {
						field[i][c] = 'O'
						field[i+1][c] = '.'
					} else {
						break
					}
				}
			}
		}
	}

	// fmt.Println()
	// for r := range field {
	// 	fmt.Println(field[r])
	// }

	sum := 0
	for r := range field {
		for c := range field[r] {
			if field[r][c] == 'O' {
				sum += len(field) - r
			}
		}
		// fmt.Println(sum)
	}

	return fmt.Sprint(sum)
}

func part2(stdin *bufio.Scanner) string {
	return "part2"
}
