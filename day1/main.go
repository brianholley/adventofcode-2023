package main

import (
    "adventofcode2023/lib"
    "bufio"
    "fmt"
    "strconv"
    "strings"
    "unicode"
)

func main() {
    result := lib.Run(part1, part2)
    fmt.Println(result)
}

type SubstrFunc func(string, string) bool

func isSpelledOut(str string, match SubstrFunc) int {
    spellings := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }
    
    for k, v := range spellings {       
        if match(str, k) {
            return v
        }
    }
    return 0
}

func firstDigit(str string, allowSpellings bool) string {
    for i, c := range str {
        if unicode.IsDigit(c) {
            return string(c)
        }
        if allowSpellings {
            digit := isSpelledOut(str[i:], strings.HasPrefix)
            // fmt.Println("P %s %d", str[i:], digit)
            if digit > 0 {
                return fmt.Sprint(digit)
            }
        }
    }
    return ""
}

func lastDigit(str string, allowSpellings bool) string {

    for i:=len(str)-1; i >= 0; i-- {
        if unicode.IsDigit(rune(str[i])) {
            return string(str[i])
        }
        if allowSpellings {
            digit := isSpelledOut(str[:i+1], strings.HasSuffix)
            // fmt.Println("S %s %d", str[:i+1], digit)
            if digit > 0 {
                return fmt.Sprint(digit)
            }
        }
    }
    return ""
}

func part1(stdin *bufio.Scanner) string {
    sum := 0
    for stdin.Scan() {
        line := stdin.Text()
        num, _ := strconv.Atoi(firstDigit(line, false) + lastDigit(line, false))
        sum += num
    }
    return fmt.Sprint(sum)
}

func part2(stdin *bufio.Scanner) string {
    sum := 0
    for stdin.Scan() {
        line := stdin.Text()
        num, _ := strconv.Atoi(firstDigit(line, true) + lastDigit(line, true))
        // fmt.Println("%d", num)
        sum += num
    }
    return fmt.Sprint(sum)
}
