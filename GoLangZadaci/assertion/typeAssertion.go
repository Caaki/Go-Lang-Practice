package assertion

import (
	"fmt"
)

type employee interface {
	getName() string
	getSalary() int
}

type testingIterfaceFuncionalites interface {
	employee
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

type test1 struct {
}

func (t test1) getSalary() int {
	return 0
}

func (c test1) getName() string {
	return ""
}

func (c contractor) getSalary() int {
	return c.hoursPerYear * c.hourlyPay
}

func (c contractor) getName() string {
	return c.name
}

func getInfo(e employee) {

	if _, ok := e.(contractor); ok == true {

		fmt.Println("This is a contractor")
		fmt.Println(e.getName())
		fmt.Println(e.getSalary())
	} else if _, ok := e.(contractor); ok == false {
		fmt.Println("This is the test Case!!!!")
	}

}

func testForTypeWithSwitch(e employee) {
	switch a := e.(type) {
	case contractor:
		fmt.Println(a.name)
	case test1:
		fmt.Println("This is a test!!!")
	}

}

func main() {

	c := contractor{
		name:         "Caki",
		hoursPerYear: 1000,
		hourlyPay:    40,
	}
	getInfo(c)
	getInfo(test1{})

	//Same but with switch

	testForTypeWithSwitch(c)
	testForTypeWithSwitch(test1{})

}
