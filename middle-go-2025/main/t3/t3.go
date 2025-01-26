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

		n1WordsIndices := addWord(n1, word)
		n2WordsIndices := addWord(n2, word[1:])

		res += len(n1WordsIndices) + len(n2WordsIndices) - getIntersectionLen(n1WordsIndices, n2WordsIndices)

		n1WordsIndices[i] = struct{}{}
		if n2WordsIndices != nil {
			n2WordsIndices[i] = struct{}{}
		}
	}

	return res, nil
}

type prefixNode struct {
	children     []*prefixNode
	wordsIndices map[int]struct{}
}

func newPrefixNode() *prefixNode {
	return &prefixNode{
		children:     make([]*prefixNode, 26),
		wordsIndices: make(map[int]struct{}),
	}
}

func addWord(node *prefixNode, word []byte) map[int]struct{} {
	if len(word) == 0 {
		return nil
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
	return node.wordsIndices
}

func getIntersectionLen(a, b map[int]struct{}) int {
	var min_, max_ map[int]struct{}
	if len(a) < len(b) {
		min_, max_ = a, b
	} else {
		min_, max_ = b, a
	}

	res := 0
	for i := range min_ {
		if _, ok := max_[i]; ok {
			res++
		}
	}

	return res
}
