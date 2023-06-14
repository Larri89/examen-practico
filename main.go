package main

func localSum(a int, b int) int {
	return a + b
}

func localRes(a int, b int) int {
	return a - b
}

func main() {
	println(localSum(5, 9))
	println(localRes(11, 20))
}
