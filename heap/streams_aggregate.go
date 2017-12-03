package main

import (
	"container/heap"
	"log"
)

type data struct {
	weight int
	raw    []byte
}

type weightedIndexItem struct {
	weight int
	index  int
}

type priorityQueue []weightedIndexItem

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(weightedIndexItem))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	if len(old) == 0 {
		return nil
	}
	var r weightedIndexItem
	r, *pq = old[len(old)-1], old[:len(old)-1]
	return r
}

func streamsAggregate(in []data, ins ...[]data) chan data {
	output := make(chan data, 1)
	input := append([][]data{in}, ins...)
	go func() {
		defer close(output)
		pq := &priorityQueue{}
		for i := range input {
			if len(input[i]) > 0 {
				heap.Push(pq, weightedIndexItem{index: i, weight: input[i][0].weight})
			}
		}
		for pq.Len() > 0 {
			item := heap.Pop(pq).(weightedIndexItem)
			datum := input[item.index]
			consume := func(d data) {
				output <- d
			}
			consume(datum[0])
			if len(datum) > 1 {
				input[item.index] = datum[1:]
				item.weight = input[item.index][0].weight
				heap.Push(pq, item)
			}
		}

	}()
	return output
}

func main() {
	for d := range streamsAggregate(
		[]data{{weight: 1, raw: []byte("abc")}, {weight: 8, raw: []byte("vwx")}, {weight: 9, raw: []byte("yz0")}},
		[]data{{weight: 2, raw: []byte("def")}, {weight: 7, raw: []byte("stu")}, {weight: 10, raw: []byte("123")}},
		[]data{{weight: 3, raw: []byte("ghi")}, {weight: 6, raw: []byte("pqr")}, {weight: 11, raw: []byte("456")}},
		[]data{{weight: 4, raw: []byte("jkl")}, {weight: 5, raw: []byte("mno")}, {weight: 12, raw: []byte("789")}},
	) {
		log.Printf("%s\n", string(d.raw))
	}
}
