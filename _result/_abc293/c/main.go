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
	H := scan()
	W := scan()

	l := make([][]int, H)
	for i := 0; i < H; i++ {
		list := make([]int, W)
		for ii := 0; ii < W; ii++ {
			list[ii] = scan()
		}

		l[i] = list
	}

	stepCount := H + W - 2
	count := 0
	f := func(indexes []int) {
		rightsSteps := make([]int, stepCount)

		for _, v := range indexes {
			rightsSteps[v] = 1
		}

		x, y := 0, 0
		m := map[int]int{l[y][x]: 1}
		addFlag := true
		for _, rightStep := range rightsSteps {
			if rightStep == 1 {
				x += 1
			} else {
				y += 1
			}

			v := l[y][x] 
			if _, ok := m[v]; !ok {
				m[v] = 1
			} else {
				addFlag = false
				break
			}
		}

		if addFlag {
			count += 1
		}
	}

	each_combination(stepCount, W - 1, f)
	fmt.Println(count)
}

func init() {
	buf := make([]byte, 10*1024)
	sc.Buffer(buf, math.MaxInt32)

	sc.Split(bufio.ScanWords)
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

func NewQueue() *Queue         { return &Queue{} }
func (q *Queue) isEmpty() bool { return len(q.v) == 0 }
func (q *Queue) len() int      { return len(q.v) }
func (q *Queue) push(i int)    { q.v = append(q.v, i) }
func (q *Queue) pop() int {
	v := q.v[0]
	q.v = q.v[1:]
	return v
}

type Stack struct{ v []int }

func NewStack(n int) *Stack    { return &Stack{} }
func (s *Stack) isEmpty() bool { return len(s.v) == 0 }
func (s *Stack) len() int      { return len(s.v) }
func (s *Stack) push(i int)    { s.v = append(s.v, i) }
func (s *Stack) pop() int {
	v := s.v[len(s.v)-1]
	s.v = s.v[:len(s.v)]
	return v
}

func each_combination(n int, k int, f func([]int)) {
	indexes := make([]int, k)
	recursive_combination(indexes, n-1, k, f)
}

func recursive_combination(indexes []int, s int, rest int, f func([]int)) {
	if rest == 0 {
		f(indexes)
	} else {
		if s < 0 {
			return
		}
		recursive_combination(indexes, s-1, rest, f)
		indexes[rest-1] = s
		recursive_combination(indexes, s-1, rest-1, f)
	}
}
