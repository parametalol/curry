package seq_test

import (
	"fmt"
	"slices"

	"github.com/parametalol/curry/seq"
)

func ExampleIndex() {
	fmt.Println(slices.Collect(
		seq.Take(5,
			seq.Generate(seq.Index(5)))))
	// Output: [5 6 7 8 9]
}
