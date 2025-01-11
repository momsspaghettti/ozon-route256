package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

func eMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func E() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var friendsCount, cardsCount int
	_, _ = fmt.Fscan(in, &friendsCount, &cardsCount)

	friendsCards := make([]Pair[int, int], 0, friendsCount)
	var friendCards int
	for i := 0; i < friendsCount; i++ {
		_, _ = fmt.Fscan(in, &friendCards)
		friendsCards = append(friendsCards, Pair[int, int]{First: friendCards, Second: i})
	}

	sort.Slice(friendsCards, func(i, j int) bool {
		return friendsCards[i].First < friendsCards[j].First
	})

	answerExists := true
	lastTakenCard := 0
	for i := range friendsCards {
		lastTakenCard = eMax(lastTakenCard+1, friendsCards[i].First+1)
		friendsCards[i].First = lastTakenCard
		if lastTakenCard > cardsCount {
			answerExists = false
			break
		}
	}

	if answerExists {
		sort.Slice(friendsCards, func(i, j int) bool {
			return friendsCards[i].Second < friendsCards[j].Second
		})

		for i, card := range friendsCards {
			_, _ = fmt.Fprint(out, card.First)
			if i+1 != len(friendsCards) {
				_, _ = fmt.Fprint(out, " ")
			}
		}
	} else {
		_, _ = fmt.Fprintln(out, -1)
	}
}

func main() {
	E()
}
