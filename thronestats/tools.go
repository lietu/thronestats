package thronestats

import (
	"log"
	"strconv"
	"runtime/debug"
)

func IsIn(list []int, value int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}

func ToInt(value string) int {
	if value == "" {
		return 0
	}

	i, err := strconv.Atoi(value)

	if err != nil {
		debug.PrintStack()
		log.Fatalf("Error parsing value as integer %s", err)
		return 0
	}

	return i
}

