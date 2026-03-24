// Package flame implements a fire simulation using the Bubble Tea TUI
// framework.
//
// The simulation is based on a simple cellular automaton where each cell's heat
// value is influenced by its neighbors. The model maintains a grid of heat
// values and corresponding symbols that represent the fire visually. The
// controller updates the model based on a timer and user input, while the view
// renders the current state of the model as a string for display.
package flame

import "time"

const (
	// BasePower is a parameter that controls the intensity of the fire,
	// affecting how many new cells are ignited each step.
	BasePower = 100
	// MinPower and MaxPower define the bounds for the fire intensity that the
	// user can adjust.
	MinPower = 10
	MaxPower = 200

	// BaseHeat is the heat value assigned to new ignited cells at the bottom of
	// the grid.
	BaseHeat = 65
	// MinHeat and MaxHeat define the bounds for the heat value that is assigned
	// to new ignited cells at the bottom of the grid.
	MinHeat = 30
	MaxHeat = 200

	// TimeStep is thåe duration between each update of the fire simulation, controlling
	// the speed of the animation.
	TimeStep = 30 * time.Millisecond
)
