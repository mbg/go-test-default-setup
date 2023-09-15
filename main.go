package main

func main() {
	x := min(1, 2, 3)
	y := max(1, 2, 3)
}

func foo[T any](a T) {
	clear(a)
}
