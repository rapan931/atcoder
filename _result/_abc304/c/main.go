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

func dfs(pos int) {
	visited[pos] = true;

	for i := 0; i < len(G[pos]); i ++ {
		next := G[pos][i]
		if visited[next] == false {
			dfs(next)
		}
	}
}

var visited = make([]bool, 2009)
var G = make([][]int, 2009)

func main() {
	N, D := scan2()

	pairs := make([][]int, N)

	for i := 0; i < N; i++ {
		pairs[i] = append(pairs[i], scan(), scan())
	}

	for i := 0; i < N; i++ {
		baseX := pairs[i][0]
		baseY := pairs[i][1]

		for ii := 0; ii < N; ii++ {
			if i == ii {
				continue
			}
			if ((baseX-pairs[ii][0])*(baseX-pairs[ii][0])+(baseY-pairs[ii][1])*(baseY-pairs[ii][1])) <= D*D {
				G[i] = append(G[i], ii)
			}
		}
	}
	dfs(0)

	for i := 0; i < N; i++ {
		if visited[i] == true {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func init() {
	buf := make([]byte, 10*1024)
	sc.Buffer(buf, math.MaxInt32)

	sc.Split(bufio.ScanWords)
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
