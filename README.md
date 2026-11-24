# Minimum Spanning Tree (Kruskal)

> Striver SDE Sheet 169 · Day 24 Graph II — take edges lightest-first, and let union-find answer "would this close a cycle?" in near-constant time. Go, stdlib only, tested.

## Problem

The MST again — the sheet assigns both constructions because they factor the problem differently: Prim grows one tree; Kruskal merges a forest.

## Approach

Sort edges by weight; walk them, accepting any edge whose endpoints are in different components and skipping cycle-closers. The accept/skip test is the disjoint-set union structure: `find` with path halving plus union by rank gives effectively O(1) amortized operations (inverse Ackermann). Having taken n−1 edges, stop — fewer means disconnected. Kruskal's real virtue is that it never needs an adjacency structure at all, just the edge list — which is why it wins on sparse graphs and streaming edges, and why its DSU core is a reusable tool far beyond MSTs (cycle detection, connectivity queries, percolation).

**Complexity:** O(E log E) for the sort; DSU ops effectively O(1).

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
