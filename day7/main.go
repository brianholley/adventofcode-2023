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

func cardPriority(card byte) int {
	order := []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	for i, v := range order {
		if v == card {
			return i
		}
	}
	panic("Unknown card")
}

func sortHand(hand string) string {
	cards := []byte(hand)
	sort.Slice(cards, func(i, j int) bool {
		return cardPriority(cards[i]) < cardPriority(cards[j])
	})
	return string(cards)
}

func scoreHand(hand string) int {
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

type player struct {
	hand   string
	sorted string
	bid    int
	score  int
}

func sortPlayers(players []player) {
	sort.Slice(players, func(a, b int) bool {
		if players[a].score < players[b].score {
			return true
		} else if players[a].score > players[b].score {
			return false
		}

		for i := 0; i < len(players[a].hand) && i < len(players[b].hand); i++ {
			aPri := cardPriority(players[a].hand[i])
			bPri := cardPriority(players[b].hand[i])
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

		sorted := sortHand(hand)
		score := scoreHand(sorted)
		players = append(players, player{hand, sorted, bid, score})
	}

	sortPlayers(players)

	score := 0
	for i, p := range players {
		// fmt.Println(p.hand, p.sorted, p.bid, p.score)
		score += (len(players) - i) * p.bid
	}
	return fmt.Sprint(score)
}

func part2(stdin *bufio.Scanner) string {
	return "part2"
}
