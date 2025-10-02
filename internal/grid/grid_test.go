package grid_test

import (
	"testing"
	"mars/internal/grid"
)

// TestCreateSpacesSquare checks CreateSpaces with one argument
func TestCreateSpacesSquare(t *testing.T) {
	spaces := grid.CreateSpaces(5)
	if spaces.Row != 5 || spaces.Col != 5 {
		t.Errorf("expected 5x5, got %dx%d", spaces.Row, spaces.Col)
	}
	if len(spaces.Obstacles) == 0 {
		t.Errorf("expected some obstacles, got none")
	}
}

// TestCreateSpacesRectangle checks CreateSpaces with two arguments
func TestCreateSpacesRectangle(t *testing.T) {
	spaces := grid.CreateSpaces(4, 6)
	if spaces.Row != 4 || spaces.Col != 6 {
		t.Errorf("expected 4x6, got %dx%d", spaces.Row, spaces.Col)
	}
}

// TestCreateSpacesPanic ensures CreateSpaces panics on invalid args
func TestCreateSpacesPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic but did not get one")
		}
	}()
	_ = grid.CreateSpaces(1, 2, 3) // invalid
}

// TestFinalToJSON ensures Final serializes correctly
func TestFinalToJSON(t *testing.T) {
	f := &grid.Final{
		FinalPosition:  [2]int{2, 3},
		FinalDirection: "E",
		Status:         "Success",
	}
	jsonStr, err := f.ToJSON()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if jsonStr == "" {
		t.Errorf("expected non-empty JSON string")
	}
	if !contains(jsonStr, `"status"`) {
		t.Errorf("expected JSON to contain status field, got %s", jsonStr)
	}
}

// helper
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (string([]byte(s)[0:len(substr)]) == substr || contains(s[1:], substr))
}
