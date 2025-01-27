package t3

import (
	"bufio"
	"fmt"
	"os"
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

	fn := newPrefixNode()
	n1 := newPrefixNode()
	n2 := newPrefixNode()

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

		c1 := addWord(n1, word)
		c2 := addWord(n2, word[1:])
		c := addFullWord(fn, word)

		res += c1 + c2
		if c2 != 0 {
			res -= c
		}
	}

	return res, nil
}

type prefixNode struct {
	children   []*prefixNode
	wordsCount int
}

func newPrefixNode() *prefixNode {
	return &prefixNode{
		children: make([]*prefixNode, 26),
	}
}

func addWord(node *prefixNode, word []byte) int {
	if len(word) == 0 {
		return 0
	}
	for i := 0; i < len(word); i += 2 {
		ind := int(word[i] - 'a')
		cNode := node.children[ind]
		if cNode == nil {
			cNode = newPrefixNode()
			node.children[ind] = cNode
		}
		node = cNode
	}

	res := node.wordsCount
	node.wordsCount++

	return res
}

func addFullWord(node *prefixNode, word []byte) int {
	for i := range word {
		ind := int(word[i] - 'a')
		cNode := node.children[ind]
		if cNode == nil {
			cNode = newPrefixNode()
			node.children[ind] = cNode
		}
		node = cNode
	}

	res := node.wordsCount
	node.wordsCount++

	return res
}
