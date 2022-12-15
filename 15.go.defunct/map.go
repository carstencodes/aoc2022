package main

import (
	"fmt"
)

type cave_map struct {
	center point
	size   int
	outer  point
	glyphs []string
}

func (cm *cave_map) build() {
	len := 1
	lines := []string{}
	min_y := cm.center.y - cm.size
	min_x := cm.center.x - cm.size
	for y := min_y; y <= cm.center.y+cm.size; y++ {
		line := ""
		for x := min_x; x <= cm.center.x+cm.size; x++ {
			min_glyph_x := cm.center.x - len/2
			max_glyph_x := cm.center.x + len/2
			glyph := "."
			if x > min_glyph_x && x < max_glyph_x {
				glyph = "#"
			}
			if x == cm.center.x && y == cm.center.y {
				glyph = "S"
			} else if x == cm.outer.x && y == cm.outer.y {
				glyph = "B"
			}

			line += glyph
		}
		if y < cm.center.y {
			len += 2
		} else {
			len -= 2
		}

		lines = append(lines, line)
	}

	cm.glyphs = lines
}

func (cm cave_map) draw() {
	text := fmt.Sprintf("S: (%d, %d) B: (%d, %d), D: %d\n", cm.center.x, cm.center.y, cm.outer.x, cm.outer.y, cm.size)

	for _, line := range cm.glyphs {
		text += line + "\n"
	}

	println(text)
}
