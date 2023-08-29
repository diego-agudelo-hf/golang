package interfacesexample

import (
	"fmt"
	"git/interfaces"
)

func ManRepirando(human interfaces.Human) {
	human.Respirar()
	fmt.Println(human.Gender())

}
