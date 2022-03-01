package main

import "fmt"

// Create twoNums which is a type
type twoNums func(int, int) int

func main() {
	area := areaFunc()
	add := addFunc()
	sub := subFunc()
	print(30, 20, area)
	print(10, 10, area)
	print(20, 20, add)
	print(50, 20, sub)
}

func print(x, y int, t twoNums) {
	fmt.Printf("total is: %d\n", t(x, y))
}

func areaFunc() twoNums {
	return func(x, y int) int {
		return x * y
	}
}

func addFunc() twoNums {
	return func(x, y int) int {
		return x + y
	}
}

func subFunc() twoNums {
	return func(x, y int) int {
		return x - y
	}
}
