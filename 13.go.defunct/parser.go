package main

import (
	"errors"
	"strconv"
)

type line_parser struct {
	line         string
	cursor       int
	signal_stack stack_of_signals
}

func (p line_parser) until_end_of_line() string {
	return p.line[p.cursor:]
}

func (p line_parser) is_empty() bool {
	return p.line[p.cursor] == '[' && p.line[p.cursor+1] == ']'
}

func (p *line_parser) parse_empty() signal {
	p.cursor += 2
	return empty_signal{}
}

func (p line_parser) is_value() bool {
	return p.line[p.cursor] != '[' && p.line[p.cursor] != ']' && p.line[p.cursor] != ','
}

func (p *line_parser) parse_value() (signal, error) {
	marker := p.cursor

	for p.line[marker] != ',' && p.line[marker] != ']' {
		marker++
		if marker > len(p.line) {
			return nil, errors.New("line overflow")
		}
	}

	token := p.line[p.cursor:marker]

	value, err := strconv.Atoi(token)

	if err != nil {
		return nil, err
	}

	p.cursor = marker

	return value_signal{value}, nil
}

func (p line_parser) starts_list() bool {
	return p.line[p.cursor] == '['
}

func (p *line_parser) begin_parse_list() {
	p.cursor++
	p.signal_stack.push()
}

func (p *line_parser) parse_list() (signal, error) {
	for {
		current, err := p.parse_signal()
		if err != nil {
			return nil, err
		}
		p.signal_stack.append(current)
		if p.line[p.cursor] == ']' {
			break
		}
	}

	return list_signal{p.signal_stack.list_of_signals}, nil
}

func (p *line_parser) end_parse_list() error {
	if p.line[p.cursor] != ']' {
		return errors.New("Not at the end of a list: " + p.until_end_of_line())
	}
	p.cursor++
	return p.signal_stack.pop()
}

func (p *line_parser) jump_to_next_list() {
	if p.cursor < len(p.line) {
		if p.line[p.cursor] == ',' {
			p.cursor++
		}
	}
}

func (p *line_parser) parse_signal() (signal, error) {
	defer p.jump_to_next_list()
	if p.is_empty() {
		parsed := p.parse_empty()

		return parsed, nil
	}
	if p.starts_list() {
		p.begin_parse_list()
		parsed, err := p.parse_list()
		if err != nil {
			return nil, err
		}
		err = p.end_parse_list()
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	if p.is_value() {
		parsed, err := p.parse_value()
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}

	return nil, errors.New("failed to parse at current position: " + p.until_end_of_line())
}
