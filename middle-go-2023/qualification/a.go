package main

import (
	"bufio"
	"fmt"
	"os"
)

func A() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	_, _ = fmt.Fprintln(out, "I am sure that I will fill out the form by 23:59 on August 27, 2023.")
}

func main() {
	A()
}
