package main

import (
	"fmt"
	"os"
)

// cmax replaces element with max element of all previous
func cmax(a []int) {
	max := a[0]
	for i := 1; i < len(a); i++ {
		prev := a[i]
		a[i] = max
		if prev > max {
			max = prev
		}
	}
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

	cmax(a)

	//After
	fmt.Println(a)

	// Output result into a1.res
	fout, _ := os.Create("../results/cmax.res")
	defer fout.Close()
	for _, c := range a {
		fmt.Fprint(fout, c, " ")
	}
	fmt.Fprintln(fout)
}
