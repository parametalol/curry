# curry

[![CI](https://github.com/parametalol/curry/actions/workflows/go.yml/badge.svg)](https://github.com/parametalol/curry/actions/workflows/go.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/parametalol/curry.svg)](https://pkg.go.dev/github.com/parametalol/curry) [![Go Report Card](https://goreportcard.com/badge/github.com/parametalol/curry)](https://goreportcard.com/report/github.com/parametalol/curry) [![License](https://img.shields.io/github/license/parametalol/curry)](./LICENSE)


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
- **Return Values Dropping**: Ignore first or last function return values.
- **Lazy Arguments Evaluation**: Convert function parameters to thunk functions.

## Installation

```bash
go get github.com/parametalol/curry
```

## Usage

See [examples](./example_test.go) for usages.


## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License.
