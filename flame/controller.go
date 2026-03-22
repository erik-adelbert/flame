package flame

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

// Update handles incoming messages and updates the model accordingly.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		m.stepFire()

		m.sample(time.Time(msg)) // update the FPS counter

		return m, tick()

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "m":
			m.mute = !m.mute

		case "-":
			m.power = max(MinPower, m.power-5)
		case "=":
			m.power = BasePower
		case "+":
			m.power = min(MaxPower, m.power+5)

		case "left":
			m.heat0 = max(MinHeat, m.heat0-5)
		case "up":
			m.heat0 = BaseHeat
		case "right":
			m.heat0 = min(MaxHeat, m.heat0+5)
		}

	case tea.WindowSizeMsg:
		h := max(1, msg.Height)
		w := max(1, msg.Width)

		m.Resize(h, w) // reserve one line for the header
	}

	return m, nil
}

// Init starts the fire simulation.
func (m *Model) Init() tea.Cmd {
	return tick()
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(TimeStep, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
