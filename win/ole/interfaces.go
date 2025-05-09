//go:build windows

package ole

import (
	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/win/co"
)

// A [COM] pointer, rooted in [IUnknown].
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type ComPtr interface {
	Releasable

	// Returns the unique COM [interface ID].
	//
	// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
	IID() co.IID

	// Returns the [IUnknown] virtual table.
	//
	// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
	Ppvt() **vt.IUnknown

	// Calls [Release], then sets a new [IUnknown] virtual table.
	//
	// If you pass nil, you effectively release the object; the owning
	// ole.Releaser will simply do nothing.
	//
	// [Release]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
	// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
	Set(ppvt **vt.IUnknown)
}

// A constructible [COM] pointer, rooted in [IUnknown].
//
// Used in functions that instantiate COM pointers, like [CoCreateInstance] and
// [QueryInterface].
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
// [CoCreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
// [QueryInterface]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
type ComCtor[T any] interface {
	*T
	ComPtr
}

// A [COM] object whose lifetime can be managed by an ole.Releaser, automating the
// cleanup.
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type Releasable interface {
	// Frees the resources of the object immediately.
	//
	// You usually don't need to call this method directly, since every function
	// which returns a [COM] object will require a Releaser to manage the object's
	// lifetime.
	//
	// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
	Release()
}
