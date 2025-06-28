package curry_test

import (
	"fmt"
	"slices"
	"strings"

	"github.com/parametalol/curry"
	"github.com/parametalol/curry/seq"
)

func ExampleF2R() {
	multiply := func(a, b int) int {
		return a * b
	}
	curriedMultiply := curry.F2R(multiply)
	multiplyByTwo := curriedMultiply(2)
	result := multiplyByTwo(5)
	fmt.Println(result) // Output: 10
}

func ExampleUn2R() {
	curriedAdd := func(a int) func(int) int {
		return func(b int) int {
			return a + b
		}
	}

	uncurriedAdd := curry.Un2R(curriedAdd)
	result := uncurriedAdd(3, 4)
	fmt.Println(result) // Output: 7
}

func ExampleBindFirstOf2R() {
	subtract := func(a, b int) int {
		return a - b
	}

	subtractFromTen := curry.BindFirstOf2R(subtract, 10)
	result := subtractFromTen(3)
	fmt.Println(result) // Output: 7
}

func ExampleDropLastOf2() {
	f := func() (int, error) {
		return 1, nil
	}
	result := curry.DropLastOf2(f())
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
	lazyF := curry.Lazy2(f)
	// Bind "first":
	lazyBound := curry.BindFirstOf2(
		lazyF, curry.Thunk("first"))

	lazyBound(curry.Thunk("second"))
	// Output: first second
}

func ExampleLazy1() {

	process := func(expensive string) {
		fmt.Println("That was", expensive)
	}
	expensive := func() string {
		fmt.Println("Computing...")
		return "expensive "
	}

	// The expensive process argument is not computed right now, but only when
	// defer executes.
	defer curry.Lazy1(process)(expensive)

	fmt.Println("The expensive process argument hasn't been computed yet.")

	// Output:
	// The expensive process argument hasn't been computed yet.
	// Computing...
	// That was expensive
}

func ExampleWrap() {
	// isValue(string) int is a function that compares a string to "value".
	isValue := curry.BindLastOf2R(strings.Compare, "value")
	isZero := curry.F2R(curry.Eq[int])(0)

	// Construct a chain of processors that returns true if a given string
	// is equal to "value" ignoring case when trimmed.
	// Return in the end is just for the next lines alignment.
	chain := curry.Wrap(curry.Wrap(curry.Wrap(curry.Return(
		strings.TrimSpace, // string -> string
	), strings.ToLower, // string -> string
	), isValue, // string -> int
	), isZero, // int -> bool
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

func ExampleNot() {
	fruits := []string{"banana", "banana", "orange", "banana"}

	isBanana := curry.BindFirstOf2R(curry.Eq, "banana")

	notBanana := curry.Wrap(isBanana, curry.Not)

	fmt.Println(slices.Collect(
		seq.Filter(slices.Values(fruits),
			notBanana),
	))
	// Output:
	// [orange]
}
