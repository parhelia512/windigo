/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"wingows/api"
	c "wingows/consts"
)

// Native check box control.
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type CheckBox struct {
	controlNativeBase
}

// Optional; returns a CheckBox with a specific control ID.
func MakeCheckBox(ctrlId c.ID) CheckBox {
	return CheckBox{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *CheckBox) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles c.WS_EX, styles c.WS, btnStyles c.BS) *CheckBox {

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles, "BUTTON", text, // check box is, in fact, a button
		styles|c.WS(btnStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx(). Position will be adjusted to the current system DPI.
// The size will be calculated to fit the text exactly.
func (me *CheckBox) CreateSimple(parent Window, x, y int32,
	text string) *CheckBox {

	x, y, _, _ = multiplyByDpi(x, y, 0, 0)
	cx, cy := calcIdealSize(parent.Hwnd(), text, true)

	cx += uint32(api.GetSystemMetrics(c.SM_CXMENUCHECK)) +
		uint32(api.GetSystemMetrics(c.SM_CXEDGE)) // https://stackoverflow.com/a/1165052/6923555

	me.controlNativeBase.create(c.WS_EX(0), "BUTTON", text,
		c.WS_CHILD|c.WS_GROUP|c.WS_VISIBLE|c.WS(c.BS_AUTOCHECKBOX),
		x, y, cx, cy, parent)
	globalUiFont.SetOnControl(me)
	return me
}

func (me *CheckBox) IsChecked() bool {
	return me.State() == c.BST_CHECKED
}

func (me *CheckBox) State() c.BST {
	return me.Hwnd().GetParent().IsDlgButtonChecked(me.CtrlId())
}