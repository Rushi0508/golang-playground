## Packages and Imports in Go Lang

- In Go, packages are used instaed of classes. There are no concepts of OOPs unlike C++ or Java.
- Each Go file must belong to some package, therefore it must start with keyword `package` followed by package name.
- `main` is special package in Go. An executable program must contain a `main` package.

- Importing a package is done usign `import` keyword followed by list of packages inside parentheses. Ex-

```
import (
    "fmt"
)
```

- `fmt` is package used for printing to console.

- In go, packages are tightly coupled with the directory so that all go files in single directory are considered to be part of same package.

- To create a new package we must create a different directory.
