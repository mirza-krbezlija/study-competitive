package loopsandarrays

import (
	"bufio"
	"fmt"
	"os"
)

type Task5 struct{}

func (t Task5) Run() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(in, &n)

	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		cnt[x]++
	}

	total := n * (n - 1) / 2
	same := 0
	for _, v := range cnt {
		if v > 1 {
			same += v * (v - 1) / 2
		}
	}

	fmt.Println(total - same)
}
