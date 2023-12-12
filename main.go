package main

import (
	//"fmt"
	//"log"
	"fmt"
	"shimodaira/functions"
)

func main(){
	// fn := functions.FizzBuzzFunc{}
	// for i := 0; i <= 100; i++{
	// 	result, err := fn.FizzBuzz(i)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	} else {
	// 		fmt.Println(result)
	// 	} 
	// } 

	//fi := functions.FibonacciFunc{}
	//fi.Fibonacci(5)

	is := functions.Palindrome{}
	fmt.Println(is.IsPalindrome(-121))

}