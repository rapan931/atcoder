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
	buf := make([]byte, 1000*1024)
	sc.Buffer(buf, math.MaxInt32)

	sc.Split(bufio.ScanWords)
}

func main() {
	H, W := scan2()
	m := make([][]byte, H)

	for i := 0; i < H; i++ {
		m[i] = scanb()
	}

	// 左上、右上、左下、右下
	var a = []int{}
	var b = []int{}
	var c = []int{}
	var d = []int{}

	for i := 0; i < H; i++ {
		for ii := 0; ii < W; ii++ {
			if m[i][ii] == '#' {
				a = append(a, i, ii)
				break
			}
		}
		if len(a) != 0 {
			break
		}
	}

	for i := 0; i < H; i++ {
		for ii := W - 1; 0 <= ii; ii-- {
			if m[i][ii] == '#' {
				b = append(b, i, ii)
				break
			}
		}
		if len(b) != 0 {
			break
		}
	}

	for i := H - 1; 0 <= i; i-- {
		for ii := 0; ii < W; ii++ {
			if m[i][ii] == '#' {
				c = append(c, i, ii)
				break
			}
		}
		if len(c) != 0 {
			break
		}
	}

	for i := H - 1; 0 <= i; i-- {
		for ii := W - 1; 0 <= ii; ii-- {
			if m[i][ii] == '#' {
				d = append(d, i, ii)
				break
			}
		}
		if len(d) != 0 {
			break
		}
	}

	if a[1] != c[1] {
		if a[1] < c[1] {
			c[1] = a[1]
		} else {
			a[1] = c[1]
		}
	}

	if b[1] != d[1] {
		if b[1] < d[1] {
			b[1] = d[1]
		} else {
			d[1] = b[1]
		}
	}

	for i := a[0]; i < c[0]+1; i++ {
		for ii := a[1]; ii < b[1]+1; ii++ {
			if m[i][ii] == '.' {
				printSliceSepSpace([]int{i + 1, ii + 1})
			}
		}

	}
}

func max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a int, b int) int {
	if b < a {
		return b
	}
	return a
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

func scanf() float64 {
	return float64(scan())
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

func repeatSlice(rep int, v int) []int {
	s := make([]int, rep)
	for i := range s {
		s[i] = v
	}
	return s
}

func contains(list []int, v int) bool {
	for _, s := range list {
		if v == s {
			return true
		}
	}
	return false
}
