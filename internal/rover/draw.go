package rover

import (
	"fmt"
	"mars/internal/grid")

func DrawSpaces(g *grid.Spaces, r *Rover) {
	for y := g.Row - 1; y >= 0; y-- {
		for x := 0; x < g.Col; x++ {
			if r.X == x && r.Y == y {
				fmt.Print("🚗 ")
				continue
			}
			hasObstacle := false
			for _, obs := range g.Obstacles {
				if obs.X == x && obs.Y == y {
					hasObstacle = true
					break
				}
			}
			if hasObstacle {
				fmt.Print("X ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
