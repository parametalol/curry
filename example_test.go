package curry_test

import (
	"fmt"
	"slices"

	"github.com/parametalol/curry"
	"github.com/parametalol/curry/seq"
)

func ExampleTwo() {
	multiply := func(a, b int) int {
		return a * b
	}
	curriedMultiply := curry.Two(multiply)
	multiplyByTwo := curriedMultiply(2)
	result := multiplyByTwo(5)
	fmt.Println(result) // Output: 10
}

func ExampleUnTwo() {
	curriedAdd := func(a int) func(int) int {
		return func(b int) int {
			return a + b
		}
	}

	uncurriedAdd := curry.UnTwo(curriedAdd)
	result := uncurriedAdd(3, 4)
	fmt.Println(result) // Output: 7
}

func ExampleBindFirstOfTwo() {
	subtract := func(a, b int) int {
		return a - b
	}

	subtractFromTen := curry.BindFirstOfTwo(subtract, 10)
	result := subtractFromTen(3)
	fmt.Println(result) // Output: 7
}

func ExampleDropLastOfTwo() {
	f := func() (int, error) {
		return 1, nil
	}
	result := curry.DropLastOfTwo(f())
	fmt.Println(result) // Output: 1
}

func ExampleAdaptOne() {
	i := 0
	odd := func() bool {
		i++
		return i&1 == 0
	}
	fmt.Println(slices.DeleteFunc([]int{1, 2, 3, 4, 5},
		curry.AdaptOne[int, bool](odd)))
	// Output:
	// [1 3 5]
}

func ExampleReturn() {
	// Return binds return value to a function without parameters.
	// It can be used for passing static value to functions that accept
	// functions.
	f := func(fn func() string) {
		fmt.Println(fn())
	}
	f(curry.Return("message"))
	// Output: message
}

func ExampleSame() {
	fmt.Println(slices.Collect(seq.Take(5, seq.Generate(curry.Same))))
	// Output: [0 1 2 3 4]
}

func ExampleLazyOne0() {

	process := func(expensive string) {
		fmt.Println("That was", expensive)
	}
	expensive := func() string {
		fmt.Println("Computing...")
		return "expensive "
	}

	defer curry.LazyOne0(process)(expensive)

	fmt.Println("The expensive process argument hasn't been computed yet.")

	// Output:
	// The expensive process argument hasn't been computed yet.
	// Computing...
	// That was expensive
}
