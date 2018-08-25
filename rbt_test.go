package rbt

import (
	"testing"
)

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

	tree.Insert(1, "1")
	tree.Insert(3, "3")
	tree.Insert(4, "4")
	tree.Insert(6, "6")
	tree.Insert(5, "5")
	tree.Insert(2, "2")
	for i := 1; i <= 6; i++ {
		tree.Delete(i)
	}
	if tree.Size() != 0 {
		t.Error("Error whilst deletion")
	}
}
