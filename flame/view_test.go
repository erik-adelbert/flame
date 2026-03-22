package flame

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkWriteHeader(b *testing.B) {
	var sb strings.Builder

	m := NewModel(24, 80)

	for b.Loop() {
		sb.Reset()
		sb.Grow(64)

		m.writeHeader(&sb)
	}
}

func BenchmarkView(b *testing.B) {
	sizes := []struct {
		w, h int
	}{
		{80, 24},
		{120, 40},
		{160, 48},
		{320, 90},
		{480, 135},
		{640, 270},
		{3829, 700},
	}

	for _, sz := range sizes {
		b.Run(
			fmt.Sprintf("%dx%d", sz.w, sz.h),

			func(b *testing.B) {
				m := NewModel(sz.h, sz.w)

				for range 10 {
					m.stepFire()
				}

				for b.Loop() {
					m.View()
				}
			},
		)
	}
}
