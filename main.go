package main

import (
	"os"
	"strconv"
	"strings"

	introduction "github.com/mirza-krbezlija/study-competitive/1-introduction"
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
	}

	tasks[taskNr].Run()
}
