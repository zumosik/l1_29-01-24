package main

import (
	"fmt"
	"os"
)

// [1, 2, 3, 4, 6, 7, 8, 9, 5, 10] -> [2, 4, 6, 8, 10,    1, 3, 5, 7, 9]
func swap(a []int) {
	j := 0

	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			a[i], a[j] = a[j], a[i]
			j++
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

	swap(a)

	//After
	fmt.Println(a)

	// Output result into a1.res
	fout, _ := os.Create("../results/swap.res")
	defer fout.Close()
	for _, c := range a {
		fmt.Fprint(fout, c, " ")
	}
	fmt.Fprintln(fout)
}
