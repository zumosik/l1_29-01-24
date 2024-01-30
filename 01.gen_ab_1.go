package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	fmt.Print("Enter n: ")
	var n int
	fmt.Scan(&n)
	f, _ := os.Create("ab1.dat")
	fmt.Fprintln(f, n)
	defer f.Close()
	for i := 0; i < n; i++ {
		fmt.Fprintf(f, "%7.2f", float64(rand.Intn(200))*0.5)
	}
	fmt.Fprintf(f, "\n")
	for i := 0; i < n; i++ {
		fmt.Fprintf(f, "%7.2f", float64(rand.Intn(200))*0.5)
	}
	fmt.Fprintf(f, "\n")
}
