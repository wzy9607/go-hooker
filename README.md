# go-hooker

[![ci](https://github.com/wzy9607/go-hooker/actions/workflows/pull-request.yml/badge.svg)](https://github.com/wzy9607/go-hooker/actions/workflows/pull-request.yml)
[![codecov](https://codecov.io/gh/wzy9607/go-hooker/graph/badge.svg?token=VVMWEWOQFO)](https://codecov.io/gh/wzy9607/go-hooker)

Golang Generic Hook

## installation

```bash
go get github.com/dingyaguang117/go-hooker
```

## How to use

```go
package main

import (
	"fmt"
	
	"github.com/dingyaguang117/go-hooker/hooker"
)

type Func func(a, b int) int

func SomeFunc(a, b int) int {
	fmt.Printf("%d + %d\n", a, b)
	return a + b
}

func Hook1(next Func) Func {
	return func(a, b int) int {
		fmt.Println("----- Hook1 before ----")
		c := next(a, b)
		fmt.Println("----- Hook1 after ----")
		return c
	}
}

func Hook2(next Func) Func {
	return func(a, b int) int {
		fmt.Println("----- Hook2 before ----")
		c := next(a, b)
		fmt.Println("----- Hook2 after ----")
		return c
	}
}

func Hook3(next Func) Func {
	return func(a, b int) int {
		fmt.Println("----- Hook3 before ----")
		c := next(a, b)
		fmt.Println("----- Hook3 after ----")
		return c
	}
}

func main() {
	h := hooker.NewHooker[Func](SomeFunc, Hook1, Hook2)
	h.AddHook(Hook3)
	h.GetWrapped()(1, 2)
}
```

output

```
----- Hook1 before ----
----- Hook2 before ----
----- Hook3 before ----
1 + 2
----- Hook3 after ----
----- Hook2 after ----
----- Hook1 after ----
```
