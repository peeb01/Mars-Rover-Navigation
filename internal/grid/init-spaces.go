package grid

import (
    "math/rand"
    "math"
)

type Obstacle struct {
	X int
	Y int
}

type Spaces struct {
	Row int
	Col int
    Obstacles []Obstacle
}

func CreateSpaces(sizes ...int) *Spaces {
    var row, col int
    if len(sizes) == 1 {
        row, col = sizes[0], sizes[0]
    } else if len(sizes) == 2 {
        row, col = sizes[0], sizes[1]
    } else {
        panic("CreateSpaces requires 1 or 2 args")
    }

    numObstacles := int(math.Ceil(float64(row*col) / 10.0))

    spaces := &Spaces{Row: row, Col: col}

    protec := map[[2]int]bool{
        {0, 0}: true, {0, 1}: true,
        {1, 0}: true, {1, 1}: true,
    }

    maxAvailable := row*col - len(protec)
    if maxAvailable < numObstacles {
        numObstacles = maxAvailable
        if numObstacles < 0 {
            numObstacles = 0
        }
    }

    occup := make(map[[2]int]bool)
    for len(spaces.Obstacles) < numObstacles {
        x := rand.Intn(col)
        y := rand.Intn(row)
        pos := [2]int{x, y}
        if !protec[pos] && !occup[pos] {
            occup[pos] = true
            spaces.Obstacles = append(spaces.Obstacles, Obstacle{X: x, Y: y})
        }
    }

    return spaces
}
