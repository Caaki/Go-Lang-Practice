package slices

import "fmt"

func pazitiNaOvo() {

	i := make([]int, 3, 8)
	fmt.Println("len of i:", len(i))
	// len of i: 3
	fmt.Println("cap of i:", cap(i))
	// cap of i: 8
	fmt.Println("appending 4 to j from i")
	// appending 4 to j from i
	j := append(i, 4)
	fmt.Println("j:", j)
	// j: [0 0 0 4]
	fmt.Println("addr of j:", &j[0])
	// addr of j: 0x454000
	fmt.Println("appending 5 to g from i")
	// appending 5 to g from i
	g := append(i, 5)
	fmt.Println("addr of g:", &g[0])
	// addr of g: 0x454000
	fmt.Println("i:", i)
	// i: [0 0 0]
	fmt.Println("j:", j)
	// j: [0 0 0 5]
	fmt.Println("g:", g)
	// g: [0 0 0 5]

	i = append(i, 7)

	fmt.Println("i:", i)
	// i: [0 0 0]
	fmt.Println("j:", j)
	// j: [0 0 0 5]
	fmt.Println("g:", g)

}
