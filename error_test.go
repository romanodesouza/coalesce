// Copyright 2018 Romano de Souza. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package coalesce_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/romanodesouza/coalesce"
)

func TestError(t *testing.T) {
	tests := []struct {
		name string
		in   []error
		want error
	}{
		{
			name: "it should return the first non-nil value",
			in:   []error{nil, errors.New("err: a"), errors.New("err: b")},
			want: errors.New("err: a"),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []error{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Error(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestErrorSlice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]error
		want []error
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]error{nil, []error{errors.New("err: a")}, []error{errors.New("err: b")}},
			want: []error{errors.New("err: a")},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]error{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.ErrorSlice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}
