package utils

import (
	"math/rand"
	"time"
)

func Random4Digit() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9000) + 1000
}
