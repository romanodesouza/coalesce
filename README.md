# coalesce [![Build Status](https://travis-ci.com/romanoaugusto88/coalesce.svg?branch=master)](https://travis-ci.com/romanoaugusto88/coalesce) [![codecov](https://codecov.io/gh/romanoaugusto88/coalesce/branch/master/graph/badge.svg)](https://codecov.io/gh/romanoaugusto88/coalesce) [![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/romanoaugusto88/coalesce) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/romanoaugusto88/coalesce/master/LICENSE)

Coalesce built-in Go types.

# motivation
Go doesn't has a [null coalescing operator](https://en.wikipedia.org/wiki/Null_coalescing_operator).
This package brings the capability to return the first non-nil value in a list, for each built-in type.

This is pretty much inspired by the `COALESCE()` sql function.

# example

```go
package main

func main() {
	// optionalValue could be a value given by user input which is not required.
	var optionalValue *int
	// defaultValue is the fallback value for the optional value.
	var defaultValue int = 10

	// Then, get the input value or the default one hassle-free.
	value := *coalesce.Int(optionalValue, &defaultValue)

	// ...
}
```
