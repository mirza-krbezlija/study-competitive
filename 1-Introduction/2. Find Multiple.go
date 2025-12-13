package introduction

import "fmt"

type Task2 struct{}

func (t Task2) Run() {
	// 1 <= a <= b <= 1000
	// 1 <= c <= 1000
	var a, b, c int
	_, err := fmt.Scanln(&a, &b, &c)
	if err != nil {
		panic(err)
	}

	for i := a; i <= b; i++ {
		if i%c == 0 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
