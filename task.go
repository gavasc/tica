package main

import (
	"time"

	//"github.com/jmoiron/sqlx"
)

type Task struct {
  Id        uint
  Code      string
  Desc      string
  totalTime time.Time
}

func (t Task) Create() {
  db := connectDb()
  defer db.Close()

  db.MustExec("INSERT INTO tasks (code, total_time) VALUES (?, ?)", t.Code, time.Now())
}
