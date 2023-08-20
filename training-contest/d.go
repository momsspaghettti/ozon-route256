package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func dFindAnswerForSet(in *bufio.Reader, out *bufio.Writer, lastSet bool) {
	_, _ = fmt.Fscan(in)
	var n, m int
	_, _ = fmt.Fscan(in, &n, &m)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			_, _ = fmt.Fscan(in, &matrix[i][j])
		}
	}

	var queriesCount, colInd, prevColInd int
	prevColInd = -1
	_, _ = fmt.Fscan(in, &queriesCount)
	for k := 0; k < queriesCount; k++ {
		_, _ = fmt.Fscan(in, &colInd)
		colInd--
		if colInd == prevColInd {
			continue
		}
		sort.SliceStable(matrix, func(i, j int) bool {
			return matrix[i][colInd] < matrix[j][colInd]
		})
		prevColInd = colInd
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			_, _ = fmt.Fprint(out, matrix[i][j])
			if j+1 != m {
				_, _ = fmt.Fprint(out, " ")
			}
		}
		_, _ = fmt.Fprintln(out)
	}

	if !lastSet {
		_, _ = fmt.Fprintln(out)
	}
}

func D() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setsCount int
	_, _ = fmt.Fscan(in, &setsCount)

	for i := 0; i < setsCount; i++ {
		dFindAnswerForSet(in, out, i+1 == setsCount)
	}
}

func main() {
	D()
}
