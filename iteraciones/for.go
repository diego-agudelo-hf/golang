package iteraciones

import "fmt"

func Invoke() {
	var count int = 100000000
	for i := 0; i < count; i += 3 {
		fmt.Println("hola ", i)
	}
}
