// Package main initializes and runs the fire simulation using the Bubble Tea
// TUI framework.
package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/erik-adelbert/flame/flame"

	"golang.org/x/term"
)

// main initializes the fire simulation and starts the Bubble Tea program.
func main() {
	w, h, err := term.GetSize(int(os.Stdin.Fd()))

	if err != nil {
		fatal("Could not get terminal size:", err)
	}

	// ensure the dimensions are strictly positive
	h = max(1, h)
	w = max(1, w)

	p := tea.NewProgram(flame.NewModel(h, w))

	if _, err := p.Run(); err != nil {
		fatal("Error running program:", err)
	}
}

func fatal(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}
