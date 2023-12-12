package functions

import(
	"fmt"
)

type FibonacciFunc struct{}

func(f *FibonacciFunc) Fibonacci (n int){
	var test []int
	test = append(test, 0, 1)
	result := 0

	for i := 0; n >= i; i++{
		fmt.Println(test[i])
		result = test[i] + test[i + 1]
		test = append(test, result)
	}
}