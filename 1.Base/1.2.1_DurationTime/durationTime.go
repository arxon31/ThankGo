package main

import (
	"fmt"
	"time"
)

func main() {
	var inputTime string
	fmt.Scan(&inputTime)
	durationTime, _ := time.ParseDuration(inputTime)
	fmt.Println(inputTime, "=", durationTime.Minutes(), "min")
}
