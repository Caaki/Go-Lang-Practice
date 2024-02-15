package higherOrderFunction

func add(x, y int) int {
	return x + y
}

func mult(x, y int) int {
	return x * y
}

func agregate(x, y, c int, arithmatic func(int, int) int) int {

	return arithmatic(arithmatic(x, y), c)
}

func main() {

}
