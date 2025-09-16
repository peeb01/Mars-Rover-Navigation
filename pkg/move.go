package pkg

import (
	// "fmt"
	"strings"
)

func MoveRover(g *Spaces, r *Rover, final *Final ,commands string) *Final {
	command := strings.Split(commands, "")
	delta := map[string][2]int{
		"N": {0, 1},
		"E": {1, 0},
		"S": {0, -1},
		"W": {-1, 0},
	}
	for _, cmd := range command {
		if cmd == "M" {
			dx, dy := delta[r.Direction][0], delta[r.Direction][1]
			// fmt.Printf("Moving from (%d, %d) to (%d, %d)\n", r.X, r.Y, r.X+dx, r.Y+dy)
			newX, newY := r.X + dx, r.Y + dy
			if newX >= 0 && newX < g.Col && newY >= 0 && newY < g.Row {
				r.X, r.Y = newX, newY
				for _, obs := range g.Obstacles {
					if r.X == obs.X && r.Y == obs.Y {
						final.FinalDirection = r.Direction
						final.FinalPosition = [2]int{r.X, r.Y}
						final.Status = "Obstacle encountered"
						return final
					} else {
						final.FinalDirection = r.Direction
						final.FinalPosition = [2]int{r.X, r.Y}
						final.Status = "Success"
					}
				}
			} else {
				r.X, r.Y = newX, newY
				final.FinalDirection = r.Direction
				final.FinalPosition = [2]int{r.X, r.Y}
				final.Status = "Out of bounds"
				return final
			}
		}
		if cmd == "L" {
			switch r.Direction {
			case "N":
				r.Direction = "W"
				final.FinalDirection = r.Direction
				final.FinalPosition = [2]int{r.X, r.Y}
				final.Status = "Success"
			case "W":
				r.Direction = "S"
				final.FinalDirection = r.Direction
				final.FinalPosition = [2]int{r.X, r.Y}
				final.Status = "Success"
			case "S":
				r.Direction = "E"
				final.FinalDirection = r.Direction
				final.FinalPosition = [2]int{r.X, r.Y}
				final.Status = "Success"
			case "E":
				r.Direction = "N"
				final.FinalDirection = r.Direction
				final.FinalPosition = [2]int{r.X, r.Y}
				final.Status = "Success"
			}
		}
		if cmd == "R" {
			switch r.Direction {
			case "N":
				r.Direction = "E"
				final.FinalDirection = r.Direction
				final.FinalPosition = [2]int{r.X, r.Y}
				final.Status = "Success"
			case "E":
				r.Direction = "S"
				final.FinalDirection = r.Direction
				final.FinalPosition = [2]int{r.X, r.Y}
				final.Status = "Success"
			case "S":
				r.Direction = "W"
				final.FinalDirection = r.Direction
				final.FinalPosition = [2]int{r.X, r.Y}
				final.Status = "Success"
			case "W":
				r.Direction = "N"
				final.FinalDirection = r.Direction
				final.FinalPosition = [2]int{r.X, r.Y}
				final.Status = "Success"
			}
		}
	}
	return final
}
