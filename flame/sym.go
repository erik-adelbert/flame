package flame

import (
	"image/color"
	"sync"

	"charm.land/lipgloss/v2"
)

// LUT is a package-level color lookup table. It is intentionally left editable.
var LUT = []color.RGBA{
	ColorCold: {R: 0x62, G: 0x72, B: 0xa4, A: 0xff}, // Dark Blue
	ColorLow:  {R: 0xd1, G: 0x46, B: 0x46, A: 0xff}, // Red
	ColorWarm: {R: 0xfa, G: 0xe0, B: 0x8c, A: 0xff}, // Yellow
	ColorHot:  {R: 0xef, G: 0xef, B: 0xef, A: 0xff}, // White
}

// Symbol represents a pair of color and character.
type Symbol struct {
	byte
	Color
}

var stylePool = sync.Pool{
	New: func() any {
		return new(lipgloss.NewStyle())
	},
}

func getStyle() *lipgloss.Style {
	style := stylePool.Get().(*lipgloss.Style)
	return style
}

func putStyle(s *lipgloss.Style) {
	stylePool.Put(s)
}

// String returns the formatted string representation of the Symbol.
func (s Symbol) String() string {
	str, ok := stringCache[s]

	if !ok {
		style := getStyle()
		str = style.Foreground(s.RGBA()).Render(string(s.byte))
		putStyle(style)

		stringCache[s] = str
	}

	return str
}

// Color represents the color of a Symbol. It is an index into the LUT.
type Color int

const (
	ColorCold Color = iota
	ColorLow
	ColorWarm
	ColorHot
)

// RGBA returns the RGBA values of the Color.
func (c Color) RGBA() color.RGBA {
	return LUT[c]
}

// String returns the name of the Color as a string.
func (c Color) String() string {
	if c > ColorHot {
		return "Unknown"
	}

	return []string{
		ColorCold: "Cold",
		ColorLow:  "Low",
		ColorWarm: "Warm",
		ColorHot:  "Hot",
	}[c]
}

// toSymbol converts a heat value to a Symbol with the appropriate character and
// color.
func toSymbol(h uint8) Symbol {
	// It is unclear who used this charset first but it is commonly used for
	// heatmaps and flames.
	char := " .:^*xsS#$"[min(9, h)]

	color := ColorCold
	switch {
	case h > 15:
		color = ColorHot
	case h > 9:
		color = ColorWarm
	case h > 4:
		color = ColorLow
	}

	return Symbol{byte: char, Color: color}
}

var stringCache = make(map[Symbol]string, 16)
