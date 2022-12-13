package main

import (
	"errors"
	"fmt"
)

type stack_of_signals struct {
	list_of_signals []signal
	previous        *stack_of_signals
}

type signal_list struct {
	items []signal_pair
}

type signal_pair struct {
	left_side  signal
	right_side signal
}

func (s *stack_of_signals) push() {
	new_stack := stack_of_signals{s.list_of_signals, s.previous}
	s.previous = &new_stack
	s.list_of_signals = []signal{}
}

func (s *stack_of_signals) append(sgn signal) {
	s.list_of_signals = append(s.list_of_signals, sgn)
}

func (s *stack_of_signals) pop() error {
	if s.previous == nil {
		return errors.New("no items left in stack")
	}

	s.list_of_signals = s.previous.list_of_signals
	s.previous = s.previous.previous
	return nil
}

func (s signal_pair) is_in_correct_order() bool {
	result, _ := compare(s.left_side, s.right_side)
	return result
}

func create_signal_stack() stack_of_signals {
	return stack_of_signals{nil, nil}
}

func create_line_parser(line string) *line_parser {
	return &line_parser{line, 0, create_signal_stack()}
}

func (s signal_list) to_string() string {
	value := ""

	for idx, signal_item := range s.items {
		value += fmt.Sprintf("\n%d:\n", idx)
		value += signal_item.to_string()
		value += "\n"
	}

	return value
}

func (s signal_pair) to_string() string {
	return fmt.Sprintf("%s\n%s", s.left_side.to_string(), s.right_side.to_string())
}
