package main

import (
	"bufio"
	"fmt"
	"os"
)

func eClearMap(m map[int]int) {
	for k := range m {
		delete(m, k)
	}
}

func eFindAnswerForSet(in *bufio.Reader, out *bufio.Writer, taskIdToLastDayMap map[int]int) {
	eClearMap(taskIdToLastDayMap)
	var n int
	_, _ = fmt.Fscan(in, &n)

	var taskId int
	res := true
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(in, &taskId)
		if !res {
			continue
		}
		prevTaskLastDay, ok := taskIdToLastDayMap[taskId]
		if !ok || prevTaskLastDay+1 == i {
			taskIdToLastDayMap[taskId] = i
		} else {
			res = false
		}
	}

	if res {
		_, _ = fmt.Fprintln(out, "YES")
	} else {
		_, _ = fmt.Fprintln(out, "NO")
	}
}

func E() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setsCount int
	_, _ = fmt.Fscan(in, &setsCount)

	taskIdToLastDayMap := make(map[int]int)
	for i := 0; i < setsCount; i++ {
		eFindAnswerForSet(in, out, taskIdToLastDayMap)
	}
}

func main() {
	E()
}
