package ejercicios

import (
	"fmt"
	"strconv"
)

func Exec1(str string) (int, string) {
	Number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ocurrio un error")

	}
	var Text string
	if Number > 100 {
		Text = "Es mayor a 100"

	} else {
		Text = "No es mayor a 100"
	}

	return Number, Text
}
