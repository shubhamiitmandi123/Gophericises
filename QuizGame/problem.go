package main

import (
	"strconv"
	"strings"
)

// Question :
type Question struct {
	Que string
	Ans int
}

func getProblemSet(Questions []string) []Question {
	Q := make([]Question, len(Questions))
	for i, v := range Questions {
		q := strings.Split(v, ",")
		Q[i].Ans, _ = strconv.Atoi(q[1])
		Q[i].Que = q[0]

	}
	return Q
}
