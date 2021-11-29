package bst

import (
	"constraints"
	// "fmt"
	"sort"
	"testing"
)

func TestEmptyBST(t *testing.T) {
	bst := NewBST[int]()
	if bst.IsEmpty() == false {
		t.Fatalf(`IsEmpty() reports false on an empty binary tree!`)
	}
}

func TestOneElementBST(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	if bst.IsEmpty() == true {
		t.Fatalf(`IsEmpty() reports true on a non-empty binary tree!`)
	}
}

func ListEq[T constraints.Ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInorderTraversal(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 90, 4}
	for _, value := range samples {
		bst.Insert(value)
	}
	result := bst.InorderTraversal()
	sort.Ints(samples)
	if ListEq(result, samples) == false {
		t.Fatalf(`InorderTraversal() returns incorrect sequence`)
	}
}

func TestInorderTraversalWithDuplicates(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 90, 3, 4, 18}
	for _, value := range samples {
		bst.Insert(value)
	}
	result := bst.InorderTraversal()
	sort.Ints(samples)
	if ListEq(result, samples) == false {
		t.Fatalf(`InorderTraversalWithDuplicates() returns incorrect sequence`)
	}
}

func TestPreorderTraversal(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 90, 4}
	for _, value := range samples {
		bst.Insert(value)
	}
	result := bst.PreorderTraversal()
	expected := []int{3, 2, 1, 52, 9, 7, 4, 18, 63, 90}
	if ListEq(result, expected) == false {
		t.Fatalf(`PreorderTraversal() returns incorrect sequence`)
	}
}

func TestPreorderTraversalWithDuplicates(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 90, 4, 9, 63}
	for _, value := range samples {
		bst.Insert(value)
	}
	result := bst.PreorderTraversal()
	expected := []int{3, 2, 1, 52, 9, 9, 7, 4, 18, 63, 63, 90}
	if ListEq(result, expected) == false {
		t.Fatalf(`PreorderTraversal() returns incorrect sequence`)
	}
}

func TestPostorderTraversal(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 90, 4}
	for _, value := range samples {
		bst.Insert(value)
	}
	result := bst.PostorderTraversal()
	expected := []int{1, 2, 4, 7, 18, 9, 90, 63, 52, 3}
	if ListEq(result, expected) == false {
		t.Fatalf(`PostorderTraversal() returns incorrect sequence`)
	}
}

func TestPostorderTraversalWithDuplicates(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 18, 90, 4, 1}
	for _, value := range samples {
		bst.Insert(value)
	}
	result := bst.PostorderTraversal()
	expected := []int{1, 1, 2, 4, 7, 18, 18, 9, 90, 63, 52, 3}
	if ListEq(result, expected) == false {
		t.Fatalf(`PostorderTraversal() returns incorrect sequence`)
	}
}

func TestMinMax(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 90, 4}
	for _, value := range samples {
		bst.Insert(value)
	}
	if max, ok := bst.Max(); !ok {
		t.Fatal(`Max() returns not ok on non-empty bst`)
	} else {
		if max != 90 {
			t.Fatalf(`Max() returns incorrect max value`)
		}
	}
	if min, ok := bst.Min(); !ok {
		t.Fatalf(`Min() returns not ok on non-empty bst`)
	} else {
		if min != 1 {
			t.Fatalf(`Min() returns incorrect min value`)
		}
	}
}

func TestSearch(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 90, 4}
	for _, value := range samples {
		bst.Insert(value)
	}
	p, ok := bst.Search(63)
	if !ok {
		t.Fatalf("Search() failed to return correct ok status")
	}
	if p == nil {
		t.Fatalf("Search() failed to find a key in bst")
	}
	if p.key != 63 {
		t.Fatalf("Search() returns incorrect pointer")
	}
}

func TestDeleteLeave(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 90, 4}
	for _, value := range samples {
		bst.Insert(value)
	}
	p, ok := bst.Search(18)
	if !ok {
		t.Fatalf("Search() failed to return correct ok status")
	}
	if p == nil {
		t.Fatalf("Search() failed to find a key in bst")
	}
	bst.Delete(p)
	p, ok = bst.Search(18)
	if ok {
		t.Fatalf("Search() failed to return correct ok status")
	}
	if p != nil {
		t.Fatalf("Search() p should be nil since we just removed it")
	}
	result := bst.InorderTraversal()
	expected := []int{1, 2, 3, 4, 7, 9, 52, 63, 90}
	if ListEq(result, expected) == false {
		t.Fatalf(`Delete() does not work in deleting a leave node`)
	}
}

func TestDeleteInternalNodeSingleLeftChild(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 90, 4}
	for _, value := range samples {
		bst.Insert(value)
	}
	p, ok := bst.Search(7)
	if !ok {
		t.Fatalf("Search() failed to return correct ok status")
	}
	if p == nil {
		t.Fatalf("Search() failed to find a key in bst")
	}
	bst.Delete(p)
	p, ok = bst.Search(7)
	if ok {
		t.Fatalf("Search() failed to return correct ok status")
	}
	if p != nil {
		t.Fatalf("Search() p should be nil since we just removed it")
	}
	result := bst.InorderTraversal()
	expected := []int{1, 2, 3, 4, 9, 18, 52, 63, 90}
	if ListEq(result, expected) == false {
		t.Fatalf(`Delete() does not work in deleting internal node with single left child`)
	}
}

func TestDeleteInternalNodeSingleRightChild(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 90, 4}
	for _, value := range samples {
		bst.Insert(value)
	}
	p, ok := bst.Search(63)
	if !ok {
		t.Fatalf("Search() failed to return correct ok status")
	}
	if p == nil {
		t.Fatalf("Search() failed to find a key in bst")
	}
	bst.Delete(p)
	p, ok = bst.Search(63)
	if ok {
		t.Fatalf("Search() failed to return correct ok status")
	}
	if p != nil {
		t.Fatalf("Search() p should be nil since we just removed it")
	}
	result := bst.InorderTraversal()
	expected := []int{1, 2, 3, 4, 7, 9, 18, 52, 90}
	if ListEq(result, expected) == false {
		t.Fatalf(`Delete() does not work in deleting internal node with single right child`)
	}
}

func TestDeleteInternalNodeTwoChildren(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 90, 4}
	for _, value := range samples {
		bst.Insert(value)
	}
	p, ok := bst.Search(52)
	if !ok {
		t.Fatalf("Search() failed to return correct ok status")
	}
	if p == nil {
		t.Fatalf("Search() failed to find a key in bst")
	}
	bst.Delete(p)
	p, ok = bst.Search(52)
	if ok {
		t.Fatalf("Search() failed to return correct ok status")
	}
	if p != nil {
		t.Fatalf("Search() p should be nil since we just removed it")
	}
	result := bst.InorderTraversal()
	expected := []int{1, 2, 3, 4, 7, 9, 18, 63, 90}
	if ListEq(result, expected) == false {
		t.Fatalf(`Delete() does not work in deleting internal node with two children`)
	}
}

func TestDeleteRoot(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 90, 4}
	for _, value := range samples {
		bst.Insert(value)
	}
	p, ok := bst.Search(3)
	if !ok {
		t.Fatalf("Search() failed to return correct ok status")
	}
	if p == nil {
		t.Fatalf("Search() failed to find a key in bst")
	}
	bst.Delete(p)
	p, ok = bst.Search(3)
	if ok {
		t.Fatalf("Search() failed to return correct ok status")
	}
	if p != nil {
		t.Fatalf("Search() p should be nil since we just removed it")
	}
	result := bst.InorderTraversal()
	expected := []int{1, 2, 4, 7, 9, 18, 52, 63, 90}
	if ListEq(result, expected) == false {
		t.Fatalf(`Delete() does not work in deleting root node`)
	}
}

func TestDeleteDuplicate(t *testing.T) {
	bst := NewBST[int]()
	samples := []int{3, 52, 63, 2, 9, 1, 18, 7, 18, 90, 4, 9}
	for _, value := range samples {
		bst.Insert(value)
	}
	p, ok := bst.Search(9)
	if !ok {
		t.Fatalf("Search() failed to return correct ok status")
	}
	if p == nil {
		t.Fatalf("Search() failed to find a key in bst")
	}
	bst.Delete(p)
	p, ok = bst.Search(9)
	if !ok {
		t.Fatalf("Search() failed to return correct ok status")
	}
	if p == nil {
		t.Fatalf("Search() p should be not be nil since 9 is duplicate")
	}
	result := bst.InorderTraversal()
	expected := []int{1, 2, 3, 4, 7, 9, 18, 18, 52, 63, 90}
	if ListEq(result, expected) == false {
		t.Fatalf(`Delete() does not work in delete duplicate node`)
	}
}
