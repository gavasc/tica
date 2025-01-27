package main

import (
	"time"
	"fmt"
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

// Gets all the tasks from the db, appends each one to a string and returns that string 
func ListHandler() string {
  tasks, err := Task{}.GetAll()
  if err != nil {
    return err.Error()
  }

  tasksStr := ""
  for _, task := range tasks {
    tasksStr += fmt.Sprintf("%s -- %s -- %dmin\n", task.Code, task.Description, task.TotalTime/60)
  }

  return tasksStr
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
  timeSpent := int(time.Now().Unix()) -  h.LastPunchTime()
  h.Punch()
  t.AddToTotal(timeSpent)
  fmt.Printf("Time spent in %s: %ds\n", t.Code, timeSpent)
}
