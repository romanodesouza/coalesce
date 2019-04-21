// Copyright 2018 Romano de Souza. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package coalesce

// Int returns the first non-nil value in a list of pointers to int.
func Int(args ...*int) *int {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// IntSlice returns the first non-nil value in a list of slice of int.
func IntSlice(args ...[]int) []int {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Int8 returns the first non-nil value in a list of pointers to int8.
func Int8(args ...*int8) *int8 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Int8Slice returns the first non-nil value in a list of slice of int8.
func Int8Slice(args ...[]int8) []int8 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Int16 returns the first non-nil value in a list of pointers to int16.
func Int16(args ...*int16) *int16 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Int16Slice returns the first non-nil value in a list of slice of int16.
func Int16Slice(args ...[]int16) []int16 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Int32 returns the first non-nil value in a list of pointers to int32.
func Int32(args ...*int32) *int32 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Int32Slice returns the first non-nil value in a list of slice of int32.
func Int32Slice(args ...[]int32) []int32 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Int64 returns the first non-nil value in a list of pointers to int64.
func Int64(args ...*int64) *int64 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Int64Slice returns the first non-nil value in a list of slice of int64.
func Int64Slice(args ...[]int64) []int64 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}
