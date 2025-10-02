# 🚀 Mars Rover Simulation

This project simulates a Mars rover moving on a 2D grid with obstacles.\
The rover can turn left, right, or move forward while avoiding obstacles
and boundaries.

------------------------------------------------------------------------

## 📂 Project Structure

    .
    ├── cmd/               # Application commands
    │   ├── run.go         # Interactive runner
    │   ├── run_once.go    # Test helper
    │   └── run_test.go    # Unit tests for cmd
    ├── internal/
    │   ├── grid/          # Grid, obstacles, and final state
    │   └── rover/         # Rover logic (movement, drawing)
    ├── main.go            # Entry point
    ├── Dockerfile         # Docker build
    ├── Makefile           # Makefile shortcuts
    └── README.md

------------------------------------------------------------------------

## 🏃 How to Run

### Run the application

``` bash
go run main.go
```

You will see the grid, the rover's position, direction, and status.\
Then you can enter commands interactively.

### Available commands

-   **M** → Move forward
-   **L** → Turn left
-   **R** → Turn right
-   **Q / QUIT / EXIT** → Quit the simulation

### Example

    ======================================================================
    {"position": [0, 0], "direction": "N", "status": "Ready"}
    ======================================================================
    . . .
    . . X
    🚗 . .
    Enter rover commands (M/L/R): M

------------------------------------------------------------------------

## 🧪 Run Tests

Run all tests with coverage:

``` bash
go test ./... -cover
```

Run tests for a specific package:

``` bash
go test ./cmd -v
```
``` bash
go test ./internal/rover -v
```
``` bash
go test ./internal/grid -v
```
------------------------------------------------------------------------

## 🐳 Run with Docker

### Build Docker image

``` bash
docker build -t mars-rover .
```

### Run container

``` bash
docker run -it mars-rover
```

This will start the interactive rover simulation inside a container.

------------------------------------------------------------------------
## Run with Makefile
``` bash
make prepare        # install dependencies (Go, golangci-lint)
make build          # build binary into ./bin/mars
make run            # run the program
make test           # run all tests
make cover          # generate text coverage report
make cover-html     # generate HTML coverage report
make lint           # run golangci-lint
make clean          # clean build artifacts
```

------------------------------------------------------------------------

## 📜 License

This project is licensed under the MIT License.
