// Copyright 2018 Romano de Souza. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package coalesce

// Bool returns the first non-nil value in a list of pointers to bool.
func Bool(args ...*bool) *bool {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// BoolSlice returns the first non-nil value in a list of slice of bool.
func BoolSlice(args ...[]bool) []bool {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}
