package main

import (
	"math"
	"sync"
)

type point struct {
	x int
	y int
}

type measurement struct {
	sensor point
	beacon point
}

func (m measurement) calculate_manhattan_distance() int {
	x_beacon := m.beacon.x
	y_beacon := m.beacon.y

	x_sensor := m.sensor.x
	y_sensor := m.sensor.y

	delta_x := x_beacon - x_sensor
	delta_x = int(math.Abs(float64(delta_x)))

	delta_y := y_beacon - y_sensor
	delta_y = int(math.Abs(float64(delta_y)))

	return delta_x + delta_y + 1
}

func (m measurement) build_map(channel chan<- cave_map, wg *sync.WaitGroup) {
	distance := m.calculate_manhattan_distance()

	cm := cave_map{m.sensor, distance, m.beacon, []string{}}
	cm.build()

	channel <- cm

	wg.Done()
}
