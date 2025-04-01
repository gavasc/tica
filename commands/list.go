package commands

import (
	"log"
	"os"
	"strconv"

	"github.com/gavasc/tica/data"
	"github.com/olekukonko/tablewriter"
)

// Gets all the tasks from the db, appends each one to a string and returns that string
func ListHandler() {
	tasks, err := data.Task{}.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Code", "Description", "Total Time"})

	for _, task := range tasks {
		table.Append([]string{task.Code, task.Description, strconv.Itoa(task.TotalTime/60) + "min"})
	}

	table.Render()
}
