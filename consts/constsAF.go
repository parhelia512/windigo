package consts

type COLOR uint8

const (
	COLOR_3DDKSHADOW              COLOR = 21
	COLOR_3DFACE                  COLOR = 15
	COLOR_3DHIGHLIGHT             COLOR = 20
	COLOR_3DHILIGHT               COLOR = 20
	COLOR_3DLIGHT                 COLOR = 22
	COLOR_3DSHADOW                COLOR = 16
	COLOR_ACTIVEBORDER            COLOR = 10
	COLOR_ACTIVECAPTION           COLOR = 2
	COLOR_APPWORKSPACE            COLOR = 12
	COLOR_BACKGROUND              COLOR = 1
	COLOR_BTNFACE                 COLOR = 15
	COLOR_BTNHIGHLIGHT            COLOR = 20
	COLOR_BTNHILIGHT              COLOR = 20
	COLOR_BTNSHADOW               COLOR = 16
	COLOR_BTNTEXT                 COLOR = 18
	COLOR_CAPTIONTEXT             COLOR = 9
	COLOR_DESKTOP                 COLOR = 1
	COLOR_GRADIENTACTIVECAPTION   COLOR = 27
	COLOR_GRADIENTINACTIVECAPTION COLOR = 28
	COLOR_GRAYTEXT                COLOR = 17
	COLOR_HIGHLIGHT               COLOR = 13
	COLOR_HIGHLIGHTTEXT           COLOR = 14
	COLOR_HOTLIGHT                COLOR = 26
	COLOR_INACTIVEBORDER          COLOR = 11
	COLOR_INACTIVECAPTION         COLOR = 3
	COLOR_INACTIVECAPTIONTEXT     COLOR = 19
	COLOR_INFOBK                  COLOR = 24
	COLOR_INFOTEXT                COLOR = 23
	COLOR_MENU                    COLOR = 4
	COLOR_MENUHILIGHT             COLOR = 29
	COLOR_MENUBAR                 COLOR = 30
	COLOR_MENUTEXT                COLOR = 7
	COLOR_SCROLLBAR               COLOR = 0
	COLOR_WINDOW                  COLOR = 5
	COLOR_WINDOWFRAME             COLOR = 6
	COLOR_WINDOWTEXT              COLOR = 8
)

type FW uint32

const (
	FW_DONTCARE   FW = 0
	FW_THIN       FW = 100
	FW_EXTRALIGHT FW = 200
	FW_ULTRALIGHT FW = FW_EXTRALIGHT
	FW_LIGHT      FW = 300
	FW_NORMAL     FW = 400
	FW_REGULAR    FW = 400
	FW_MEDIUM     FW = 500
	FW_SEMIBOLD   FW = 600
	FW_DEMIBOLD   FW = FW_SEMIBOLD
	FW_BOLD       FW = 700
	FW_EXTRABOLD  FW = 800
	FW_ULTRABOLD  FW = FW_EXTRABOLD
	FW_HEAVY      FW = 900
	FW_BLACK      FW = FW_HEAVY
)
