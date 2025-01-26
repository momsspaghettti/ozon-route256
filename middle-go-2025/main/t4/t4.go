package t4

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type box struct {
	Name     string
	Children []*box
	Area     int
}

func (b *box) writeToMap(m map[string]any) {
	if len(b.Name) == 0 {
		return
	}
	if len(b.Children) == 0 {
		m[b.Name] = b.Area
		return
	}

	cMap := make(map[string]any, len(b.Children))
	m[b.Name] = cMap
	for _, child := range b.Children {
		child.writeToMap(cMap)
	}
}

func Task4() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func() {
		_ = out.Flush()
	}()

	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		panic(fmt.Errorf("failed to read t: %w", err))
	}

	res := make([]map[string]any, 0, t)
	for range t {
		boxes, err := readBoxes(in)
		if err != nil {
			panic(err)
		}
		res = append(res, writeToMap(boxes))
	}

	if err := json.NewEncoder(out).Encode(res); err != nil {
		panic(err)
	}
}

func readBoxes(in *bufio.Reader) ([]*box, error) {
	var n, m int
	if _, err := fmt.Fscan(in, &n, &m); err != nil {
		return nil, err
	}

	_, _ = in.ReadBytes('\n')

	matrix := make([][]byte, 0, n)
	for range n {
		row, err := in.ReadBytes('\n')
		if err != nil {
			return nil, err
		}
		matrix = append(matrix, row[:m])
	}

	return parseMatrix(matrix, 0, 0, n, m), nil
}

func parseMatrix(matrix [][]byte, fi, fj, n, m int) []*box {
	res := make([]*box, 0)
	for i := fi; i < n; i++ {
		for j := fj; j < m; j++ {
			if matrix[i][j] == '+' {
				res = append(res, parseBox(matrix, i, j))
			}
		}
	}
	return res
}

func parseBox(matrix [][]byte, i, j int) *box {
	matrix[i][j] = '.'

	lj := j + 1
	for matrix[i][lj] != '+' {
		matrix[i][lj] = '.'
		lj++
	}
	matrix[i][lj] = '.'

	ib := i + 1
	for matrix[ib][lj] != '+' {
		matrix[ib][lj] = '.'
		ib++
	}
	matrix[ib][lj] = '.'

	matrix[ib][j] = '.'

	name := make([]byte, 0)
	ni := i + 1
	nj := j + 1
	for matrix[ni][nj] != '.' && matrix[ni][nj] != '+' && matrix[ni][nj] != '-' && matrix[ni][nj] != '|' {
		name = append(name, matrix[ni][nj])
		nj++
	}

	children := parseMatrix(matrix, i+1, j+1, ib, lj)
	b := &box{
		Name: string(name),
	}

	if len(children) > 0 {
		b.Children = children
	} else {
		b.Area = (lj - j - 1) * (ib - i - 1)
	}

	return b
}

func writeToMap(boxes []*box) map[string]any {
	res := make(map[string]any, len(boxes))
	for _, b := range boxes {
		b.writeToMap(res)
	}
	return res
}
