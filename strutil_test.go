// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package strutil

import "testing"

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
	}

	for _, test := range tests {
		if v := InitialToUpper(test[0]); v != test[1] {
			t.Errorf("expected %q, got %q", test[1], v)
		}
	}
}
