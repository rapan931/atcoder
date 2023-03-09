package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	buf := make([]byte, 10*1024)
	sc.Buffer(buf, math.MaxInt32)

	sc.Split(bufio.ScanWords)
}

func main() {
}

func scan() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
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

func printSliceSepSpace(nums []int) {
	fmt.Println(strings.Trim(fmt.Sprint(nums), "[]"))
}

type Queue struct{ v []int }

func NewQueue(n int) *Queue    { return &Queue{v: make([]int, n)} }
func (q *Queue) isEmpty() bool { return len(q.v) == 0 }
func (q *Queue) len() int      { return len(q.v) }
func (q *Queue) push(i int)    { q.v = append(q.v, i) }
func (q *Queue) pop() int {
	v := q.v[0]
	q.v = q.v[1:]
	return v
}

type Stack struct{ v []int }

func NewStack(n int) *Stack    { return &Stack{v: make([]int, n)} }
func (s *Stack) isEmpty() bool { return len(s.v) == 0 }
func (s *Stack) len() int      { return len(s.v) }
func (s *Stack) push(i int)    { s.v = append(s.v, i) }
func (s *Stack) pop() int {
	v := s.v[len(s.v)-1]
	s.v = s.v[:len(s.v)]
	return v
}
