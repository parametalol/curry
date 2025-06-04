package curry_test

import (
	"fmt"

	"github.com/parametalol/curry"
)

func ExampleTwo() {
	multiply := func(a, b int) int {
		return a * b
	}
	curriedMultiply := curry.Two(multiply)
	multiplyByTwo := curriedMultiply(2)
	result := multiplyByTwo(5)
	fmt.Println(result) // Outputs: 10
}

func ExampleUnTwo() {
	curriedAdd := func(a int) func(int) int {
		return func(b int) int {
			return a + b
		}
	}

	uncurriedAdd := curry.UnTwo(curriedAdd)
	result := uncurriedAdd(3, 4)
	fmt.Println(result) // Outputs: 7
}

func ExampleBindFirstOfTwo() {
	subtract := func(a, b int) int {
		return a - b
	}

	subtractFromTen := curry.BindFirstOfTwo(subtract, 10)
	result := subtractFromTen(3)
	fmt.Println(result) // Outputs: 7
}

func ExampleDropLastOfTwo() {
	f := func() (int, error) {
		return 1, nil
	}
	result := curry.DropLastOfTwo(f())
	fmt.Println(result) // Outputs: 1
}
