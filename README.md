# coalesce [![Build Status](https://travis-ci.com/romanodesouza/coalesce.svg?branch=master)](https://travis-ci.com/romanoaugusto88/coalesce) [![codecov](https://codecov.io/gh/romanoaugusto88/coalesce/branch/master/graph/badge.svg)](https://codecov.io/gh/romanoaugusto88/coalesce) [![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/romanoaugusto88/coalesce) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/romanoaugusto88/coalesce/master/LICENSE)

Coalesce built-in Go types.

# motivation
Go doesn't have a [null coalescing operator](https://en.wikipedia.org/wiki/Null_coalescing_operator), so this package brings the capability to return the first non-nil value in a list, pretty much inspired by the `COALESCE()` SQL function.

# install
```
go get github.com/romanoaugusto88/coalesce
```

# example

```go
package main

import (
	"github.com/romanoaugusto88/coalesce"
)

// OptionalParams holds values that are not required.
type OptionalParams struct {
	Int    *int
	String *string
}

// Params represents the "final" structure to be used.
type Params struct {
	Int    int
	String string
}

// DefaultParams holds the optional params fallback values.
var DefaultParams = Params{Int: 10, String: "default string"}

func main() {
	// OptionalParams could be given by user input containing non-required field values.
	optional := OptionalParams{}
	// The coalesce functions are handy when initializing structs:
	params := Params{
		Int:    *coalesce.Int(optional.Int, &DefaultParams.Int),
		String: *coalesce.String(optional.String, &DefaultParams.String),
	}

	// Check out the godoc reference for more functions available, such as:
	_ = *coalesce.Bool(args...)
	_ = coalesce.Error(args...)
	_ = coalesce.StringSlice(args...)
}
```
