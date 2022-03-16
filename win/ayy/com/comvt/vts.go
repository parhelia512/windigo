package comvt

// IUnknown virtual table, base to all COM virtual tables.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknown struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}
