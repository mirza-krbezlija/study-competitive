package introduction

import "fmt"

type Task1 struct{}

func (t Task1) Run() {
	// 1 <= a < b <= 10
	var a, b int
	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		panic(err)
	}

	if b-a == 1 || (a == 1 && b == 10) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
