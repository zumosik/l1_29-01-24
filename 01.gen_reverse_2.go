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
	f, _ := os.Create("a2.dat")
	defer f.Close()
	for i := 1; i < n; i++ {
		fmt.Fprint(f, rand.Intn(100) + 1, " ")
	}
	fmt.Fprintln(f, rand.Intn(100) + 1)
}
