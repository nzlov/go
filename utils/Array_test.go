package utils

import (
	"testing"
)

func TestNewArray(t *testing.T) {
	array := NewArray()
	if array.Size() != 0 {
		t.Error("NewArray size ", array.Size(), "<> 0")
	}
}

func TestArrayByCap(t *testing.T) {
	array := NewArrayByCap(10)
	if array.Size() != 0 {
		t.Error("NewArray size ", array.Size(), "<> 0")
	}
}

func TestNewArrayByArray(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	array := NewArrayByArray(tA)
	if array.Size() != tA.Size() {
		t.Error("NewArray size ", array.Size(), "<>", tA.Size())
	}
}

func TestNewArrayByIs(t *testing.T) {
	array := NewArrayByIs(1, 2, 3)
	if array.Size() != 3 {
		t.Error("NewArray size ", array.Size(), "<>", 3)
	}
}

func TestNewArrayByArrayE(t *testing.T) {
	a := []interface{}{1, 2, 3}
	array := NewArrayByArrayE(a, 0, len(a))
	if array.Size() != len(a) {
		t.Error("NewArray size ", array.Size(), "<>", len(a))
	}
}

func TestAdd(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	if 2 != tA.Size() {
		t.Error("NewArray size ", 2, "<>", tA.Size())
	}
}

func TestAddAllArray(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	tB := NewArray()
	tB.AddAllArray(tA)
	if tB.Size() != tA.Size() {
		t.Error("NewArray size ", tB.Size(), "<>", tA.Size())
	}
}

func TestAddAllIs(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	a := []interface{}{1, 2, 3}
	tA.AddAllIs(a)
	if len(a)+2 != tA.Size() {
		t.Error("NewArray size ", len(a)+2, "<>", tA.Size())
	}
}

func TestGet(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	if v, _ := tA.Get(0); v != 1 {
		t.Error("NewArray get(0)  ", v, "<>", 1)
	}
}

func TestSet(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	tA.Set(1, 1)
	if v, _ := tA.Get(1); v != 1 {
		t.Error("NewArray set(1,1) ", v, "<>", 1)
	}
}

func TestInsert(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	tA.Insert(1, 1)
	if v, _ := tA.Get(0); v != 1 {
		t.Error("NewArray Add ", 1, "<>", v)
	}
	if v, _ := tA.Get(1); v != 1 {
		t.Error("NewArray Insert ", 1, "<>", v)
	}
	if v, _ := tA.Get(2); v != 2 {
		t.Error("NewArray Add ", 2, "<>", v)
	}
}

func TestSwap(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	tA.Swap(0, 1)
	if v, _ := tA.Get(0); v != 2 {
		t.Error("NewArray TestSwap ", v, "<>", 1)
	}
	if v, _ := tA.Get(1); v != 1 {
		t.Error("NewArray TestSwap ", v, "<>", 2)
	}
}

func TestContains(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	if !tA.Contains(1) {
		t.Error("NewArray Contains(1) false")
	}
}

func TestIndexOf(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	if tA.IndexOf(2) != 1 {
		t.Error("NewArray IndexOf(2) not index :1")
	}
}

func TestLastIndexOf(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	if tA.LastIndexOf(1) != 0 {
		t.Error("NewArray LastIndexOf(1) not index :0")
	}
}

func TestRemoveIndex(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	if v, _ := tA.RemoveIndex(1); v != 2 {
		t.Error("NewArray RemoveIndex(1) not value :2")
	}
	if tA.Size() != 1 {
		t.Error("NewArray RemoveIndex(1) size <> 1")
	}
}

func TestRemoveValue(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	if !tA.RemoveValue(1) {
		t.Error("NewArray RemoveValue(1) false")
	}
	if tA.Items[0] != 2 {
		t.Error("NewArray RemoveValue(1) tA.Items[0] != 2")
	}
}

func TestRemoveAll(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	tB := NewArray()
	tB.Add(1)
	tB.Add(2)
	tB.Add(3)
	tB.RemoveAll(tA)
	if tB.Items[0] != 3 {
		t.Error("NewArray RemoveAll tB.Items[0] != 3")
	}
}

func TestClear(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	tA.Clear()
	if tA.Size() != 0 {
		t.Error("NewArray Clear tA.Size()!=0")
	}
}

func TestPop(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	if v, _ := tA.Pop(); v != 2 {
		t.Error("NewArray Pop() ", v, " != 2")
	}
	if tA.Size() != 1 {
		t.Error("NewArray Pop() tA.Size() != 1")
	}
}

func TestPeek(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	if v, _ := tA.Peek(); v != 2 {
		t.Error("NewArray Pop() ", v, " != 2")
	}
	if tA.Size() != 2 {
		t.Error("NewArray Pop() tA.Size() != 2")
	}
}

func TestFirst(t *testing.T) {
	tA := NewArray()
	tA.Add(1)
	tA.Add(2)
	if v, _ := tA.First(); v != 1 {
		t.Error("NewArray Pop() ", v, " != 1")
	}
	if tA.Size() != 2 {
		t.Error("NewArray Pop() tA.Size() != 2")
	}
}
