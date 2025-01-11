package t1

import (
	"bufio"
	"fmt"
	"os"
)

func Task1() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func() {
		_ = out.Flush()
	}()

	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		panic(fmt.Errorf("failed to read t: %w", err))
	}

	dr := &digitRemover{}

	for i := range t {
		var num string
		if _, err := fmt.Fscan(in, &num); err != nil {
			panic(fmt.Errorf("failed to read num #%d: %w", i, err))
		}

		resNum := dr.RemoveOneDigit(num)
		if _, err := fmt.Fprintln(out, resNum); err != nil {
			panic(fmt.Errorf("failed to write num #%d answer '%s': %w", i, resNum, err))
		}
	}
}

type digitRemover struct {
	buff []byte
}

func (d *digitRemover) RemoveOneDigit(num string) string {
	if len(num) < 2 {
		return "0"
	}

	d.prepare(num)

	found := false
	for i := 0; i+1 < len(num); i++ {
		if found {
			d.buff = append(d.buff, num[i])
			continue
		}

		if int(num[i]-'0') < int(num[i+1]-'0') {
			found = true
		} else {
			d.buff = append(d.buff, num[i])
		}
	}

	if found {
		d.buff = append(d.buff, num[len(num)-1])
	}

	return string(d.buff)
}

func (d *digitRemover) prepare(num string) {
	if cap(d.buff) < len(num) {
		d.buff = make([]byte, 0, len(num))
	} else {
		d.buff = d.buff[:0]
	}
}
