package pkg_test

import (
	"bytes"
	"mars/pkg"
	"os"
	"testing"
)

// Test all rotations (L and R through all 4 directions)
func TestMoveRoverAllRotations(t *testing.T) {
	spaces := &pkg.Spaces{Row: 5, Col: 5}
	r := pkg.InitRover()

	// 4 left turns should bring rover back to N
	for i := 0; i < 4; i++ {
		pkg.MoveRover(spaces, r, &pkg.Final{}, "L")
	}
	if r.Direction != "N" {
		t.Errorf("expected N after 4 left turns, got %s", r.Direction)
	}

	// 4 right turns should bring rover back to N
	for i := 0; i < 4; i++ {
		pkg.MoveRover(spaces, r, &pkg.Final{}, "R")
	}
	if r.Direction != "N" {
		t.Errorf("expected N after 4 right turns, got %s", r.Direction)
	}
}

// Test CreateSpaces creates correct dimensions and obstacles
func TestCreateSpaces(t *testing.T) {
	spaces := pkg.CreateSpaces(10, 20)
	if spaces.Row != 10 || spaces.Col != 20 {
		t.Errorf("expected 10x20, got %dx%d", spaces.Row, spaces.Col)
	}
	if len(spaces.Obstacles) == 0 {
		t.Errorf("expected some obstacles, got none")
	}
	// protected zone should not contain obstacles
	protected := map[[2]int]bool{{0, 0}: true, {0, 1}: true, {1, 0}: true, {1, 1}: true}
	for _, obs := range spaces.Obstacles {
		if protected[[2]int{obs.X, obs.Y}] {
			t.Errorf("obstacle placed in protected zone (%d,%d)", obs.X, obs.Y)
		}
	}
}

// Test DrawSpaces prints car and obstacles
func TestDrawSpaces(t *testing.T) {
	spaces := &pkg.Spaces{
		Row: 2, Col: 2,
		Obstacles: []pkg.Obstacle{{X: 1, Y: 1}},
	}
	r := &pkg.Rover{X: 0, Y: 0}

	old := os.Stdout
	rPipe, wPipe, _ := os.Pipe()
	os.Stdout = wPipe

	pkg.DrawSpaces(spaces, r)

	wPipe.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(rPipe)
	out := buf.String()

	if !bytes.Contains([]byte(out), []byte("🚗")) {
		t.Errorf("expected car emoji, got %s", out)
	}
	if !bytes.Contains([]byte(out), []byte("X")) {
		t.Errorf("expected obstacle, got %s", out)
	}
}

// Test PrintFinalState wraps ToJSON correctly
func TestPrintFinalState(t *testing.T) {
	f := &pkg.Final{
		FinalPosition:  [2]int{3, 4},
		FinalDirection: "S",
		Status:         "Success",
	}
	jsonStr, err := pkg.PrintFinalState(f)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if jsonStr == "" {
		t.Errorf("expected JSON output, got empty string")
	}
}

// TestMoveRoverPlainSuccess ensures rover moves correctly without obstacles
func TestMoveRoverPlainSuccess(t *testing.T) {
	spaces := &pkg.Spaces{Row: 3, Col: 3}
	r := &pkg.Rover{X: 1, Y: 1, Direction: "N"}

	_ = pkg.MoveRover(spaces, r, &pkg.Final{}, "M")

	if r.X != 1 || r.Y != 2 {
		t.Errorf("expected rover at (1,2) but got (%d,%d)", r.X, r.Y)
	}
	if r.Direction != "N" {
		t.Errorf("expected rover facing N but got %s", r.Direction)
	}
}


// TestMoveRoverSequence tests a sequence of commands
func TestMoveRoverSequence(t *testing.T) {
	spaces := &pkg.Spaces{Row: 5, Col: 5}
	r := pkg.InitRover()

	final := pkg.MoveRover(spaces, r, &pkg.Final{}, "MMRML")

	if r.X != 1 || r.Y != 2 {
		t.Errorf("expected rover at (1,2), got (%d,%d)", r.X, r.Y)
	}
	if final.FinalDirection != "N" {
		t.Errorf("expected direction N, got %s", final.FinalDirection)
	}
}

// TestCreateSpacesSquare checks single-arg CreateSpaces
func TestCreateSpacesSquare(t *testing.T) {
	spaces := pkg.CreateSpaces(5)
	if spaces.Row != 5 || spaces.Col != 5 {
		t.Errorf("expected 5x5, got %dx%d", spaces.Row, spaces.Col)
	}
}

// TestCreateSpacesPanic checks CreateSpaces panics on invalid args
func TestCreateSpacesPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic but did not get one")
		}
	}()
	_ = pkg.CreateSpaces(1, 2, 3) // invalid call
}


// TestMoveRoverMultiStep covers multiple commands in one string
func TestMoveRoverMultiStep(t *testing.T) {
	spaces := &pkg.Spaces{
		Row: 5, Col: 5,
		Obstacles: []pkg.Obstacle{{X: 2, Y: 2}}, // obstacle in path
	}
	r := pkg.InitRover()

	final := pkg.MoveRover(spaces, r, &pkg.Final{}, "MMRMM")

	if final.Status != "Obstacle encountered" {
		t.Errorf("expected Obstacle encountered, got %s", final.Status)
	}
}

// TestCreateSpacesSquareArg covers single argument usage
func TestCreateSpacesSquareArg(t *testing.T) {
	spaces := pkg.CreateSpaces(4)
	if spaces.Row != 4 || spaces.Col != 4 {
		t.Errorf("expected 4x4, got %dx%d", spaces.Row, spaces.Col)
	}
}

// TestCreateSpacesInvalidArgs ensures panic on invalid args
func TestCreateSpacesInvalidArgs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic but got none")
		}
	}()
	_ = pkg.CreateSpaces(1, 2, 3) // invalid call
}

// TestPrintFinalStateKeys ensures JSON has required fields
func TestPrintFinalStateKeys(t *testing.T) {
	f := &pkg.Final{
		FinalPosition:  [2]int{0, 0},
		FinalDirection: "N",
		Status:         "Ready",
	}
	jsonStr, _ := pkg.PrintFinalState(f)
	if !bytes.Contains([]byte(jsonStr), []byte("status")) {
		t.Errorf("expected JSON to contain status, got %s", jsonStr)
	}
	if !bytes.Contains([]byte(jsonStr), []byte("final_position")) {
		t.Errorf("expected JSON to contain final_position, got %s", jsonStr)
	}
}

// TestDrawSpacesOverlap checks when rover stands on obstacle
func TestDrawSpacesOverlap(t *testing.T) {
	spaces := &pkg.Spaces{
		Row: 2, Col: 2,
		Obstacles: []pkg.Obstacle{{X: 0, Y: 0}}, // same as rover
	}
	r := &pkg.Rover{X: 0, Y: 0}

	old := os.Stdout
	rPipe, wPipe, _ := os.Pipe()
	os.Stdout = wPipe

	pkg.DrawSpaces(spaces, r)

	wPipe.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(rPipe)
	out := buf.String()

	// Ensure car emoji takes precedence over obstacle
	if !bytes.Contains([]byte(out), []byte("🚗")) {
		t.Errorf("expected car emoji, got %s", out)
	}
}
