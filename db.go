package main

import (
  "github.com/jmoiron/sqlx"
  "github.com/mattn/go-sqlite3"
)

const (
  createDbQuery = `CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY,
    code TEXT,
    description TEXT,
    total_time INTEGER
  );
  CREATE TABLE IF NOT EXISTS history (
    id INTEGER PRIMARY KEY,
    timestamp INTEGER,
    punch_type INTEGER,
    task_id, INTEGER,
    FOREIGN KEY (task_id) REFERENCES tasks(id)
  );`
)

func connectDb() *sqlx.DB {
  var db *sqlx.DB

  db = sqlx.MustConnect("go-sqlite3", "data/data.db")
  db.MustExec(createDbQuery)

  return db
}

