package main

import "errors"

func delete(orig []int, index int) ([]int, error) {
	if index < 0 || index >= len(orig) {
		return nil, errors.New("Index out of the range")
	}
	orig = append(orig[:index], orig[index+1:]...)
	return orig, nil
}
