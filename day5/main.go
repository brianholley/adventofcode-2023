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
    stdin.Scan()
    seedLine := stdin.Text()
    seeds := lib.ParseStringOfIntsSpaceDelimited(seedLine[strings.Index(seedLine, ":")+1:])
    stdin.Scan() // empty line

    // fmt.Println(seeds)

    current := []int{}
    next := seeds

    for stdin.Scan() {
        line := stdin.Text()

        if len(line) == 0 {
            continue
        }

        if strings.Contains(line, "map:") {
            next = append(next, current...)
            // fmt.Println(line, next)
            current = next
            next = []int{}
        } else {
            mapping := lib.ParseStringOfIntsSpaceDelimited(line)
            dest := mapping[0]
            source := mapping[1]
            length := mapping[2]

            for i := 0; i < len(current); {
                if current[i] >= source && current[i] < source + length {
                    mapped := (current[i] - source) + dest
                    fmt.Println(current[i], "->", mapped)
                    next = append(next, mapped)
                    current = append(current[:i], current[i+1:]...)
                } else {
                    i++
                }
            }
        }
    }

    next = append(next, current...)
    // fmt.Println("Locations", next)
    return fmt.Sprint(lib.ArrayMin(next))
}

type seq struct {
    start int
    len int
}

func seqIntersects(a seq, b seq) seq {
    start := lib.Max(a.start, b.start)
    end := lib.Min(a.start + a.len - 1, b.start + b.len - 1)
    return seq{start, end - start + 1}
}

func part2(stdin *bufio.Scanner) string {
    stdin.Scan()
    seedLine := stdin.Text()
    seedRanges := lib.ParseStringOfIntsSpaceDelimited(seedLine[strings.Index(seedLine, ":")+1:])
    seedSeq := []seq{}
    for i:=0; i < len(seedRanges); i+=2 {
        seedSeq = append(seedSeq, seq{seedRanges[i], seedRanges[i+1]})
    }
    stdin.Scan() // empty line

    // fmt.Println(seedSeq)

    current := []seq{}
    next := seedSeq

    for stdin.Scan() {
        line := stdin.Text()

        if len(line) == 0 {
            continue
        }

        if strings.Contains(line, "map:") {
            next = append(next, current...)
            // fmt.Println(line, next)
            current = next
            next = []seq{}
        } else {
            mapping := lib.ParseStringOfIntsSpaceDelimited(line)
            dest := mapping[0]
            source := mapping[1]
            length := mapping[2]

            for i := 0; i < len(current); {
                isect := seqIntersects(current[i], seq{source, length})
                if isect.len > 0 {
                    mapped := (isect.start - source) + dest
                    // fmt.Println(current[i], "->", mapped)
                    next = append(next, seq{mapped, isect.len})
                    if isect.start > current[i].start {
                        current = append(current, seq{current[i].start, isect.start - current[i].start})
                    }
                    if isect.start + isect.len < current[i].start + current[i].len {
                        current = append(current, seq{isect.start + isect.len, current[i].start + current[i].len - (isect.start + isect.len)})
                    }
                    current = append(current[:i], current[i+1:]...)
                } else {
                    i++
                }
            }
        }
    }
    
    next = append(next, current...)
    // fmt.Println("Locations", next)

    min := next[0].start
    for i := 1; i < len(next); i++ {
        if next[i].start < min {
            min = next[i].start
        }
    }
    return fmt.Sprint(min)
}
