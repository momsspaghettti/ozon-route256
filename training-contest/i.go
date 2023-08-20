package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Processor struct {
	Consumption   int64
	AvailableTime int64
}

type MinHeap struct {
	Processors []*Processor
	less       func(f, s *Processor) bool
}

func (h *MinHeap) Len() int {
	return len(h.Processors)
}

func (h *MinHeap) Less(i, j int) bool {
	return h.less(h.Processors[i], h.Processors[j])
}

func (h *MinHeap) Swap(i, j int) {
	h.Processors[i], h.Processors[j] = h.Processors[j], h.Processors[i]
}

func (h *MinHeap) Push(x any) {
	h.Processors = append(h.Processors, x.(*Processor))
}

func (h *MinHeap) Pop() any {
	n := len(h.Processors)
	var res any
	res, h.Processors = h.Processors[n-1], h.Processors[:n-1]
	return res
}

func I() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	_, _ = fmt.Fscan(in, &n, &m)

	minConsumptionHeap := &MinHeap{
		Processors: make([]*Processor, 0, n),
		less: func(f, s *Processor) bool {
			return f.Consumption < s.Consumption
		},
	}
	minAvailableTimeHeap := &MinHeap{
		Processors: make([]*Processor, 0, n),
		less: func(f, s *Processor) bool {
			return f.AvailableTime < s.AvailableTime
		},
	}

	var consumption int64
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(in, &consumption)
		minConsumptionHeap.Processors = append(minConsumptionHeap.Processors, &Processor{Consumption: consumption})
	}
	heap.Init(minConsumptionHeap)

	var totalUsedEnergy int64
	var taskStart, taskDuration int64
	for j := 0; j < m; j++ {
		_, _ = fmt.Fscan(in, &taskStart, &taskDuration)
		for minConsumptionHeap.Len() > 0 && minConsumptionHeap.Processors[0].AvailableTime > taskStart {
			heap.Push(minAvailableTimeHeap, heap.Pop(minConsumptionHeap))
		}

		for minAvailableTimeHeap.Len() > 0 && minAvailableTimeHeap.Processors[0].AvailableTime <= taskStart {
			heap.Push(minConsumptionHeap, heap.Pop(minAvailableTimeHeap))
		}

		if minConsumptionHeap.Len() == 0 {
			continue
		}

		processor := (heap.Pop(minConsumptionHeap)).(*Processor)
		totalUsedEnergy += taskDuration * processor.Consumption
		processor.AvailableTime = taskStart + taskDuration
		heap.Push(minAvailableTimeHeap, processor)
	}

	_, _ = fmt.Fprintln(out, totalUsedEnergy)
}

func main() {
	I()
}
