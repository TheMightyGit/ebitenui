// +build android

package input

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	touchScrolling                           bool
	lastTouchScrollPosX, lastTouchScrollPosY int
)

// Update updates the input system. This is called by the UI.
func Update() {
	LeftMouseButtonPressed = false
	touchScrolling = false

	touches := ebiten.TouchIDs()
	if len(touches) == 1 {
		// first touchs emulates mouse and left button
		CursorX, CursorY = ebiten.TouchPosition(touches[0])
		LeftMouseButtonPressed = true
	} else if len(touches) == 2 {
		// two fingers for scroll
		if !touchScrolling {
			// start of scroll touch
			lastTouchScrollPosX, lastTouchScrollPosY = ebiten.TouchPosition(touches[1])
			log.Println("start scrolling", WheelX, WheelY)
		} else {
			// continue scrolling
			touchScrollPosX, touchScrollPosY := ebiten.TouchPosition(touches[1])
			WheelX += float64(touchScrollPosX - lastTouchScrollPosX)
			WheelY += float64(touchScrollPosY - lastTouchScrollPosY)
			log.Println("continue scrolling", WheelX, WheelY)
			lastTouchScrollPosX, lastTouchScrollPosY = touchScrollPosX, touchScrollPosY
		}
		touchScrolling = true
	}

	AnyKeyPressed = false
}
