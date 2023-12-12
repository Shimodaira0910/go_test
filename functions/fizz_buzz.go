package functions

import (
	"fmt"
	"errors"
)

type FizzBuzzFunc struct {}

func (f *FizzBuzzFunc)FizzBuzz(n int)(string, error){
	if n > 100 {
		return "", errors.New("Arguments must be less than 100.")
	}
	
	if n % 3 == 0 && n % 5 == 0 {
		return "FizzBuzz", nil
	} else if n % 3 == 0 {
		return "Fizz", nil
	} else if n % 5 == 0 {
		return "Buzz", nil
	} else {
		return fmt.Sprintf("%d", n), nil
	}
}