package files

import (
	"bufio"
	"fmt"
	"git/ejercicios"
	"io/ioutil"
	"os"
)

var dir string = "./files/txt/tabla.txt"

func Invoke() {
	var tmp string = ejercicios.Invoke()
	archivo, err := os.Create(dir)
	if err != nil {
		fmt.Println("ocurrio un error " + err.Error())
		return
	}
	fmt.Fprintln(archivo, tmp)
	archivo.Close()
}

func Addstring() {
	var tmp string = ejercicios.Invoke()
	if !Append(dir, tmp) {
		fmt.Println("Error al concatenar el archivo...")
	}
}

func Append(dir string, tmp string) bool {
	archivo, err := os.OpenFile(dir, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {

		fmt.Println("Error en append..." + err.Error())
		return false
	}
	_, err = archivo.WriteString(tmp)
	if err != nil {

		fmt.Println("Error en WriteString..." + err.Error())
		return false
	}
	archivo.Close()
	return true
}

func ReadFIle() {
	archivo, err := ioutil.ReadFile(dir)

	if err != nil {
		fmt.Println("Error al leer..." + err.Error())
	}
	fmt.Println(string(archivo))
}

func ReadFIleOs() {
	archivo, err := os.Open(dir)

	if err != nil {
		fmt.Println("Error al leer 2..." + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("> " + line)
	}
	archivo.Close()
}
