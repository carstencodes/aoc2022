package main

import (
	"fmt"
	"strings"
)

func compare(left signal, right signal) (bool, bool) {
	outcome, same := left.is_in_correct_order_to(right)
	fmt.Printf("- %s vs. %s\n", left.to_string(), right.to_string())
	fmt.Printf("   %t-%t\n", outcome, same)
	return outcome, same
}

func make_signal_pair(left, right string) (*signal_pair, error) {
	signal_left, err_left := create_line_parser(left).parse_signal()
	signal_right, err_right := create_line_parser(right).parse_signal()

	if err_left != nil {
		return nil, err_left
	}
	if err_right != nil {
		return nil, err_right
	}

	return &signal_pair{signal_left, signal_right}, nil
}

func make_list(data string) (*signal_list, error) {
	lines := strings.Split(data, "\n")

	signals := []signal_pair{}

	for i := 0; i < len(lines); i += 3 {
		current := i
		next := i + 1

		current_line := lines[current]
		next_line := lines[next]

		pair, err := make_signal_pair(current_line, next_line)
		if err != nil {
			return nil, err
		}

		signals = append(signals, *pair)
	}

	return &signal_list{signals}, nil
}
