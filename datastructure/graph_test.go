package datastructure

import (
	"container/list"
)

// Graph 图的存储
type Graph struct {
	VertexNum     int                // 顶点的个数
	AdjacencyList map[int]*list.List // 邻接表存储关系
}

func NewGraph(v int) *Graph {
	graph := &Graph{
		VertexNum:     v,
		AdjacencyList: make(map[int]*list.List),
	}
	for i := 0; i < v; i++ {
		graph.AdjacencyList[i] = list.New()
	}
	return graph
}

func (graph *Graph) AddEdge(a, b int) {
	graph.AdjacencyList[a].PushBack(b)
	graph.AdjacencyList[b].PushBack(a)
}
