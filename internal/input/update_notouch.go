// +build !android

package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Update updates the input system. This is called by the UI.
func Update() {
	LeftMouseButtonPressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	MiddleMouseButtonPressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle)
	RightMouseButtonPressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	CursorX, CursorY = ebiten.CursorPosition()

	wx, wy := ebiten.Wheel()
	WheelX += wx
	WheelY += wy

	InputChars = append(InputChars, ebiten.InputChars()...)

	AnyKeyPressed = false
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		p := ebiten.IsKeyPressed(k)
		KeyPressed[k] = p

		if p {
			AnyKeyPressed = true
		}
	}
}
