package main

type Task struct {
  Id        uint
  Code      string
  Desc      string
  totalTime int //stored in seconds
}

func (t Task) Create() {
  db := connectDb()
  defer db.Close()

  db.MustExec("INSERT INTO tasks (code, total_time) VALUES (?, 0)", t.Code)
}

func (t *Task) GetIdByCode() error {
  db := connectDb()
  defer db.Close()

  err := db.Get(&t.Id, "SELECT id FROM tasks WHERE code=?", t.Code)
  if err != nil {
    return err
  }

  return nil
}

func (t Task) Exists() bool {
  db := connectDb()
  defer db.Close()

  err := t.GetIdByCode()
  if err != nil {
    return false
  }

  return true
}

func (t Task) AddToTotal(seconds int) error {
  db := connectDb()
  defer db.Close()
  
  var current int 
  err := db.Get(&current, "SELECT total_time FROM tasks WHERE id=?;", t.Id)
  if err != nil {
    return err
  }

  newTotal := current + seconds
  db.MustExec("UPDATE tasks SET total_time=? WHERE id=?", newTotal, t.Id)

  return nil
}
