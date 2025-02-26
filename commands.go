package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

func PunchHandler(args []string) {
	t := Task{Code: args[0]}
	if !t.Exists() {
		fmt.Println("Creating task ", t.Code)
		t.Create()
	}

	err := t.GetIdByCode()
	if err != nil {
		print(err.Error())
		return
	}

	h := History{
		TaskId:    t.Id,
		timeStamp: int(time.Now().Unix()),
	}
	if h.LastPunchType() == In {
		fmt.Println("Punching out task", t.Code)
		punchOut(h, t)
	} else {
		fmt.Println("Punching in task", t.Code)
		punchIn(h)
	}
}

// Gets all the tasks from the db, appends each one to a string and returns that string
func ListHandler() {
	tasks, err := Task{}.GetAll()
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

func DeleteHandler(task string) {
	t := Task{Code: task}
	if t.Exists() {
		t.Delete()
		fmt.Printf("Task %s deleted!\n", task)
	}
}

// "private" functions

func punchIn(h History) {
	h.Type = In
	h.Punch()
}

func punchOut(h History, t Task) {
	h.Type = Out
	timeSpent := int(time.Now().Unix()) - h.LastPunchTime()
	h.Punch()
	t.AddToTotal(timeSpent)
	fmt.Printf("Time spent in %s: %ds\n", t.Code, timeSpent)
}
