package main

import (
	"time"
)

func greeting() string {
	return "Hello world at: " + time.Now().String()
}
