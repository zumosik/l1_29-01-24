package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type  (
	tRow []int
	tData2 []tRow
)	

func main() {
	var (
		a	tData2
		c	tRow
		row int
	)

	// Чтение данных из файла 2D.dat
	fin, _ := os.Open("2D.dat")
	defer fin.Close()

	for scanner := bufio.NewScanner(fin); scanner.Scan(); row++ {
		if scanner.Text() == "" {
			break
		}
		c = make([]int, 0)
		for _, snum := range strings.Fields(scanner.Text()) {
			if x, err := strconv.Atoi(snum); err == nil {
				c = append(c, x)
			}
		}
		fmt.Printf("Row #%2d: %v\n", row, c)
		a = append(a, c)
	}
	fmt.Println(a)
}
