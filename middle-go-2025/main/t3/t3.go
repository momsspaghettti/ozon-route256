package t3

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var (
	h1Sb = &bytes.Buffer{}
	h2Sb = &bytes.Buffer{}
)

func Task3() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func() {
		_ = out.Flush()
	}()

	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		panic(fmt.Errorf("failed to read t: %w", err))
	}

	h1Sb.Grow(1_000_000)
	h2Sb.Grow(1_000_000)

	for i := range t {
		pairsCount, err := countSimilarWordsPairs(in)
		if err != nil {
			panic(fmt.Errorf("failed to count pairs for #%d: %w", i, err))
		}
		if _, err = fmt.Fprintln(out, pairsCount); err != nil {
			panic(fmt.Errorf("failed to write pairs count for #%d: %w", i, err))
		}
	}
}

func countSimilarWordsPairs(in *bufio.Reader) (int, error) {
	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return 0, fmt.Errorf("failed to read n: %w", err)
	}

	_, _ = in.ReadBytes('\n')

	fullWordsCountMap := make(map[string]int, n)
	h1WordsCountMap := make(map[string]int, n)
	h2WordsCountMap := make(map[string]int, n)

	var (
		word []byte
		err  error
	)
	res := 0
	for i := range n {
		word, err = in.ReadBytes('\n')
		if err != nil {
			return 0, fmt.Errorf("failed to read word #%d: %w", i, err)
		}
		word = word[:len(word)-1]

		h1Sb.Reset()
		h2Sb.Reset()
		for j := range word {
			if j%2 == 0 {
				h1Sb.WriteByte(word[j])
			} else {
				h2Sb.WriteByte(word[j])
			}
		}

		fullWord := string(word)
		h1Word := h1Sb.String()
		h2Word := h2Sb.String()

		c1 := h1WordsCountMap[h1Word]
		c2 := h2WordsCountMap[h2Word]
		c := fullWordsCountMap[fullWord]

		h1WordsCountMap[h1Word]++
		if len(h2Word) > 0 {
			h2WordsCountMap[h2Word]++
		}
		fullWordsCountMap[fullWord]++

		res += c1 + c2
		if c2 != 0 {
			res -= c
		}
	}

	return res, nil
}
