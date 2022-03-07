
package result_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/result"
)

func ExampleResult_Unwrap() {
	fmt.Println(result.Ok(4).Unwrap())
	// Output: 4
}

func ExampleResult_UnwrapOr() {
	fmt.Println(result.Ok(4).UnwrapOr(3))
	fmt.Println(result.Err[int](errors.New("oops")).UnwrapOr(3))
	// Output:
	// 4
	// 3
}

func ExampleResult_UnwrapOrElse() {
	fmt.Println(result.Ok(4).UnwrapOrElse(func() int {
		return 3
	}))

	fmt.Println(result.Err[int](errors.New("oops")).UnwrapOrElse(func() int {
		return 3
	}))

	// Output:
	// 4
	// 3
}

func ExampleResult_UnwrapOrZero() {
	fmt.Println(result.Ok(4).UnwrapOrZero())
	fmt.Println(result.Err[int](errors.New("oops")).UnwrapOrZero())

	// Output
	// 4
	// 0
}

func ExampleResult_IsOk() {
	fmt.Println(result.Ok(4).IsOk())
	fmt.Println(result.Err[int](errors.New("oops")).IsOk())

	// Output:
	// true
	// false
}

func ExampleResult_IsErr() {
	fmt.Println(result.Ok(4).IsErr())
	fmt.Println(result.Err[int](errors.New("oops")).IsErr())

	// Output:
	// false
	// true
}

func ExampleResult_Value() {
	value, ok := result.Ok(4).Value()
	fmt.Println(value)
	fmt.Println(ok)

	// Output:
	// 4
	// <nil>
}
