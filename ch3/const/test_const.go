package main

import (
	"fmt"
	"math"
	"time"
)

const (
	PI = 3.14
	E  = 2.555
)

var a time.Duration

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp Flags = 1 << (iota * 10)
	FlagBroadcast
	FlagLoopBack
	FlagPoint2Point
	FlagMulticast
)

func main() {
	var x float32 = math.Pi
	fmt.Println(x)

	const PI = math.Pi
	fmt.Println(PI)
}
