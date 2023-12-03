package main

import (
    "adventofcode2023/lib"
    "bufio"
    "fmt"
    "strconv"
    "unicode"
)

func main() {
    result := lib.Run(part1, part2)
    fmt.Println(result)
}

type part struct {
    row int
    col int
    value int
    partLen int
}

type symbol struct {
    row int
    col int
    value rune
}

func readSchematic(stdin *bufio.Scanner) ([]part, []symbol) {
    parts := []part{}
    symbols := []symbol{}

    r := 0
    for stdin.Scan() {
        line := stdin.Text()
        currentPart := part{-1, -1, 0, -1}
        for c:=0; c < len(line); c++  {
            if unicode.IsDigit(rune(line[c])) {
                d, _ := strconv.Atoi(string(line[c]))
                currentPart.value = currentPart.value * 10 + d
                if currentPart.row == -1 {
                    currentPart.row = r
                    currentPart.col = c
                }
            } else {
                if rune(line[c]) != '.' {
                    symbols = append(symbols, symbol{r, c, rune(line[c])})
                }
                if currentPart.value > 0 {
                    currentPart.partLen = len(strconv.Itoa(currentPart.value))
                    parts = append(parts, currentPart)
                    currentPart = part{-1, -1, 0, -1}
                }
            }
        }
        if currentPart.value > 0 {
            currentPart.partLen = len(strconv.Itoa(currentPart.value))
            parts = append(parts, currentPart)
            currentPart = part{-1, -1, 0, -1}
        }
        r++
    }
    return parts, symbols
}

func part1(stdin *bufio.Scanner) string {
    parts, symbols := readSchematic(stdin)

    // fmt.Println("Parts", len(parts))
    // fmt.Println("Symbols", len(symbols))

    sum := 0
    for _, p := range parts {
        // fmt.Println(sum)
        // fmt.Println(p.value, p.row, p.col)

        for _, s := range symbols {
            if s.row >= p.row - 1 && s.row <= p.row + 1 && s.col >= p.col - 1 && s.col <= p.col + p.partLen {
                // fmt.Println("Match", p.value, string(s.value))
                sum += p.value
                break
            }
        }

    }
    return fmt.Sprint(sum)
}

func part2(stdin *bufio.Scanner) string {
    parts, symbols := readSchematic(stdin)

    sum := 0
    for _, s := range symbols {
        adjacent := 0
        product := 1

        for _, p := range parts {
            if s.row >= p.row - 1 && s.row <= p.row + 1 && s.col >= p.col - 1 && s.col <= p.col + p.partLen {
                adjacent++
                product *= p.value
            }
        }

        if adjacent == 2 {
            // fmt.Println("Gear at", s.row, s.col)
            sum += product
        }

    }
    return fmt.Sprint(sum)
}
