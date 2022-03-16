package com

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/ayy/com/comco"
	"github.com/rodrigocfd/windigo/win/ayy/com/comvt"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// IUnknown COM interface, base to all COM interfaces.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknown interface {
	// Returns the IID for this COM interface.
	//
	// This is a static method, which can be called from an uninitialized
	// object.
	Iid() co.IID

	// Returns the underlying pointer to pointer to the COM virtual table.
	//
	// This method is used internally by the library, don't use it unless you
	// know what you're doing. Improper use can cause resource leaks.
	Ppv() ***comvt.IUnknown

	// Releases the COM pointer and sets the internal pointer to nil.
	//
	// Never fails, can be called any number of times.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
	Release() uint32
}

type _IUnknown struct{ ppv **comvt.IUnknown }

func (me *_IUnknown) Iid() co.IID {
	return comco.IID_IUnknown
}

func (me *_IUnknown) Ppv() ***comvt.IUnknown {
	return &me.ppv
}

func (me *_IUnknown) Release() uint32 {
	ret := uintptr(0)
	if *me.Ppv() != nil {
		ret, _, _ = syscall.Syscall((**me.ppv).Release, 1,
			uintptr(unsafe.Pointer(*me.ppv)), 0, 0)
		if ret == 0 { // COM pointer was released
			*me.ppv = nil
		}
	}
	return uint32(ret)
}

// Creates a COM object from its CLSID. Panics if the COM object cannot be
// created.
//
// ⚠️ You must defer IUnknown.Release() on the returned COM object. If iUnkOuter
// is not null, you must defer IUnknown.Release() on it too.
//
// Example for an ordinary COM object:
//
//  taskbl := com.CoCreateInstance[shell.ITaskbarList](
//      shellco.CLSID_TaskbarList,
//      nil,
//      comco.CLSCTX_INPROC_HANDLER)
//  defer taskbl.Release()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateInstance[T IUnknown](
	rclsid co.CLSID,
	iUnkOuter *IUnknown,
	dwClsContext comco.CLSCTX) T {

	var ppvQueried **comvt.IUnknown

	var pppvOuter ***comvt.IUnknown
	if iUnkOuter != nil { // was the outer pointer requested?
		(*iUnkOuter).Release() // release if existing
		var ppvOuterBuf **comvt.IUnknown
		pppvOuter = &ppvOuterBuf // we'll request the outer pointer too
	}

	var newObj T

	ret, _, _ := syscall.Syscall6(proc.CoCreateInstance.Addr(), 5,
		uintptr(unsafe.Pointer(win.GuidFromClsid(rclsid))),
		uintptr(unsafe.Pointer(pppvOuter)),
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(win.GuidFromIid(newObj.Iid()))),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}

	*newObj.Ppv() = ppvQueried
	return newObj
}
