package curry_test

import (
	"fmt"
	"github.com/parametalol/curry"
)

func ExampleCurry2() {
	multiply := func(a, b int) int {
		return a * b
	}
	curriedMultiply := curry.Curry2(multiply)
	multiplyByTwo := curriedMultiply(2)
	result := multiplyByTwo(5)
	fmt.Println(result) // Outputs: 10
}

func ExampleUnCurry2() {
	curriedAdd := func(a int) func(int) int {
		return func(b int) int {
			return a + b
		}
	}

	uncurriedAdd := curry.UnCurry2(curriedAdd)
	result := uncurriedAdd(3, 4)
	fmt.Println(result) // Outputs: 7
}

func ExampleBind1st2() {
	subtract := func(a, b int) int {
		return a - b
	}

	subtractFromTen := curry.Bind1st2(subtract, 10)
	result := subtractFromTen(3)
	fmt.Println(result) // Outputs: 7
}

func ExampleDropLast2() {
	f := func() (int, error) {
		return 1, nil
	}
	result := curry.DropLast2(f())
	fmt.Println(result) // Outputs: 1
}
