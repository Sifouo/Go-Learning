package main

import (
	"errors"
)

func insert(orig []int, index int, value int) ([]int, error) {
	if index < 0 {
		return nil, errors.New("Index cannot be less than 0 ")
	}
	if index >= len(orig) {
		return append(orig, value), nil
	}
	orig = append(orig[:index], orig[index:]...)
	orig[index] = value
	return orig, nil
}
