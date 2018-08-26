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
