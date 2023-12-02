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
    for stdin.Scan() {
        fmt.Println(stdin.Text())
    }
    return "part1"
}

func part2(stdin *bufio.Scanner) string {
    return "part2"
}
