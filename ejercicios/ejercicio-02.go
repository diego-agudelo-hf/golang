package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number int
var err error
var conten string

func Invoke() string {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese tabla a calcular")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())
			if err == nil {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {

		conten += fmt.Sprintf("%d x %d =%d -", number, i, number*i)

	}

	return conten

}
