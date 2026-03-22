package flame

import (
	"image/color"
	"strings"
	"testing"
)

func TestColorRGBA(t *testing.T) {
	tests := []struct {
		color    Color
		expected color.RGBA
	}{
		{ColorCold, color.RGBA{R: 0x62, G: 0x72, B: 0xa4, A: 0xff}},
		{ColorLow, color.RGBA{R: 0xd1, G: 0x46, B: 0x46, A: 0xff}},
		{ColorWarm, color.RGBA{R: 0xfa, G: 0xe0, B: 0x8c, A: 0xff}},
		{ColorHot, color.RGBA{R: 0xef, G: 0xef, B: 0xef, A: 0xff}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := tt.color.RGBA()

			if got != tt.expected {
				t.Errorf("RGBA() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestToSymbol(t *testing.T) {
	tests := []struct {
		heat      uint8
		wantColor Color
		wantChar  byte
	}{
		{0, ColorCold, ' '},
		{4, ColorCold, '*'},
		{5, ColorLow, 'x'},
		{9, ColorLow, '$'},
		{10, ColorWarm, '$'},
		{15, ColorWarm, '$'},
		{16, ColorHot, '$'},
		{255, ColorHot, '$'},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {

			got := toSymbol(tt.heat)

			if got.Color != tt.wantColor || got.byte != tt.wantChar {
				t.Errorf(
					"toSymbol(%d) = {%v, %c}, want {%v, %c}",
					tt.heat,
					got.Color,
					got.byte,
					tt.wantColor,
					tt.wantChar,
				)
			}
		})
	}
}

func TestSymbolString(t *testing.T) {
	sym := Symbol{Color: ColorHot, byte: '$'}
	str := sym.String()

	if str == "" {
		t.Error("String() returned empty string")
	}

	if len(str) == 1 || !strings.Contains(str, "$") {
		t.Errorf("String() = %q, want a color string containing '$'", str)
	}
}

func TestColorString(t *testing.T) {
	tests := []struct {
		color    Color
		expected string
	}{
		{ColorCold, "cold"},
		{ColorLow, "low"},
		{ColorWarm, "warm"},
		{ColorHot, "hot"},
		{Color(99), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			got := tt.color.String()

			if strings.ToLower(got) != tt.expected {
				t.Errorf("String() = %q, want %q", got, tt.expected)
			}
		})
	}
}
