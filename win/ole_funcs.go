package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Used to retrieve class IDs to create COM Automation objects.
//
// If the progId is invalid, error returns errco.CO_E_CLASSSTRING.
//
// Example:
//
//  clsId, _ := win.CLSIDFromProgID("Excel.Application")
//
//  mainObj := win.CoCreateInstance(
//      clsId, nil, co.CLSCTX_SERVER, co.IID_IUNKNOWN)
//  defer mainObj.Release()
//
//  excel := mainObj.QueryInterface(automco.IID_IDispatch)
//  defer excel.Release()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-clsidfromprogid
func CLSIDFromProgID(progId string) (co.CLSID, error) {
	var guid GUID
	ret, _, _ := syscall.Syscall(proc.CLSIDFromProgID.Addr(), 2,
		uintptr(unsafe.Pointer(Str.ToNativePtr(progId))),
		uintptr(unsafe.Pointer(&guid)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return co.CLSID(guid.String()), nil
	} else {
		return "", hr
	}
}

// Loads the COM module. This needs to be done only once in your application.
// Typically uses COINIT_APARTMENTTHREADED.
//
// ⚠️ You must defer CoUninitialize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
func CoInitializeEx(coInit co.COINIT) {
	ret, _, _ := syscall.Syscall(proc.CoInitializeEx.Addr(), 2,
		0, uintptr(coInit), 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}

// ⚠️ You must defer CoTaskMemFree().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
func CoTaskMemAlloc(size int) uintptr {
	ret, _, _ := syscall.Syscall(proc.CoTaskMemAlloc.Addr(), 1,
		uintptr(size), 0, 0)
	if ret == 0 {
		panic("CoTaskMemAlloc() failed.")
	}
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemfree
func CoTaskMemFree(pv uintptr) {
	syscall.Syscall(proc.CoTaskMemFree.Addr(), 1,
		pv, 0, 0)
}

// ⚠️ You must defer CoTaskMemFree().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemrealloc
func CoTaskMemRealloc(pv uintptr, size int) uintptr {
	ret, _, _ := syscall.Syscall(proc.CoTaskMemRealloc.Addr(), 2,
		pv, uintptr(size), 0)
	if ret == 0 {
		panic("CoTaskMemRealloc() failed.")
	}
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-couninitialize
func CoUninitialize() {
	syscall.Syscall(proc.CoUninitialize.Addr(), 0, 0, 0, 0)
}
