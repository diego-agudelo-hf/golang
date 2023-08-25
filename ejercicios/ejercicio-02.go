package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number int
var err error

func Invoke() {

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

	for i := 0; i <= 10; i++ {
		fmt.Println(number, "X", i, "=", number*i)
	}

}
