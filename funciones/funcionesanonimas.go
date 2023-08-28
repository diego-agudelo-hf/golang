package funciones

import "fmt"

func Calculate() {
	var num1 int = 1
	var num2 int = 44
	sum := func(n1 int, n2 int) int {
		return n1 + n2 + num1 + num2
	}
	fmt.Println(sum(3, 4))

}
