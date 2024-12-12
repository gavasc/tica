/*
Copyright © 2024 gavasc <gavascdev@gmail.com>
*/
package cmd

import (
	"fmt"
	"time"

  data "github.com/gavasc/tica/data"
	"github.com/spf13/cobra"
)

// puchCmd represents the puch command
var punchCmd = &cobra.Command{
	Use:   "punch",
	Short: "Punches the current time in a task",
	Long: `Punches the current time in a task and, if the task does not exist,
it creates it an starts its timer`,
	Run: func(cmd *cobra.Command, args []string) {
    task := args[0]
    now := time.Now()
    fmt.Printf("Punching %s - %s\n", task, now.Format("02/01/2006 15:04")) 
    data.Task{ Code: task }.Create()
	},
}

func init() {
	rootCmd.AddCommand(punchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// puchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// puchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
