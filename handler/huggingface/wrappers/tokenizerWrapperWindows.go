//go:build windows

package wrappers

/*
#cgo windows LDFLAGS: -L"${SRCDIR}/../lib/win_x64" -ltokenizers -lm -lstdc++
*/
import "C"
