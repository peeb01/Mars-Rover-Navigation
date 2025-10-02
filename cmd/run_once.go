package cmd

import (
	rv "mars/internal/rover"
	gd "mars/internal/grid"
)

// RunOnceForTest executes one step for testing purposes
func RunOnceForTest(spaces *gd.Spaces, rover *rv.Rover, commands string) *gd.Final {
	final := &gd.Final{
		FinalPosition:  [2]int{rover.X, rover.Y},
		FinalDirection: rover.Direction,
		Status:         "Ready",
	}
	return rv.MoveRover(spaces, rover, final, commands)
}
