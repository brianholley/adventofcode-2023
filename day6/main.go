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

func part1(stdin *bufio.Scanner) string {
	stdin.Scan()
	line := stdin.Text()
	times := lib.ParseStringOfIntsSpaceDelimited(line[strings.Index(line, ":")+1:])

	stdin.Scan()
	line = stdin.Text()
	distances := lib.ParseStringOfIntsSpaceDelimited(line[strings.Index(line, ":")+1:])

	result := 1

	for i, _ := range times {
		wins := 0
		for t := 1; t < times[i]; t++ {
			dist := t * (times[i] - t)
			if dist > distances[i] {
				wins++
			}
		}

		result *= wins
	}

	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	stdin.Scan()
	line := stdin.Text()
	time, _ := strconv.Atoi(strings.ReplaceAll(line[strings.Index(line, ":")+1:], " ", ""))

	stdin.Scan()
	line = stdin.Text()
	distance, _ := strconv.Atoi(strings.ReplaceAll(line[strings.Index(line, ":")+1:], " ", ""))

	wins := 0
	for t := 1; t < time; t++ {
		dist := t * (time - t)
		if dist > distance {
			wins++
		}
	}

	return fmt.Sprint(wins)
}
