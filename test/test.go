package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	year, month, day := time.Now().Date()
	t1 := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	u1 := t1.UnixMilli()
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	year1, month1, day1 := sevenDaysAgo.Date()
	t2 := time.Date(year1, month1, day1, 0, 0, 0, 0, time.Local)
	u2 := t2.UnixMilli()
	fmt.Println(strconv.FormatInt(u2, 10))
	fmt.Println(strconv.FormatInt(u1, 10))
}
