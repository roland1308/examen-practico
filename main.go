package main

func sumInt(a int, b int) int {
	return a + b
}

func subInt(a int, b int) int {
	return a - b
}

func main() {
	println("La suma entre 4 y 2 es:", sumInt(4, 2))
	println("La resta entre 4 y 2 es:", subInt(4, 2))
}
