/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

// Returns *uint16, wrapper to syscall.UTF16PtrFromString(). Panics on error.
func StrToPtr(s string) *uint16 {
	// We won't return an uintptr right away because it has no pointer semantics,
	// it's just a number, so pointed memory can be garbage-collected.
	// https://stackoverflow.com/a/51188315
	pstr, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		panic(fmt.Sprintf("StrToPtr failed \"%s\": %s",
			s, err))
	}
	return pstr
}

// Returns a null-terminated []uint16, wrapper to syscall.UTF16FromString().
// Panics on error.
func StrToSlice(s string) []uint16 {
	sli, err := syscall.UTF16FromString(s)
	if err != nil {
		panic(fmt.Sprintf("StrToSlice failed \"%s\": %s",
			s, err))
	}
	return sli
}

// Returns *uint16, or nil of empty string, wrapper to
// syscall.UTF16PtrFromString(). Panics on error.
func StrToPtrBlankIsNil(s string) *uint16 {
	if s != "" {
		return StrToPtr(s)
	}
	return nil
}

//------------------------------------------------------------------------------

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycaret
func DestroyCaret() {
	ret, _, lerr := syscall.Syscall(proc.DestroyCaret.Addr(), 0, 0, 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "DestroyCaret").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dispatchmessage
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.Syscall(proc.DispatchMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)), 0, 0)
	return ret
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endmenu
func EndMenu() {
	ret, _, lerr := syscall.Syscall(proc.EndMenu.Addr(), 0, 0, 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "EndMenu").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumwindows
func EnumWindows(
	lpEnumFunc func(hwnd HWND, lParam LPARAM) bool,
	lParam LPARAM) {

	ret, _, lerr := syscall.Syscall(proc.EnumWindows.Addr(), 2,
		syscall.NewCallback(
			func(hwnd HWND, lParam LPARAM) int32 {
				return boolToInt32(lpEnumFunc(hwnd, lParam))
			}),
		uintptr(lParam), 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "EnumWindow").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getasynckeystate
func GetAsyncKeyState(virtKeyCode co.VK) uint16 {
	ret, _, _ := syscall.Syscall(proc.GetAsyncKeyState.Addr(), 1,
		uintptr(virtKeyCode), 0, 0)
	return uint16(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcaretpos
func GetCaretPos() *RECT {
	rc := &RECT{}
	ret, _, lerr := syscall.Syscall(proc.GetCaretPos.Addr(), 1,
		uintptr(unsafe.Pointer(rc)), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "GetCaretPos").Error())
	}
	return rc
}

// https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthreadid
func GetCurrentThreadId() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetCurrentThreadId.Addr(), 0,
		0, 0, 0)
	return uint32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorpos
func GetCursorPos() *POINT {
	pt := &POINT{}
	ret, _, lerr := syscall.Syscall(proc.GetCursorPos.Addr(), 1,
		uintptr(unsafe.Pointer(pt)), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "GetCursorPos").Error())
	}
	return pt
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdpiforsystem
//
// Available in Windows 10, version 1607.
func GetDpiForSystem() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetDpiForSystem.Addr(), 0,
		0, 0, 0)
	return uint32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagew
func GetMessage(msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax uint32) int32 {
	ret, _, lerr := syscall.Syscall6(proc.GetMessage.Addr(), 4,
		uintptr(unsafe.Pointer(msg)), uintptr(hWnd),
		uintptr(msgFilterMin), uintptr(msgFilterMax),
		0, 0)
	if int(ret) == -1 {
		panic(NewWinError(co.ERROR(lerr), "GetMessage").Error())
	}
	return int32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/commdlg/nf-commdlg-getopenfilenamew
func GetOpenFileName(ofn *OPENFILENAME) bool {
	ofn.LStructSize = uint32(unsafe.Sizeof(*ofn)) // safety
	ret, _, _ := syscall.Syscall(proc.GetOpenFileName.Addr(), 1,
		uintptr(unsafe.Pointer(ofn)), 0, 0)

	if ret == 0 {
		ret, _, _ := syscall.Syscall(proc.CommDlgExtendedError.Addr(), 0,
			0, 0, 0)
		if ret != 0 {
			panic(fmt.Sprintf("GetOpenFileName failed: %d.", ret))
		} else {
			return false // user cancelled
		}
	}
	return true // user clicked OK
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getphysicalcursorpos
func GetPhysicalCursorPos() *POINT {
	pt := &POINT{}
	ret, _, lerr := syscall.Syscall(proc.GetPhysicalCursorPos.Addr(), 1,
		uintptr(unsafe.Pointer(pt)), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "GetPhysicalCursorPos").Error())
	}
	return pt
}

// https://docs.microsoft.com/en-us/windows/win32/api/commdlg/nf-commdlg-getsavefilenamew
func GetSaveFileName(ofn *OPENFILENAME) bool {
	ofn.LStructSize = uint32(unsafe.Sizeof(*ofn)) // safety
	ret, _, _ := syscall.Syscall(proc.GetSaveFileName.Addr(), 1,
		uintptr(unsafe.Pointer(ofn)), 0, 0)

	if ret == 0 {
		ret, _, _ := syscall.Syscall(proc.CommDlgExtendedError.Addr(), 0,
			0, 0, 0)
		if ret != 0 {
			panic(fmt.Sprintf("GetSaveFileName failed: %d.", ret))
		} else {
			return false // user cancelled
		}
	}
	return true // user clicked OK
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
func GetSystemMetrics(index co.SM) int32 {
	ret, _, _ := syscall.Syscall(proc.GetSystemMetrics.Addr(), 1,
		uintptr(index), 0, 0)
	return int32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrols
func InitCommonControls() {
	syscall.Syscall(proc.InitCommonControls.Addr(), 0, 0, 0, 0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isguithread
//
// Warning: passing true will force current thread to GUI, and it may deadlock.
func IsGUIThread(bConvertToGuiThread bool) bool {
	ret, _, _ := syscall.Syscall(proc.IsGUIThread.Addr(), 1,
		boolToUintptr(bConvertToGuiThread), 0, 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsversionorgreater
func IsWindowsVersionOrGreater(majorVersion, minorVersion uint32,
	servicePackMajor uint16) bool {

	ovi := OSVERSIONINFOEX{
		DwMajorVersion:    majorVersion,
		DwMinorVersion:    minorVersion,
		WServicePackMajor: servicePackMajor,
	}
	ovi.DwOsVersionInfoSize = uint32(unsafe.Sizeof(ovi))

	conditionMask := VerSetConditionMask(
		VerSetConditionMask(
			VerSetConditionMask(0, co.VER_MAJORVERSION, co.VER_COND_GREATER_EQUAL),
			co.VER_MINORVERSION, co.VER_COND_GREATER_EQUAL),
		co.VER_SERVICEPACKMAJOR, co.VER_COND_GREATER_EQUAL)

	ret, _ := VerifyVersionInfo(&ovi,
		co.VER_MAJORVERSION|co.VER_MINORVERSION|co.VER_SERVICEPACKMAJOR,
		conditionMask)
	return ret
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows10orgreater
func IsWindows10OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(hiByte(uint16(co.WIN32_WINNT_WINTHRESHOLD))),
		uint32(loByte(uint16(co.WIN32_WINNT_WINTHRESHOLD))),
		0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows7orgreater
func IsWindows7OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(hiByte(uint16(co.WIN32_WINNT_WIN7))),
		uint32(loByte(uint16(co.WIN32_WINNT_WIN7))),
		0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows8orgreater
func IsWindows8OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(hiByte(uint16(co.WIN32_WINNT_WIN8))),
		uint32(loByte(uint16(co.WIN32_WINNT_WIN8))),
		0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindows8point1orgreater
func IsWindows8Point1OrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(hiByte(uint16(co.WIN32_WINNT_WINBLUE))),
		uint32(loByte(uint16(co.WIN32_WINNT_WINBLUE))),
		0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsvistaorgreater
func IsWindowsVistaOrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(hiByte(uint16(co.WIN32_WINNT_VISTA))),
		uint32(loByte(uint16(co.WIN32_WINNT_VISTA))),
		0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/versionhelpers/nf-versionhelpers-iswindowsxporgreater
func IsWindowsXpOrGreater() bool {
	return IsWindowsVersionOrGreater(
		uint32(hiByte(uint16(co.WIN32_WINNT_WINXP))),
		uint32(loByte(uint16(co.WIN32_WINNT_WINXP))),
		0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-muldiv
func MulDiv(number, numerator, denominator int32) int32 {
	ret, _, _ := syscall.Syscall(proc.MulDiv.Addr(), 3,
		uintptr(number), uintptr(numerator), uintptr(denominator))
	return int32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int32) {
	syscall.Syscall(proc.PostQuitMessage.Addr(), 1, uintptr(exitCode), 0, 0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postthreadmessagew
func PostThreadMessage(
	idThread uint32, Msg co.WM, wParam WPARAM, lParam LPARAM) {

	ret, _, lerr := syscall.Syscall6(proc.PostThreadMessage.Addr(), 4,
		uintptr(idThread), uintptr(Msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "PostThreadMessage").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func RegisterClassEx(wcx *WNDCLASSEX) (ATOM, *WinError) {
	wcx.CbSize = uint32(unsafe.Sizeof(*wcx)) // safety
	ret, _, lerr := syscall.Syscall(proc.RegisterClassEx.Addr(), 1,
		uintptr(unsafe.Pointer(wcx)), 0, 0)
	if ret == 0 {
		return ATOM(0), NewWinError(co.ERROR(lerr), "RegisterClassEx")
	}
	return ATOM(ret), nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerwindowmessagew
func RegisterWindowMessage(lpString string) (uint32, *WinError) {
	ret, _, lerr := syscall.Syscall(proc.RegisterWindowMessage.Addr(), 1,
		uintptr(unsafe.Pointer(StrToPtr(lpString))), 0, 0)
	if ret == 0 {
		return 0, NewWinError(co.ERROR(lerr), "RegisterWindowMessage")
	}
	return uint32(ret), nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-replymessage
func ReplyMessage(lResult uintptr) bool {
	ret, _, _ := syscall.Syscall(proc.ReplyMessage.Addr(), 1,
		lResult, 0, 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiawarenesscontext
//
// Available in Windows 10, version 1703.
func SetProcessDpiAwarenessContext(value co.DPI_AWARE_CTX) {
	ret, _, lerr := syscall.Syscall(proc.SetProcessDpiAwarenessContext.Addr(), 1,
		uintptr(value), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SetProcessDpiAwarenessContext").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiaware
//
// Available in Windows Vista.
func SetProcessDPIAware() {
	ret, _, _ := syscall.Syscall(proc.SetProcessDPIAware.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic("SetProcessDPIAware failed.")
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowshookexw
func SetWindowsHookEx(idHook co.WH,
	lpfn func(code int32, wp WPARAM, lp LPARAM) uintptr,
	hmod HINSTANCE, dwThreadId uint32) HHOOK {

	ret, _, lerr := syscall.Syscall6(proc.SetWindowsHookEx.Addr(), 4,
		uintptr(idHook), syscall.NewCallback(lpfn),
		uintptr(hmod), uintptr(dwThreadId), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SetWindowsHookEx").Error())
	}
	return HHOOK(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
//
// Depends of CoInitializeEx().
func SHGetFileInfo(pszPath string, dwFileAttributes co.FILE_ATTRIBUTE,
	uFlags co.SHGFI) *SHFILEINFO {

	shfi := &SHFILEINFO{}
	ret, _, _ := syscall.Syscall6(proc.SHGetFileInfo.Addr(), 5,
		uintptr(unsafe.Pointer(StrToPtr(pszPath))),
		uintptr(dwFileAttributes), uintptr(unsafe.Pointer(shfi)),
		unsafe.Sizeof(*shfi), uintptr(uFlags), 0)

	if (uFlags&co.SHGFI_EXETYPE) == 0 || (uFlags&co.SHGFI_SYSICONINDEX) == 0 {
		if ret == 0 {
			panic(NewWinError(co.ERROR_E_UNEXPECTED, "SHGetFileInfo").Error())
		}
	}

	if (uFlags & co.SHGFI_EXETYPE) != 0 {
		if ret == 0 {
			panic("SHGetFileInfo failed.")
		}
	}

	return shfi
}

// https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-sleep
func Sleep(dwMilliseconds uint32) {
	syscall.Syscall(proc.Sleep.Addr(), 1,
		uintptr(dwMilliseconds), 0, 0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
func SystemParametersInfo(uiAction co.SPI, uiParam uint32,
	pvParam unsafe.Pointer, fWinIni uint32) {

	ret, _, lerr := syscall.Syscall6(proc.SystemParametersInfo.Addr(), 4,
		uintptr(uiAction), uintptr(uiParam), uintptr(pvParam), uintptr(fWinIni),
		0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SystemParametersInfo").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translatemessage
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(proc.TranslateMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)), 0, 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-verifyversioninfow
func VerifyVersionInfo(ovi *OSVERSIONINFOEX, typeMask co.VER,
	conditionMask uint64) (bool, co.ERROR) {

	ret, _, lerr := syscall.Syscall(proc.VerifyVersionInfo.Addr(), 3,
		uintptr(unsafe.Pointer(ovi)),
		uintptr(typeMask), uintptr(conditionMask))
	return ret != 0, co.ERROR(lerr)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-versetconditionmask
func VerSetConditionMask(conditionMask uint64, typeMask co.VER,
	condition co.VER_COND) uint64 {

	ret, _, _ := syscall.Syscall(proc.VerSetConditionMask.Addr(), 3,
		uintptr(conditionMask), uintptr(typeMask), uintptr(condition))
	return uint64(ret)
}
