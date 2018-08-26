package coalesce_test

import (
	"reflect"
	"testing"

	"github.com/romanoaugusto88/coalesce"
)

func TestInt(t *testing.T) {
	tests := []struct {
		name string
		in   []*int
		want *int
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*int{nil, ptrint(1), ptrint(2)},
			want: ptrint(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*int{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Int(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestIntSlice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		want []int
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]int{nil, []int{1}, []int{2}},
			want: []int{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]int{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.IntSlice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	tests := []struct {
		name string
		in   []*int8
		want *int8
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*int8{nil, ptrint8(1), ptrint8(2)},
			want: ptrint8(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*int8{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Int8(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestInt8Slice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int8
		want []int8
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]int8{nil, []int8{1}, []int8{2}},
			want: []int8{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]int8{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Int8Slice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	tests := []struct {
		name string
		in   []*int16
		want *int16
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*int16{nil, ptrint16(1), ptrint16(2)},
			want: ptrint16(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*int16{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Int16(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestInt16Slice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int16
		want []int16
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]int16{nil, []int16{1}, []int16{2}},
			want: []int16{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]int16{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Int16Slice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	tests := []struct {
		name string
		in   []*int32
		want *int32
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*int32{nil, ptrint32(1), ptrint32(2)},
			want: ptrint32(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*int32{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Int32(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestInt32Slice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int32
		want []int32
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]int32{nil, []int32{1}, []int32{2}},
			want: []int32{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]int32{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Int32Slice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	tests := []struct {
		name string
		in   []*int64
		want *int64
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*int64{nil, ptrint64(1), ptrint64(2)},
			want: ptrint64(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*int64{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Int64(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestInt64Slice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int64
		want []int64
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]int64{nil, []int64{1}, []int64{2}},
			want: []int64{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]int64{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Int64Slice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func ptrint(i int) *int {
	return &i
}

func ptrint8(i int8) *int8 {
	return &i
}

func ptrint16(i int16) *int16 {
	return &i
}

func ptrint32(i int32) *int32 {
	return &i
}

func ptrint64(i int64) *int64 {
	return &i
}
