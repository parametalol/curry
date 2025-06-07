package curry

import (
	"testing"
)

func TestDrop(t *testing.T) {
	Assert(t, Equal(1, DropLastOfTwo(1, 2)))
	Assert(t, Equal(1, DropLastOfTwo(DropLastOfThree(1, 2, 3))))
	Assert(t, Equal(1, DropLastOfTwo(DropLastOfThree(DropLastOfFour(1, 2, 3, 4)))))

	Assert(t, Equal(2, DropFirstOfTwo(1, 2)))
	Assert(t, Equal(3, DropFirstOfTwo(DropFirstOfThree(1, 2, 3))))
	Assert(t, Equal(4, DropFirstOfTwo(DropFirstOfThree(DropFirstOfFour(1, 2, 3, 4)))))
}
