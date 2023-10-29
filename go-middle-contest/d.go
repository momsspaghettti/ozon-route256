package main

import (
	"bufio"
	"fmt"
	"os"
)

func dFindAnswerForSet(in *bufio.Reader, out *bufio.Writer) {
	var mountainsCount, n, m int
	_, _ = fmt.Fscan(in, &mountainsCount, &n, &m)

	var dotChar byte = '.'

	image := make([][]byte, n)
	for i := 0; i < n; i++ {
		image[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			image[i][j] = dotChar
		}
	}

	var imageRow string
	var i, j int
	for k := 0; k < mountainsCount; k++ {
		for i = 0; i < n; i++ {
			_, _ = fmt.Fscan(in, &imageRow)
			for j = 0; j < m; j++ {
				if image[i][j] == dotChar {
					image[i][j] = imageRow[j]
				}
			}
		}
		if k+1 != mountainsCount {
			_, _ = fmt.Fscan(in)
		}
	}

	for i = 0; i < n; i++ {
		_, _ = fmt.Fprintln(out, string(image[i]))
	}

	_, _ = fmt.Fprintln(out)
}

func D() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setsCount int
	_, _ = fmt.Fscan(in, &setsCount)

	for i := 0; i < setsCount; i++ {
		dFindAnswerForSet(in, out)
	}
}

func main() {
	D()
}
