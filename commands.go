package main

import (
	"fmt"
	"time"
)

func punchHandler(args []string) {
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

  h := History {
    TaskId: t.Id,
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

func punchIn(h History) {
  h.Type = In
  h.Punch()
}

func punchOut(h History, t Task) {
  h.Type = Out
  timeSpent := int(time.Now().Unix()) -  h.LastPunchTime()
  h.Punch()
  t.AddToTotal(timeSpent)
  fmt.Printf("Time spent in %s: %ds", t.Code, timeSpent)
}
