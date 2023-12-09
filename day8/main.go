package main

import (
	"adventofcode2023/lib"
	"bufio"
	"fmt"
	"strings"
)

type node struct {
	name  string
	left  string
	right string
}

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func part1(stdin *bufio.Scanner) string {
	stdin.Scan()
	instructions := stdin.Text()
	stdin.Scan() // empty line

	graph := map[string]node{}

	for stdin.Scan() {
		line := stdin.Text()

		name := line[:3]
		left := line[7:10]
		right := line[12:15]

		graph[name] = node{name, left, right}
	}

	pos := "AAA"
	i := 0
	steps := 0
	for pos != "ZZZ" {
		if instructions[i] == 'L' {
			pos = graph[pos].left
		} else {
			pos = graph[pos].right
		}
		steps++
		i++
		if i >= len(instructions) {
			i = 0
		}
	}

	return fmt.Sprint(steps)
}

// func part2BruteForce(stdin *bufio.Scanner) string {
// 	stdin.Scan()
// 	instructions := stdin.Text()
// 	stdin.Scan() // empty line

// 	graph := map[string]node{}

// 	pos := []string{}

// 	for stdin.Scan() {
// 		line := stdin.Text()

// 		name := line[:3]
// 		left := line[7:10]
// 		right := line[12:15]

// 		graph[name] = node{name, left, right}

// 		if strings.HasSuffix(name, "A") {
// 			pos = append(pos, name)
// 		}
// 	}

// 	i := 0
// 	steps := 0
// 	for {
// 		for p, _ := range pos {
// 			if instructions[i] == 'L' {
// 				pos[p] = graph[pos[p]].left
// 			} else {
// 				pos[p] = graph[pos[p]].right
// 			}
// 		}
// 		steps++
// 		i++
// 		if i >= len(instructions) {
// 			i = 0
// 		}

// 		// Break condition
// 		ended := 0
// 		for p, _ := range pos {
// 			if strings.HasSuffix(pos[p], "Z") {
// 				ended++
// 			}
// 		}
// 		if ended == len(pos) {
// 			break
// 		}

// 		if ended > 2 {
// 			fmt.Println(steps, ended, pos)
// 		}
// 	}

func GCD(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a int, b int, c ...int) int {
	lcm := a * b / GCD(a, b)

	for i := 0; i < len(c); i++ {
		lcm = LCM(lcm, c[i])
	}

	return lcm
}

func part2(stdin *bufio.Scanner) string {
	stdin.Scan()
	instructions := stdin.Text()
	stdin.Scan() // empty line

	graph := map[string]node{}

	pos := []string{}

	for stdin.Scan() {
		line := stdin.Text()

		name := line[:3]
		left := line[7:10]
		right := line[12:15]

		graph[name] = node{name, left, right}

		if strings.HasSuffix(name, "A") {
			pos = append(pos, name)
		}
	}

	cycle := []int{}
	for p, _ := range pos {
		i := 0
		steps := 0
		for !strings.HasSuffix(pos[p], "Z") {
			if instructions[i] == 'L' {
				pos[p] = graph[pos[p]].left
			} else {
				pos[p] = graph[pos[p]].right
			}

			steps++
			i++
			if i >= len(instructions) {
				i = 0
			}
		}

		cycle = append(cycle, steps)
	}

	fmt.Println(cycle)

	lcm := LCM(cycle[0], cycle[1], cycle[2:]...)

	return fmt.Sprint(lcm)
}
