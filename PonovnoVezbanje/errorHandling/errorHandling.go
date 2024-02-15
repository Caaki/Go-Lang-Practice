package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	if costToCustomer, errMsg := sendSMS(msgToCustomer); errMsg != nil {
		return 0, errMsg
	} else if costTpSpuse, errMsg := sendSMS(msgToSpouse); errMsg != nil {
		return 0, errMsg
	} else {

		return costToCustomer + costTpSpuse, nil
	}

}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

type user struct {
	name string
}

type userError struct {
	name string
}

func (e userError) Error() string {

	return fmt.Sprintf("There was an error with the user: %v", e.name)
}

func main2() {

	someError := userError{name: "User1"}

	fmt.Println(someError.Error())

	text, err := sendSMSToCouple("aaaaa", "aaaaaa")

	fmt.Printf("The count %v The error %v\n", text, err)
	a, b := sendSMSToCouple("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "")
	fmt.Printf("The count %v The error %v", a, b)
}
