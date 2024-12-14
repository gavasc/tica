package main

type Task struct {
  Id          uint
  Code        string
  Description string
  TotalTime   int `db:"total_time"` //stored in seconds
}

func (t Task) Create() {
  db := connectDb()
  defer db.Close()

  db.MustExec("INSERT INTO tasks (code, description, total_time) VALUES (?, '(no description)', 0)", t.Code)
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

func (Task) GetAll() ([]Task, error) {
  db := connectDb()
  defer db.Close()
  
  tasks := []Task{}

  err := db.Select(&tasks, "SELECT * FROM tasks;")

  return tasks, err
}

func (t Task) Delete() {
  db := connectDb()
  defer db.Close()

  db.MustExec("DELETE FROM tasks WHERE code=?", t.Code)
}
