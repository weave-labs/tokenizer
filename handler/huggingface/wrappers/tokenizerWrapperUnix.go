//go:build !windows

package wrappers

/*
#cgo !windows LDFLAGS: -ltokenizers -ldl -lm -lstdc++
*/
import "C"
