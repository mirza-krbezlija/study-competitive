package loopsandarrays

import (
	"bufio"
	"fmt"
	"os"
)

type Task6 struct{}

func (t Task6) Run() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(in, &n)

	var suma int64 = 0
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		var elm int
		if i == n-1 {
			fmt.Fscanln(in, &elm)
		} else {
			fmt.Fscan(in, &elm)
		}
		arr[i] = elm
		suma += int64(elm)
	}

	var x int64
	fmt.Fscanln(in, &x)

	div := int64(x / suma)
	var suma2 int64 = div * suma
	var idx int = n * int(div)
	for i := 0; i < n; i++ {
		suma2 += int64(arr[i])
		idx++
		if suma2 > x {
			break
		}
	}

	fmt.Println(idx)
}
