// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package strutil

import (
	"math/rand"
	"time"
	"unicode"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

const defaultRandomSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Random generates a fixed length random string from the default random set.
func Random(length int) string {
	return RandomSet(length, defaultRandomSet)
}

// RandomSet generates a fixed length random string from the given set.
func RandomSet(length int, set string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = set[rand.Intn(len(set))]
	}
	return string(b)
}
