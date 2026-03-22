package flame

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func TestNewModel(t *testing.T) {
	w := rand.IntN(100) + 1
	h := max(1, w/5)

	m := NewModel(h, w)

	if m.h != h || m.w != w {
		t.Errorf("Expected dimensions %dx%d, got %dx%d", h, w, m.h, m.w)
	}

	if m.Size() != h*w {
		t.Errorf("Expected size 200, got %d", m.Size())
	}
}

func TestResize(t *testing.T) {
	m := NewModel(5, 5)

	m.Resize(10, 10)

	if m.h != 10 || m.w != 10 {
		t.Errorf("Expected dimensions 10x10, got %dx%d", m.h, m.w)
	}

	if m.Size() != 100 {
		t.Errorf("Expected size 100, got %d", m.Size())
	}
}

func TestIsEOL(t *testing.T) {
	m := NewModel(5, 5)

	tests := []struct {
		idx      int
		expected bool
	}{
		{4, true},  // last col
		{9, true},  // last col
		{3, false}, // not last col
		{8, false}, // not last col
	}
	for _, tt := range tests {
		if m.IsEOL(tt.idx) != tt.expected {
			t.Errorf("IsEOL(%d) = %v, want %v", tt.idx, m.IsEOL(tt.idx), tt.expected)
		}
	}
}

func TestStepFire(t *testing.T) {
	m := NewModel(3, 3)

	m.heats[4] = 100

	m.stepFire()

	if len(m.grid) != 9 {
		t.Errorf("Expected grid size 9, got %d", len(m.grid))
	}
}

func BenchmarkStepFire(b *testing.B) {
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

				for b.Loop() {
					m.stepFire()
				}
			},
		)
	}
}
