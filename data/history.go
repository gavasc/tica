package data

import (
	"time"

	"gorm.io/gorm"
)

type PunchType int
const (
  In PunchType = iota
  Out
)

type History struct {
  gorm.Model
  Task Task
  Type PunchType
  timeStamp time.Time
}
