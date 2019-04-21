package coalesce_test

import (
	"reflect"
	"testing"

	"github.com/romanodesouza/coalesce"
)

func TestFloat32(t *testing.T) {
	tests := []struct {
		name string
		in   []*float32
		want *float32
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*float32{nil, ptrfloat32(1), ptrfloat32(2)},
			want: ptrfloat32(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*float32{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Float32(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestFloat32Slice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]float32
		want []float32
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]float32{nil, []float32{1}, []float32{2}},
			want: []float32{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]float32{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Float32Slice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		name string
		in   []*float64
		want *float64
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*float64{nil, ptrfloat64(1), ptrfloat64(2)},
			want: ptrfloat64(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*float64{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Float64(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestFloat64Slice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]float64
		want []float64
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]float64{nil, []float64{1}, []float64{2}},
			want: []float64{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]float64{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Float64Slice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func ptrfloat32(f float32) *float32 {
	return &f
}

func ptrfloat64(f float64) *float64 {
	return &f
}
