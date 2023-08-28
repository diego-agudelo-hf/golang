package funciones

import "fmt"

func Tabla(v int) func() int {
	num := v
	sec := 0
	return func() int {
		sec++
		return num * sec
	}
}

func Invoke() {
	tabla2 := 2
	rtabla2 := Tabla(tabla2)

	for i := 1; i <= 10; i++ {
		fmt.Println(rtabla2())
	}

}
