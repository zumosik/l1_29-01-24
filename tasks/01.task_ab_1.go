package main

import (
	"fmt"
	"math"
	"os"
)

func processing(a []float64, b []float64) {
	if len(a) != len(b) {
		return
	}
	for i := 0; i < len(a); i++ {
		a[i], b[i] = (a[i]+b[i])/2, math.Sqrt(a[i]*b[i])
	}
}

func main() {

	fin, _ := os.Open("../ab1.dat")
	defer fin.Close()

	// Input data from ab1.dat
	var n int
	fmt.Fscanln(fin, &n)
	a := make([]float64, n, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Fscan(fin, &a[i]); err != nil {
			break
		}
	}
	b := make([]float64, n, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Fscan(fin, &b[i]); err != nil {
			break
		}
	}

	processing(a, b)

	// Output result into ab1.res
	fout, _ := os.Create("../results/ab1.res")
	defer fout.Close()
	for _, c := range a {
		fmt.Fprintf(fout, "%9.4f", c)
	}
	fmt.Fprintln(fout)
	for _, c := range b {
		fmt.Fprintf(fout, "%9.4f", c)
	}
	fmt.Fprintln(fout)
}
