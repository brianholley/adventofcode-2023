package main

import (
	"adventofcode2023/lib"
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func cardPriorityPart1(card byte) int {
	order := []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	for i, v := range order {
		if v == card {
			return i
		}
	}
	panic("Unknown card")
}

func cardPriorityPart2(card byte) int {
	order := []byte{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
	for i, v := range order {
		if v == card {
			return i
		}
	}
	panic("Unknown card")
}

func sortHand(hand string, priority func(card byte) int) string {
	cards := []byte(hand)
	sort.Slice(cards, func(i, j int) bool {
		return priority(cards[i]) < priority(cards[j])
	})
	return string(cards)
}

func scoreHandPart1(hand string) int {
	groups := []int{0, 0, 0, 0}
	card := hand[0]
	count := 1
	for i := 1; i < len(hand); i++ {
		if hand[i] == card {
			count++
		} else {
			if count > 1 {
				groups[count-2]++
			}
			card = hand[i]
			count = 1
		}
	}
	if count > 1 {
		groups[count-2]++
	}

	// fmt.Println(hand, groups)

	// Five of a kind
	if groups[3] > 0 {
		return 1
	}
	// Four of a kind
	if groups[2] > 0 {
		return 2
	}
	// Full house
	if groups[1] > 0 && groups[0] > 0 {
		return 3
	}
	// Three of a kind
	if groups[1] > 0 {
		return 4
	}
	// Two pair
	if groups[0] > 1 {
		return 5
	}
	// One pair
	if groups[0] > 0 {
		return 6
	}
	// High card
	return 7
}

func scoreHandPart2(hand string) int {
	groups := []int{0, 0, 0, 0}
	card := hand[0]
	count := 1
	jokers := 0
	if hand[0] == 'J' {
		jokers++
	}
	for i := 1; i < len(hand); i++ {
		if hand[i] == 'J' {
			jokers++
			continue
		}
		if hand[i] == card {
			count++
		} else {
			if count > 1 {
				groups[count-2]++
			}
			card = hand[i]
			count = 1
		}
	}
	if count > 1 {
		groups[count-2]++
	}

	fmt.Println(hand, groups, jokers)

	for i := len(groups) - 1; i >= 0; i-- {
		if groups[i] > 0 {
			groups[i+jokers]++
			groups[i]--
			jokers = 0
			break
		}
	}
	if jokers == 5 {
		groups[3]++
		jokers = 0
	}
	if jokers > 0 {
		groups[jokers-1]++
		jokers = 0
	}

	fmt.Println(hand, groups, jokers)

	// Five of a kind
	if groups[3] > 0 {
		return 1
	}
	// Four of a kind
	if groups[2] > 0 {
		return 2
	}
	// Full house
	if groups[1] > 0 && groups[0] > 0 {
		return 3
	}
	// Three of a kind
	if groups[1] > 0 {
		return 4
	}
	// Two pair
	if groups[0] > 1 {
		return 5
	}
	// One pair
	if groups[0] > 0 {
		return 6
	}
	// High card
	return 7
}

type player struct {
	hand   string
	sorted string
	bid    int
	score  int
}

func sortPlayers(players []player, priority func(card byte) int) {
	sort.Slice(players, func(a, b int) bool {
		if players[a].score < players[b].score {
			return true
		} else if players[a].score > players[b].score {
			return false
		}

		for i := 0; i < len(players[a].hand) && i < len(players[b].hand); i++ {
			aPri := priority(players[a].hand[i])
			bPri := priority(players[b].hand[i])
			if aPri < bPri {
				return true
			} else if bPri < aPri {
				return false
			}
		}
		panic("Same hand?")
	})
}

func part1(stdin *bufio.Scanner) string {
	players := []player{}
	for stdin.Scan() {
		line := strings.Split(stdin.Text(), " ")
		hand := line[0]
		bid, _ := strconv.Atoi(line[1])

		sorted := sortHand(hand, cardPriorityPart1)
		score := scoreHandPart1(sorted)
		players = append(players, player{hand, sorted, bid, score})
	}

	sortPlayers(players, cardPriorityPart1)

	score := 0
	for i, p := range players {
		// fmt.Println(p.hand, p.sorted, p.bid, p.score)
		score += (len(players) - i) * p.bid
	}
	return fmt.Sprint(score)
}

func part2(stdin *bufio.Scanner) string {
	players := []player{}
	for stdin.Scan() {
		line := strings.Split(stdin.Text(), " ")
		hand := line[0]
		bid, _ := strconv.Atoi(line[1])

		sorted := sortHand(hand, cardPriorityPart2)
		score := scoreHandPart2(sorted)
		players = append(players, player{hand, sorted, bid, score})
	}

	sortPlayers(players, cardPriorityPart2)

	score := 0
	for i, p := range players {
		fmt.Println(p.hand, p.sorted, p.bid, p.score)
		score += (len(players) - i) * p.bid
	}
	return fmt.Sprint(score)
}
