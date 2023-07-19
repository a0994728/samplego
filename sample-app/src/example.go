package main

func test1(x, y int) int {
	return x + y
}

func adder(t int) func(int) int {
	sum := 0
	return func(x int) int {
		sum += x + t
		return sum
	}
}

func test2(x int) func() int {
	y := 0
	return func() int {
		y += y + x
		return y
	}
}
