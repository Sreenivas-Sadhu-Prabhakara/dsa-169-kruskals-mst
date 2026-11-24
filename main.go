package main

import (
	"fmt"
	"sort"
)

// Edge is an undirected weighted edge.
type Edge struct {
	U, V, W int
}

// dsu is a disjoint-set union with path compression and union by rank.
type dsu struct {
	parent, rank []int
}

func newDSU(n int) *dsu {
	d := &dsu{parent: make([]int, n), rank: make([]int, n)}
	for i := range d.parent {
		d.parent[i] = i
	}
	return d
}

func (d *dsu) find(x int) int {
	for d.parent[x] != x {
		d.parent[x] = d.parent[d.parent[x]] // path halving
		x = d.parent[x]
	}
	return x
}

// union merges the sets of a and b; false if they were already joined.
func (d *dsu) union(a, b int) bool {
	ra, rb := d.find(a), d.find(b)
	if ra == rb {
		return false
	}
	if d.rank[ra] < d.rank[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	if d.rank[ra] == d.rank[rb] {
		d.rank[ra]++
	}
	return true
}

// KruskalMST returns the minimum spanning tree's total weight: edges taken
// lightest-first, skipping any whose endpoints are already connected.
// ok is false when the graph is disconnected.
func KruskalMST(n int, edges []Edge) (total int, ok bool) {
	sorted := append([]Edge{}, edges...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].W < sorted[j].W })
	d := newDSU(n)
	used := 0
	for _, e := range sorted {
		if d.union(e.U, e.V) {
			total += e.W
			used++
			if used == n-1 {
				break
			}
		}
	}
	return total, used >= n-1
}

func main() {
	edges := []Edge{{0, 1, 2}, {0, 2, 1}, {1, 2, 1}, {2, 3, 2}, {3, 4, 1}, {4, 2, 2}}
	fmt.Println(KruskalMST(5, edges))
}
