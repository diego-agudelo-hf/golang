package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var Status bool
var Salary float32
var Date time.Time

func ShowAny() {
	Name = "pedro"
	Status = true
	Salary = 25000000
	Date = time.Now()
	fmt.Println(Name + " Agudelo")
	fmt.Println(Status)
	fmt.Println(Salary)
	fmt.Println(Date)
}

func ConvertToString(number int) (bool, string) {
	Text := strconv.Itoa((number))
	return true, Text
}
