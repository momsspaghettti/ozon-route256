package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func cFindAnswerForSet(in *bufio.Reader, out *bufio.Writer, developerRanks []int, lastSet bool) {
	developerRanks = developerRanks[:0]
	var developersCount int
	_, _ = fmt.Fscan(in, &developersCount)

	var developerRank int
	for i := 0; i < developersCount; i++ {
		_, _ = fmt.Fscan(in, &developerRank)
		developerRanks = append(developerRanks, developerRank)
	}

	var firstDevInd, secondDevInd, minAbsDiff, absDiff int
	firstDevInd = 0
	for i := 0; i < developersCount/2; i++ {
		for developerRanks[firstDevInd] == -1 {
			firstDevInd++
		}

		minAbsDiff = -1
		for j := firstDevInd + 1; j < developersCount; j++ {
			if developerRanks[j] == -1 {
				continue
			}
			absDiff = abs(developerRanks[firstDevInd] - developerRanks[j])
			if minAbsDiff == -1 || absDiff < minAbsDiff {
				minAbsDiff = absDiff
				secondDevInd = j
			}
		}

		_, _ = fmt.Fprintln(out, firstDevInd+1, secondDevInd+1)
		developerRanks[firstDevInd] = -1
		developerRanks[secondDevInd] = -1
	}

	if !lastSet {
		_, _ = fmt.Fprintln(out)
	}
}

func C() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setsCount int
	_, _ = fmt.Fscan(in, &setsCount)

	developerRanks := make([]int, 0, 50)
	for i := 0; i < setsCount; i++ {
		cFindAnswerForSet(in, out, developerRanks, i+1 == setsCount)
	}
}

func main() {
	C()
}
