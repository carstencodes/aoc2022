package main

import (
	"math"
	"strings"
)

type map_of_maps struct {
	maps   []cave_map
	glyphs []string
}

func (mom *map_of_maps) build() {
	header := mom.get_header()
	glyphs := []string{header}

	var (
		min_x int = 0
		max_x int = 0
		min_y int = 0
		max_y int = 0
	)
	for _, cm := range mom.maps {
		min_x = int(math.Min(float64(min_x), float64(cm.center.x-cm.size)))
		max_x = int(math.Max(float64(max_x), float64(cm.center.x+cm.size)))
		min_y = int(math.Min(float64(min_y), float64(cm.center.y-cm.size)))
		max_y = int(math.Max(float64(max_y), float64(cm.center.y+cm.size)))
	}

	max_x++
	max_y++

	single_runes := [][]byte{}

	for y := min_y; y < max_y; y++ {
		runes := []byte{}
		for x := min_x; x < max_x; x++ {
			runes = append(runes, '.')
		}
		single_runes = append(single_runes, runes)
	}

	for _, cm := range mom.maps {
		single_runes = transfer_map_to_runes(cm, single_runes, min_x, min_y)
	}

	for _, rune_line := range single_runes {
		builder := new(strings.Builder)
		for _, item := range rune_line {
			_ = builder.WriteByte(item)
		}
		line := builder.String()
		glyphs = append(glyphs, line)
	}

	mom.glyphs = glyphs
}

func (mom map_of_maps) draw() {
	println(mom.to_string())
}

func (mom map_of_maps) to_string() string {
	text := ""

	for _, line := range mom.glyphs {
		text += line + "\n"
	}

	return text
}

func (mom map_of_maps) get_header() string {
	text := ""
	var (
		min_x int = 0
		max_x int = 0
		min_y int = 0
		max_y int = 0
	)
	for _, cm := range mom.maps {
		min_x = int(math.Min(float64(min_x), float64(cm.center.x-cm.size)))
		max_x = int(math.Max(float64(max_x), float64(cm.center.x+cm.size)))
		min_y = int(math.Min(float64(min_y), float64(cm.center.y-cm.size)))
		max_y = int(math.Max(float64(max_y), float64(cm.center.y+cm.size)))
	}

	max_x++
	max_y++

	for x := min_x; x <= max_x; x++ {
		sym := " "
		if x%5 == 0 {
			sym = "5"
			if x%10 == 0 {
				sym = "0"
			} else if x < 0 {
				sym = "\b-" + sym
			}
		}
		text += sym
	}

	return text
}

func transfer_map_to_runes(cm cave_map, runes [][]byte, x_min int, y_min int) [][]byte {
	len := 1
	center_translated := point{cm.center.x - x_min, cm.center.y - y_min}
	outer_translated := point{cm.outer.x - x_min, cm.outer.y - y_min}
	translated_map := cave_map{center_translated, cm.size, outer_translated, []string{}}
	translated_map.build()

	min_y := translated_map.center.y - translated_map.size
	min_x := translated_map.center.x - translated_map.size
	for y := min_y; y <= translated_map.center.y+translated_map.size; y++ {
		for x := min_x; x <= translated_map.center.x+translated_map.size; x++ {
			min_glyph_x := translated_map.center.x - len/2
			max_glyph_x := translated_map.center.x + len/2
			glyph := runes[x][y]
			if glyph == '.' {
				if x > min_glyph_x && x < max_glyph_x {
					glyph = '#'
				}
			}
			if glyph == '.' || glyph == '#' {
				if x == translated_map.center.x && y == translated_map.center.y {
					glyph = 'S'
				} else if x == translated_map.outer.x && y == translated_map.outer.y {
					glyph = 'B'
				}
			}
			runes[x][y] = glyph
		}
		if y < cm.center.y {
			len += 2
		} else {
			len -= 2
		}
	}

	return runes
}
