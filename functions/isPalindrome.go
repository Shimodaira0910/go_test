package functions

import (
	//"fmt"
	"strconv"
	"strings"
)

type Palindrome struct{}

func (p *Palindrome)IsPalindrome(n int) bool{
	toStringN := strconv.Itoa(n)
	toArrayNum := strings.Split(toStringN, "")

	var reversedNumArray []string
	for i := len(toArrayNum) - 1; i >= 0; i-- {
		reversedNumArray = append(reversedNumArray, toArrayNum[i])
	}

	toReverseStringN := strings.Join(reversedNumArray, "")
	
	return toStringN == toReverseStringN
}	