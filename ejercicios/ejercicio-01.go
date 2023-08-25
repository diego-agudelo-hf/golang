package ejercicios

import (
	"strconv"
)

func Exec1(str string) (int, string) {
	Number, err := strconv.Atoi(str)
	if err != nil {
		return Number, "Ocurrio un error " + err.Error()

	}
	var Text string
	if Number > 100 {
		Text = "Es mayor a 100"

	} else {
		Text = "No es mayor a 100"
	}

	return Number, Text
}
