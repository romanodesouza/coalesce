// Copyright 2018 Romano de Souza. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package coalesce

// Error returns the first non-nil value in a list of error.
func Error(args ...error) error {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// ErrorSlice returns the first non-nil value in a list of error.
func ErrorSlice(args ...[]error) []error {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}
