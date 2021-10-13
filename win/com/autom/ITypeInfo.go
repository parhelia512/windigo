package autom

import (
	"github.com/rodrigocfd/windigo/win"
)

// ITypeInfo virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypeinfo
type ITypeInfoVtbl struct {
	win.IUnknownVtbl
	GetTypeAttr          uintptr
	GetTypeComp          uintptr
	GetFuncDesc          uintptr
	GetVarDesc           uintptr
	GetNames             uintptr
	GetRefTypeOfImplType uintptr
	GetImplTypeFlags     uintptr
	GetIDsOfNames        uintptr
	Invoke               uintptr
	GetDocumentation     uintptr
	GetDllEntry          uintptr
	GetRefTypeInfo       uintptr
	AddressOfMember      uintptr
	CreateInstance       uintptr
	GetMops              uintptr
	GetContainingTypeLib uintptr
	ReleaseTypeAttr      uintptr
	ReleaseFuncDesc      uintptr
	ReleaseVarDesc       uintptr
}

//------------------------------------------------------------------------------

// ITypeInfo COM interface.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypeinfo
type ITypeInfo struct {
	win.IUnknown // Base IUnknown.
}