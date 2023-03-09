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
	users := []int{}
	for i := 0; i < n; i++ {
		users = append(users, 0)
	}

	q := scan()
	for i := 0; i < q; i++ {
		event := scan()
		userIndex := scan() - 1
		switch event {
		case 1:
			users[userIndex] += 1
		case 2:
			users[userIndex] += 2
		case 3:
			if users[userIndex] >= 2 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
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
