package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{
		[]string{"Maruska", "15", "10/20", "(10.32, 56.21, 30.25)"},
		[]string{"Nastenka", "30", "30/50", "(1,1,1)"},
		[]string{"Sengina", "21", "80/80", "(1,1,1)"},
		[]string{"Jednorozec", "8", "30/40", "(1,1,1)"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NPC", "Speed", "Power", "Location"})
	table.AppendBulk(data)
	table.Render()
}
