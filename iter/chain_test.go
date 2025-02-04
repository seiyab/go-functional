package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleChain() {
	fmt.Println(iter.Collect[int](iter.Chain[int](iter.Lift([]int{1, 2}), iter.Lift([]int{3, 4}), iter.Lift([]int{0, 9}))))
	// Output: [1 2 3 4 0 9]
}

func TestChainMultiple(t *testing.T) {
	items := iter.Chain[int](iter.Lift([]int{1, 2}), iter.Lift([]int{3, 4}))
	assert.Equal(t, items.Next().Unwrap(), 1)
	assert.Equal(t, items.Next().Unwrap(), 2)
	assert.Equal(t, items.Next().Unwrap(), 3)
	assert.Equal(t, items.Next().Unwrap(), 4)
	assert.True(t, items.Next().IsNone())
}

func TestChainSingle(t *testing.T) {
	items := iter.Chain[int](iter.Lift([]int{1, 2}))
	assert.Equal(t, items.Next().Unwrap(), 1)
	assert.Equal(t, items.Next().Unwrap(), 2)
	assert.True(t, items.Next().IsNone())
}

func TestChainEmpty(t *testing.T) {
	assert.True(t, iter.Chain[int]().Next().IsNone())
}
