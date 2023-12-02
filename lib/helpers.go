package lib

import (
    "bufio"
	"os"
	"strconv"
	"strings"
)

type PartFn func(*bufio.Scanner) string

func Run(part1 PartFn, part2 PartFn) string {
	if len(os.Args[1:]) == 0 {
		panic("Expected part number argument")
	}

	scanner := bufio.NewScanner(os.Stdin)
 
    switch part := os.Args[1]; part {
	case "1":
		return part1(scanner)
	case "2":
		return part2(scanner)
	default:
		panic("Unknown part: " + part)
	}
}

func Read2dArray(scanner bufio.Scanner) [][]int {
	array := make([][]int, 0)
	for scanner.Scan() {
        s := strings.Split(scanner.Text(), " ")
		row := make([]int, len(s))
		for i, v := range s {
			row[i], _ = strconv.Atoi(v)
		}
		array = append(array, row)
    }
	return array
}

func Create2dArray(rows int, cols int, defaultValue int) [][]int {
	arr := make([][]int, rows)
	for i:=0; i < rows; i++ {
		arr = append(arr, make([]int, cols, defaultValue))
	}
	return arr
}

func SumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}