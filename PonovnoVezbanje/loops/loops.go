package loops

import (
	"errors"
	"fmt"
)

func fizzbuzz() {

	for i := 1; i <= 100; i++ {

		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == "pro" {
		return allMessages[:], nil
	} else if plan == "free" {

		return allMessages[0:2], nil
	}
	return nil, errors.New("Unsuported plan")
}

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func main() {
	fizzbuzz()
	fmt.Println(getMessageWithRetriesForPlan("pro"))
	fmt.Println(getMessageWithRetriesForPlan("freee"))

}
