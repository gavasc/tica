package data

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
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
    task_id INTEGER,
    FOREIGN KEY (task_id) REFERENCES tasks(id)
  );`
)

func CreateDb() {
	db := connectDb()
	defer db.Close()

	db.MustExec(createDbQuery)
}

func connectDb() *sqlx.DB {
	homeDir, _ := os.UserHomeDir()
	db := sqlx.MustConnect("sqlite", homeDir+"/.tica/data.db")

	return db
}
