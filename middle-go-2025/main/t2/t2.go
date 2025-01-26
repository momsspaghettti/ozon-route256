package t2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	for i := range t {
		valid, err := validateInput(in)
		if err != nil {
			panic(fmt.Errorf("failed to validate input #%d: %w", i, err))
		}

		var output string
		if valid {
			output = "YES"
		} else {
			output = "NO"
		}

		if _, err := fmt.Fprintln(out, output); err != nil {
			panic(fmt.Errorf("failed to write output #%d: %w", i, err))
		}
	}
}

func validateInput(in *bufio.Reader) (bool, error) {
	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return false, fmt.Errorf("failed to read n: %w", err)
	}

	nameToPrise := make(map[string]int, n)
	var (
		name  string
		price int
	)
	for i := range n {
		if _, err := fmt.Fscan(in, &name, &price); err != nil {
			return false, fmt.Errorf("failed to read name and price #%d: %w", i, err)
		}
		nameToPrise[name] = price
	}

	_, _ = in.ReadBytes('\n')

	inputToValidate, err := in.ReadString('\n')
	if err != nil {
		return false, fmt.Errorf("failed to read input to validate: %w", err)
	}

	namesPrices := strings.Split(strings.TrimRight(inputToValidate, "\n"), ",")
	seenPrices := make(map[int]struct{}, len(namesPrices))
	for _, namePrice := range namesPrices {
		namePriceSlice := strings.Split(namePrice, ":")
		if len(namePriceSlice) != 2 {
			return false, nil
		}

		name = namePriceSlice[0]
		if !validateName(name) {
			return false, nil
		}

		price, valid := validatePrice(namePriceSlice[1])
		if !valid {
			return false, nil
		}

		expectedPrice, ok := nameToPrise[name]
		if !ok || expectedPrice != price {
			return false, nil
		}

		_, seenPrice := seenPrices[price]
		if seenPrice {
			return false, nil
		}

		seenPrices[price] = struct{}{}
	}

	for _, price := range nameToPrise {
		if _, seen := seenPrices[price]; !seen {
			return false, nil
		}
	}

	return true, nil
}

func validateName(name string) bool {
	if len(name) == 0 || len(name) > 10 {
		return false
	}

	for i := range len(name) {
		if name[i] < 'a' || name[i] > 'z' {
			return false
		}
	}

	return true
}

func validatePrice(price string) (int, bool) {
	priceInt, pErr := strconv.Atoi(price)
	if pErr != nil {
		return 0, false
	}

	if priceInt < 1 || priceInt > 1_000_000_000 {
		return 0, false
	}

	if price[0] == '0' {
		return 0, false
	}

	return priceInt, true
}
