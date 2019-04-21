package coalesce_test

import (
	"reflect"
	"testing"

	"github.com/romanodesouza/coalesce"
)

func TestUint(t *testing.T) {
	tests := []struct {
		name string
		in   []*uint
		want *uint
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*uint{nil, ptruint(1), ptruint(2)},
			want: ptruint(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*uint{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Uint(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestUintSlice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]uint
		want []uint
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]uint{nil, []uint{1}, []uint{2}},
			want: []uint{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]uint{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.UintSlice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestUint8(t *testing.T) {
	tests := []struct {
		name string
		in   []*uint8
		want *uint8
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*uint8{nil, ptruint8(1), ptruint8(2)},
			want: ptruint8(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*uint8{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Uint8(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestUint8Slice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]uint8
		want []uint8
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]uint8{nil, []uint8{1}, []uint8{2}},
			want: []uint8{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]uint8{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Uint8Slice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	tests := []struct {
		name string
		in   []*uint16
		want *uint16
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*uint16{nil, ptruint16(1), ptruint16(2)},
			want: ptruint16(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*uint16{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Uint16(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestUint16Slice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]uint16
		want []uint16
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]uint16{nil, []uint16{1}, []uint16{2}},
			want: []uint16{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]uint16{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Uint16Slice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	tests := []struct {
		name string
		in   []*uint32
		want *uint32
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*uint32{nil, ptruint32(1), ptruint32(2)},
			want: ptruint32(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*uint32{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Uint32(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestUint32Slice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]uint32
		want []uint32
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]uint32{nil, []uint32{1}, []uint32{2}},
			want: []uint32{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]uint32{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Uint32Slice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	tests := []struct {
		name string
		in   []*uint64
		want *uint64
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*uint64{nil, ptruint64(1), ptruint64(2)},
			want: ptruint64(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*uint64{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Uint64(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestUint64Slice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]uint64
		want []uint64
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]uint64{nil, []uint64{1}, []uint64{2}},
			want: []uint64{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]uint64{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Uint64Slice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func ptruint(i uint) *uint {
	return &i
}

func ptruint8(i uint8) *uint8 {
	return &i
}

func ptruint16(i uint16) *uint16 {
	return &i
}

func ptruint32(i uint32) *uint32 {
	return &i
}

func ptruint64(i uint64) *uint64 {
	return &i
}
