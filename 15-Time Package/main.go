package main

import (
	"fmt"
	"time"
)

func main() {
    currentTime := time.Now()
    fmt.Println("Current Time:", currentTime)

    time.Sleep(2 * time.Second)
    fmt.Println("After 2 seconds:", time.Now())

    fmt.Println("Elapsed Time:", time.Since(currentTime))

    fmt.Println("Current Time:", currentTime.Format("2006-01-02 15:04:05"))

    timeValue := "2021-07-01"
    parsedTime, _ := time.Parse("2006-01-02", timeValue)
    fmt.Println("Parsed Time:", parsedTime)

    futureTime := currentTime.Add(2 * time.Hour)
    fmt.Println("Future Time:", futureTime)

    duration := futureTime.Sub(currentTime)
    fmt.Println("Duration:", duration)
}