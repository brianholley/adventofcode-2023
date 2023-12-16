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

func shiftNorth(field [][]rune) [][]rune {
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
	return field
}

func shiftEast(field [][]rune) [][]rune {
	for c := len(field[0]) - 1; c >= 0; c-- {
		for r := range field {
			if field[r][c] == 'O' {
				for i := c + 1; i < len(field[0]); i++ {
					if field[r][i] == '.' {
						field[r][i] = 'O'
						field[r][i-1] = '.'
					} else {
						break
					}
				}
			}
		}
	}
	return field
}

func shiftSouth(field [][]rune) [][]rune {
	for r := len(field) - 1; r >= 0; r-- {
		for c := range field[r] {
			if field[r][c] == 'O' {
				for i := r + 1; i < len(field); i++ {
					if field[i][c] == '.' {
						field[i][c] = 'O'
						field[i-1][c] = '.'
					} else {
						break
					}
				}
			}
		}
	}
	return field
}

func shiftWest(field [][]rune) [][]rune {
	for c := range field[0] {
		for r := range field {
			if field[r][c] == 'O' {
				for i := c - 1; i >= 0; i-- {
					if field[r][i] == '.' {
						field[r][i] = 'O'
						field[r][i+1] = '.'
					} else {
						break
					}
				}
			}
		}
	}
	return field
}

func weight(field [][]rune) int {
	sum := 0
	for r := range field {
		for c := range field[r] {
			if field[r][c] == 'O' {
				sum += len(field) - r
			}
		}
	}
	return sum
}

func part1(stdin *bufio.Scanner) string {
	field := [][]rune{}
	for stdin.Scan() {
		field = append(field, []rune(stdin.Text()))
	}

	// for r := range field {
	// 	fmt.Println(field[r])
	// }

	field = shiftNorth(field)

	// fmt.Println()
	// for r := range field {
	// 	fmt.Println(field[r])
	// }

	sum := weight(field)

	return fmt.Sprint(sum)
}

func hash(field [][]rune) int {
	h := 0
	for r := range field {
		for c := range field[r] {
			if field[r][c] == 'O' {
				h += c + r*len(field[r])
			}
		}
	}
	return h
}

func part2(stdin *bufio.Scanner) string {
	field := [][]rune{}
	for stdin.Scan() {
		field = append(field, []rune(stdin.Text()))
	}

	// for r := range field {
	// 	fmt.Println(field[r])
	// }

	hashes := []int{}

	goalIteration := -1
	for iter := 1; iter <= 1000000000; iter++ {
		field = shiftNorth(field)
		field = shiftWest(field)
		field = shiftSouth(field)
		field = shiftEast(field)

		if iter == goalIteration {
			break
		}

		h := hash(field)
		hashes = append(hashes, h)
		// fmt.Println(iter, weight(field), h)

		if goalIteration == -1 {

			hits := []int{}
			for i := range hashes {
				if hashes[i] == h {
					hits = append(hits, i)
				}
			}

			if len(hits) >= 3 && hits[1]-hits[0] == hits[2]-hits[1] {
				length := hits[1] - hits[0]
				i := hits[0]
				for ; i < hits[1]; i++ {
					if hashes[i] != hashes[i+length] {
						break
					}
				}

				if i == hits[1] {
					// We've found a recurring sequence
					// fmt.Println("SEQUENCE! Iterations ", hits[0], "to", hits[1])

					nextSequenceStart := int(math.Ceil(float64(iter-hits[0])/float64(length)))*length + hits[0]
					goalSequenceOffset := (1000000000 - hits[0]) % length
					goalIteration = nextSequenceStart + goalSequenceOffset

					// fmt.Println("Sequence start:", hits[0])
					// fmt.Println("Sequence len:", length)
					// fmt.Println("Next sequence start:", nextSequenceStart)
					// fmt.Println("Goal sequence offset:", goalSequenceOffset)
					// fmt.Println("Goal iteration:", goalIteration)
				}
			}
		}
	}

	// fmt.Println()
	// for r := range field {
	// 	fmt.Println(field[r])
	// }

	sum := weight(field)

	return fmt.Sprint(sum)
}
