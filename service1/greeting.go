package main

import (
	"time"
)

func greeting() string {
	return "hello " + time.Now().String()
}
