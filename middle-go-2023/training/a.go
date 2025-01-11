package main

import (
	"bufio"
	"fmt"
	"os"
)

func A() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	_, _ = fmt.Fscan(in, &n)

	var a, b int
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(in, &a, &b)
		_, _ = fmt.Fprintln(out, a+b)
	}
}

func main() {
	A()
}
