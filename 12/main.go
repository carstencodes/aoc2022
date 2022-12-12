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
	size    int
	mutex   *sync.Mutex
}

type path struct {
	current point
	origin  *path
	len     int
}

func (p point) hash() int {
	return p.row_id*1024 + p.col_id
}

func (f field) is_in_matrix(p point) bool {
	row_id := p.row_id
	col_id := p.col_id

	return (row_id >= 0) && (row_id < len(f.symbols)) &&
		(col_id >= 0) && col_id < len(f.symbols[row_id])
}

func (f field) is_reachable(from point, to point) bool {
	if f.is_in_matrix(from) && f.is_in_matrix(to) {
		symbol_from := f.symbols[from.row_id][from.col_id]
		symbol_to := f.symbols[to.row_id][to.col_id]

		delta := symbol_to - symbol_from

		return delta <= 1
	}

	return false
}

func (p path) to_str() string {
	if p.origin != nil {
		return fmt.Sprintf("%s -> %s", p.origin.to_str(), p.current.to_str())
	} else {
		return p.current.to_str()
	}
}

func (p path) contains(pnt point) bool {
	//return contains(p.points, pnt.hash())
	return p.current.equals(pnt) || (p.origin != nil && p.origin.contains(pnt))
}

func (p point) to_str() string {
	return fmt.Sprintf("(%d, %d)", p.row_id, p.col_id)
}

func (p point) equals(another point) bool {
	return p.row_id == another.row_id && p.col_id == another.col_id
}

func (data field) get_neighbors(p point) []point {
	p_up := point{p.row_id - 1, p.col_id}
	p_down := point{p.row_id + 1, p.col_id}
	p_left := point{p.row_id, p.col_id - 1}
	p_right := point{p.row_id, p.col_id + 1}

	result := make([]point, 0)
	if data.is_reachable(p, p_up) {
		result = append(result, p_up)
	}
	if data.is_reachable(p, p_down) {
		result = append(result, p_down)
	}
	if data.is_reachable(p, p_left) {
		result = append(result, p_left)
	}
	if data.is_reachable(p, p_right) {
		result = append(result, p_right)
	}

	return result
}

func (f field) get_symbol(p point) byte {
	if f.is_in_matrix(p) {
		return f.symbols[p.row_id][p.col_id]
	}

	return 0
}

func create_paths(p path, data field, result chan path, wg *sync.WaitGroup) {
	defer wg.Done()

	if data.end.equals(p.current) {
		result <- p
		return
	}

	neighbors := data.get_neighbors(p.current)

	for _, item := range neighbors {
		if !p.contains(item) {
			wg.Add(1)
			next_path := path{item, &p, p.len + 1}

			go create_paths(next_path, data, result, wg)
		}
	}
}

func (data field) find_route_from_path() path {

	result := make(chan path, data.size)

	origin := path{data.start, nil, 0}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go create_paths(origin, data, result, wg)

	wg.Wait()
	close(result)

	var shortest_path *path = nil
	for found_path := range result {
		if shortest_path != nil {
			if found_path.len < shortest_path.len {
				shortest_path = &found_path
			}
		} else {
			shortest_path = &found_path
		}
	}

	if shortest_path == nil {
		return origin
	} else {
		return *shortest_path
	}
}

func make_field(data string) field {
	var lines = strings.Split(data, "\n")

	var matrix = make(map[int][]byte, 0)

	const start_symbol = 'S'
	const end_symbol = 'E'

	var start_position = point{-1, -1}
	var end_position = point{-1, -1}

	var nodes = 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		row := make([]byte, 0)
		visitors := make([]byte, 0)
		for j := 0; j < len(line); j++ {
			nodes++
			chr := line[j]
			p := point{i, j}

			if chr == start_symbol {
				start_position = p
				chr = 'a'
			}
			if chr == end_symbol {
				end_position = p
				chr = 'z'
			}

			row = append(row, chr)
			visitors = append(visitors, 0)
		}

		matrix[i] = row
	}

	var my_field = field{
		start_position,
		end_position,
		matrix,
		nodes,
		&sync.Mutex{},
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
	var data = sample

	my_field := make_field(data)

	path := my_field.find_route_from_path()

	println(path.to_str())

	println(path.len)
}
