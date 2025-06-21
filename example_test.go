package curry_test

import (
	"fmt"
	"slices"
	"strings"

	"github.com/parametalol/curry"
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
	// It can be used for binding a value to lazy functions.
	f := func(a string, b string) {
		fmt.Println(a, b)
	}
	lazyF := curry.LazyTwo0(f)
	// Bind "first":
	lazyBound := curry.BindFirstOfTwo0(
		lazyF, curry.Return("first"))

	lazyBound(curry.Return("second"))
	// Output: first second
}

func ExampleLazyOne0() {

	process := func(expensive string) {
		fmt.Println("That was", expensive)
	}
	expensive := func() string {
		fmt.Println("Computing...")
		return "expensive "
	}

	// The expensive process argument is not computed right now, but only when
	// defer executes.
	defer curry.LazyOne0(process)(expensive)

	fmt.Println("The expensive process argument hasn't been computed yet.")

	// Output:
	// The expensive process argument hasn't been computed yet.
	// Computing...
	// That was expensive
}

func ExampleWrap() {
	// isValue(string) int is a function that compares a string to "value".
	isValue := curry.BindLastOfTwo(strings.Compare, "value")

	// Construct a chain of processors that returns true if a given string
	// is equal to "value" ignoring case when trimmed.
	// Pass in the end is just for the next lines alignment.
	chain := curry.Wrap(curry.Wrap(curry.Wrap(curry.Pass(
		strings.TrimSpace, // string -> string
	), strings.ToLower, // string -> string
	), isValue, // string -> int
	), curry.Eq(0), // int -> bool
	)

	// This is equal to:
	//	 chain := func(s string) bool {
	//	 	return strings.Compare(
	//	 		strings.ToLower(
	//	 			strings.TrimSpace(s)), "value") == 0
	//	 }

	fmt.Println(chain("test"), chain("  VALUE  "))
	// Output:
	// false true
}
