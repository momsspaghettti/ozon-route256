package t5

import (
	"bufio"
	"fmt"
	"os"
)

type cellType byte

const (
	barrier    cellType = '#'
	empty      cellType = '.'
	robotA     cellType = 'A'
	robotB     cellType = 'B'
	robotAPath cellType = 'a'
	robotBPath cellType = 'b'
)

func Task5() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func() {
		_ = out.Flush()
	}()

	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		panic(fmt.Errorf("failed to read t: %w", err))
	}

	for i := range t {
		matrix, err := readWarehouseScheme(in)
		if err != nil {
			panic(fmt.Errorf("failed to read warehouse scheme #%d: %w", i, err))
		}

		solveRobotsPaths(matrix)

		if err = writeAnswer(out, matrix); err != nil {
			panic(fmt.Errorf("failed to write answer #%d: %w", i, err))
		}
	}
}

func solveRobotsPaths(matrix [][]cellType) {
	ai, aj := getRobotCoordinates(matrix, robotA)
	bi, bj := getRobotCoordinates(matrix, robotB)

	if tryPavePathTopLeft(matrix, ai, aj, robotAPath) {
		if tryPavePathBottomRight(matrix, bi, bj, robotBPath) {
			return
		}
		clearPath(matrix, robotAPath)
	}

	if tryPavePathTopLeft(matrix, bi, bj, robotBPath) {
		if tryPavePathBottomRight(matrix, ai, aj, robotAPath) {
			return
		}
	}

	out := bufio.NewWriter(os.Stderr)
	defer func() {
		_ = out.Flush()
	}()
	err := writeAnswer(out, matrix)
	panic(fmt.Errorf("failed to solve robots paths; write error: %w", err))
}

func tryPavePathTopLeft(matrix [][]cellType, fi, fj int, path cellType) bool {
	if fi == 0 && fj == 0 {
		return true
	}

	if fi > 0 && matrix[fi-1][fj] == empty {
		matrix[fi-1][fj] = path
		if tryPavePathTopLeft(matrix, fi-1, fj, path) {
			return true
		}
		matrix[fi-1][fj] = empty
	}
	if fj > 0 && matrix[fi][fj-1] == empty {
		matrix[fi][fj-1] = path
		if tryPavePathTopLeft(matrix, fi, fj-1, path) {
			return true
		}
		matrix[fi][fj-1] = empty
	}

	return false
}

func tryPavePathBottomRight(matrix [][]cellType, fi, fj int, path cellType) bool {
	if fi == len(matrix)-1 && fj == len(matrix[fi])-1 {
		return true
	}

	if fi+1 < len(matrix) && matrix[fi+1][fj] == empty {
		matrix[fi+1][fj] = path
		if tryPavePathBottomRight(matrix, fi+1, fj, path) {
			return true
		}
		matrix[fi+1][fj] = empty
	}
	if fj+1 < len(matrix[fi]) && matrix[fi][fj+1] == empty {
		matrix[fi][fj+1] = path
		if tryPavePathBottomRight(matrix, fi, fj+1, path) {
			return true
		}
		matrix[fi][fj+1] = empty
	}

	return false
}

func clearPath(matrix [][]cellType, ct cellType) {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == ct {
				matrix[i][j] = empty
			}
		}
	}
}

func getRobotCoordinates(matrix [][]cellType, rType cellType) (x, y int) {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == rType {
				return i, j
			}
		}
	}
	panic(fmt.Errorf("failed to find robot '%s' coordinate", string([]byte{byte(rType)})))
}

func writeAnswer(out *bufio.Writer, matrix [][]cellType) error {
	for _, row := range matrix {
		for _, cell := range row {
			if err := out.WriteByte(byte(cell)); err != nil {
				return err
			}
		}
		if err := out.WriteByte('\n'); err != nil {
			return err
		}
	}
	return nil
}

func readWarehouseScheme(in *bufio.Reader) ([][]cellType, error) {
	var n, m int
	if _, err := fmt.Fscan(in, &n, &m); err != nil {
		return nil, fmt.Errorf("failed to read n and m: %w", err)
	}
	_, err := in.ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read new line after n and m: %w", err)
	}

	matrix := make([][]cellType, 0, n)
	for i := range n {
		line := make([]cellType, 0, m)
		var b byte
		for j := range m {
			b, err = in.ReadByte()
			if err != nil {
				return nil, fmt.Errorf("failed to read byte #%d for line #%d: %w", j, i, err)
			}
			line = append(line, cellType(b))
		}
		matrix = append(matrix, line)

		_, err = in.ReadBytes('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read new line after line %d: %w", i, err)
		}
	}

	return matrix, nil
}
