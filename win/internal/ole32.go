/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package proc

import (
	"syscall"
)

var (
	dllOle32 = syscall.NewLazyDLL("ole32.dll")

	CoCreateInstance = dllOle32.NewProc("CoCreateInstance")
	CoInitializeEx   = dllOle32.NewProc("CoInitializeEx")
	CoTaskMemFree    = dllOle32.NewProc("CoTaskMemFree")
	CoUninitialize   = dllOle32.NewProc("CoUninitialize")
)