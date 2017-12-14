package main

import (
	"time"
)

func greeting() string {
	return "service1 " + time.Now().String()
}
