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

	fin, _ := os.Open("../a2.dat")
	defer fin.Close()

	// Input data from a2.dat
	var a []int
	for {
		var c int
		if _, err := fmt.Fscan(fin, &c); err != nil {
			break
		}
		a = append(a, c)
	}
	//Before
	fmt.Println(a)

	reverse(a)

	//After
	fmt.Println(a)

	// Запись перевёрнутого массива в файл a2.res
	fout, _ := os.Create("../results/a2.res")
	defer fout.Close()
	for _, c := range a {
		fmt.Fprint(fout, c, " ")
	}
	fmt.Fprintln(fout)
}
