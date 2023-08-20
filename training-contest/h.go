package main

import (
	"bufio"
	"fmt"
	"os"
)

type void struct{}

type set map[int]void

func getCellNumber(i, j, m int) int {
	return i*m + j
}

type move struct {
	di, dj int
}

func hClearMap(m map[byte]set) {
	for k := range m {
		delete(m, k)
	}
}

func traversePolygon(matrix []string, i, j, n, m int, regionToCellsMap map[byte]set, moves []move) {
	var emptyCell byte = 46
	cell := matrix[i][j]
	cellNumber := getCellNumber(i, j, m)
	regionToCellsMap[cell][cellNumber] = void{}

	for _, move_ := range moves {
		newI, newJ := i+move_.di, j+move_.dj
		if newI < 0 || newI >= n {
			continue
		}
		if newJ < 0 || newJ >= m {
			continue
		}
		if matrix[newI][newJ] == emptyCell || matrix[newI][newJ] != cell {
			continue
		}
		newCellNumber := getCellNumber(newI, newJ, m)
		if _, ok := regionToCellsMap[cell][newCellNumber]; ok {
			continue
		}

		traversePolygon(matrix, newI, newJ, n, m, regionToCellsMap, moves)
	}
}

func hFindAnswerForSet(in *bufio.Reader, out *bufio.Writer, regionToCellsMap map[byte]set, moves []move) {
	hClearMap(regionToCellsMap)
	var n, m int
	_, _ = fmt.Fscan(in, &n, &m)

	matrix := make([]string, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(in, &matrix[i])
	}

	res := true
	var emptyCell byte = 46
MainLoop:
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == emptyCell {
				continue
			}

			if _, ok := regionToCellsMap[matrix[i][j]]; !ok {
				regionToCellsMap[matrix[i][j]] = make(map[int]void)
				traversePolygon(matrix, i, j, n, m, regionToCellsMap, moves)
				continue
			}

			if _, ok := regionToCellsMap[matrix[i][j]][getCellNumber(i, j, m)]; !ok {
				res = false
				break MainLoop
			}
		}
	}

	if res {
		_, _ = fmt.Fprintln(out, "YES")
	} else {
		_, _ = fmt.Fprintln(out, "NO")
	}
}

func H() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setsCount int
	_, _ = fmt.Fscan(in, &setsCount)

	regionToCellsMap := make(map[byte]set)
	moves := []move{
		{0, -2}, {0, 2},
		{-1, -1}, {-1, 1},
		{1, -1}, {1, 1},
	}

	for i := 0; i < setsCount; i++ {
		hFindAnswerForSet(in, out, regionToCellsMap, moves)
	}
}

func main() {
	H()
}
