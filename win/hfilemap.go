package win

import (
	"fmt"
	"syscall"
	"wingows/co"
	"wingows/win/proc"
)

// These types don't exist in Win32, we're wrapping HANDLE just to have a proper
// scope on the functions.
type (
	HFILEMAP     HANDLE  // Returned by HFILE.CreateFileMapping().
	HFILEMAPADDR uintptr // Returned by HFILEMAP.MapViewOfFile(), just a BYTE pointer to memory address.
)

func (hmap HFILEMAP) CloseHandle() {
	ret, _, lerr := syscall.Syscall(proc.CloseHandle.Addr(), 1,
		uintptr(hmap), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("CloseHandle failed: %d %s\n",
			lerr, lerr.Error()))
	}
}

func (hmap HFILEMAP) MapViewOfFile(desiredAccess co.FILE_MAP,
	offset uint32, numBytesToMap uintptr) HFILEMAPADDR {

	ret, _, lerr := syscall.Syscall6(proc.MapViewOfFile.Addr(), 5,
		uintptr(hmap), uintptr(desiredAccess), 0, uintptr(offset),
		numBytesToMap, 0)
	if ret == 0 {
		panic(fmt.Sprintf("MapViewOfFile failed: %d %s\n",
			lerr, lerr.Error()))
	}
	return HFILEMAPADDR(ret)
}

func (mappedPtr HFILEMAPADDR) UnmapViewOfFile() {
	ret, _, lerr := syscall.Syscall(proc.UnmapViewOfFile.Addr(), 1,
		uintptr(mappedPtr), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("UnmapViewOfFile failed: %d %s\n",
			lerr, lerr.Error()))
	}
}