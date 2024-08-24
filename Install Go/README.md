## Installation

- Download Go lang from here https://go.dev/doc/install
- Verify the verison using command
  `go version`

### First Go Program

```
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

Start the program with a package name.

To run a go program use command `go run <filename>.go`

To run a Go program we must initialize a `go.mod` file at the root of the folder where the go files are in, using command `go mod init <module-name>`
