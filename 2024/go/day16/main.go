package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func main() {
	score, _ := Dijkstra(parseinput())
	fmt.Println(score)
}

type Point struct {
	x, y int
}

type Direction uint8

func (d Direction) clockwise() Direction         { return (d + 4 + 1) % 4 }
func (d Direction) counter_clockwise() Direction { return (d + 4 - 1) % 4 }

const (
	E Direction = iota
	S
	W
	N
)

var Delta = map[Direction]Point{N: {-1, 0}, E: {0, 1}, S: {1, 0}, W: {0, -1}}

type Node struct {
	Pos       Point
	Direction Direction
}

func Dijkstra(m map[Point]struct{}, start Point, end Point) (int, map[Node][]Node) {
	var (
		minScore = math.MaxInt64
		parent   = map[Node][]Node{}
	)

	startNode := Node{start, E}

	pq := &priorityQueue{}

	heap.Init(pq)
	heap.Push(pq, pqNode{startNode, 0})

	scores := map[Node]int{startNode: 0}

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(pqNode)

		if curr.Pos == end {
			if curr.Score > minScore {
				return minScore, parent
			}

			minScore = curr.Score
		}
		for _, n := range neighbours(m, curr) {
			if _, ok := scores[n.Node]; !ok || n.Score <= scores[n.Node] {
				parent[n.Node] = append(parent[n.Node], curr.Node)
				scores[n.Node] = n.Score
				heap.Push(pq, n)
			}
		}
	}
	panic("no path found")
}

func parseinput() (map[Point]struct{}, Point, Point) {
	var (
		m     = map[Point]struct{}{}
		start Point
		end   Point
	)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		for j, ch := range scanner.Text() {
			switch ch {
			case 'S':
				m[Point{i, j}] = struct{}{}
				start = Point{i, j}
			case 'E':
				m[Point{i, j}] = struct{}{}
				start = Point{i, j}
			case '.':
				m[Point{i, j}] = struct{}{}
			case '#':
			}

		}
	}
	return m, start, end
}

func neighbours(m map[Point]struct{}, curr pqNode) []pqNode {
	list := []pqNode{
		{Node{curr.Pos, curr.Direction.clockwise()}, curr.Score + 1000},
		{Node{curr.Pos, curr.Direction.counter_clockwise()}, curr.Score + 1000},
	}

	delta := Delta[curr.Direction]

	if _, ok := m[Point{curr.Pos.x + delta.x, curr.Pos.y + delta.y}]; ok {
		list = append(list, pqNode{Node{
			Pos:       Point{curr.Pos.x + delta.x, curr.Pos.y + delta.y},
			Direction: curr.Direction,
		}, curr.Score + 1})
	}
	return list
}
