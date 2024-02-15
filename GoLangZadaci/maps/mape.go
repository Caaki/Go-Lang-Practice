package maps

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Bad input")
	}

	forReturning := make(map[string]user)
	for i := 0; i < len(names); i++ {
		forReturning[names[i]] = user{name: names[i], phoneNumber: phoneNumbers[i]}
	}

	return forReturning, nil
}

type user struct {
	name        string
	phoneNumber int
}

func main() {
	ages := map[string]int{
		"nesto": 34,
	}

	ages["novie"] = 2 // Dodavanje

	delete(ages, "novi") // brisanje

	provera, ok := ages["novi"] // Provera postojanja vraca value i boolean

	if ok == true {
		fmt.Println(provera)
	} else {
		fmt.Println("Doslo je do greske")
	}

	fmt.Println(ages)

	for k, v := range getNameCounts([]string{"bob", "bobi", "bob", "alen", "al", "sam"}) {
		fmt.Println(string(rune(k)))
		fmt.Println(v)
	}

}

func getNameCounts(names []string) map[rune]map[string]int {

	toBeReturned := make(map[rune]map[string]int)
	for _, name := range names {
		if _, ok := toBeReturned[rune(name[0])]; !ok {
			toBeReturned[rune(name[0])] = make(map[string]int)
		}
		toBeReturned[rune(name[0])][name]++
	}
	return toBeReturned
}
