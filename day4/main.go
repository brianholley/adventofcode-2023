package main

import (
	"adventofcode2023/lib"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func part1(stdin *bufio.Scanner) string {
	i := 0
	sum := 0
	for stdin.Scan() {
		line := stdin.Text()
		line = line[strings.Index(line, ": ")+2:]

		parts := strings.Split(line, " | ")
		winners := lib.ParseStringOfIntsSpaceDelimited(parts[0])
		cards := lib.ParseStringOfIntsSpaceDelimited(parts[1])

		matches := 0
		for _, c := range cards {
			if lib.ArrayContains(winners, c) {
				matches++
			}
		}

		if matches > 0 {
			sum += 1 << (matches - 1)
		}
		// fmt.Println(i+1, ":", matches, sum)
		i++
	}
	return fmt.Sprint(sum)
}

func part2(stdin *bufio.Scanner) string {
	i := 0
	sum := 0
	cardCounts := []int{}
	for stdin.Scan() {
		line := stdin.Text()
		line = line[strings.Index(line, ": ")+2:]

		if i < len(cardCounts) {
			cardCounts[i]++
		} else {
			cardCounts = append(cardCounts, 1)
		}

		parts := strings.Split(line, " | ")
		winners := lib.ParseStringOfIntsSpaceDelimited(parts[0])
		cards := lib.ParseStringOfIntsSpaceDelimited(parts[1])

		matches := 0
		for _, c := range cards {
			if lib.ArrayContains(winners, c) {
				matches++
			}
		}

		for m := i + 1; m <= i+matches; m++ {
			if m < len(cardCounts) {
				cardCounts[m] += cardCounts[i]
			} else {
				cardCounts = append(cardCounts, cardCounts[i])
			}
		}

		sum += cardCounts[i]
		// fmt.Println(i+1, ":", cardCounts)
		i++
	}
	return fmt.Sprint(sum)
}
