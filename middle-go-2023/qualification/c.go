package main

import (
	"bufio"
	"fmt"
	"os"
)

func cMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func cFindAnswerForSet(in *bufio.Reader, out *bufio.Writer) {
	var employeesCount int
	_, _ = fmt.Fscan(in, &employeesCount)

	suitableRangeFrom := 15
	suitableRangeTo := 30
	suitableRangeExists := true

	leftRangeConstraintSign := ">="
	var rangeConstraintSign string
	var rangeConstraintValue int
	for i := 0; i < employeesCount; i++ {
		_, _ = fmt.Fscan(in, &rangeConstraintSign, &rangeConstraintValue)
		if !suitableRangeExists {
			_, _ = fmt.Fprintln(out, -1)
			continue
		}

		if rangeConstraintSign == leftRangeConstraintSign {
			if rangeConstraintValue > suitableRangeTo {
				suitableRangeExists = false
				_, _ = fmt.Fprintln(out, -1)
				continue
			}

			suitableRangeFrom = cMax(suitableRangeFrom, rangeConstraintValue)
			_, _ = fmt.Fprintln(out, suitableRangeFrom)
			continue
		}

		if rangeConstraintValue < suitableRangeFrom {
			suitableRangeExists = false
			_, _ = fmt.Fprintln(out, -1)
			continue
		}

		suitableRangeTo = cMin(suitableRangeTo, rangeConstraintValue)
		_, _ = fmt.Fprintln(out, suitableRangeFrom)
	}

	_, _ = fmt.Fprintln(out)
}

func C() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setsCount int
	_, _ = fmt.Fscan(in, &setsCount)

	for i := 0; i < setsCount; i++ {
		cFindAnswerForSet(in, out)
	}
}

func main() {
	C()
}
