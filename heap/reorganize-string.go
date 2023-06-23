package main

import (
	"container/heap"
	"fmt"
)

type heapItem struct {
    letter string
    cnt int
}

type Heap []*heapItem

func (h Heap) Len() int {
    return len(h)
}

func (h Heap) Less(i, j int) bool {
    return h[i].cnt > h[j].cnt
}

func (h Heap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(*heapItem))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getHeap(s string) Heap {
    d := make(map[string]int)

    for _, c := range s {
        d[string(c)]++
    }

    result := make(Heap, len(d))
    i := 0
    for l, cnt := range d {
        result[i] = &heapItem{letter: l, cnt: cnt}

        i++
    }

    return result
}

func reorganizeString(s string) string {
    h := getHeap(s)

	heap.Init(&h)
    var result string
    var prev *heapItem
	for h.Len() > 0 {
        item := (heap.Pop(&h)).(*heapItem)

        result += item.letter
        item.cnt--
        if prev != nil && prev.cnt > 0 {
            heap.Push(&h, prev)
        }

        prev = item
	}

    if len(s) != len(result) {
        return ""
    }

    return result
}
