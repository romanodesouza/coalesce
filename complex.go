package coalesce

// Complex64 returns the first non-nil value in a list of pointers to complex64.
func Complex64(args ...*complex64) *complex64 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Complex64Slice returns the first non-nil value in a list of slice of complex64.
func Complex64Slice(args ...[]complex64) []complex64 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Complex128 returns the first non-nil value in a list of pointers to complex128.
func Complex128(args ...*complex128) *complex128 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Complex128Slice returns the first non-nil value in a list of slice of complex128.
func Complex128Slice(args ...[]complex128) []complex128 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}
