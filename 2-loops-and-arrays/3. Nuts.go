package loopsandarrays

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task3 struct{}

func (t Task3) Run() {
	in := bufio.NewReader(os.Stdin)

	var n int
	_, err := fmt.Fscanln(in, &n)
	if err != nil {
		panic(err)
	}
	line, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	lineArr := strings.Split(strings.TrimSpace(line), " ")

	nuts := 0
	for _, nutsStr := range lineArr {
		nutsI, _ := strconv.Atoi(nutsStr)
		if nutsI > 10 {
			nuts += (nutsI - 10)
		}
	}

	fmt.Println(nuts)
}
