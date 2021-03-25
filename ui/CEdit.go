package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native edit control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/about-edit-controls
type Edit interface {
	AnyControl

	// Exposes all the Edit notifications the can be handled.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-edit-control-reference-notifications
	On() *_EditEvents

	LimitText(maxChars int)                 // Limits the length of the text.
	ReplaceSelection(text string)           // Replaces the current text selection with the given text.
	SelectedRange() (int, int)              // Retrieves the index of first and last selected chars.
	SetSelectedRange(idxFirst, idxLast int) // Sets the currently selected chars.
	SetText(text string)                    // Sets the text.
	Text() string                           // Retrieves the text.
}

//------------------------------------------------------------------------------

type _Edit struct {
	_NativeControlBase
	events _EditEvents
}

// Creates a new Edit specifying all options, which will be passed to the
// underlying CreateWindowEx().
func NewEditOpts(parent AnyParent, opts EditOpts) Edit {
	opts.fillBlankValuesWithDefault()

	me := _Edit{}
	me._NativeControlBase.new(parent, opts.CtrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(_CreateOrInitDialog(parent), func(_ wm.Any) {
		_MultiplyDpi(&opts.Position, &opts.Size)

		me._NativeControlBase.createWindow(opts.ExStyles,
			"EDIT", opts.Text, opts.Styles|co.WS(opts.EditStyles),
			opts.Position, opts.Size, win.HMENU(opts.CtrlId))

		me.Hwnd().SendMessage(co.WM_SETFONT, win.WPARAM(_globalUiFont), 1)
	})

	return &me
}

// Creates a new Edit from a dialog resource.
func NewEditDlg(parent AnyParent, ctrlId int) Edit {
	me := _Edit{}
	me._NativeControlBase.new(parent, ctrlId)
	me.events.new(&me._NativeControlBase)

	parent.internalOn().addMsgZero(co.WM_INITDIALOG, func(_ wm.Any) {
		me._NativeControlBase.assignDlgItem()
	})

	return &me
}

func (me *_Edit) On() *_EditEvents {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the Edit is created.")
	}
	return &me.events
}

func (me *_Edit) LimitText(maxChars int) {
	me.Hwnd().SendMessage(co.EM_LIMITTEXT, win.WPARAM(maxChars), 0)
}

func (me *_Edit) ReplaceSelection(replacementText string) {
	me.Hwnd().SendMessage(co.EM_REPLACESEL,
		1, win.LPARAM(unsafe.Pointer(win.Str.ToUint16Ptr(replacementText))))
}

func (me *_Edit) SelectedRange() (int, int) {
	idxFirst, idxLast := uint32(0), uint32(0)
	me.Hwnd().SendMessage(co.EM_GETSEL,
		win.WPARAM(unsafe.Pointer(&idxFirst)),
		win.LPARAM(unsafe.Pointer(&idxLast)))
	return int(idxFirst), int(idxLast)
}

func (me *_Edit) SetSelectedRange(idxFirst, idxLast int) {
	me.Hwnd().SendMessage(co.EM_SETSEL,
		win.WPARAM(idxFirst), win.LPARAM(idxLast))
}

func (me *_Edit) SetText(text string) {
	me.Hwnd().SetWindowText(text)
}

func (me *_Edit) Text() string {
	return me.Hwnd().GetWindowText()
}

//------------------------------------------------------------------------------

// Options for NewEditOpts().
type EditOpts struct {
	// Control ID.
	// Defaults to an auto-generated ID.
	CtrlId int

	// Text to appear in the edit, passed to CreateWindowEx().
	// Defaults to empty string.
	Text string
	// Position within parent's client area in pixels.
	// Defaults to 0x0. Will be adjusted to the current system DPI.
	Position win.POINT
	// Control size in pixels.
	// Defaults to 100x21. Will be adjusted to the current system DPI.
	Size win.SIZE
	// Edit control styles, passed to CreateWindowEx().
	// Defaults to ES_AUTOHSCROLL | ES_NOHIDESEL.
	EditStyles co.ES
	// Window styles, passed to CreateWindowEx().
	// Defaults to WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE.
	Styles co.WS
	// Extended window styles, passed to CreateWindowEx().
	// Defaults to WS_EX_CLIENTEDGE.
	ExStyles co.WS_EX
}

func (opts *EditOpts) fillBlankValuesWithDefault() {
	if opts.CtrlId == 0 {
		opts.CtrlId = _NextCtrlId()
	}

	if opts.Size.Cx == 0 {
		opts.Size.Cx = 100
	}
	if opts.Size.Cy == 0 {
		opts.Size.Cy = 21
	}

	if opts.EditStyles == 0 {
		opts.EditStyles = co.ES_AUTOHSCROLL | co.ES_NOHIDESEL
	}
	if opts.Styles == 0 {
		opts.Styles = co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE
	}
	if opts.ExStyles == 0 {
		opts.ExStyles = co.WS_EX_CLIENTEDGE
	}
}

//------------------------------------------------------------------------------

// Edit control notifications.
type _EditEvents struct {
	ctrlId int
	events *_EventsNfy
}

func (me *_EditEvents) new(ctrl *_NativeControlBase) {
	me.ctrlId = ctrl.CtrlId()
	me.events = ctrl.parent.On()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-align-ltr-ec
func (me *_EditEvents) EnAlignLtrEc(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_ALIGN_LTR_EC, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-align-rtl-ec
func (me *_EditEvents) EnAlignRtlEc(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_ALIGN_RTL_EC, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-change
func (me *_EditEvents) EnChange(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_CHANGE, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-errspace
func (me *_EditEvents) EnErrSpace(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_ERRSPACE, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-hscroll
func (me *_EditEvents) EnHScroll(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_HSCROLL, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-killfocus
func (me *_EditEvents) EnKillFocus(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_KILLFOCUS, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-maxtext
func (me *_EditEvents) EnMaxText(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_MAXTEXT, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-setfocus
func (me *_EditEvents) EnSetFocus(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_SETFOCUS, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-update
func (me *_EditEvents) EnUpdate(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_UPDATE, func(_ wm.Command) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/en-vscroll
func (me *_EditEvents) EnVScroll(userFunc func()) {
	me.events.addCmdZero(me.ctrlId, co.EN_VSCROLL, func(_ wm.Command) {
		userFunc()
	})
}