package generici

/*

ASSIGNMENT
At Mailio we often store all the emails of a given email campaign in memory as a slice.
We store payments for a single user in the same way.

Complete the getLast() function. It should be a generic function that returns the last element from a slice,
no matter the types stored in the slice. If the slice is empty, it should return the zero value of the type.

TIP: ZERO VALUE OF A TYPE
Creating a variable that's the zero value of a type is easy:

var myZeroInt int
Copy icon
It's the same with generics, we just have a variable that represents the type:

var myZero T

*/
import "fmt"

func getLast[T any](s []T) T {
	if len(s) < 1 {
		var zero T
		return zero
	} else {
		return s[len(s)-1]
	}
}

// don't edit below this line

type email struct {
	message        string
	senderEmail    string
	recipientEmail string
}

type payment struct {
	amount         int
	senderEmail    string
	recipientEmail string
}

func main() {
	test([]email{}, "email")
	test([]email{
		{
			"Hi Margo",
			"janet@example.com",
			"margo@example.com",
		},
		{
			"Hey Margo I really wanna chat",
			"janet@example.com",
			"margo@example.com",
		},
		{
			"ANSWER ME",
			"janet@example.com",
			"margo@example.com",
		},
	}, "email")
	test([]payment{
		{
			5,
			"jane@example.com",
			"sally@example.com",
		},
		{
			25,
			"jane@example.com",
			"mark@example.com",
		},
		{
			1,
			"jane@example.com",
			"sally@example.com",
		},
		{
			16,
			"jane@example.com",
			"margo@example.com",
		},
	}, "payment")
}

func test[T any](s []T, desc string) {
	last := getLast(s)
	fmt.Printf("Getting last %v from slice of length: %v\n", desc, len(s))
	for i, v := range s {
		fmt.Printf("Item #%v: %v\n", i+1, v)
	}
	fmt.Printf("Last item in list: %v\n", last)
	fmt.Println(" --- ")
}
