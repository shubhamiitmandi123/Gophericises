package main

import (
	"fmt"
	"os"
	"path/filepath"

	"task_app/cmd"
	"task_app/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir() //string and err
	dbpath := filepath.Join(home, "tasks.db")
	must(db.Init(dbpath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
