package main

import "fmt"

type signal interface {
	is_in_correct_order_to(another signal) (bool, bool)
	length() int
	to_string() string
}

type value_signal struct {
	value int
}

type list_signal struct {
	signals []signal
}

type empty_signal struct {
}

func (e empty_signal) is_in_correct_order_to(another signal) (bool, bool) {
	switch another.(type) {
	case empty_signal:
		{
			return true, true
		}
	default:
		{
			return another.length() > 0, false // left side ran out of items
		}
	}
}

func (e empty_signal) length() int {
	return 0
}

func (e empty_signal) to_string() string {
	return "[]"
}

func (v value_signal) is_in_correct_order_to(another signal) (bool, bool) {
	switch x := another.(type) {
	case empty_signal:
		{
			return false, false // right side ran out of items
		}
	case list_signal:
		{
			return compare(list_signal{[]signal{v}}, (another))
		}
	case value_signal:
		{
			return v.value <= x.value, v.value == x.value
		}
	}

	return false, false // Invalid signal
}

func (v value_signal) length() int {
	return 1
}

func (v value_signal) to_string() string {
	return fmt.Sprintf("%d", v.value)
}

func (l list_signal) is_in_correct_order_to(another signal) (bool, bool) {
	switch x := another.(type) {
	case empty_signal:
		{
			return false, false // right side ran out of items
		}
	case value_signal:
		{
			if len(l.signals) > 1 {
				return compare(l.signals[0], list_signal{[]signal{x}})
			}
			if len(l.signals) == 0 {
				return true, false // left side ran out of items
			}

			return compare(l.signals[0], another)
		}
	case list_signal:
		{
			result := false
			same := false
			for i, item := range l.signals {
				if i >= len(x.signals) {
					return false, false // Right side ran out of items
				}
				result, same = compare(item, x.signals[i])
				if result {
					if same {
						continue
					}
				} else {
					return false, false
				}
			}

			return result, same
		}
	}

	return false, false // invalid signal
}

func (l list_signal) length() int {
	return len(l.signals)
}

func (l list_signal) to_string() string {
	value := "["

	for idx, item := range l.signals {
		if idx > 0 {
			value += ", "
		}

		value += item.to_string()
	}

	value += "]"

	return value
}
