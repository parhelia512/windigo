/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type (
	_IDispatchImpl struct{ _IUnknownImpl }

	// https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch
	//
	// IDispatch > IUnknown.
	IDispatch struct{ _IDispatchImpl }

	_IDispatchVtbl struct {
		_IUnknownVtbl
		GetTypeInfoCount uintptr
		GetTypeInfo      uintptr
		GetIDsOfNames    uintptr
		Invoke           uintptr
	}
)
