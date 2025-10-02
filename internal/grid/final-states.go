package grid

import (
    "encoding/json"
)

type Final struct {
    FinalPosition  [2]int `json:"final_position"`
    FinalDirection string `json:"final_direction"` // N, E, S, W
    Status         string `json:"status"`          // Success, Obstacle encountered, Out of bounds
}

func (f *Final) ToJSON() (string, error) {
    b, err := json.MarshalIndent(f, "", "  ")
    if err != nil {
        return "", err
    }
    return string(b), nil
}

func PrintFinalState(f *Final) (string, error) {
	return f.ToJSON()
}