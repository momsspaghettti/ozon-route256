package main

import (
	"bufio"
	"fmt"
	"os"
)

func B() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var originalSticker string
	_, _ = fmt.Fscan(in, &originalSticker)
	originalStickerChars := []byte(originalSticker)

	var patchesCount int
	_, _ = fmt.Fscan(in, &patchesCount)

	var from, to int
	var patch string
	for i := 0; i < patchesCount; i++ {
		_, _ = fmt.Fscan(in, &from, &to, &patch)
		for j := range patch {
			originalStickerChars[from-1+j] = patch[j]
		}
	}

	_, _ = fmt.Fprintln(out, string(originalStickerChars))
}

func main() {
	B()
}
