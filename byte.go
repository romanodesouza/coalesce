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
