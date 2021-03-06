// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package strutil

import (
	"strings"
	"testing"
)

func TestInitialToLower(t *testing.T) {
	tests := [][]string{
		{"", ""},
		{"foo", "foo"},
		{"fOo", "fOo"},
		{"foO", "foO"},
		{"fOO", "fOO"},
		{"Foo", "foo"},
		{"FOo", "fOo"},
		{"FoO", "foO"},
		{"FOO", "fOO"},
		{"Foo bar", "foo bar"},
		{"中文", "中文"},
	}

	for _, test := range tests {
		if v := InitialToLower(test[0]); v != test[1] {
			t.Errorf("expected %q, got %q", test[1], v)
		}
	}
}

func TestInitialToUpper(t *testing.T) {
	tests := [][]string{
		{"", ""},
		{"foo", "Foo"},
		{"fOo", "FOo"},
		{"foO", "FoO"},
		{"fOO", "FOO"},
		{"Foo", "Foo"},
		{"FOo", "FOo"},
		{"FoO", "FoO"},
		{"FOO", "FOO"},
		{"foo bar", "Foo bar"},
		{"中文", "中文"},
	}

	for _, test := range tests {
		if v := InitialToUpper(test[0]); v != test[1] {
			t.Errorf("expected %q, got %q", test[1], v)
		}
	}
}

func TestRandom(t *testing.T) {
	for _, length := range []int{1, 2, 4, 8, 16, 32} {
		s := Random(length)
		if len(s) != length {
			t.Errorf("expected length %d, got %d", length, len(s))
		}
		for _, r := range s {
			if !strings.ContainsRune(defaultRandomSet, r) {
				t.Errorf("unexpected %s", string(r))
			}
		}
	}
}
