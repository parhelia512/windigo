/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

type PAGE uint32 // CreateFileMapping flProtect

const (
	PAGE_NONE                   PAGE = 0
	PAGE_NOACCESS               PAGE = 0x01
	PAGE_READONLY               PAGE = 0x02
	PAGE_READWRITE              PAGE = 0x04
	PAGE_WRITECOPY              PAGE = 0x08
	PAGE_EXECUTE                PAGE = 0x10
	PAGE_EXECUTE_READ           PAGE = 0x20
	PAGE_EXECUTE_READWRITE      PAGE = 0x40
	PAGE_EXECUTE_WRITECOPY      PAGE = 0x80
	PAGE_GUARD                  PAGE = 0x100
	PAGE_NOCACHE                PAGE = 0x200
	PAGE_WRITECOMBINE           PAGE = 0x400
	PAGE_ENCLAVE_THREAD_CONTROL PAGE = 0x80000000
	PAGE_REVERT_TO_FILE_MAP     PAGE = 0x80000000
	PAGE_TARGETS_NO_UPDATE      PAGE = 0x40000000
	PAGE_TARGETS_INVALID        PAGE = 0x40000000
	PAGE_ENCLAVE_UNVALIDATED    PAGE = 0x20000000
	PAGE_ENCLAVE_DECOMMIT       PAGE = 0x10000000
)

type PRF uint32 // WM_PRINT

const (
	PRF_CHECKVISIBLE PRF = 0x00000001
	PRF_NONCLIENT    PRF = 0x00000002
	PRF_CLIENT       PRF = 0x00000004
	PRF_ERASEBKGND   PRF = 0x00000008
	PRF_CHILDREN     PRF = 0x00000010
	PRF_OWNED        PRF = 0x00000020
)

type PT uint8 // PolyDraw

const (
	PT_CLOSEFIGURE PT = 0x01
	PT_LINETO      PT = 0x02
	PT_BEZIERTO    PT = 0x04
	PT_MOVETO      PT = 0x06
)