package main

import "strconv"

func strToInt(str string) int {
	if i, err := strconv.Atoi(str); err == nil {
		return i
	}
	return -1
}
