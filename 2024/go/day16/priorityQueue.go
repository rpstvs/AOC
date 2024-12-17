package main

type pqNode struct {
	Node
	Score int
}

type priorityQueue []pqNode

func (n priorityQueue) len() int {
	return len(n)
}

func (n priorityQueue) less(i, j int) bool {
	return n[i].Score < n[j].Score
}

func (n priorityQueue) swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *priorityQueue) push(x interface{}) {
	*n = append(*n, x.(pqNode))
}

func (n *priorityQueue) pop() interface{} {
	old := *n
	l := len(old)
	x := old[l-1]
	*n = old[0 : l-1]
	return x
}
