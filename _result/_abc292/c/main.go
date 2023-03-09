package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

type tuple struct{ x, y int }

func main() {
	n := scan()

	counts := map[int]int{}
	for a := 1; a < n; a++ {
		for b := 1; b < n; b++ {
			s := a*b
			if s >= n {
				break
			}
			if _, ok := counts[s]; ok {
				counts[s] += 1
			} else {
				counts[s] = 1
			}
		}
	}

	total := 0
	for i := 1; i < n; i++ {
		if _, ok := counts[i]; !ok {
			continue
		}

		if _, ok := counts[n - i]; !ok {
			continue
		}
		total += counts[i] * counts[n - i]
	}
	fmt.Println(total)
}

func scan() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func scans() string {
	sc.Scan()
	return sc.Text()
}

func scanb() []byte {
	sc.Scan()
	return sc.Bytes()
}
