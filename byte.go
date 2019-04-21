// Copyright 2018 Romano de Souza. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package coalesce

// Byte returns the first non-nil value in a list of pointers to byte.
func Byte(args ...*byte) *byte {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// ByteSlice returns the first non-nil value in a list of slice of byte.
func ByteSlice(args ...[]byte) []byte {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}
