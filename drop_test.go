package curry

import (
	"testing"
)

func TestDrop(t *testing.T) {
	DropOne("dropped")
	Assert(t,
		Equal(1, DropLastOfTwo(1, 2)),
		Equal(1, DropLastOfTwo(DropLastOfThree(1, 2, 3))),
		Equal(1, DropLastOfTwo(DropLastOfThree(DropLastOfFour(1, 2, 3, 4)))),

		Equal(2, DropFirstOfTwo(1, 2)),
		Equal(3, DropFirstOfTwo(DropFirstOfThree(1, 2, 3))),
		Equal(4, DropFirstOfTwo(DropFirstOfThree(DropFirstOfFour(1, 2, 3, 4)))),
	)
}
