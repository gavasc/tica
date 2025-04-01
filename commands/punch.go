package commands

import (
	"fmt"
	"time"

	"github.com/gavasc/tica/data"
)

func PunchHandler(args []string) {
	t := data.Task{Code: args[0]}
	if !t.Exists() {
		fmt.Println("Creating task ", t.Code)
		t.Create()
	}

	err := t.GetIdByCode()
	if err != nil {
		print(err.Error())
		return
	}

	h := data.History{
		TaskId: t.Id,
	}
	if h.LastPunchType() == data.In {
		fmt.Println("Punching out task", t.Code)
		punchOut(h, t)
	} else {
		fmt.Println("Punching in task", t.Code)
		punchIn(h)
	}
}

// "private" functions

func punchIn(h data.History) {
	h.Type = data.In
	h.Punch()
}

func punchOut(h data.History, t data.Task) {
	h.Type = data.Out
	timeSpent := int(time.Now().Unix()) - h.LastPunchTime()
	h.Punch()
	t.AddToTotal(timeSpent)
	fmt.Printf("Time spent in %s: %ds\n", t.Code, timeSpent)
}
