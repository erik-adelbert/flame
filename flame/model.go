package flame

import (
	"math/rand/v2"
	"slices"
	"time"
)

// Model represents the state of the fire simulation, including the heat values
// and the corresponding symbols for rendering.
type Model struct {
	fps

	h, w  int
	power int

	row0 int

	grid []Symbol

	heats []byte
	heat0 byte

	mute bool
}

// NewModel creates and returns a new Model with the given height and width.
func NewModel(h, w int) *Model {
	hmax := min(40, h)

	return &Model{
		h: h,
		w: w,

		grid:  make([]Symbol, hmax*w),
		heats: make([]uint8, (hmax*w)+(w+1)),

		heat0: BaseHeat,
		power: BasePower,
	}
}

// Resize changes the dimensions of the Model and resets the heat and grid
// slices.
func (m *Model) Resize(h, w int) {
	hmax := min(40, h)

	if h*w > m.Size() {
		m.heats = make([]uint8, (hmax*w)+(w+1))
		m.grid = make([]Symbol, hmax*w)
	}

	m.h, m.w = h, w

	m.heats = m.heats[:(hmax*w)+(w+1)]
	m.grid = m.grid[:hmax*w]
}

// Size returns the total number of cells in the grid.
func (m Model) Size() int {
	hmax := min(40, m.h)

	return hmax * m.w
}

// IsEOL checks if the given index is at the end of a row in the grid.
func (m Model) IsEOL(i int) bool {
	return (i+1)%m.w == 0
}

// ignite randomly ignites new cells at the bottom of the grid to keep the fire
// going.
func (m *Model) ignite() {
	r := float64(m.w) * float64(m.power) / 500
	n := max(1, int(r)) // spark count based on power level

	hmax := min(40, m.h)

	lo, hi := (hmax-1)*m.w, hmax*m.w // index of the first and last cell in the bottom row
	ω := m.heats[lo:hi:hi]           // bottom row

	clear(ω) // clear the bottom row before igniting new cells

	for range n {
		x := rand.IntN(m.w)
		ω[x] = m.heat0
	}
}

// stepFire updates the heat values of the grid based on the current state and
// the fire spreading algorithm.
func (m *Model) stepFire() {

	m.ignite()

	for i := range m.Size() {
		// average the heat of the current cell and its three neighbors
		m.heats[i] += m.heats[i+1] + m.heats[i+m.w] + m.heats[i+m.w+1]
		m.heats[i] /= 4

		m.grid[i] = toSymbol(m.heats[i])
	}

	// track the minimum row with heat for rendering optimization
	i := slices.IndexFunc(m.heats, func(h uint8) bool { return h != 0 })
	m.row0 = i / m.w
}

type fps struct {
	frames int
	last   time.Time
	cur    float64
}

func (f *fps) FPS() float64 {
	return f.cur
}

func (f *fps) sample(now time.Time) {
	f.frames++

	if f.last.IsZero() {
		f.last = now
	}

	dt := now.Sub(f.last)

	if dt >= time.Second {
		f.cur = float64(f.frames) / dt.Seconds()
		f.frames = 0
		f.last = now
	}
}
