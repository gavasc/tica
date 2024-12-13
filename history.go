package main

import (
	"strconv"
)

type PunchType int
const (
  In PunchType = iota + 1
  Out
)

type History struct {
  TaskId      uint
  Type      PunchType
  timeStamp int
}

func (h History) Punch() {
  db := connectDb()
  defer db.Close()

  db.MustExec("INSERT INTO history (task_id, punch_type, timestamp) VALUES (?, ?, ?)", h.TaskId, h.Type, h.timeStamp)
}

func (h History) LastPunchType() PunchType {
  db := connectDb()
  defer db.Close()

  db.Get(&h.Type, "SELECT punch_type FROM history WHERE task_id=? ORDER BY id DESC LIMIT 1;", h.TaskId)
  return h.Type
}

func (h History) LastPunchTime() int {
  db := connectDb()
  defer db.Close()

  var lastStr string
  db.Get(&lastStr, "SELECT timestamp FROM history WHERE task_id=? ORDER BY id DESC LIMIT 1;", h.TaskId)
  num, _ := strconv.Atoi(lastStr)
  return num
}
