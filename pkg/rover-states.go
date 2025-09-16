package pkg

// State of the rover
type Rover struct {
    X, Y      int
    Direction string // N, E, S, W
}

func InitRover() *Rover {
    return &Rover{
        X:         0,
        Y:         0,
        Direction: "N",
    }
}