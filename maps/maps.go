package maps

import "fmt"

func Invoke() {
	countries := make(map[string]string)
	countries["Colombia"] = "Bogota"
	countries["Mexico"] = "Cdmx"
	countries["Argentina"] = "Buenos aires"
	fmt.Println(countries)
	fmt.Println(countries["Colombia"])

	department := map[string]int{
		"quetame": 1,
		"fosca":   2,
		"fosca2":  2,
	}
	fmt.Println(department)

	for key, value := range department {
		fmt.Println(key, value)
		if key == "fosca2" {
			delete(department, key)
		}
	}
	fmt.Println(department)

	value, exist := department["fosca"]
	fmt.Println(value, exist)
}
