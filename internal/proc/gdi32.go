package proc

import (
	"syscall"
)

var (
	gdi32 = syscall.NewLazyDLL("gdi32.dll")

	AbortPath              = gdi32.NewProc("AbortPath")
	AngleArc               = gdi32.NewProc("AngleArc")
	Arc                    = gdi32.NewProc("Arc")
	ArcTo                  = gdi32.NewProc("ArcTo")
	BeginPath              = gdi32.NewProc("BeginPath")
	BitBlt                 = gdi32.NewProc("BitBlt")
	CancelDC               = gdi32.NewProc("CancelDC")
	Chord                  = gdi32.NewProc("Chord")
	CloseFigure            = gdi32.NewProc("CloseFigure")
	CombineRgn             = gdi32.NewProc("CombineRgn")
	CreateBitmap           = gdi32.NewProc("CreateBitmap")
	CreateBitmapIndirect   = gdi32.NewProc("CreateBitmapIndirect")
	CreateBrushIndirect    = gdi32.NewProc("CreateBrushIndirect")
	CreateCompatibleBitmap = gdi32.NewProc("CreateCompatibleBitmap")
	CreateCompatibleDC     = gdi32.NewProc("CreateCompatibleDC")
	CreateDIBSection       = gdi32.NewProc("CreateDIBSection")
	CreateEllipticRgn      = gdi32.NewProc("CreateEllipticRgn")
	CreateFontIndirect     = gdi32.NewProc("CreateFontIndirectW")
	CreateHatchBrush       = gdi32.NewProc("CreateHatchBrush")
	CreatePatternBrush     = gdi32.NewProc("CreatePatternBrush")
	CreatePen              = gdi32.NewProc("CreatePen")
	CreatePenIndirect      = gdi32.NewProc("CreatePenIndirect")
	CreateRectRgnIndirect  = gdi32.NewProc("CreateRectRgnIndirect")
	CreateRoundRectRgn     = gdi32.NewProc("CreateRoundRectRgn")
	CreateSolidBrush       = gdi32.NewProc("CreateSolidBrush")
	DeleteDC               = gdi32.NewProc("DeleteDC")
	DeleteObject           = gdi32.NewProc("DeleteObject")
	Ellipse                = gdi32.NewProc("Ellipse")
	EndPath                = gdi32.NewProc("EndPath")
	FillPath               = gdi32.NewProc("FillPath")
	FillRect               = gdi32.NewProc("FillRect")
	FillRgn                = gdi32.NewProc("FillRgn")
	FlattenPath            = gdi32.NewProc("FlattenPath")
	FrameRect              = gdi32.NewProc("FrameRect")
	FrameRgn               = gdi32.NewProc("FrameRgn")
	GdiFlush               = gdi32.NewProc("GdiFlush")
	GetCurrentPositionEx   = gdi32.NewProc("GetCurrentPositionEx")
	GetDCBrushColor        = gdi32.NewProc("GetDCBrushColor")
	GetDCPenColor          = gdi32.NewProc("GetDCPenColor")
	GetDeviceCaps          = gdi32.NewProc("GetDeviceCaps")
	GetObject              = gdi32.NewProc("GetObject")
	GetPolyFillMode        = gdi32.NewProc("GetPolyFillMode")
	GetTextExtentPoint32   = gdi32.NewProc("GetTextExtentPoint32W")
	GetTextFace            = gdi32.NewProc("GetTextFaceW")
	GetTextMetrics         = gdi32.NewProc("GetTextMetricsW")
	GetViewportExtEx       = gdi32.NewProc("GetViewportExtEx")
	GetViewportOrgEx       = gdi32.NewProc("GetViewportOrgEx")
	GetWindowExtEx         = gdi32.NewProc("GetWindowExtEx")
	GetWindowOrgEx         = gdi32.NewProc("GetWindowOrgEx")
	IntersectClipRect      = gdi32.NewProc("IntersectClipRect")
	InvertRect             = gdi32.NewProc("InvertRect")
	InvertRgn              = gdi32.NewProc("InvertRgn")
	LineTo                 = gdi32.NewProc("LineTo")
	LPtoDP                 = gdi32.NewProc("LPtoDP")
	MaskBlt                = gdi32.NewProc("MaskBlt")
	MoveToEx               = gdi32.NewProc("MoveToEx")
	OffsetRgn              = gdi32.NewProc("OffsetRgn")
	PaintRgn               = gdi32.NewProc("PaintRgn")
	PathToRegion           = gdi32.NewProc("PathToRegion")
	Pie                    = gdi32.NewProc("Pie")
	PolyDraw               = gdi32.NewProc("PolyDraw")
	Polygon                = gdi32.NewProc("Polygon")
	Polyline               = gdi32.NewProc("Polyline")
	PolylineTo             = gdi32.NewProc("PolylineTo")
	PolyPolygon            = gdi32.NewProc("PolyPolygon")
	PolyPolyline           = gdi32.NewProc("PolyPolyline")
	PtVisible              = gdi32.NewProc("PtVisible")
	Rectangle              = gdi32.NewProc("Rectangle")
	RestoreDC              = gdi32.NewProc("RestoreDC")
	RoundRect              = gdi32.NewProc("RoundRect")
	SaveDC                 = gdi32.NewProc("SaveDC")
	SelectClipPath         = gdi32.NewProc("SelectClipPath")
	SelectClipRgn          = gdi32.NewProc("SelectClipRgn")
	SelectObject           = gdi32.NewProc("SelectObject")
	SetArcDirection        = gdi32.NewProc("SetArcDirection")
	SetBkColor             = gdi32.NewProc("SetBkColor")
	SetBkMode              = gdi32.NewProc("SetBkMode")
	SetPolyFillMode        = gdi32.NewProc("SetPolyFillMode")
	SetStretchBltMode      = gdi32.NewProc("SetStretchBltMode")
	SetTextAlign           = gdi32.NewProc("SetTextAlign")
	StretchBlt             = gdi32.NewProc("StretchBlt")
	StrokeAndFillPath      = gdi32.NewProc("StrokeAndFillPath")
	StrokePath             = gdi32.NewProc("StrokePath")
	TextOut                = gdi32.NewProc("TextOutW")
	TransparentBlt         = gdi32.NewProc("TransparentBlt")
	WidenPath              = gdi32.NewProc("WidenPath")
)
