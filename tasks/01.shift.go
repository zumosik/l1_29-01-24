package main

import (
	"fmt"
	"os"
)

func shift(a []int) {
	first := a[0]
	for i := 0; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
	a[len(a)-1] = first
}

func main() {
	fin, _ := os.Open("../a1.dat") // we use same data as 01.task_reverse_2.go
	defer fin.Close()

	// Input data from a1.dat
	var n int
	fmt.Fscanln(fin, &n)
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Fscan(fin, &a[i]); err != nil {
			break
		}
	}
	//Before
	fmt.Println(a)

	shift(a)

	//After
	fmt.Println(a)

	// Output result into a1.res
	fout, _ := os.Create("../results/shift.res")
	defer fout.Close()
	for _, c := range a {
		fmt.Fprint(fout, c, " ")
	}
	fmt.Fprintln(fout)
}
