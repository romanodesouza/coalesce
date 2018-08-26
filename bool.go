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
