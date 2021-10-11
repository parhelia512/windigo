package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Handle to a process.
type HPROCESS HANDLE

// ⚠️ You must defer HPROCESS.CloseHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentprocess
func GetCurrentProcess() HPROCESS {
	ret, _, _ := syscall.Syscall(proc.GetCurrentProcess.Addr(), 0,
		0, 0, 0)
	return HPROCESS(ret)
}

// ⚠️ You must defer HPROCESS.CloseHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-openprocess
func OpenProcess(
	desiredAccess co.PROCESS,
	inheritHandle bool, processId uint32) (HPROCESS, error) {

	ret, _, err := syscall.Syscall(proc.OpenProcess.Addr(), 3,
		uintptr(desiredAccess), util.BoolToUintptr(inheritHandle), uintptr(processId))
	if ret == 0 {
		return HPROCESS(0), errco.ERROR(err)
	}
	return HPROCESS(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hProcess HPROCESS) CloseHandle() error {
	ret, _, err := syscall.Syscall(proc.CloseHandle.Addr(), 1,
		uintptr(hProcess), 0, 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getexitcodeprocess
func (hProcess HPROCESS) GetExitCodeProcess() uint32 {
	exitCode := uint32(0)
	ret, _, err := syscall.Syscall(proc.GetExitCodeProcess.Addr(), 2,
		uintptr(hProcess), uintptr(unsafe.Pointer(&exitCode)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return exitCode
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessid
func (hProcess HPROCESS) GetProcessId() uint32 {
	ret, _, err := syscall.Syscall(proc.GetProcessId.Addr(), 1,
		uintptr(hProcess), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocesstimes
func (hProcess HPROCESS) GetProcessTimes(
	creationTime, exitTime, kernelTime, userTime *FILETIME) {

	ret, _, err := syscall.Syscall6(proc.GetProcessTimes.Addr(), 5,
		uintptr(hProcess), uintptr(unsafe.Pointer(creationTime)),
		uintptr(unsafe.Pointer(exitTime)), uintptr(unsafe.Pointer(kernelTime)),
		uintptr(unsafe.Pointer(userTime)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// Pass -1 for infinite timeout.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject
func (hProcess HPROCESS) WaitForSingleObject(milliseconds uint32) co.WAIT {
	ret, _, err := syscall.Syscall(proc.WaitForSingleObject.Addr(), 2,
		uintptr(hProcess), uintptr(milliseconds), 0)
	if co.WAIT(ret) == co.WAIT_FAILED {
		panic(errco.ERROR(err))
	}
	return co.WAIT(ret)
}
