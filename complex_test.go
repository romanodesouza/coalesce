// Copyright 2018 Romano de Souza. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package coalesce_test

import (
	"reflect"
	"testing"

	"github.com/romanodesouza/coalesce"
)

func TestComplex64(t *testing.T) {
	tests := []struct {
		name string
		in   []*complex64
		want *complex64
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*complex64{nil, ptrcomplex64(1), ptrcomplex64(2)},
			want: ptrcomplex64(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*complex64{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Complex64(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestComplex64Slice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]complex64
		want []complex64
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]complex64{nil, []complex64{1}, []complex64{2}},
			want: []complex64{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]complex64{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Complex64Slice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128(t *testing.T) {
	tests := []struct {
		name string
		in   []*complex128
		want *complex128
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*complex128{nil, ptrcomplex128(1), ptrcomplex128(2)},
			want: ptrcomplex128(1),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*complex128{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Complex128(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestComplex128Slice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]complex128
		want []complex128
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]complex128{nil, []complex128{1}, []complex128{2}},
			want: []complex128{1},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]complex128{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Complex128Slice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func ptrcomplex64(f complex64) *complex64 {
	return &f
}

func ptrcomplex128(f complex128) *complex128 {
	return &f
}
