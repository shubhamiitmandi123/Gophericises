package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readCsv(fileName string) []string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Eror ocured ", err)
	}
	Datast := string(data)
	Questions := strings.Split(Datast, "\n")
	return Questions
}
