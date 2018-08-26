package coalesce

// Uint returns the first non-nil value in a list of pointers to uint.
func Uint(args ...*uint) *uint {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// UintSlice returns the first non-nil value in a list of slice of uint.
func UintSlice(args ...[]uint) []uint {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Uint8 returns the first non-nil value in a list of pointers to uint8.
func Uint8(args ...*uint8) *uint8 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Uint8Slice returns the first non-nil value in a list of slice of uint8.
func Uint8Slice(args ...[]uint8) []uint8 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Uint16 returns the first non-nil value in a list of pointers to uint16.
func Uint16(args ...*uint16) *uint16 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Uint16Slice returns the first non-nil value in a list of slice of uint16.
func Uint16Slice(args ...[]uint16) []uint16 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Uint32 returns the first non-nil value in a list of pointers to uint32.
func Uint32(args ...*uint32) *uint32 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Uint32Slice returns the first non-nil value in a list of slice of uint32.
func Uint32Slice(args ...[]uint32) []uint32 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Uint64 returns the first non-nil value in a list of pointers to uint64.
func Uint64(args ...*uint64) *uint64 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}

// Uint64Slice returns the first non-nil value in a list of slice of uint64.
func Uint64Slice(args ...[]uint64) []uint64 {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}
