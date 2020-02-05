// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package strutil

import (
	"unicode"
)

// InitialToLower converts initial to lower.
func InitialToLower(s string) string {
	for _, r := range s {
		u := string(unicode.ToLower(r))
		return u + s[len(u):]
	}

	return s
}

// InitialToUpper converts initial to upper.
func InitialToUpper(s string) string {
	for _, r := range s {
		u := string(unicode.ToUpper(r))
		return u + s[len(u):]
	}

	return ""
}
