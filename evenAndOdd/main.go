package main

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range ints {
		isEven := v%2 == 0
		print(v, " is ")
		if isEven {
			println("even")
		} else {
			println("odd")
		}
	}
}
