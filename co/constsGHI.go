/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

// GetAncestor() gaFlags.
type GA uint32

const (
	GA_PARENT    GA = 1
	GA_ROOT      GA = 2
	GA_ROOTOWNER GA = 3
)

// GetDeviceCaps() index; originally has no prefix.
type GDC int32

const (
	GDC_DRIVERVERSION   GDC = 0
	GDC_TECHNOLOGY      GDC = 2
	GDC_HORZSIZE        GDC = 4
	GDC_VERTSIZE        GDC = 6
	GDC_HORZRES         GDC = 8
	GDC_VERTRES         GDC = 10
	GDC_BITSPIXEL       GDC = 12
	GDC_PLANES          GDC = 14
	GDC_NUMBRUSHES      GDC = 16
	GDC_NUMPENS         GDC = 18
	GDC_NUMMARKERS      GDC = 20
	GDC_NUMFONTS        GDC = 22
	GDC_NUMCOLORS       GDC = 24
	GDC_PDEVICESIZE     GDC = 26
	GDC_CURVECAPS       GDC = 28
	GDC_LINECAPS        GDC = 30
	GDC_POLYGONALCAPS   GDC = 32
	GDC_TEXTCAPS        GDC = 34
	GDC_CLIPCAPS        GDC = 36
	GDC_RASTERCAPS      GDC = 38
	GDC_ASPECTX         GDC = 40
	GDC_ASPECTY         GDC = 42
	GDC_ASPECTXY        GDC = 44
	GDC_LOGPIXELSX      GDC = 88
	GDC_LOGPIXELSY      GDC = 90
	GDC_SIZEPALETTE     GDC = 104
	GDC_NUMRESERVED     GDC = 106
	GDC_COLORRES        GDC = 108
	GDC_PHYSICALWIDTH   GDC = 110
	GDC_PHYSICALHEIGHT  GDC = 111
	GDC_PHYSICALOFFSETX GDC = 112
	GDC_PHYSICALOFFSETY GDC = 113
	GDC_SCALINGFACTORX  GDC = 114
	GDC_SCALINGFACTORY  GDC = 115
	GDC_VREFRESH        GDC = 116
	GDC_DESKTOPVERTRES  GDC = 117
	GDC_DESKTOPHORZRES  GDC = 118
	GDC_BLTALIGNMENT    GDC = 119
	GDC_SHADEBLENDCAPS  GDC = 120
	GDC_COLORMGMTCAPS   GDC = 121
)

// CreateFile() dwDesiredAccess.
type GENERIC uint32

const (
	GENERIC_READ    GENERIC = 0x80000000
	GENERIC_WRITE   GENERIC = 0x40000000
	GENERIC_EXECUTE GENERIC = 0x20000000
	GENERIC_ALL     GENERIC = 0x10000000
)

// GetMenuDefaultItem() gmdiFlags.
type GMDI uint32

const (
	GMDI_USEDISABLED  GMDI = 0x0001
	GMDI_GOINTOPOPUPS GMDI = 0x0002
)

// GetWindow() uCmd.
type GW uint32

const (
	GW_HWNDFIRST    GW = 0
	GW_HWNDLAST     GW = 1
	GW_HWNDNEXT     GW = 2
	GW_HWNDPREV     GW = 3
	GW_OWNER        GW = 4
	GW_CHILD        GW = 5
	GW_ENABLEDPOPUP GW = 6
	GW_MAX          GW = 6
)

// Get/SetWindowLongPtr() nIndex.
type GWLP int32

const (
	GWLP_STYLE      GWLP = -16
	GWLP_EXSTYLE    GWLP = -20
	GWLP_WNDPROC    GWLP = -4
	GWLP_HINSTANCE  GWLP = -6
	GWLP_HWNDPARENT GWLP = -8
	GWLP_USERDATA   GWLP = -21
	GWLP_ID         GWLP = -12
)

// SetWindowsHookEx() callback hook codes.
type HC int32

const (
	HC_ACTION      HC = 0
	HC_GETNEXT     HC = 1
	HC_SKIP        HC = 2
	HC_NOREMOVE    HC = 3
	HC_NOREM       HC = HC_NOREMOVE
	HC_SYSMODALON  HC = 4
	HC_SYSMODALOFF HC = 5
)

// SetWindowsHookEx() callback CBT hook codes.
type HCBT int32

const (
	HCBT_MOVESIZE     HCBT = 0
	HCBT_MINMAX       HCBT = 1
	HCBT_QS           HCBT = 2
	HCBT_CREATEWND    HCBT = 3
	HCBT_DESTROYWND   HCBT = 4
	HCBT_ACTIVATE     HCBT = 5
	HCBT_CLICKSKIPPED HCBT = 6
	HCBT_KEYSKIPPED   HCBT = 7
	HCBT_SYSCOMMAND   HCBT = 8
	HCBT_SETFOCUS     HCBT = 9
)

// List view header message.
type HDM WM

const (
	_HDM_FIRST HDM = 0x1200

	HDM_GETITEMCOUNT HDM = _HDM_FIRST + 0
	HDM_INSERTITEM   HDM = _HDM_FIRST + 10
	HDM_DELETEITEM   HDM = _HDM_FIRST + 2
	HDM_GETITEM      HDM = _HDM_FIRST + 11
	HDM_SETITEM      HDM = _HDM_FIRST + 12
	HDM_LAYOUT       HDM = _HDM_FIRST + 5
)

// HELPINFO iContextType.
type HELPINFO int32

const (
	HELPINFO_WINDOW   HELPINFO = 0x0001
	HELPINFO_MENUITEM HELPINFO = 0x0002
)

// RegOpenKeyEx() hKey.
type HKEY uintptr

const (
	HKEY_CLASSES_ROOT                HKEY = 0x80000000
	HKEY_CURRENT_USER                HKEY = 0x80000001
	HKEY_LOCAL_MACHINE               HKEY = 0x80000002
	HKEY_USERS                       HKEY = 0x80000003
	HKEY_PERFORMANCE_DATA            HKEY = 0x80000004
	HKEY_PERFORMANCE_TEXT            HKEY = 0x80000050
	HKEY_PERFORMANCE_NLSTEXT         HKEY = 0x80000060
	HKEY_CURRENT_CONFIG              HKEY = 0x80000005
	HKEY_DYN_DATA                    HKEY = 0x80000006
	HKEY_CURRENT_USER_LOCAL_SETTINGS HKEY = 0x80000007
)

// TVINSERTSTRUCT hInsertAfter.
type HTREEITEM uintptr

const (
	HTREEITEM_ROOT  HTREEITEM = 0x10000
	HTREEITEM_FIRST HTREEITEM = 0x0FFFF
	HTREEITEM_LAST  HTREEITEM = 0x0FFFE
	HTREEITEM_SORT  HTREEITEM = 0x0FFFD
)

// MessageBox() return value.
type MBID int32

const (
	MBID_OK       MBID = 1
	MBID_CANCEL   MBID = 2
	MBID_ABORT    MBID = 3
	MBID_RETRY    MBID = 4
	MBID_IGNORE   MBID = 5
	MBID_YES      MBID = 6
	MBID_NO       MBID = 7
	MBID_TRYAGAIN MBID = 10
	MBID_CONTINUE MBID = 11
)

// LoadCursor() lpCursorName.
type IDC uintptr

const (
	IDC_ARROW       IDC = 32512
	IDC_IBEAM       IDC = 32513
	IDC_WAIT        IDC = 32514
	IDC_CROSS       IDC = 32515
	IDC_UPARROW     IDC = 32516
	IDC_SIZENWSE    IDC = 32642
	IDC_SIZENESW    IDC = 32643
	IDC_SIZEWE      IDC = 32644
	IDC_SIZENS      IDC = 32645
	IDC_SIZEALL     IDC = 32646
	IDC_NO          IDC = 32648
	IDC_HAND        IDC = 32649
	IDC_APPSTARTING IDC = 32650
	IDC_HELP        IDC = 32651
	IDC_PIN         IDC = 32671
	IDC_PERSON      IDC = 32672
)

// LoadIcon() lpIconName.
type IDI uintptr

const (
	IDI_APPLICATION IDI = 32512
	IDI_HAND        IDI = 32513
	IDI_QUESTION    IDI = 32514
	IDI_EXCLAMATION IDI = 32515
	IDI_ASTERISK    IDI = 32516
	IDI_WINLOGO     IDI = 32517
	IDI_SHIELD      IDI = 32518
	IDI_WARNING     IDI = IDI_EXCLAMATION
	IDI_ERROR       IDI = IDI_HAND
	IDI_INFORMATION IDI = IDI_ASTERISK
)

// WM_HOTKEY identifier.
type IDHOT int32

const (
	IDHOT_SNAPWINDOW  IDHOT = -1
	IDHOT_SNAPDESKTOP IDHOT = -2
)

// ImageList_Create() flags.
type ILC uint32

const (
	ILC_MASK             ILC = 0x00000001
	ILC_COLOR            ILC = 0x00000000
	ILC_COLORDDB         ILC = 0x000000FE
	ILC_COLOR4           ILC = 0x00000004
	ILC_COLOR8           ILC = 0x00000008
	ILC_COLOR16          ILC = 0x00000010
	ILC_COLOR24          ILC = 0x00000018
	ILC_COLOR32          ILC = 0x00000020
	ILC_PALETTE          ILC = 0x00000800
	ILC_MIRROR           ILC = 0x00002000
	ILC_PERITEMMIRROR    ILC = 0x00008000
	ILC_ORIGINALSIZE     ILC = 0x00010000 // Vista
	ILC_HIGHQUALITYSCALE ILC = 0x00020000 // Vista
)

// ImageList_Draw() flags.
type ILD uint32

const (
	ILD_NORMAL        ILD = 0x00000000
	ILD_TRANSPARENT   ILD = 0x00000001
	ILD_MASK          ILD = 0x00000010
	ILD_IMAGE         ILD = 0x00000020
	ILD_ROP           ILD = 0x00000040
	ILD_BLEND25       ILD = 0x00000002
	ILD_BLEND50       ILD = 0x00000004
	ILD_OVERLAYMASK   ILD = 0x00000F00
	ILD_PRESERVEALPHA ILD = 0x00001000
	ILD_SCALE         ILD = 0x00002000
	ILD_DPISCALE      ILD = 0x00004000
	ILD_ASYNC         ILD = 0x00008000 // Vista
	ILD_SELECTED      ILD = ILD_BLEND50
	ILD_FOCUS         ILD = ILD_BLEND25
	ILD_BLEND         ILD = ILD_BLEND50
)

// Image list state.
type ILS uint32

const (
	ILS_NORMAL   ILS = 0x00000000
	ILS_GLOW     ILS = 0x00000001
	ILS_SHADOW   ILS = 0x00000002
	ILS_SATURATE ILS = 0x00000004
	ILS_ALPHA    ILS = 0x00000008
)
