package cmd

import (
	"fmt"
	"mars/pkg"
	"strings"
)

func Run(sizes ... int) {
	spaces := pkg.CreateSpaces(sizes...)
	rover := pkg.InitRover()
	var commands string

	final := &pkg.Final{FinalPosition: [2]int{rover.X, rover.Y}, FinalDirection: rover.Direction, Status: "Ready"}

	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("======================================================================")
		fmt.Printf(
			`{"position": [%d, %d], "direction": "%s", "status": "%s"}`+"\n",
			final.FinalPosition[0], final.FinalPosition[1],
			final.FinalDirection, final.Status,
		)
		fmt.Println("======================================================================")
		pkg.DrawSpaces(spaces, rover)


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

		final = pkg.MoveRover(spaces, rover, &pkg.Final{}, commands)

		fmt.Println("======================================================================")

		if final.Status == "Obstacle encountered" || final.Status == "Out of bounds" {
			fmt.Println(final.Status)
			break
		}
	}
}