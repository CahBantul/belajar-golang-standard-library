package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now)

	var utc time.Time = time.Date(2024, time.February, 8, 21, 30, 0, 0, time.UTC)
	fmt.Println(utc)

	formatter := "2006-01-02 15:04:05"

	value := "2022-12-12 12:12:12"

	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

}
