package arrays

import "fmt"

var table_slice []int = []int{1, 2, 3, 4, 5}
var table_to_slice [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func Invoke_s() {
	fmt.Println(table_slice)
	porcion := table_to_slice[3:]
	fmt.Println(porcion)
	porcion2 := table_to_slice[2:5]
	fmt.Println(porcion2)
	porcion3 := table_to_slice[:3]
	fmt.Println(porcion3)
	fmt.Println(table_to_slice)
}

func Capacity() {
	elements := make([]int, 5, 20)

	fmt.Println(len(elements), cap(elements))
	Nums := make([]int, 0, 0)
	for i := 0; i < 1000; i++ {
		Nums = append(Nums, i)
	}
	fmt.Println(len(Nums), cap(Nums))

}
