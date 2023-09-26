package main

import (
	"errors"
)

func GetCharacterAtPosition(str string, position int) (rune, error) {
	n := 0
	for _ = range str {
		n++
	}
	if (position < 0 || position > n) {
		return rune(0), errors.New("position out of range")
	}
	return []rune(str)[position],nil
}