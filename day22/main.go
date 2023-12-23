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

type pos struct {
	x, y, z int
}

type brick struct {
	index       int
	cubes       []pos
	falling     bool
	supports    []int
	supportedBy []int
}

func parsePos(coord string) pos {
	parts := strings.Split(coord, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])
	return pos{x, y, z}
}

func parseBricks(stdin *bufio.Scanner) []brick {
	bricks := []brick{}
	for stdin.Scan() {
		line := stdin.Text()
		parts := strings.Split(line, "~")
		start, end := parsePos(parts[0]), parsePos(parts[1])

		cubes := []pos{}
		if start.x != end.x {
			for x := lib.Min(start.x, end.x); x <= lib.Max(start.x, end.x); x++ {
				cubes = append(cubes, pos{x, start.y, start.z})
			}
		} else if start.y != end.y {
			for y := lib.Min(start.y, end.y); y <= lib.Max(start.y, end.y); y++ {
				cubes = append(cubes, pos{start.x, y, start.z})
			}
		} else if start.z != end.z {
			for z := lib.Min(start.z, end.z); z <= lib.Max(start.z, end.z); z++ {
				cubes = append(cubes, pos{start.x, start.y, z})
			}
		} else {
			cubes = append(cubes, start)
		}

		b := brick{len(bricks), cubes, true, []int{}, []int{}}
		bricks = append(bricks, b)
	}
	return bricks
}

func settleBricks(bricks []brick) {
	falling := true
	for falling {
		falling = false

		// Check for anything on the floor or supported by a brick no longer falling
		changed := true
		for changed {
			changed = false

			for i := range bricks {
				if bricks[i].falling {
					if isOnFloor(bricks[i]) {
						bricks[i].falling = false
						// fmt.Println("Brick", i, "hit the floor")
						changed = true
						continue
					}

					under := supportedBy(bricks[i], bricks)
					if len(under) > 0 {
						bricks[i].falling = false
						// fmt.Println("Brick", i, "now supported")
						changed = true
						continue
					}
				}
			}
		}

		// Drop all the falling bricks 1 space
		for i := range bricks {
			if bricks[i].falling {
				// fmt.Println("Brick", i, "drops")
				for c := range bricks[i].cubes {
					bricks[i].cubes[c].z--
				}
				falling = true
			}
		}
	}

	// fmt.Println("END")
	for i := range bricks {
		// fmt.Println("Brick", i, bricks[i].cubes[0], "~", bricks[i].cubes[len(bricks[i].cubes)-1])

		under := supportedBy(bricks[i], bricks)
		for _, u := range under {
			bricks[i].supportedBy = append(bricks[i].supportedBy, u.index)
			bricks[u.index].supports = append(bricks[u.index].supports, i)
		}
	}
}

func isOnFloor(b brick) bool {
	for _, c := range b.cubes {
		if c.z == 1 {
			return true
		}
	}
	return false
}

// Does a support b?
func supports(a brick, b brick) bool {
	if a.falling {
		return false
	}
	for i := range a.cubes {
		for j := range b.cubes {
			if a.cubes[i].x == b.cubes[j].x && a.cubes[i].y == b.cubes[j].y && a.cubes[i].z+1 == b.cubes[j].z {
				return true
			}
		}
	}
	return false
}

// Note: this could use bounding cubes to drastically speed up collision detection
func supportedBy(b brick, bricks []brick) []brick {
	result := []brick{}
	for i := range bricks {
		if i != b.index && supports(bricks[i], b) {
			result = append(result, bricks[i])
		}
	}
	return result
}

func part1(stdin *bufio.Scanner) string {
	bricks := parseBricks(stdin)
	settleBricks(bricks)

	result := 0
	for i := range bricks {
		// fmt.Println("Brick", i, "supports", bricks[i].supports)
		soleSupport := false
		for _, s := range bricks[i].supports {
			// fmt.Println("  brick", s, "is supported by", bricks[s].supportedBy)
			if len(bricks[s].supportedBy) == 1 {
				soleSupport = true
			}
		}
		if !soleSupport {
			// fmt.Println("  can be disintegrated")
			result++
		} else {
			// fmt.Println("  CANNOT be disintegrated")
		}
	}
	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	bricks := parseBricks(stdin)
	settleBricks(bricks)

	result := 0
	for i := range bricks {
		// fmt.Println("-- Checking brick", i)
		bricksToCheck := []int{i}
		fallingBricks := []int{i}
		for len(bricksToCheck) > 0 {
			b := bricksToCheck[0]
			bricksToCheck = bricksToCheck[1:]
			// fmt.Println("Brick", b, "supports", bricks[b].supports)
			for _, s := range bricks[b].supports {
				if lib.IndexOf(fallingBricks, s) == -1 {
					reducedSupports := make([]int, len(bricks[s].supportedBy))
					copy(reducedSupports, bricks[s].supportedBy)
					// fmt.Println("Brick", s, "supported by", reducedSupports)
					for _, f := range fallingBricks {
						reducedSupports = lib.ArrayRemoveItem(reducedSupports, f)
						// fmt.Println("Removing", f, "is", reducedSupports)
					}
					if len(reducedSupports) == 0 {
						fallingBricks = append(fallingBricks, s)
						bricksToCheck = append(bricksToCheck, s)
						// fmt.Println("Brick", s, "will fall also")
					}
				}
			}
		}

		// fmt.Println("Brick", i, "would cause", len(fallingBricks)-1, "bricks to fall")
		result += len(fallingBricks) - 1
	}
	return fmt.Sprint(result)
}
