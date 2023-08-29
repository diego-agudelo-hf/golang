package arrays

import "fmt"

var table [10]int = [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1}
var matriz [10][10]int

func Invoke() {
	table[8] = 3333333
	table[2] = 2
	table2 := [10]string{"diego a"}
	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
	matriz[1][2] = 5
	fmt.Println(table, table2, matriz)
}
