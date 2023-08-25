package variables

import "fmt"

func ShowInt() {
	var intcomun int
	intv32 := int32(10)
	intv64 := int64(100)
	fmt.Println(intcomun, intv32, intv64)
}
