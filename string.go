// Copyright 2018 Romano de Souza. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package coalesce

// String returns the first non-nil value in a list of pointers to string.
func String(args ...*string) *string {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// StringSlice returns the first non-nil value in a list of slice of string.
func StringSlice(args ...[]string) []string {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}
