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

	red := 12
	green := 13
	blue := 14

	validGames := []int{}
	game := 1
	for stdin.Scan() {
		line := stdin.Text()

		line = line[strings.Index(line, ": ")+2:]
		// fmt.Println(line)

		validGame := true
		plays := strings.Split(line, "; ")
		for _, p := range plays {
			cubes := strings.Split(p, ", ")
			for _, c := range cubes {
				parts := strings.Split(c, " ")
				count, _ := strconv.Atoi(parts[0])
				color := parts[1]
				if (color == "red" && count > red) ||
					(color == "green" && count > green) ||
					(color == "blue" && count > blue) {
					validGame = false
				}

			}
		}
		if validGame {
			validGames = append(validGames, game)
		}
		game++
	}

	return fmt.Sprint(lib.ArraySum(validGames))
}

func part2(stdin *bufio.Scanner) string {
	powers := []int{}
	game := 1
	for stdin.Scan() {
		line := stdin.Text()

		line = line[strings.Index(line, ": ")+2:]
		// fmt.Println(line)

		red, green, blue := 0, 0, 0

		plays := strings.Split(line, "; ")
		for _, p := range plays {
			cubes := strings.Split(p, ", ")
			for _, c := range cubes {
				parts := strings.Split(c, " ")
				count, _ := strconv.Atoi(parts[0])
				color := parts[1]
				switch color {
				case "red":
					if count > red {
						red = count
					}
				case "green":
					if count > green {
						green = count
					}
				case "blue":
					if count > blue {
						blue = count
					}
				}
			}
		}
		powers = append(powers, red*green*blue)
		game++
	}

	return fmt.Sprint(lib.ArraySum(powers))
}
