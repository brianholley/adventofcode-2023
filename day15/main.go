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
	stdin.Scan()
	steps := strings.Split(stdin.Text(), ",")
	for _, step := range steps {
		h := runHASH(step)
		// fmt.Println(step, h, sum)
		sum += h
	}
	return fmt.Sprint(sum)
}

type lens struct {
	label string
	focal int
}

func printBoxes(boxes [][]lens) {
	for i := range boxes {
		if len(boxes[i]) > 0 {
			fmt.Print("Box", i, ":")

			for l := range boxes[i] {
				prt := "[" + boxes[i][l].label + " " + fmt.Sprint(boxes[i][l].focal) + "]"
				fmt.Print(prt)
			}
			fmt.Println()
		}
	}
}

func part2(stdin *bufio.Scanner) string {

	boxes := make([][]lens, 256)
	for i := range boxes {
		boxes[i] = []lens{}
	}

	// lenses := map[string]int{}

	stdin.Scan()
	steps := strings.Split(stdin.Text(), ",")
	for _, step := range steps {
		if strings.HasSuffix(step, "-") {
			sep := strings.Index(step, "-")
			label := step[:sep]
			box := runHASH(label)

			for i := range boxes[box] {
				if boxes[box][i].label == label {
					if i < len(boxes[box])-1 {
						boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
					} else {
						boxes[box] = boxes[box][:i]
					}
					break
				}
			}
		} else {
			sep := strings.Index(step, "=")
			label := step[:sep]
			box := runHASH(label)
			focal, _ := strconv.Atoi(step[sep+1:])

			replaced := false
			for i := range boxes[box] {
				if boxes[box][i].label == label {
					boxes[box][i].focal = focal
					replaced = true
					break
				}
			}
			if !replaced {
				boxes[box] = append(boxes[box], lens{label, focal})
			}
		}
		// fmt.Println("Step", step)
		// printBoxes(boxes)
		// fmt.Println()
	}

	sum := 0
	for i := range boxes {
		for l := range boxes[i] {
			sum += (i + 1) * (l + 1) * boxes[i][l].focal
		}
	}

	return fmt.Sprint(sum)
}
