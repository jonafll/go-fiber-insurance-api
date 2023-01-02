package utils

import (
	"math/rand"
	"time"
)

func SimulateRestCall() {
	rand.Seed(time.Now().UnixNano())
	duration := 100 + rand.Intn(401)
	time.Sleep(time.Duration(duration) * time.Millisecond)
}
