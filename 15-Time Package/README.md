## Time Package in Go

The `time` package in Go provides functionality for measuring and displaying time. It provides a way to measure time in seconds, milliseconds, and nanoseconds. It also provides a way to format time in a human-readable format.

### Time Package Functions

The `time` package provides the following functions:

1. `Now()`: This function returns the current time.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    currentTime := time.Now()
    fmt.Println("Current Time:", currentTime)
    // Output: Current Time: 2021-07-01 15:04:05.999999999 +0000 UTC m=+0.000000001
}
```

2. `Sleep(d time.Duration)`: This function pauses the current goroutine for the given duration.

```go
fmt.Println("Start")
time.Sleep(2 * time.Second)
fmt.Println("End")
// Output: Start
//         End
```

3. `Since(t time.Time)`: This function returns the time elapsed since the given time.

```go
startTime := time.Now()
time.Sleep(2 * time.Second)
elapsedTime := time.Since(startTime)
fmt.Println("Elapsed Time:", elapsedTime)
// Output: Elapsed Time: 2.000000001s
```

4. `Format(layout string)`: This function formats the time according to the given layout.

`2006-01-02` is a special date layout that represents the date `January 2, 2006`.

```go
currentTime := time.Now()
fmt.Println("Current Time:", currentTime.Format("dd-MM-yyyy"))
// Output: Current Time: dd-MM-yyyy
fmt.Println("Current Time:", currentTime.Format("2006-01-02"))
// Output: Current Time: 2021-07-01
```

5. `Parse(layout, value string)`: This function parses the time value according to the given layout.

```go
timeValue := "2021-07-01"
parsedTime, _ := time.Parse("2006-01-02", timeValue)
fmt.Println("Parsed Time:", parsedTime)
// Output: Parsed Time: 2021-07-01 00:00:00 +0000 UTC
```

6. `Add(d time.Duration)`: This function adds the given duration to the time.

```go
currentTime := time.Now()
fmt.Println("Current Time:", currentTime)
// Output: Current Time: 2021-07-01 15:04:05.999999999 +0000 UTC m=+0.000000001
futureTime := currentTime.Add(2 * time.Hour)
fmt.Println("Future Time:", futureTime)
// Output: Future Time: 2021-07-01 17:04:05.999999999 +0000 UTC m=+7200.000000001
```

7. `Sub(t time.Time)`: This function returns the duration between two times.

```go
currentTime := time.Now()
futureTime := currentTime.Add(2 * time.Hour)
duration := futureTime.Sub(currentTime)
fmt.Println("Duration:", duration)
// Output: Duration: 2h0m0s
```

### Time Package Constants

The `time` package provides the following constants:

1. `Nanosecond`: Represents one nanosecond.
2. `Microsecond`: Represents one microsecond.
3. `Millisecond`: Represents one millisecond.
4. `Second`: Represents one second.
5. `Minute`: Represents one minute.
6. `Hour`: Represents one hour.

```go
fmt.Println("Nanosecond:", time.Nanosecond)
// Output: Nanosecond: 1
fmt.Println("Microsecond:", time.Microsecond)
// Output: Microsecond: 1000
fmt.Println("Millisecond:", time.Millisecond)
// Output: Millisecond: 1000000
fmt.Println("Second:", time.Second)
// Output: Second: 1000000000
fmt.Println("Minute:", time.Minute)
// Output: Minute: 60000000000
fmt.Println("Hour:", time.Hour)
// Output: Hour: 3600000000000
```
