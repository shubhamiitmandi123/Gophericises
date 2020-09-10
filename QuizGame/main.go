package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	// Aurgument Colletion
	fileName := flag.String("fname", "problem.csv", "Csv File name")
	timeOut := flag.Int("t", 30, "Quiz time")
	flag.Parse()

	//File Read
	qData := readCsv(*fileName)

	//Making Problem set
	problemSet := getProblemSet(qData)
	numberOfProblems := len(problemSet)

	//Confirm To start
	fmt.Print("Would you like to start Quiz? Enter any key:")
	var ch string
	fmt.Scanf("%s", &ch)
	fmt.Println("Time has started now!")

	//Conduct Quiz and Declear Result
	timer := time.NewTimer(time.Duration(*timeOut) * time.Second)
	ansC := make(chan int)
	cnt := 0
	for i, v := range problemSet {
		fmt.Printf("Q %v: %s = ? : ", i+1, v.Que)
		go func() {
			var ans int
			fmt.Scanf("%d", &ans)
			ansC <- ans
		}()
		select {
		case ans := <-ansC:
			if ans == v.Ans {
				cnt++
			}
		case <-timer.C:
			fmt.Println("\nYour Time is Expired !")
			fmt.Printf("You Solved %v out of %v\n", cnt, numberOfProblems)
			return
		}
	}
	fmt.Printf("You Solved %v out of %v\n", cnt, numberOfProblems)
}
