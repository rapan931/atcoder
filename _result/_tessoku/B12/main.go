package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	N := sF()
	L, R := 0.0, 100.0

	for L < R {
		M := (L + R) / 2.0

		result := M * M * M + M
		if math.Abs(result - N) < 0.001 {
			fmt.Println(M)
			return
		} else if  result < N {
			L = M
		} else {
			R = M
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

func sI() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func sIs(n int) []int {
	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = sI()
	}

	return l
}

func sI2() (int, int) {
	return sI(), sI()
}

func sI3() (int, int, int) {
	return sI(), sI(), sI()
}

func sI4() (int, int, int, int) {
	return sI(), sI(), sI(), sI()
}

func sF() float64 {
	return float64(sI())
}

func sS() string {
	sc.Scan()
	return sc.Text()
}

func sB() []byte {
	return []byte(sS())
}

func sGrid(h, w int) [][]int {
	g := make([][]int, h)
	for i := 0; i < h; i++ {
		g[i] = sIs(w)
	}
	return g
}

func sByteGrid(h int) [][]byte {
	g := make([][]byte, h)
	for i := 0; i < h; i++ {
		g[i] = sB()
	}
	return g
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

// x ** y
func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// x ** y
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

// SEE: https://zenn.dev/johniel/articles/f1028e37f91489
// SEE: https://rtoch.com/posts/golang-segment-tree/
type SegmentTree struct {
	data []int
	n    int
	t    SegmentTreeType
}

const (
	SEGMENT_MAX = 100000000
	SEGMENT_MIN = -100000000
)

type SegmentTreeType int

const (
	SegmentTreeTypeSum SegmentTreeType = iota
	SegmentTreeTypeMin
	SegmentTreeTypeMax
)

func newSegmentTree(data []int, t SegmentTreeType) *SegmentTree {
	n := len(data)

	segTree := new(SegmentTree)
	segTree.n = 1
	for segTree.n < n {
		segTree.n *= 2
	}

	segTree.data = make([]int, segTree.n*2-1)
	segTree.t = t
	switch t {
	case SegmentTreeTypeMin:
		for i := 0; i < segTree.n*2-1; i++ {
			segTree.data[i] = SEGMENT_MAX
		}
	case SegmentTreeTypeMax:
		for i := 0; i < segTree.n*2-1; i++ {
			segTree.data[i] = SEGMENT_MIN
		}
	}

	return segTree
}

// ある頂点の表す区間が [a, b) であるならその子の表す区間はそれぞれ、
// [a, (a+b)/2) と [(a+b)/2, b) です。
//
// * [a, b) は a以上b未満の区間(bは含まない区間)を表す
func (segTree *SegmentTree) query(begin, end, idx, a, b int) int {
	if b <= begin || end <= a {
		// クエリと関係のない区間
		switch segTree.t {
		case SegmentTreeTypeSum:
			return 0
		case SegmentTreeTypeMin:
			return SEGMENT_MAX
		case SegmentTreeTypeMax:
			return SEGMENT_MIN
		default:
			// ありえない
			return 0
		}
	}

	if begin <= a && b <= end {
		// 全体がクエリの対象になる区間
		return segTree.data[idx]
	}

	// 一部がクエリの対象にならない場合は子に尋ねる
	v1 := segTree.query(begin, end, idx*2+1, a, (a+b)/2)
	v2 := segTree.query(begin, end, idx*2+2, (a+b)/2, b)

	switch segTree.t {
	case SegmentTreeTypeSum:
		return v1 + v2
	case SegmentTreeTypeMin:
		return min(v1, v2)
	case SegmentTreeTypeMax:
		return max(v1, v2)
	default:
		// ありえない
		return 0
	}
}

// [begin, end)
func (segTree *SegmentTree) Query(begin, end int) int {
	return segTree.query(begin, end, 0, 0, segTree.n)
}

// idx += segTree.n - 1 は、末端の葉を探している。
// +---------------+
// |       0       |
// +-------+-------+
// |   1   |   2   |
// +---+---+-------+
// | 3 | 4 | 5 | 6 |
// +-+-+---+-+-+-+-+
// |7|8|9|A|B|C|D|E|
// +-+-+-+-+-+-+-+-+
func (segTree *SegmentTree) update(idx, x int) {
	idx += segTree.n - 1
	segTree.data[idx] = x
	for 0 < idx {
		idx = (idx - 1) / 2

		switch segTree.t {
		case SegmentTreeTypeSum:
			segTree.data[idx] = segTree.data[idx*2+1] + segTree.data[idx*2+2]
		case SegmentTreeTypeMin:
			segTree.data[idx] = min(segTree.data[idx*2+1], segTree.data[idx*2+2])
		case SegmentTreeTypeMax:
			segTree.data[idx] = max(segTree.data[idx*2+1], segTree.data[idx*2+2])
		}
	}
}

func (segTree *SegmentTree) Build(arr []int) {
	for i, v := range arr {
		segTree.update(i, v)
	}
}

func binStrToInt(str string) uint64 {
	v, error := strconv.ParseUint(str, 2, 64)
	if error != nil {
		panic("")
	}
	return v
}

// index, valueを返す
// 見つけられなかったら、invdexはlength、valueは0を返す
func lowerBound(data []int, v int) (int, int) {
	lb := sort.Search(len(data), func(i int) bool { return data[i] >= v })
	if lb < len(data) {
		return lb, data[lb]
	} else {
		return len(data), 0
	}
}

// index, valueを返す
// 見つけられなかったら、invdexはlength、valueは0を返す
func upperBound(data []int, v int) (int, int) {
	lb := sort.Search(len(data), func(i int) bool { return data[i] > v })
	if lb < len(data) {
		return lb, data[lb]
	} else {
		return len(data), 0
	}
}

// index,を返す
// 見つけられなかったら、invdexはlength
func binarySearch(data []int, v int) int {
	L, R := 0, len(data)-1

	for L <= R {
		M := (L + R + 1) / 2
		if data[M] == v {
			return M
		}
		if data[M] < v {
			L = M + 1
		} else {
			R = M - 1
		}
	}

	return len(data)
}
