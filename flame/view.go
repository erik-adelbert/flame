package flame

import (
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
)

// View renders the current state of the model as a string that can be displayed
// in the terminal.
func (m *Model) View() tea.View {
	var (
		i  int
		sb strings.Builder
	)

	// preallocate the string builder
	size := max(32768, 2*m.Size())
	sb.Grow(size)

	if !m.mute {
		m.writeHeader(&sb)
	}

	for i < m.Size() {
		// skip empty rows at the top of the grid
		if i < m.row0*m.w {
			sb.WriteByte('\n')

			i += m.w

			continue
		}

		// render the symbol for the current cell
		sym := m.grid[i]

		sb.WriteString(sym.String())

		// track the end of the row
		if m.IsEOL(i) {
			sb.WriteByte('\n')
		}

		i++
	}

	v := tea.NewView(sb.String())
	v.AltScreen = true

	return v
}

// writeHeader writes the header information to the provided string builder.
func (m *Model) writeHeader(sb *strings.Builder) {
	const (
		header0 = "fps [q]uit [m]ute [-][=][+]: "
		header1 = " [←][↑][→]: "
	)

	var buf [24]byte

	s := strconv.AppendFloat(buf[:0], m.FPS(), 'f', 0, 64)
	if len(s) < 2 {
		sb.WriteByte(' ')
	}
	sb.Write(s)
	sb.WriteString(header0)

	sb.Write(strconv.AppendInt(buf[:0], int64(m.power), 10))
	sb.WriteString(header1)

	sb.Write(strconv.AppendInt(buf[:0], int64(m.heat0), 10))
	sb.WriteByte('\n')
}
