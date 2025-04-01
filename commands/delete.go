package commands

import (
	"fmt"

	"github.com/gavasc/tica/data"
)

func DeleteHandler(task string) {
	t := data.Task{Code: task}
	if t.Exists() {
		t.Delete()
		fmt.Printf("Task %s deleted!\n", task)
	}
}
