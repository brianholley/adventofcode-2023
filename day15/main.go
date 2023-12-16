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

func runHASH(str string) int {
	current := 0
	for _, r := range str {
		current += int(byte(r))
		current *= 17
		current = current % 256
	}
	return current
}

func part1(stdin *bufio.Scanner) string {
	sum := 0
	for stdin.Scan() {
		steps := strings.Split(stdin.Text(), ",")
		for _, step := range steps {
			h := runHASH(step)
			// fmt.Println(step, h, sum)
			sum += h
		}
	}
	return fmt.Sprint(sum)
}

func part2(stdin *bufio.Scanner) string {
	return "part2"
}
