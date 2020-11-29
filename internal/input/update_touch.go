// +build android

package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Update updates the input system. This is called by the UI.
func Update() {
	LeftMouseButtonPressed = false

	// all touches emulate mouse and left button
	touches := ebiten.TouchIDs()
	for _, id := range touches {
		CursorX, CursorY = ebiten.TouchPosition(id)
		LeftMouseButtonPressed = true
	}

	wx, wy := ebiten.Wheel()
	WheelX += wx
	WheelY += wy

	AnyKeyPressed = false
}
