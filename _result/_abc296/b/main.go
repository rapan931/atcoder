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
	list := [][]string{
		{"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8"},
		{"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7"},
		{"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6"},
		{"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5"},
		{"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4"},
		{"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3"},
		{"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2"},
		{"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1"},
	}

	H := 8
	W := 8
	for i := 0; i < H; i++ {
		s := scans()
		for ii := 0; ii < W; ii++ {
			if s[ii] == '*' {
				fmt.Println(list[i][ii])
				return
			}
		}
	}
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

func scan2() (int, int) {
	return scan(), scan()
}

func scan3() (int, int, int) {
	return scan(), scan(), scan()
}

func scan4() (int, int, int, int) {
	return scan(), scan(), scan(), scan()
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
func (q *Queue) first() int    { return q.v[0] }
func (q *Queue) isEmpty() bool { return len(q.v) == 0 }
func (q *Queue) len() int      { return len(q.v) }
func (q *Queue) push(i int)    { q.v = append(q.v, i) }
func (q *Queue) pop() int {
	v := q.v[0]
	q.v = q.v[1:]
	return v
}

type Stack struct{ v []int }

func NewStack() *Stack         { return &Stack{} }
func (s *Stack) first() int    { return s.v[0] }
func (s *Stack) isEmpty() bool { return len(s.v) == 0 }
func (s *Stack) len() int      { return len(s.v) }
func (s *Stack) push(i int)    { s.v = append(s.v, i) }
func (s *Stack) pop() int {
	v := s.v[len(s.v)-1]
	s.v = s.v[:len(s.v)-1]
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

func each_permutation(n int, r int) {
	if n == r {
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			nums[i] = i
		}
		recursive_permutation(len(nums), nums, func(indexes []int) {
			fmt.Println(indexes)
		})
	} else {
		each_combination(n, r, func(indexes []int) {
			recursive_permutation(len(indexes), indexes, func(indexes2 []int) {
				fmt.Println(indexes2)
			})
		})
	}

}

func recursive_permutation(n int, pat []int, f func([]int)) {
	if n == 1 {
		f(pat)
	} else {
		for i := 0; i < n; i++ {
			recursive_permutation(n-1, pat, f)
			if n%2 == 0 {
				pat[i], pat[n-1] = pat[n-1], pat[i]
			} else {
				pat[0], pat[n-1] = pat[n-1], pat[0]
			}
		}
	}
}

func factorialNum(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}

func powInt64(x, y uint64) uint64 {
	return uint64(math.Pow(float64(x), float64(y)))
}
