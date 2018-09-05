package rbt

import (
	"strconv"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("not equal, a=%s,b=%s", a, b)
	}
}

func TestTraverse(t *testing.T) {
	tree := NewTree()

	tree.Insert(1, "1")
	tree.Insert(3, "3")
	tree.Insert(4, "4")
	tree.Insert(6, "6")
	tree.Insert(5, "5")
	tree.Insert(2, "2")
	if tree.Size() != 6 {
		t.Error("Error whilst insertion")
	}
}

func TestSearch(t *testing.T) {

	tree := NewTree()

	tree.Insert(1, "1")
	tree.Insert(2, "2")
	tree.Insert(3, "3")

	n := tree.Search(1)
	if n.value != "1" {
		t.Error("Wrong value")
	}
}

func TestDelete(t *testing.T) {
	tree := NewTree()

	for i := 1; i < 1000; i++ {
		tree.Insert(int64(i), strconv.Itoa(i))
	}

	for i := 1; i <= 1000; i++ {
		tree.Delete(int64(i))
	}
	if tree.Size() != 0 {
		t.Error("Error whilst deletion")
	}
}

func Test_Nearest(t *testing.T) {
	tree := NewTree()
	tree.Insert(1, "")
	tree.Insert(2, "")
	tree.Insert(3, "")
	tree.Insert(5, "")
	tree.Insert(8, "")
	tree.Insert(9, "")
	tree.Insert(12, "")
	y := tree.Nearest(int64(10))
	assertEqual(t, y.key, int64(9))
}

func Test_Circular_Queue(t *testing.T) {
	tree := NewTree()
	tree.Insert(1, "")
	tree.Insert(2, "")
	tree.Insert(3, "")
	tree.Insert(5, "")
	tree.Insert(8, "")
	tree.Insert(9, "")
	tree.Insert(12, "")
	h := FindMinimum(tree)
	assertEqual(t, h.key, int64(1))
	assertEqual(t, h.successor().key, int64(2))
}
