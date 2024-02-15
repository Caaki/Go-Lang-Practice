package closures

/*

Keeping track of how many emails we send is mission-critical at Mailio. Complete the adder() enclosing function.
It should return a function that adds its input (an int) to an enclosed sum value,
then return the new sum. In other words, it keeps a running total of the sum variable within a closure.
*/

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

// don't touch below this line

type emailBill struct {
	costInPennies int
}

func test(bills []emailBill) {
	defer fmt.Println("====================================")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
	}
}

func main() {
	test([]emailBill{
		{45},
		{32},
		{43},
		{12},
		{34},
		{54},
	})

	test([]emailBill{
		{12},
		{12},
		{976},
		{12},
		{543},
	})

	test([]emailBill{
		{743},
		{13},
		{8},
	})
}
