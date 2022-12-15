package main

import (
	"sync"
)

func main() {
	data := SAMPLE

	p := parser{data}

	items, err := p.parse()

	if err != nil {
		panic(err)
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(items))

	maps := make(chan cave_map, len(items))

	for _, measurement := range items {
		go measurement.build_map(maps, wg)
	}

	wg.Wait()
	close(maps)

	built_maps := []cave_map{}
	for mp := range maps {
		built_maps = append(built_maps, mp)
	}

	mom := map_of_maps{built_maps, []string{}}
	mom.build()

	mom.draw()
}
