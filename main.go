package main

import (
	"os"
	"strconv"
	"strings"

	introduction "github.com/mirza-krbezlija/study-competitive/1-introduction"
	loopsandarrays "github.com/mirza-krbezlija/study-competitive/2-loops-and-arrays"
)

type Task interface {
	Run()
}

func main() {
	taskNrStr := os.Args[1]
	taskNr, err := strconv.Atoi(strings.TrimSpace(taskNrStr))
	if err != nil {
		panic(err)
	}

	tasks := map[int]Task{
		1: introduction.Task1{},
		2: introduction.Task2{},
		3: loopsandarrays.Task3{},
		4: loopsandarrays.Task4{},
		5: loopsandarrays.Task5{},
		6: loopsandarrays.Task6{},
	}

	tasks[taskNr].Run()
}
