package main

import (
	"fmt"
	"strings"
	"sync"
)

type point struct {
	row_id int
	col_id int
}

type field struct {
	start   point
	end     point
	symbols map[int][]byte
	points  []point
}

type route struct {
	current point
	next    *route
	end     *point
}

func is_in_matrix(p point, matrix map[int][]byte) bool {
	row_id := p.row_id
	col_id := p.col_id

	return (row_id >= 0) && (row_id < len(matrix)) &&
		(col_id >= 0) && col_id < len(matrix[row_id])
}

func is_reachable(from point, to point, matrix map[int][]byte) bool {
	if is_in_matrix(from, matrix) && is_in_matrix(to, matrix) {
		symbol_from := matrix[from.row_id][from.col_id]
		symbol_to := matrix[to.row_id][to.col_id]

		if symbol_from > symbol_to {
			return false
		}

		delta := symbol_to - symbol_from

		return delta == 0 || delta == 1
	}

	return false
}

func (r route) to_slice() []point {
	points := make([]point, 0)
	points = append(points, r.current)
	if r.next == nil {
		return points
	}

	return append(points, r.next.to_slice()...)
}

func (r route) to_str() string {
	if r.next != nil {
		return fmt.Sprintf("%s -> %s", r.current.to_str(), r.next.to_str())
	} else {
		return r.current.to_str()
	}
}

func (p point) to_str() string {
	return fmt.Sprintf("(%d, %d)", p.row_id, p.col_id)
}

func (r route) determine_end() *point {
	if r.end != nil {
		return r.end
	}

	if r.next != nil {
		r.end = r.next.determine_end()
		return r.end
	}

	r.end = &r.current
	return r.end
}

func (r route) ends_at(p point) bool {
	return p.equals(*r.determine_end())
}

func (r route) clone() route {
	if r.next == nil {
		return route{r.current, nil, nil}
	}

	clone := r.next.clone()
	return route{r.current, &clone, clone.end}
}

func (p point) equals(another point) bool {
	return p.row_id == another.row_id && p.col_id == another.col_id
}

func contains(haystack []point, needle point) bool {
	for index := range haystack {
		item := haystack[index]
		if item.equals(needle) {
			return true
		}
	}

	return false
}

func (data field) get_neighbors(p point) []point {
	p_up := point{p.row_id - 1, p.col_id}
	p_down := point{p.row_id + 1, p.col_id}
	p_left := point{p.row_id, p.col_id - 1}
	p_right := point{p.row_id, p.col_id + 1}

	result := make([]point, 0)
	if p_up.row_id >= 0 {
		result = append(result, p_up)
	}
	if p_left.col_id >= 0 {
		result = append(result, p_left)
	}
	if p_down.row_id < len(data.symbols) {
		result = append(result, p_down)
	}
	if p_right.col_id < len(data.symbols[p_right.row_id]) {
		result = append(result, p_right)
	}

	return result
}

func (f field) get_symbol(p point) byte {
	if is_in_matrix(p, f.symbols) {
		return f.symbols[p.row_id][p.col_id]
	}

	return 0
}

func get_routes(position point, visited []point, data field, send chan route, wg_in *sync.WaitGroup) {
	neighbors := data.get_neighbors(position)

	route_to_position := route{position, nil, &position}

	if len(neighbors) != 0 {
		visited_items := 0
		for _, item := range neighbors {
			if (!contains(visited, item)) && is_reachable(position, item, data.symbols) {
				visited_items++
			}
		}

		if visited_items != 0 {
			receive := make(chan route, len(data.points))
			wg_out := new(sync.WaitGroup)
			for _, item := range neighbors {
				if (!contains(visited, item)) && is_reachable(position, item, data.symbols) {
					wg_out.Add(1)
					new_visited := append(visited, position)

					go get_routes(item, new_visited, data, receive, wg_out)
				}
			}

			wg_out.Wait()
			close(receive)

			for new_route := range receive {
				if new_route.ends_at(data.end) {
					next_route := route{position, &new_route, new_route.determine_end()}
					send <- next_route
				} else {
					send <- route_to_position
				}
			}
		} else {
			send <- route_to_position
		}
	} else {
		send <- route_to_position
	}

	wg_in.Done()
}

func (data field) find_route() route {
	start := data.start
	visited := make([]point, 0)

	cn := make(chan route, len(data.points))

	wg_out := new(sync.WaitGroup)
	wg_out.Add(1)
	go get_routes(start, visited, data, cn, wg_out)

	wg_out.Wait()
	close(cn)

	var shortest_route *route = nil
	for rt := range cn {
		if shortest_route != nil {
			if len(rt.to_slice()) < len(shortest_route.to_slice()) {
				shortest_route = &rt
			}
		} else {
			shortest_route = &rt
		}
	}

	if shortest_route == nil {
		return route{start, nil, nil}
	} else {
		return *shortest_route
	}
}

func make_field(data string) field {
	var lines = strings.Split(data, "\n")

	var matrix = make(map[int][]byte, 0)

	const start_symbol = 'S'
	const end_symbol = 'E'

	var start_position = point{-1, -1}
	var end_position = point{-1, -1}

	var unvisited = make([]point, 0)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		row := make([]byte, 0)
		for j := 0; j < len(line); j++ {
			chr := line[j]
			p := point{i, j}
			unvisited = append(unvisited, p)

			if chr == start_symbol {
				start_position = p
				chr = 'a'
			}
			if chr == end_symbol {
				end_position = p
				chr = 'z'
			}

			row = append(row, chr)
		}
		matrix[i] = row
	}

	var my_field = field{
		start_position,
		end_position,
		matrix,
		unvisited,
	}

	return my_field
}

const (
	sample = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	puzzle = `abccccccccccccccccccccccccccaaaaaaaaacccccccccccaaacccccccccccccccccccccccccaaaaaaaaccccccccaaaaaaccaaccccccccccccccccccccccccaaaaacaacaaaacccccccccccccccccccccccccccccccccccccaaaaa
	abccaaacccccccccccccccccccccaaaaaaaaacccccccccaaaaaacccccccccccccccccccccaaaaaaaaaaaccccccccaaaaaaccaaaaaacccaacaaccccccccccccaaaaaaaacaaaaaaccccccccccccccccccccccccccccccccccaaaaaa
	abccaaaaccccccccccccccccccccaaaaaaaaccccccccccaaaaaaccccaaaaaccccccccccccaaaaaaaaaaacccccccccaaaaaccaaaaaccccaaaacccccccccccccccaaaaacccaaaaaccccccccccccccccccccaaccccccccccccaaaaaa
	abccaaaacccccccaaaccccccccccaaaaaaacccccccccccaaaaaaccccaaaaaccccccccccccacaaaaaaaaaacccccccaaaaacaaaaaaacccaaaaaccccccccccccccaaaaacccaaaaaccccccccccccccccccccaaaaccccccccccccccaaa
	abccaaaccccccaaaaaacccccccccaccaaaccccccccccccaaaaacccccaaaaaaccccaaaacccccaaaaaaaaaaaccccccaaaaacaaaaaaacccaaaaaacccccccccccccaacaaaccaccaacccccccccccccccaaaccaaaaccccccccccccccaaa
	abcccccccccccaaaaaacccccccccccccaaacccccccccccaaaaacccccaaaaaaccccaaaaccccccaaaaacaaaaccccccccccccccaaaaaaccacaaaaccccaaaaacccccccaaccccccccccccccccccccccaaaackkkaccccccccccccccccaa
	abcccccccccccaaaaaacccccccccccccccaaacccccccccccccccccccaaaaaaccccaaaacccccaaaaaccccaaccccccccccccccaaccaaccccaaccccccaaaaaccccccccccccccccccccccccccccccccaakkkkkkkccccccccccccccccc
	abaccccccccccaaaaaccccccccccccccccaaaaccccccccccccccccccccaaaccccccaaccccccccaaaccccccccccccccccccccaacccccccccccccccaaaaaacccccccccccccccccccccaaacccccccccjkkkkkkkkccccccccaacccccc
	abaccccccccccaaaaacccccaacccccccccaaaaccccccccccccaacccccccccccccccccccccccccaaacccccccccccccccccccccaaccccccaccaccccaaaaaaaaaccaccccccccccccccaaaaccccccccjjkkoopkkkkaccaacaaacccccc
	abaccccccccccccccccaaaaaacccccccccaaacccccccccccccaaaaaacccccccccccccccccccccccccccccccccccccccccccccaaaaaaccaaaaccccaaaaaaaaacaaccccccccccccccaaaacccccccjjjkoooppkkkaccaaaaaaaacccc
	abcccccccccccccccccaaaaaaaaccccccccccccccccaccccccaaaaaaccccccccccccccccccccccccaaccaacccccccccccccccaaaaaacaaaaacccccaaacccaaaaacccccccccccccccaaacccccjjjjjoooppppkklccaaaaaaaacccc
	abcccccccccccccccccaaaaaaaacccccccccccccccaaacccaaaaaaaccccaacccacccccccccccccccaaaaaaccccccccaaaccaaaaaaaccaaaaaaccccccccaaaaaacccccccccccccccccccccjjjjjjjoooouuppplllccaccaaaacccc
	abccccccccccccccccccaaaaaaaccccaacccccaaacaaacccaaaaaaaccccaaacaacccccccccccccccaaaaaacccccccccaaaaaaaaaaaacaaaaaaccccccccaaaaaaaaccccccccccccccccciijjjjjjooouuuupppllllcccccccccccc
	abccccccccccccccccccaaaaaccccccaacccccaaaaaaaaaaccaaaaaaccccaaaaaccccccccccccccaaaaaaacccccccaaaaaaaaaaaaaacccaaccccccccccaacaaaaacccccccccccccccciiiijoooooouuuuuuppplllllcccccccccc
	abcccccccccccccccccaaaaaacccaacaaaaacccaaaaaaaaaccaaccaaccaaaaaacccccccccccccccaaaaaaaaccccccaaaaacccaaaaaacccaccccccccccccccaacccccccccccccccccciiiinnoooooouuxuuuupppplllllcccccccc
	abcccccccccccccccccccccaacccaaaaaaaaccccaaaaaaccccaaccccccaaaaaaaaccaaaccccccccaaaaaaaacccccccaaaaaccaacaaaaaaaacccaaccccccccacccccccccaaaccccccciiinnnnntttuuuxxuuuppppqqllllccccccc
	abccccccccccccaacccccccccccccaaaaaccccccaaaaaacccccaaaccccaaaaaaaaccaaaacccccccccaaaccccccccccaaccaccccccaaaaaacccaaaaaaccccccccccccccaaaaccccaaiiinnnntttttuuxxxxuuvpqqqqqllmmcccccc
	abccccccccccaaaaaaccccccccccccaaaaaccccaaaaaaacccccaaacccccccaacccccaaaaccccaaccccaacccccccccccccccccccccaaaaaaccccaaaaaccccccccccccccaaaaccccaaiiinnnttttxxxxxxxyuvvvvvqqqqmmmcccccc
	abccaaacccccaaaaaacccccccccccaaacaaccccaaacaaacccccaaaaaaacccaccccccaaaccccaaaacccccccccccccccccccccccccaaaaaaaacaaaaaaacccccccccaaacccaaaccccaaaiinnntttxxxxxxxxyyyyvvvvqqqmmmcccccc
	abcaaaacccccaaaaaccccccccccccaaaccaccccccccccacccaaaaaaaaaaccccccccccccccccaaaaccccccccccccccccccccccccaaaaaaaaaaaaaaaaaaccccccccaaaaaacccccaaaaaiiinnnttxxxxxxxyyyyyyvvvqqqmmmcccccc
	SbcaaaaccccccaaaaacccccaaaccccccaaacccccccccaaccaaaaaaaaaaaacccccccccccccccaaaaccccccccccccccccccccccccaaaaaaaaaaaaaaaaaacccccccaaaaaaacccccaaaaaiiinnntttxxxxEzzyyyyvvvqqqmmmdddcccc
	abccaaaccccccaaaaacccccaaaaccccaaaaaaccccccaaaaccaaaaaaacaaacccccaaccccccccccccccccccccccccccccccccccccacaaaaacccccaaacacccccccaaaaaaaccccccaaaaaahhhnnntttxxxyyyyyyvvvvqqmmmmdddcccc
	abcccccccccccccccccccccaaaaccccaaaaaaccccccaaaaccccaaaaaaaaaaaaaaaacacccccccccccccccccccccccccccccccccccccaaaacccccaaccccccccccaaaaaaaccccccccaaaahhhnnnnttxxxyyyyyvvvqqqqmmmdddccccc
	abcccccccccccccccaacaacaaacccccaaaaaaccccccaaaacccaaaaaaaaaaaaaaaaaaacccccccccccccccccaacaaccccccccccccccccaaccccccccccccccccccccaaaaaacccccccaaccchhhmmmttxwyyyyyyvvrqqqmmmddddccccc
	abcccccccccccccccaaaacccccccccccaaaaacccccccccccaaaaaaaaaaaaaaccaaaaccccaacccccaacccccaaaaaccccccccccccccccccccccccccccccccaaaaccaaaaaacccccccaaccahhhmmssswwywwwyyyvvrqmmmmdddcccccc
	abccccccccccccccaaaaacccccccccccaacaacccccccccccaaaaaacaaaaaacccaaaaaaacaaccccaaaccccccaaaaacccccccacccccccccccccccccccccccaaaaccaacccccccccccaaaaahhhmmsswwwwwwwwywwvrrnnmdddccccccc
	abccccccccccccccaaaaaaccaaccccccccccccccccccccccaaaaaaaaaaaaacccacaacaaaaaccccaaacaaacaaaaaacaaacaaacccccccacccaaccccaaccccaaaacccccccccaaaccccaaaahhhmmssswwwwswwwwwwrrnnndddccccccc
	abaaccccccccccccacaaaacaaaaaaaccccccccccaacccccccaaaaaaaacaaaccccccccaaaaaaaaaaaaaaaacaaaacccaaaaaaacccccccaacaaaaaaaaacccccaaccccccccccaaacccaaaaahhhmmsssswsssrrwwwrrrnneddaccccccc
	abaaccccccccaaccccaaccccaaaaacccccccccccaacccccccaaaacccccccacccccccaaaaaaaaaaaaaaaaacccaaccccaaaaaacccccccaaaaacaaaaaaacccccccccccccaaaaaaaaaaaaaahhhmmssssssssrrrrrrrrnneedaaaacccc
	abacccccccccaaaaccccccaaaaaaaccccccccaaaaaaaacccccccccccccccccccccccaaaaaaaacaaaaaacccaaacccccaaaaaaaacccccaaaaaacaaaaaaaccccccccccccaaaaaaaaaaaaaahhhmmmsssssllrrrrrrrnnneeeaaaacccc
	abaaacccccaaaaaaccccccaaaaaaaacccaaccaaaaaaaaccccaaaaaccccccccccccccaaaaaacccaaaaaaccccaaaccaaaaaaaaaaccccaaaaaaaaaaaaaaaccccccccccccccaaaaaccccaachhgmmmmmlllllllrrrrrnnneeeaaaacccc
	abaaacccccaaaaaccccaccaaaaaaaacaaaaaaccaaaaccccccaaaaacccccccccccccccccaaacccaaaaaaaccaaaaaaaaaaaaaaaaccccaaaaaaaaaaaaacccccccccccccccaaaaaaccccaaccgggmmmllllllllllnnnnneeeaaaaccccc
	abcccccccccaaaaacccaaacaaaacaacaaaaaacaaaaaccccccaaaaaaccccccccccccccccccccccaaacaaacccaaaaaaaacaaaccccccccccaaccaaaaaacccccccccccccccaaaaaacaacccccgggggmlllfffflllnnnnneeeaaaaccccc
	abcccccccccaaccacccaaaaaaacccccaaaaaacaacaaacccccaaaaaaccccccccccccccccccccccaccaaccaaaaaaaaacccaaaccccccccccaaccccccaacccccaaaaccccccaccaaaaaaccccccggggggggfffffflnnneeeeeacaaacccc
	abaaaccccccccccccccaaaaaacccccccaaaaacacccaaaacccaaaaaacccccccccccccccccccaaacccaaccaaaaaaaaacccaaccccccccccccccccccccccccccaaaaccccaaaccaaaaaacccccccgggggggfffffffffeeeeeaacccccccc
	abaaaaacccccccccccaaaaaaaaccccccaaaacccccccaaacccccaacccccccaaacccccccccccaaaacaaacccaaaaaaaacccccccccccccccccccccccccccccccaaaaccccaaacccaaaaaaacccccccccgccaaaafffffeeeeccccccccccc
	abaaaaaccccccaaccaaaaaaaaaacccccaaacaaaaaacaaaccccccccccccaaaaaccccccccccaaaaacaaacccccaaaaaaacccccccccccccccccccccccccccccccaacccccaaaaaaaaaaaaaccccccccccccaaaacaafffecccccccccccaa
	abaaaacccaaacaaccaaaaaaaaaacccccaaaaaaaaaaaaaaaaacccccccccaaaaaaccccccccccaaaaaaaaaaacaaacccacccaacccccccccaaaaccccaaacccccccccccaaaaaaaaaaaaaaacccccccccccccaaaccccaaccccccccccccaaa
	abaaaaacccaaaaacccccaaacaaacccccaaaaaaccaaaaaaaaacccccccccaaaaaaccaaccacccaaaaaaaaaaaaaaaccccccaaacccccccccaaaaccccaaaaccccccccccaaaaaaaaaaaaaaccccccccccccccaaaccccccccccccccccccaaa
	abaaaaacccaaaaaaacccaaaccccccccaaaaaaaaccaaaaaaaccccccccccaaaaacccaaaaaccccaaaaaaaaaaaaacccccaaaaaaaaccccccaaaaccccaaaacccccccccccaaaaaaacccaaaccccccccccccccaaacccccccccccccccaaaaaa
	abcccccccaaaaaaaaccccaacccccccaaaaaaaaaaaaaaaaacccccccccccaaaaaccaaaaaacccccaaaaaaaaaaaaaccccaaaaaaaacccccccaacccccaaacccccccccccccaaaaaaacccccccccccccccccccccccccccccccccccccaaaaaa`
)

func main() {
	var data = puzzle

	my_field := make_field(data)

	route := my_field.find_route()

	// println(route.to_str())

	route_slice := route.to_slice()

	route_len := len(route_slice)

	route_len--

	println(route_len)
}
