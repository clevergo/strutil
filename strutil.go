// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package strutil

import "unicode"

// InitialToLower converts initial to lower.
func InitialToLower(s string) string {
	for i, v := range s {
		return string(unicode.ToLower(v)) + s[i+1:]
	}

	return ""
}

// InitialToUpper converts initial to upper.
func InitialToUpper(s string) string {
	for i, v := range s {
		return string(unicode.ToUpper(v)) + s[i+1:]
	}

	return ""
}
