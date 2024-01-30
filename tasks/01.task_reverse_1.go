package main

import (
	"fmt"
	"os"
)

func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

/*
func reverse(a []int) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}

*/

func main() {

	fin, _ := os.Open("../a1.dat")
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

	reverse(a)

	//After
	fmt.Println(a)

	// Output result into a1.res
	fout, _ := os.Create("../results/a1.res")
	defer fout.Close()
	for _, c := range a {
		fmt.Fprint(fout, c, " ")
	}
	fmt.Fprintln(fout)
}
