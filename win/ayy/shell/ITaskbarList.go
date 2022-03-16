package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/ayy/com"
	"github.com/rodrigocfd/windigo/win/ayy/shell/shellco"
	"github.com/rodrigocfd/windigo/win/ayy/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Example:
//
//  taskbl := com.CoCreateInstance[shell.ITaskbarList](
//      shellco.CLSID_TaskbarList,
//      nil,
//      comco.CLSCTX_INPROC_HANDLER)
//  defer taskbl.Release()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
type ITaskbarList interface {
	com.IUnknown

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-activatetab
	ActivateTab(hWnd win.HWND)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-addtab
	AddTab(hWnd win.HWND)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-deletetab
	DeleteTab(hWnd win.HWND)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-hrinit
	HrInit()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-setactivealt
	SetActiveAlt(hWnd win.HWND)
}

type _ITaskbarList struct{ com.IUnknown }

func (me *_ITaskbarList) Iid() co.IID {
	return shellco.IID_ITaskbarList
}

func (me *_ITaskbarList) ActivateTab(hWnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList)(unsafe.Pointer(**me.Ppv())).ActivateTab, 2,
		uintptr(unsafe.Pointer(*me.Ppv())),
		uintptr(hWnd), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList) AddTab(hWnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList)(unsafe.Pointer(**me.Ppv())).AddTab, 2,
		uintptr(unsafe.Pointer(*me.Ppv())),
		uintptr(hWnd), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList) DeleteTab(hWnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList)(unsafe.Pointer(**me.Ppv())).DeleteTab, 2,
		uintptr(unsafe.Pointer(*me.Ppv())),
		uintptr(hWnd), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList) HrInit() {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList)(unsafe.Pointer(**me.Ppv())).HrInit, 1,
		uintptr(unsafe.Pointer(*me.Ppv())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList) SetActiveAlt(hWnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList)(unsafe.Pointer(**me.Ppv())).SetActiveAlt, 2,
		uintptr(unsafe.Pointer(*me.Ppv())),
		uintptr(hWnd), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
