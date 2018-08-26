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
