package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

	var a, b []float64

	// Input data from ab2.dat
	fin, _ := os.Open("../ab2.dat")
	defer fin.Close()

	scanner := bufio.NewScanner(fin)

	_ = scanner.Scan()
	for _, snum := range strings.Fields(scanner.Text()) {
		if c, err := strconv.ParseFloat(snum, 64); err == nil {
			a = append(a, c)
		}
	}
	_ = scanner.Scan()
	for _, snum := range strings.Fields(scanner.Text()) {
		if c, err := strconv.ParseFloat(snum, 64); err == nil {
			b = append(b, c)
		}
	}

	processing(a, b)

	// Output result into ab2.res
	fout, _ := os.Create("../results/ab2.res")
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
