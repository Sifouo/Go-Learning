package main

import (
	"fmt"
)

type DigitCounter interface {
	EvenorOdd() (int, int)
}

type DigitString string

func Printt() {
	s := DigitString("123456789")
	fmt.Println(s.EvenorOdd())

}
func (ds DigitString) EvenorOdd() (int, int) {

	odds, evens := 0, 0
	for _, digit := range ds {
		if digit%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return odds, evens

}
