package cmd

import (
	"fmt"
	"strings"
	rv "mars/internal/rover"
	gd "mars/internal/grid"
)

func Run(sizes ... int) {
	spaces := gd.CreateSpaces(sizes...)
	rover := rv.InitRover()
	var commands string

	final := &gd.Final{FinalPosition: [2]int{rover.X, rover.Y}, FinalDirection: rover.Direction, Status: "Ready"}

	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("======================================================================")
		fmt.Printf(
			`{"position": [%d, %d], "direction": "%s", "status": "%s"}`+"\n",
			final.FinalPosition[0], final.FinalPosition[1],
			final.FinalDirection, final.Status,
		)
		fmt.Println("======================================================================")
		rv.DrawSpaces(spaces, rover)


		fmt.Print("Enter rover commands (M/L/R): ")
		_, err := fmt.Scan(&commands)
		if err != nil {
			fmt.Println("error reading input:", err)
			continue
		}
		commands = strings.ToUpper(commands)

		if commands == "EXIT" || commands == "QUIT" || commands == "Q" {
			fmt.Println("Exiting...")
			break
		}

		final = rv.MoveRover(spaces, rover, &gd.Final{}, commands)

		fmt.Println("======================================================================")

		if final.Status == "Obstacle encountered" || final.Status == "Out of bounds" {
			fmt.Println(final.Status)
			break
		}
	}
}