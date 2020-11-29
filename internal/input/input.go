package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	LeftMouseButtonPressed   bool
	MiddleMouseButtonPressed bool
	RightMouseButtonPressed  bool
	CursorX                  int
	CursorY                  int
	WheelX                   float64
	WheelY                   float64

	LeftMouseButtonJustPressed   bool
	MiddleMouseButtonJustPressed bool
	RightMouseButtonJustPressed  bool

	LastLeftMouseButtonPressed   bool
	LastMiddleMouseButtonPressed bool
	LastRightMouseButtonPressed  bool

	InputChars    []rune
	KeyPressed    = map[ebiten.Key]bool{}
	AnyKeyPressed bool
)

// Draw updates the input system. This is called by the UI.
func Draw() {
	LeftMouseButtonJustPressed = LeftMouseButtonPressed && LeftMouseButtonPressed != LastLeftMouseButtonPressed
	MiddleMouseButtonJustPressed = MiddleMouseButtonPressed && MiddleMouseButtonPressed != LastMiddleMouseButtonPressed
	RightMouseButtonJustPressed = RightMouseButtonPressed && RightMouseButtonPressed != LastRightMouseButtonPressed

	LastLeftMouseButtonPressed = LeftMouseButtonPressed
	LastMiddleMouseButtonPressed = MiddleMouseButtonPressed
	LastRightMouseButtonPressed = RightMouseButtonPressed
}

// AfterDraw updates the input system after the Ebiten Draw function has been called. This is called by the UI.
func AfterDraw() {
	InputChars = InputChars[:0]
	WheelX, WheelY = 0, 0
}
