package main

import (
	"strings"
	"time"

	"github.com/gonutz/input"
	"github.com/gonutz/w32/v2"
)

func main() {
	var zoom w32.HWND
	w32.EnumWindows(func(window w32.HWND) bool {
		title := w32.GetWindowText(window)
		if strings.Contains(title, "Zoom Meeting") {
			zoom = window
			return false
		}
		return true
	})
	if zoom != 0 {
		r := w32.GetWindowRect(zoom)
		x := int(r.Left)
		y := int(r.Bottom - 1)
		input.MoveMouseTo(x+777, y-40)
		input.LeftClick()
		time.Sleep(100 * time.Millisecond)
		input.MoveMouseTo(x+697, y-165)
		input.LeftClick()
	}
}
