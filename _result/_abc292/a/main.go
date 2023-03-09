package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

type tuple struct{ x, y int }

func main() {
	s := scans()
	fmt.Println(strings.ToUpper(s))
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
