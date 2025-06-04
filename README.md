# curry

A lightweight Go library that brings the power of function currying to your codebase.

## What is Currying?

Currying is a functional programming technique where a function with multiple arguments is transformed into a sequence of functions, each taking a single argument.

**Example:**

```go
// Regular function
func add(a, b int) int {
    return a + b
}

// Curried version
func add(a int) func(int) int {
    return func(b int) int {
        return a + b
    }
}

addTwo := add(2)
result := addTwo(3) // result is 5
```

This approach allows for more flexible and reusable code, enabling partial application of functions.

## Features

- **Currying Functions**: Transform standard Go functions into their curried counterparts.
- **Uncurrying Functions**: Revert curried functions back to their original form.
- **Partial Application**: Create new functions by pre-filling some arguments.
- **Argument Dropping**: Generate functions that ignore certain arguments.

## Installation

```bash
go get github.com/parametalol/curry
```

## Usage

### Currying a Function

```go
package main

import (
    "fmt"
    "github.com/parametalol/curry"
)

func multiply(a, b int) int {
    return a * b
}

func main() {
    curriedMultiply := curry.Curry2(multiply)
    multiplyByTwo := curriedMultiply(2)
    result := multiplyByTwo(5)
    fmt.Println(result) // Outputs: 10
}
```

### Uncurrying a Function

```go
package main

import (
    "fmt"
    "github.com/parametalol/curry"
)

func curriedAdd(a int) func(int) int {
    return func(b int) int {
        return a + b
    }
}

func main() {
    uncurriedAdd := curry.Uncurry2(curriedAdd)
    result := uncurriedAdd(3, 4)
    fmt.Println(result) // Outputs: 7
}
```

### Partial Application

```go
package main

import (
    "fmt"
    "github.com/parametalol/curry"
)

func subtract(a, b int) int {
    return a - b
}

func main() {
    subtractFromTen := curry.Bind1(subtract, 10)
    result := subtractFromTen(3)
    fmt.Println(result) // Outputs: 7
}
```

### Dropping Arguments

```go
package main

import (
    "fmt"
    "github.com/parametalol/curry"
)

func greet(name string) string {
    return "Hello, " + name
}

func main() {
    greetWithoutName := curry.Drop1(greet)
    result := greetWithoutName()
    fmt.Println(result) // Outputs: Hello, 
}
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License.
