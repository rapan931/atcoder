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
	N := scan()
	M := scan()

	graph := make([][]int, N)
	deg := make([]int, N)

	for i := 0; i < M; i++ {
		a := scan()
		scanb()
		c := scan()
		scanb()

		// 処理しやすいように頂点を0始まりの値に修正
		a--
		c--

		graph[a] = append(graph[a], c)
		graph[c] = append(graph[c], a)

		// 次数アップ
		deg[a] += 1
		deg[c] += 1
	}

	circleCount := 0
	nonCircleCount := 0
	used := make([]int, N)

	for i := 0; i < N; i++ {
		if used[i] == 0 {
			queue := NewQueue()
			used[i] = 1
			queue.push(i)
			f := true
			for !queue.isEmpty() {
				q := queue.pop()
				if deg[q] != 2 {
					f = false
				}

				for _, v := range graph[q] {
					if used[v] == 0 {
						queue.push(v)
						used[v] = 1
					}
				}
			}
			if f {
				circleCount++
			} else {
				nonCircleCount++
			}
		}
	}
	fmt.Println(circleCount, nonCircleCount)
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

func NewStack() *Stack         { return &Stack{} }
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

// func factorialNum(n int) int {
// 	ret := 1
// 	for i := 2; i <= n; i++ {
// 		ret *= i
// 	}
// 	return ret
// }
//
// func assignColor(vertexColors []int, adjacencyMatrix [][]int) {
// 	id := 0
// 	for i := 0; i < len(adjacencyMatrix); i++ {
// 		if vertexColors[i] == 0 {
// 			dfs(i, id, vertexColors, adjacencyMatrix)
// 			id += 1
// 		}
// 	}
// }
//
// func dfs(vertexId int, colorId int, vertexColors []int, adjacencyMatrix [][]int) {
// 	stack := NewStack()
// 	stack.push(vertexId)
// 	vertexColors[vertexId] = colorId
//
// 	for !stack.isEmpty() {
// 		u := stack.pop()
// 		for i := 0; i < len(adjacencyMatrix[u]); i++ {
// 			if adjacencyMatrix[u][i] == 1 {
// 				vertexColors[i] = colorId
// 				stack.push(i)
// 			}
// 		}
// 	}
// }