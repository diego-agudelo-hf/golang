package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func Minombrelento(nombre string, chan1 chan bool) {
	letras := strings.Split(nombre, "")

	for _, value := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(value)
	}
	chan1 <- true

}
