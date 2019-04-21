// Copyright 2018 Romano de Souza. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package coalesce_test

import (
	"reflect"
	"testing"

	"github.com/romanodesouza/coalesce"
)

func TestString(t *testing.T) {
	tests := []struct {
		name string
		in   []*string
		want *string
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*string{nil, ptrstring("a"), ptrstring("b")},
			want: ptrstring("a"),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*string{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.String(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestStringSlice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]string
		want []string
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]string{nil, []string{"a"}, []string{"b"}},
			want: []string{"a"},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]string{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.StringSlice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func ptrstring(s string) *string {
	return &s
}
