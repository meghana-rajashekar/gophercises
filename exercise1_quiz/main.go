package main

import (
        "encoding/csv"
        "fmt"
	"math/rand"
        "time"
	"bufio"
        "os"
        "flag"
        "log"
	"strings"
)

type Problem struct {
	question string
	solution string
}
func readCsv() []Problem{
        f, err := os.Open(fileName)

        if err != nil {
                log.Fatal(err)
        }
        defer f.Close()

        csvReader := csv.NewReader(f)
        problems, err := csvReader.ReadAll()
        if err != nil {
                log.Fatal(err)
        }

	var problem_list []Problem
	for _, problem := range problems {
		problem_list = append(problem_list, Problem{question:problem[0], solution:problem[1]})
	}
	return problem_list

}

func randomize(problems []Problem) []Problem {
        for i := range problems {
            j := rand.Intn(i + 1)
            problems[i], problems[j] = problems[j], problems[i]
        }
	return problems
}


func quiz(problems []Problem) (int, int) {
        var correctCounter int
        var problemCounter int

	problemCounter = len(problems)
	var problemNumber int
        for _, each_problem := range problems {
		problemNumber++
                fmt.Printf("Problem #%v: %v= ", problemNumber, each_problem.question)
                newtimer := time.NewTimer(time.Duration(timeLimit) * time.Second)
                go func(problemCounter int, correctCounter int) {
                        <-newtimer.C

                        fmt.Println("Sorry! You ran out of time!")
                        fmt.Printf("Your score: %d/%d \n", correctCounter, problemCounter)
                        os.Exit(0)
                    }(problemCounter, correctCounter)
                var userAnswer string
                fmt.Scanln(&userAnswer)
                newtimer.Stop()
                userAnswer = strings.TrimSpace(userAnswer)
                if strings.ToLower(userAnswer) == strings.ToLower(each_problem.solution) {
                        correctCounter++
                }

        }
	return problemCounter, correctCounter
}

var (
	fileName string
	timeLimit int
	shuffleFlag bool
)

func init() {

        flag.StringVar(&fileName, "file", "problems.csv", "Problem file name")
        flag.IntVar(&timeLimit, "timeLimit", 30, "Time limit")
        flag.BoolVar(&shuffleFlag, "shuffle", false, "Shuffle quiz")
        flag.Parse()

}

func main() {

	problems := readCsv()

	fmt.Printf("You will have %d seconds to attempt each problem. \n Good luck!", timeLimit)
	fmt.Print("Press 'Enter' to start the quiz...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
        if shuffleFlag {
                randomize(problems)
        }

	problemCounter, correctCounter := quiz(problems)
        fmt.Printf("Your score: %d/%d \n", correctCounter, problemCounter)
        if correctCounter > (problemCounter/2) {
                fmt.Println("You did great! COngrats :)")
        } else {
                fmt.Println("You can do better next time. :)")
        }
}

