package t4

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"slices"
)

func Task4() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func() {
		_ = out.Flush()
	}()

	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		panic(fmt.Errorf("failed to read t: %w", err))
	}

	qr := &queryReader{}
	s := &ordersScheduler{}

	for i := range t {
		q, err := qr.Read(in)
		if err != nil {
			panic(fmt.Errorf("failed to read query #%d: %w", i, err))
		}

		schedule := s.GetSchedule(q)
		if err = writeSchedule(out, schedule); err != nil {
			panic(fmt.Errorf("failed to write schedule for query #%d: %w", i, err))
		}
	}
}

func writeSchedule(out io.Writer, s []int) error {
	for i := range s {
		if _, err := fmt.Fprint(out, s[i]); err != nil {
			return err
		}
		if i+1 < len(s) {
			if _, err := fmt.Fprint(out, " "); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fprintln(out); err != nil {
		return err
	}
	return nil
}

type order struct {
	i, arrivalTime int
}

type machineAvailability struct {
	ind, start, end, capacity int
}

type minStartQueue struct {
	machines []*machineAvailability
}

func (q *minStartQueue) Len() int {
	return len(q.machines)
}

func (q *minStartQueue) Less(i, j int) bool {
	if q.machines[i].start == q.machines[j].start {
		return q.machines[i].ind < q.machines[j].ind
	}
	return q.machines[i].start < q.machines[j].start
}

func (q *minStartQueue) Swap(i, j int) {
	q.machines[i], q.machines[j] = q.machines[j], q.machines[i]
}

func (q *minStartQueue) Push(x any) {
	m := x.(*machineAvailability)
	q.machines = append(q.machines, m)
}

func (q *minStartQueue) Pop() any {
	m := q.machines[len(q.machines)-1]
	q.machines = q.machines[:len(q.machines)-1]
	return m
}

type maxEndQueue struct {
	machines []*machineAvailability
}

func (q *maxEndQueue) Len() int {
	return len(q.machines)
}

func (q *maxEndQueue) Less(i, j int) bool {
	return q.machines[i].end > q.machines[j].end
}

func (q *maxEndQueue) Swap(i, j int) {
	q.machines[i], q.machines[j] = q.machines[j], q.machines[i]
}

func (q *maxEndQueue) Push(x any) {
	m := x.(*machineAvailability)
	q.machines = append(q.machines, m)
}

func (q *maxEndQueue) Pop() any {
	m := q.machines[len(q.machines)-1]
	q.machines = q.machines[:len(q.machines)-1]
	return m
}

type query struct {
	orders   []order
	machines []machineAvailability
}

type ordersScheduler struct {
	arrivalsSchedule []int
}

func (s *ordersScheduler) GetSchedule(q *query) []int {
	s.prepare(len(q.orders))

	slices.SortFunc(q.orders, func(a, b order) int {
		return a.arrivalTime - b.arrivalTime
	})

	minStart := &minStartQueue{machines: make([]*machineAvailability, 0, len(q.machines))}
	for i := range q.machines {
		minStart.machines = append(minStart.machines, &q.machines[i])
	}
	heap.Init(minStart)

	maxEnd := &maxEndQueue{machines: make([]*machineAvailability, 0, len(q.machines))}
	heap.Init(maxEnd)

	for _, o := range q.orders {
		found := false
		for maxEnd.Len() > 0 {
			if maxEnd.machines[0].end < o.arrivalTime {
				break
			}
			heap.Push(minStart, heap.Pop(maxEnd))
		}

		for minStart.Len() > 0 {
			if minStart.machines[0].start > o.arrivalTime {
				break
			}
			if minStart.machines[0].end >= o.arrivalTime {
				found = true
				s.arrivalsSchedule[o.i] = minStart.machines[0].ind + 1

				minStart.machines[0].capacity--
				if minStart.machines[0].capacity == 0 {
					heap.Pop(minStart)
				}

				break
			}
			heap.Push(maxEnd, heap.Pop(minStart))
		}

		if !found {
			s.arrivalsSchedule[o.i] = -1
		}
	}

	return s.arrivalsSchedule
}

func (s *ordersScheduler) prepare(n int) {
	for len(s.arrivalsSchedule) < n {
		s.arrivalsSchedule = append(s.arrivalsSchedule, 0)
	}
	s.arrivalsSchedule = s.arrivalsSchedule[:n]
}

type queryReader struct {
	ordersBuff   []order
	machinesBuff []machineAvailability
}

func (q *queryReader) Read(in *bufio.Reader) (*query, error) {
	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return nil, fmt.Errorf("failed to read n: %w", err)
	}
	q.prepareOrders(n)

	for i := range n {
		var arrivalTime int
		if _, err := fmt.Fscan(in, &arrivalTime); err != nil {
			return nil, fmt.Errorf("failed to read arrival time #%d: %w", i, err)
		}
		q.ordersBuff = append(q.ordersBuff, order{i: i, arrivalTime: arrivalTime})
	}

	var m int
	if _, err := fmt.Fscan(in, &m); err != nil {
		return nil, fmt.Errorf("failed to read m: %w", err)
	}
	q.prepareMachines(m)

	for i := range m {
		var start, end, capacity int
		if _, err := fmt.Fscan(in, &start, &end, &capacity); err != nil {
			return nil, fmt.Errorf("failed to read machines line #%d: %w", i, err)
		}
		q.machinesBuff = append(
			q.machinesBuff,
			machineAvailability{ind: i, start: start, end: end, capacity: capacity},
		)
	}

	return &query{orders: q.ordersBuff, machines: q.machinesBuff}, nil
}

func (q *queryReader) prepareOrders(n int) {
	if cap(q.ordersBuff) < n {
		q.ordersBuff = make([]order, 0, n)
	} else {
		q.ordersBuff = q.ordersBuff[:0]
	}
}

func (q *queryReader) prepareMachines(m int) {
	if cap(q.machinesBuff) < m {
		q.machinesBuff = make([]machineAvailability, 0, m)
	} else {
		q.machinesBuff = q.machinesBuff[:0]
	}
}
