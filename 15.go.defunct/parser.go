package main

import (
	"strconv"
	"strings"
)

type parser struct {
	input string
}

func (p parser) parse() ([]measurement, error) {
	measurements := []measurement{}

	lines := strings.Split(p.input, "\n")
	for _, line := range lines {
		var sensor_data string
		var beacon_data string

		replaced_data := strings.ReplaceAll(line, "Sensor at ", "")
		replaced_data = strings.ReplaceAll(replaced_data, " closest beacon is at ", "")
		data := strings.Split(replaced_data, ":")
		sensor_data = data[0]
		beacon_data = data[1]

		sensor, errSensor := p.parse_point(sensor_data)
		beacon, errBeacon := p.parse_point(beacon_data)

		if errSensor != nil {
			return nil, errSensor
		}

		if errBeacon != nil {
			return nil, errBeacon
		}

		measurement := measurement{*sensor, *beacon}
		measurements = append(measurements, measurement)
	}

	return measurements, nil
}

func (p parser) parse_point(point_data string) (*point, error) {
	x_y := strings.Split(point_data, ", ")
	x_assignment := x_y[0]
	y_assignment := x_y[1]

	x_assignment = strings.Trim(x_assignment, " \n\t\r")
	y_assignment = strings.Trim(y_assignment, " \n\t\r")

	x_value := strings.ReplaceAll(x_assignment, "x=", "")
	y_value := strings.ReplaceAll(y_assignment, "y=", "")

	x, errX := strconv.Atoi(x_value)
	y, errY := strconv.Atoi(y_value)

	if errX != nil {
		return nil, errX
	}

	if (errY) != nil {
		return nil, errY
	}

	return &point{x, y}, nil
}
