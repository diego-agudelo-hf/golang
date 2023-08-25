package main

import (
	"fmt"
	"git/variables"
)

func main() {
	fmt.Println(variables.Status)
	variables.ShowInt()
	fmt.Println("end")
	variables.ShowAny()

	fmt.Println(variables.Status)
	status, text := variables.ConvertToString(122)
	fmt.Println(status)
	fmt.Println(text)

}
