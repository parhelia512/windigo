package com

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/ayy/com/comco"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Loads the COM module. This needs to be done only once in your application.
// Typically uses COINIT_APARTMENTTHREADED.
//
// ⚠️ You must defer CoUninitialize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
func CoInitializeEx(coInit comco.COINIT) {
	ret, _, _ := syscall.Syscall(proc.CoInitializeEx.Addr(), 2,
		0, uintptr(coInit), 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-couninitialize
func CoUninitialize() {
	syscall.Syscall(proc.CoUninitialize.Addr(), 0, 0, 0, 0)
}
