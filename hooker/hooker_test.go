package hooker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHooker_AddHook(t *testing.T) {
	t.Parallel()
	type handler func(a int) int
	hooker := NewHooker[handler](func(a int) int {
		assert.Equal(t, 6, a)
		return a + 4
	})
	hook1 := func(next handler) handler {
		return func(a int) int {
			assert.Equal(t, 0, a)
			return next(a + 1)
		}
	}
	hook2 := func(next handler) handler {
		return func(a int) int {
			assert.Equal(t, 1, a)
			return next(a + 2)
		}
	}
	hook3 := func(next handler) handler {
		return func(a int) int {
			assert.Equal(t, 3, a)
			return next(a + 3)
		}
	}
	hooker.AddHook(hook1, hook2)
	hooker.AddHook(hook3)

	b := hooker.GetWrapped()(0)
	assert.Equal(t, 10, b)

	c := hooker.GetOrigin()(6)
	assert.Equal(t, 10, c)
}

func ExampleHooker_AddHook() {
	type handler func(a, b int) int
	hooker := NewHooker[handler](func(a, b int) int {
		fmt.Printf("%d + %d\n", a, b)
		return a + b
	})
	hook1 := func(next handler) handler {
		return func(a, b int) int {
			fmt.Println("----- hook1 before ----")
			c := next(a, b)
			fmt.Println("----- hook1 after ----")
			return c
		}
	}
	hook2 := func(next handler) handler {
		return func(a, b int) int {
			fmt.Println("----- hook2 before ----")
			c := next(a, b)
			fmt.Println("----- hook2 after ----")
			return c
		}
	}
	hook3 := func(next handler) handler {
		return func(a, b int) int {
			fmt.Println("----- hook3 before ----")
			c := next(a, b)
			fmt.Println("----- hook3 after ----")
			return c
		}
	}
	hooker.AddHook(hook1, hook2)
	hooker.AddHook(hook3)
	c := hooker.GetWrapped()(1, 2)
	fmt.Println(c)
	// Output:
	// ----- hook1 before ----
	// ----- hook2 before ----
	// ----- hook3 before ----
	// 1 + 2
	// ----- hook3 after ----
	// ----- hook2 after ----
	// ----- hook1 after ----
	// 3
}

func TestHooker_GetHooks(t *testing.T) {
	t.Parallel()
	type handler func(a int) int
	hooker := NewHooker[handler](func(a int) int { return a })
	hooker.AddHook(func(next handler) handler { return func(a int) int { return next(a) } })
	hooker.AddHook(func(next handler) handler { return func(a int) int { return next(a) } })
	hooks := hooker.GetHooks()
	assert.Len(t, hooks, 2)

	// test hooks are cloned
	hooks[1] = func(next handler) handler { return func(a int) int { return next(a + 1) } }
	hooker.AddHook(func(next handler) handler { return func(a int) int { return next(a) } })
	c := hooker.GetWrapped()(0)
	assert.Equal(t, 0, c)
}
