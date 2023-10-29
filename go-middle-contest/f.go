package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type InputCategory struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Parent *int   `json:"parent,omitempty"`
}

type CategoryTree struct {
	Id   int             `json:"id"`
	Name string          `json:"name"`
	Next []*CategoryTree `json:"next,omitempty"`
}

func buildCategoryTree(in *bufio.Reader, buff bytes.Buffer) CategoryTree {
	buff.Reset()
	var rowsCount int
	var row string
	_, _ = fmt.Fscanf(in, "%d\n", &rowsCount)
	for i := 0; i < rowsCount; i++ {
		row, _ = in.ReadString('\n')
		buff.WriteString(row)
	}

	var categories []InputCategory
	_ = json.Unmarshal(buff.Bytes(), &categories)

	categoriesTreeMap := make(map[int]*CategoryTree)
	for _, category := range categories {
		categoriesTreeMap[category.Id] = &CategoryTree{
			Id:   category.Id,
			Name: category.Name,
			Next: make([]*CategoryTree, 0, 400),
		}
	}

	for _, category := range categories {
		if category.Parent == nil {
			continue
		}
		categoriesTreeMap[*category.Parent].Next = append(
			categoriesTreeMap[*category.Parent].Next,
			categoriesTreeMap[category.Id],
		)
	}

	return *categoriesTreeMap[0]
}

func F() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setsCount int
	_, _ = fmt.Fscanf(in, "%d\n", &setsCount)

	result := make([]CategoryTree, 0, setsCount)
	buff := bytes.Buffer{}
	buff.Grow(100000)
	for i := 0; i < setsCount; i++ {
		result = append(result, buildCategoryTree(in, buff))
	}

	writer := json.NewEncoder(out)
	_ = writer.Encode(result)
}

func main() {
	F()
}
