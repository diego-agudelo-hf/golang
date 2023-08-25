package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var text string
var err error

func Invoke() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese numero 1")
	if scanner.Scan() {
		number1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("dato incorrecto " + err.Error())
		}
	}
	fmt.Println("Ingrese numero 2")
	if scanner.Scan() {
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("dato incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese texto")
	if scanner.Scan() {
		text = scanner.Text()
	}

	fmt.Println(text, number1*number2)

}
