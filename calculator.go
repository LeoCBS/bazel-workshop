package calculator

import "fmt"

func SumInt(value ...int) int {
	for k, v := range value {
		fmt.Println(k)
		fmt.Println(v)
	}
	return 0
}
