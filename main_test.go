package main

import "testing"

func TestKruskalMST(t *testing.T) {
	edges := []Edge{{0, 1, 2}, {0, 2, 1}, {1, 2, 1}, {2, 3, 2}, {3, 4, 1}, {4, 2, 2}}
	total, ok := KruskalMST(5, edges)
	if !ok || total != 5 {
		t.Errorf("KruskalMST = (%d,%v), want 5", total, ok)
	}

	// Agreement with the known small case: cycle-closing edges are skipped.
	total, ok = KruskalMST(3, []Edge{{0, 1, 1}, {1, 2, 1}, {0, 2, 100}})
	if !ok || total != 2 {
		t.Errorf("triangle = (%d,%v), want 2", total, ok)
	}

	if _, ok := KruskalMST(3, []Edge{{0, 1, 1}}); ok {
		t.Error("disconnected graph must report not-ok")
	}
	if total, ok := KruskalMST(1, nil); !ok || total != 0 {
		t.Errorf("single vertex = (%d,%v)", total, ok)
	}
}

func TestDSU(t *testing.T) {
	d := newDSU(4)
	if !d.union(0, 1) || !d.union(2, 3) {
		t.Error("fresh unions must succeed")
	}
	if d.union(1, 0) {
		t.Error("repeated union must report already-joined")
	}
	if d.find(0) != d.find(1) || d.find(2) != d.find(3) {
		t.Error("sets not merged")
	}
	if d.find(0) == d.find(2) {
		t.Error("separate sets merged spuriously")
	}
}
