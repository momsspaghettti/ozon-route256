package t2

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func Task2() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func() {
		_ = out.Flush()
	}()

	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		panic(fmt.Errorf("failed to read t: %w", err))
	}

	v := &outputValidator{}

	for i := range t {
		valid, err := v.Validate(in)
		if err != nil {
			panic(fmt.Errorf("failed to validate #%d data set: %w", i, err))
		}

		var res string
		if valid {
			res = "yes"
		} else {
			res = "no"
		}
		if _, err = fmt.Fprintln(out, res); err != nil {
			panic(fmt.Errorf("failed to write #%d answer '%s': %w", i, res, err))
		}
	}
}

type outputValidator struct {
	inputBuff  []int
	outputBuff []byte
}

func (v *outputValidator) Validate(in *bufio.Reader) (bool, error) {
	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return false, fmt.Errorf("failed to read n: %w", err)
	}

	v.prepare(n)

	for i := range n {
		var num int
		if _, err := fmt.Fscan(in, &num); err != nil {
			return false, fmt.Errorf("failed to read num #%d: %w", i, err)
		}

		v.inputBuff = append(v.inputBuff, num)
	}

	slices.Sort(v.inputBuff)

	for i, num := range v.inputBuff {
		if i > 0 {
			v.outputBuff = append(v.outputBuff, ' ')
		}
		v.outputBuff = append(v.outputBuff, []byte(strconv.Itoa(num))...)
	}

	_, err := in.ReadBytes('\n')
	if err != nil {
		return false, fmt.Errorf("failed to read input new line: %w", err)
	}

	outputToValidate, err := in.ReadBytes('\n')
	if err != nil {
		return false, fmt.Errorf("failed to read output: %w", err)
	}
	if len(outputToValidate) > 0 {
		outputToValidate = outputToValidate[:len(outputToValidate)-1]
	}

	return slices.Equal(v.outputBuff, outputToValidate), nil
}

func (v *outputValidator) prepare(n int) {
	if cap(v.inputBuff) < n {
		v.inputBuff = make([]int, 0, n)
	} else {
		v.inputBuff = v.inputBuff[:0]
	}

	maxExpectedOutput := 11 * n
	if cap(v.outputBuff) < maxExpectedOutput {
		v.outputBuff = make([]byte, 0, maxExpectedOutput)
	} else {
		v.outputBuff = v.outputBuff[:0]
	}
}
