package curry

import (
	"fmt"
	"io"
	"testing"
)

func TestUnCurry(t *testing.T) {
	curriedJoin := Four(join4)

	Assert(t,
		Equal("abcd", UnFour(curriedJoin)("a", "b", "c", "d")),
		Equal("abcd", UnThree(curriedJoin)("a", "b", "c")("d")),
		Equal("abcd", UnTwo(curriedJoin)("a", "b")("c")("d")),
	)

	curriedJoinE := Four2(join4e)
	Assert(t,
		Equal("abcd", DropLastOfTwo(UnFour2(curriedJoinE)("a", "b", "c", "d"))),
		Equal("abcd", DropLastOfTwo(UnThree(curriedJoinE)("a", "b", "c")("d"))),
		Equal("abcd", DropLastOfTwo(UnTwo(curriedJoinE)("a", "b")("c")("d"))),
	)
}

func TestUnCurrySlice(t *testing.T) {
	Assert(t,
		Equal("a-b-c",
			UnTwoSlice(TwoSlice(fmt.Sprintf))("%s-%s-%s", "a", "b", "c")),

		Equal(13, DropLastOfTwo(
			UnTwoSlice2(TwoSlice2(fmt.Printf))("- %s - %s -", "abc", "def"))),

		Equal("5: [a b]",
			UnThreeSlice(ThreeSlice(join2slice))("%d: %s", 5, "a", "b")),

		Equal(5, DropLastOfTwo(
			UnThreeSlice2(ThreeSlice2(fmt.Fprintf))(io.Discard, "%s-%s-%s", "a", "b", "c"))),

		Equal("5: [a b]",
			UnTwo(ThreeSlice(join2slice))("%d: %s", 5)("a", "b")),
	)
}
