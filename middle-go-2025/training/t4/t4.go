package t4

import (
	"bufio"
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

	slices.SortFunc(q.machines, func(a, b machineAvailability) int {
		if a.start == b.start {
			return a.ind - b.ind
		}
		return a.start - b.start
	})

	ordersInd := 0
	for i, machine := range q.machines {
		for ordersInd < len(q.orders) && q.orders[ordersInd].arrivalTime < machine.start {
			s.arrivalsSchedule[q.orders[ordersInd].i] = -1
			ordersInd++
		}

		for ordersInd < len(q.orders) && q.orders[ordersInd].arrivalTime <= machine.end && q.machines[i].capacity > 0 {
			s.arrivalsSchedule[q.orders[ordersInd].i] = machine.ind + 1
			q.machines[i].capacity--
			ordersInd++
		}

		if ordersInd >= len(q.orders) {
			break
		}
	}

	for ordersInd < len(q.orders) {
		s.arrivalsSchedule[q.orders[ordersInd].i] = -1
		ordersInd++
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
