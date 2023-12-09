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

func calcDeltasForHistory(history []int) [][]int {
	deltas := [][]int{history}
	for {
		latest := deltas[len(deltas)-1]
		current := []int{}
		nonzero := false
		for i := 0; i < len(latest)-1; i++ {
			d := latest[i+1] - latest[i]
			current = append(current, d)
			if d != 0 {
				nonzero = true
			}
		}

		// fmt.Println(current)
		deltas = append(deltas, current)
		if !nonzero {
			break
		}
	}
	return deltas
}

func part1(stdin *bufio.Scanner) string {
	sum := 0
	for stdin.Scan() {
		history := lib.ParseStringOfIntsSpaceDelimited(stdin.Text())

		deltas := calcDeltasForHistory(history)

		next := 0
		for level := len(deltas) - 2; level >= 0; level-- {
			next += lib.ArrayLast(deltas[level])
		}

		// fmt.Println(next)
		sum += next
	}
	return fmt.Sprint(sum)
}

func part2(stdin *bufio.Scanner) string {
	sum := 0
	for stdin.Scan() {
		history := lib.ParseStringOfIntsSpaceDelimited(stdin.Text())

		deltas := calcDeltasForHistory(history)

		// prev + [level+1] = [level]
		prev := 0
		for level := len(deltas) - 2; level >= 0; level-- {
			prev = deltas[level][0] - prev
			// fmt.Println(prev, deltas[level])
		}

		// fmt.Println(next)
		sum += prev
		// fmt.Println()
	}
	return fmt.Sprint(sum)
}
