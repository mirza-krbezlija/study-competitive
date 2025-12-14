package loopsandarrays

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task4 struct{}

func (t Task4) Run() {
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

	a := make(map[int]bool, n)
	for _, a_iStr := range lineArr {
		a_i, _ := strconv.Atoi(a_iStr)
		a[a_i] = true
	}
	for i := 1; i <= n; i++ {
		if !a[i] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
