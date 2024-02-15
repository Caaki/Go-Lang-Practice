package slices

import "fmt"

func getMessageCosts(messages []string) []float64 {
	toReturn := make([]float64, len(messages))

	for i := 0; i < len(messages); i++ {
		toReturn[i] = 0.01 * float64(len((messages)[i]))
	}
	return toReturn
}

func sum(nums ...int) (all int) {

	for _, num := range nums {
		all += num
	}
	// radilo bi i samo sa return smo gore rekli all i zna da vrati to
	return all
}

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	max := 0
	for _, num := range costs {
		if max < num.day-1 {
			max = num.day + 1
		}
	}
	newSlice := make([]float64, max, max)
	for _, num := range costs {

		newSlice[num.day] += num.value
	}
	return newSlice
}

func createMatrix(rov, cols int) [][]int {

	matrix := make([][]int, 0)

	for i := 0; i < rov; i++ {
		matrix = append(matrix, []int{})
		for j := 0; j < cols; j++ {
			matrix[i] = append(matrix[i], i*j)
		}
	}

	return matrix
}

func main() {
	arr := [3]int{1, 2, 3}

	sli := make([]int, 5, 10)
	sli = arr[:]
	sli = append(sli, 4)
	sli[1] = 50

	fmt.Println(sli)
	fmt.Println(arr)

	fmt.Println(sum(1, 2, 3))

	messaages := []string{"aaa", "bbbbbb", "cccccc"}

	fmt.Println(getMessageCosts(messaages))

	fmt.Println(getCostsByDay([]cost{cost{day: 0, value: 2}, {0, 3}, {5, 1}, {1, 2}}))

	fmt.Println(createMatrix(10, 10))
}
