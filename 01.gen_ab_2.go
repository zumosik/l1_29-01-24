package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	n := rand.Intn(12) + 9
	f, _ := os.Create("ab2.dat")
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
