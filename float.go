// Copyright 2018 Romano de Souza. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package coalesce

// Float32 returns the first non-nil value in a list of pointers to float32.
func Float32(args ...*float32) *float32 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Float32Slice returns the first non-nil value in a list of slice of float32.
func Float32Slice(args ...[]float32) []float32 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Float64 returns the first non-nil value in a list of pointers to float64.
func Float64(args ...*float64) *float64 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Float64Slice returns the first non-nil value in a list of slice of float64.
func Float64Slice(args ...[]float64) []float64 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}
