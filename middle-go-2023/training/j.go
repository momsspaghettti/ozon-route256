package main

import (
	"bufio"
	"fmt"
	"os"
)

type PrefixTreeNode struct {
	Char        *byte
	Children    map[byte]*PrefixTreeNode
	isEndOfWord bool
}

func (node *PrefixTreeNode) AddWord(word []byte) {
	currNode := node
	for i, char := range word {
		charCopy := char
		child, ok := currNode.Children[char]
		if !ok {
			child = &PrefixTreeNode{
				Char:     &charCopy,
				Children: make(map[byte]*PrefixTreeNode),
			}
			currNode.Children[char] = child
		}
		if i+1 == len(word) {
			child.isEndOfWord = true
		} else {
			currNode = child
		}
	}
}

func reverse[T any](arr []T) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}

func getQueryAnswer(query string, prefixTree *PrefixTreeNode) string {
	queryChars := []byte(query)
	reverse(queryChars)

	answerNodes := make([]*PrefixTreeNode, 0, 10)
	currNode := prefixTree
	for _, char := range queryChars {
		child, ok := currNode.Children[char]
		if !ok {
			break
		}
		answerNodes = append(answerNodes, child)
		currNode = child
	}

	if len(answerNodes) < len(queryChars) {
		return getIncompletePrefixAnswer(answerNodes, currNode)
	}
	return getFullPrefixAnswer(answerNodes, prefixTree)
}

func getFullPrefixAnswer(answerNodes []*PrefixTreeNode, prefixTree *PrefixTreeNode) string {
	n := len(answerNodes)
	lastNode := answerNodes[n-1]
	if !lastNode.isEndOfWord {
		return getIncompletePrefixAnswer(answerNodes, lastNode)
	}

	if len(lastNode.Children) > 0 {
		for _, child := range lastNode.Children {
			answerNodes = append(answerNodes, child)
			break
		}
		return getIncompletePrefixAnswer(answerNodes, answerNodes[len(answerNodes)-1])
	}

	i := 1
	found := false
	for ; i < n; i++ {
		if answerNodes[n-1-i].isEndOfWord {
			return getAnswer(answerNodes[:n-i])
		}
		if len(answerNodes[n-1-i].Children) > 1 {
			for _, child := range answerNodes[n-1-i].Children {
				if child.Char != answerNodes[n-i].Char {
					answerNodes[n-i] = child
					break
				}
			}
			found = true
			break
		}
	}

	if found {
		return getIncompletePrefixAnswer(answerNodes[:n+1-i], answerNodes[n-i])
	}

	for _, child := range prefixTree.Children {
		if child.Char != answerNodes[0].Char {
			answerNodes[0] = child
			break
		}
	}

	return getIncompletePrefixAnswer(answerNodes[:1], answerNodes[0])
}

func getIncompletePrefixAnswer(answerNodes []*PrefixTreeNode, currNode *PrefixTreeNode) string {
	for !currNode.isEndOfWord {
		for _, child := range currNode.Children {
			answerNodes = append(answerNodes, child)
			currNode = child
			break
		}
	}
	return getAnswer(answerNodes)
}

func getAnswer(answerNodes []*PrefixTreeNode) string {
	n := len(answerNodes)
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = *answerNodes[n-1-i].Char
	}
	return string(res)
}

func J() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var wordsCount int
	_, _ = fmt.Fscan(in, &wordsCount)

	prefixTree := &PrefixTreeNode{Children: make(map[byte]*PrefixTreeNode)}
	var word string
	var wordChars []byte
	for i := 0; i < wordsCount; i++ {
		_, _ = fmt.Fscan(in, &word)
		wordChars = []byte(word)
		reverse(wordChars)
		prefixTree.AddWord(wordChars)
	}

	var queriesCount int
	_, _ = fmt.Fscan(in, &queriesCount)

	var query string
	for i := 0; i < queriesCount; i++ {
		_, _ = fmt.Fscan(in, &query)
		_, _ = fmt.Fprintln(out, getQueryAnswer(query, prefixTree))
	}
}

func main() {
	J()
}
