// Copyright 2018 Romano de Souza. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package coalesce_test

import (
	"reflect"
	"testing"

	"github.com/romanodesouza/coalesce"
)

func TestByte(t *testing.T) {
	tests := []struct {
		name string
		in   []*byte
		want *byte
	}{
		{
			name: "it should return the first non-nil value",
			in:   []*byte{nil, ptrbyte(byte('a')), ptrbyte(byte('b'))},
			want: ptrbyte('a'),
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   []*byte{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.Byte(tt.in...)
			if tt.want != nil && *got != *tt.want {
				t.Errorf("got %v; want %v", *got, *tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", *got, tt.want)
			}
		})
	}
}

func TestByteSlice(t *testing.T) {
	tests := []struct {
		name string
		in   [][]byte
		want []byte
	}{
		{
			name: "it should return the first non-nil value",
			in:   [][]byte{nil, []byte{byte('a')}, []byte{byte('b')}},
			want: []byte{byte('a')},
		},
		{
			name: "it should return nil if a non-nil value has not been found",
			in:   [][]byte{nil, nil, nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coalesce.ByteSlice(tt.in...)
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
			if tt.want == nil && got != nil {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func ptrbyte(b byte) *byte {
	return &b
}
