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
	result := 0
	for stdin.Scan() {
		line := stdin.Text()
		fmt.Println(line)
	}
	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	return fmt.Sprint(result)
}
