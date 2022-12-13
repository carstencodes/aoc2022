package main

func main() {
	data := puzzle

	list, err := make_list(data)
	if err != nil {
		panic(err)
	}

	println(list.to_string())

	sum := 0
	for idx, pair := range list.items {
		if pair.is_in_correct_order() {
			print(idx + 1)
			println("=true")
			//fmt.Printf("%s at index %d is in order to %s\n", pair.left_side.to_string(), idx+1, pair.right_side.to_string())
			sum += idx + 1
		} else {
			print(idx + 1)
			println("=false")
			//fmt.Printf("%s at index %d is in NOT order to %s\n", pair.left_side.to_string(), idx+1, pair.right_side.to_string())
		}
	}

	println(sum)
}
