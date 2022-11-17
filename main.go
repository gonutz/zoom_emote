package main

import (
	_ "embed"

	"bytes"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
	"time"

	"github.com/gonutz/auto"
)

var positions = map[string]position{
	"clap":         pos(-112, -113),
	"thumbs up":    pos(-71, -113),
	"cry laughing": pos(-31, -113),
	"open mouth":   pos(10, -113),
	"heart":        pos(49, -113),
	"party":        pos(86, -113),
	"yes":          pos(-104, -71),
	"no":           pos(-48, -71),
	"slow down":    pos(8, -71),
	"speed up":     pos(63, -71),
	"away":         pos(119, -71),
	"raise hand":   pos(0, -30),
}

func pos(x, y int) position {
	return position{x: x, y: y}
}

type position struct {
	x, y int
}

func main() {
	arg := strings.Join(os.Args[1:], " ")
	pos, ok := positions[arg]
	if !ok {
		return
	}

	windows, err := auto.Windows()
	if err != nil {
		return
	}
	for _, window := range windows {
		if window.Visible && !window.Minimized &&
			strings.Contains(window.Title, "Zoom Meeting") {
			if !emote(window, pos) {
				// If the icon was not found, the icon bar might not be
				// expanded, so expand it with the Alt shortcut.
				window.BringToForeground()
				auto.TypeKey(auto.KeyLeftAlt)
				time.Sleep(500 * time.Millisecond)
				emote(window, pos)
			}
			return
		}
	}
}

//go:embed reactions_icon.png
var reactionsIconPng []byte

func emote(window auto.Window, relative position) bool {
	r := window.Content
	zoom, err := auto.CaptureScreen(r.X, r.Y+r.Height-100, r.Width, 100)
	if err != nil {
		panic(err)
	}

	pattern, err := png.Decode(bytes.NewReader(reactionsIconPng))
	if err != nil {
		panic(err)
	}

	x, y, found := searchImageForPattern(zoom, pattern)
	if !found {
		return false
	}

	startX, startY, _ := auto.MousePosition()
	auto.ClickLeftMouseAt(x+20, y+20)
	time.Sleep(100 * time.Millisecond)
	auto.ClickLeftMouseAt(x+relative.x, y+relative.y)
	auto.MoveMouseTo(startX, startY)
	return true
}

func searchImageForPattern(space, pattern image.Image) (x, y int, found bool) {
	patternBounds := pattern.Bounds()
	patternW, patternH := patternBounds.Dx(), patternBounds.Dy()

	spaceBounds := space.Bounds()

	x, y, found = func() (int, int, bool) {
		for y := spaceBounds.Min.Y; y < spaceBounds.Max.Y; y++ {
			for x := spaceBounds.Min.X; x < spaceBounds.Max.X; x++ {
				found := func() bool {
					for yy := 0; yy < patternH; yy++ {
						for xx := 0; xx < patternW; xx++ {
							sx := x + xx
							sy := y + yy
							px := patternBounds.Min.X + xx
							py := patternBounds.Min.Y + yy
							if !same(space.At(sx, sy), pattern.At(px, py)) {
								return false
							}
						}
					}
					return true
				}()
				if found {
					return x, y, true
				}
			}
		}
		return 0, 0, false
	}()
	return
}

func same(c1, c2 color.Color) bool {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}
