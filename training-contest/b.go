package main

import (
	"bufio"
	"fmt"
	"os"
)

func bClearMap(m map[int]int) {
	for k := range m {
		delete(m, k)
	}
}

func bFindAnswerForSet(in *bufio.Reader, out *bufio.Writer, priceToCountMap map[int]int) {
	bClearMap(priceToCountMap)

	var itemsCount int
	_, _ = fmt.Fscan(in, &itemsCount)

	var itemPrice int
	for i := 0; i < itemsCount; i++ {
		_, _ = fmt.Fscan(in, &itemPrice)
		priceToCountMap[itemPrice] += 1
	}

	totalPrice := 0
	for price, count := range priceToCountMap {
		totalPrice += price * (count/3*2 + count%3)
	}
	_, _ = fmt.Fprintln(out, totalPrice)
}

func B() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setsCount int
	_, _ = fmt.Fscan(in, &setsCount)

	priceToCountMap := make(map[int]int)
	for i := 0; i < setsCount; i++ {
		bFindAnswerForSet(in, out, priceToCountMap)
	}
}

func main() {
	B()
}
