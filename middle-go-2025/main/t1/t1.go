package t1

import (
	"bufio"
	"fmt"
	"os"
)

type lightPosition struct {
	x, y int
	d    string
}

func Task1() {
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
		positions, err := getLightsPositions(in)
		if err != nil {
			panic(fmt.Errorf("failed to get light for #%d: %w", i, err))
		}
		if _, err = fmt.Fprintln(out, len(positions)); err != nil {
			panic(fmt.Errorf("failed to write lights count for #%d: %w", i, err))
		}
		for _, p := range positions {
			if _, err = fmt.Fprintf(out, "%d %d %s\n", p.x, p.y, p.d); err != nil {
				panic(fmt.Errorf("failed to write light for #%d: %w", i, err))
			}
		}
	}
}

func getLightsPositions(in *bufio.Reader) ([]lightPosition, error) {
	var n, m int
	if _, err := fmt.Fscan(in, &n, &m); err != nil {
		return nil, fmt.Errorf("failed to read n and m: %w", err)
	}

	if m > n {
		return getHorizontalLightsPositions(n, m), nil
	}
	return getVerticalLightsPositions(n, m), nil
}

func getVerticalLightsPositions(n, m int) []lightPosition {
	res := make([]lightPosition, 0, 2)
	res = append(res, lightPosition{1, 1, "D"})
	if m == 1 {
		return res
	}
	return append(res, lightPosition{n, m, "U"})
}

func getHorizontalLightsPositions(n, m int) []lightPosition {
	res := make([]lightPosition, 0, 2)
	res = append(res, lightPosition{1, 1, "R"})
	if n == 1 {
		return res
	}
	return append(res, lightPosition{n, m, "L"})
}
