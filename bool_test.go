// Copyright 2018 Romano de Souza. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package coalesce_test

import (
	"reflect"
	"testing"

	"github.com/romanodesouza/coalesce"
)

func TestBool(t *testing.T) {
	tests := []struct {
		name string
		in   []*bool
		want *bool
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*bool{nil, ptrbool(true), ptrbool(false)},
			want: ptrbool(true),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*bool{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Bool(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestBoolSlice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]bool
		want []bool
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]bool{nil, []bool{true}, []bool{false}},
			want: []bool{true},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]bool{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.BoolSlice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func ptrbool(b bool) *bool {
	return &b
}
