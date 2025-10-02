package cmd_test

import (
	"bytes"
	"os"
	"testing"
	rv "mars/internal/rover"
	gd "mars/internal/grid"
	"mars/cmd"
)

func runWithInput(t *testing.T, input string, sizes ...int) string {
	// Mock stdin
	r, w, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }()

	go func() {
		w.Write([]byte(input))
		w.Close()
	}()

	// Capture stdout
	oldStdout := os.Stdout
	rPipe, wPipe, _ := os.Pipe()
	os.Stdout = wPipe

	done := make(chan struct{})
	go func() {
		cmd.Run(sizes...)
		close(done)
	}()

	<-done

	wPipe.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(rPipe)
	return buf.String()
}

func TestRunQuit(t *testing.T) {
	out := runWithInput(t, "Q\n", 3)
	if !bytes.Contains([]byte(out), []byte("Exiting")) {
		t.Errorf("expected 'Exiting' in output, got:\n%s", out)
	}
}

func TestRunMoveAndQuit(t *testing.T) {
	out := runWithInput(t, "M\nQ\n", 3)
	if !bytes.Contains([]byte(out), []byte(`"position":`)) {
		t.Errorf("expected position update, got:\n%s", out)
	}
}

func TestRunRotateLeftAndQuit(t *testing.T) {
	out := runWithInput(t, "L\nQ\n", 3)
	if !bytes.Contains([]byte(out), []byte(`"direction":`)) {
		t.Errorf("expected direction update, got:\n%s", out)
	}
}

func TestRunRotateRightAndQuit(t *testing.T) {
	out := runWithInput(t, "R\nQ\n", 3)
	if !bytes.Contains([]byte(out), []byte(`"direction":`)) {
		t.Errorf("expected direction update, got:\n%s", out)
	}
}

func TestRunInvalidAndQuit(t *testing.T) {
	out := runWithInput(t, "INVALID\nQ\n", 3)
	if !bytes.Contains([]byte(out), []byte(`"status":`)) {
		t.Errorf("expected status in output, got:\n%s", out)
	}
}

// RunOnceForTest
func RunOnceForTest(spaces *gd.Spaces, rover *rv.Rover, commands string) *gd.Final {
    final := &gd.Final{
        FinalPosition:  [2]int{rover.X, rover.Y},
        FinalDirection: rover.Direction,
        Status:         "Ready",
    }
    return rv.MoveRover(spaces, rover, final, commands)
}

func TestRunOnceObstacle(t *testing.T) {
    spaces := &gd.Spaces{
        Row: 2, Col: 2,
        Obstacles: []gd.Obstacle{{X: 0, Y: 1}},
    }
    rover := rv.InitRover()
    final := cmd.RunOnceForTest(spaces, rover, "M")
    if final.Status != "Obstacle encountered" {
        t.Errorf("expected obstacle, got %s", final.Status)
    }
}

func TestRunOnceOutOfBounds(t *testing.T) {
    spaces := &gd.Spaces{Row: 1, Col: 1}
    rover := rv.InitRover()
    final := cmd.RunOnceForTest(spaces, rover, "M")
    if final.Status != "Out of bounds" {
        t.Errorf("expected out of bounds, got %s", final.Status)
    }
}
