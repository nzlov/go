package utils

import (
	"errors"
)

type Array struct {
	Items []interface{}
}

func NewArray() *Array {
	return NewArrayByCap(16)
}

func NewArrayByCap(cap int) *Array {
	a := &Array{}
	a.Items = make([]interface{}, 0, cap)
	return a
}

func NewArrayByArray(array *Array) *Array {
	a := &Array{}
	a.Items = make([]interface{}, array.Size())
	for i := 0; i < array.Size(); i++ {
		a.Items[i] = array.Items[i]
	}
	return a
}

func NewArrayByIs(is ...interface{}) *Array {
	a := &Array{}
	a.Items = make([]interface{}, len(is))
	for i := 0; i < len(is); i++ {
		a.Items[i] = is[i]
	}
	return a
}
func NewArrayByArrayE(array []interface{}, start, count int) *Array {
	a := &Array{}
	a.Items = make([]interface{}, count)
	for i := 0; i < count; i++ {
		a.Items[i] = array[i+start]
	}
	return a
}

func (this *Array) Add(v interface{}) {
	this.Items = append(this.Items, v)

}

func (this *Array) AddAllArray(array *Array) {
	this.AddAllArrayByPos(array, 0, array.Size())
}

func (this *Array) AddAllArrayByPos(array *Array, offset, length int) error {
	if offset+length > array.Size() {
		return errors.New("offset + length must be <= size: " + string(offset) + " + " + string(length) + " <= " + string(array.Size()))
	}
	this.AddAllIsByPos(array.Items, offset, length)
	return nil
}

func (this *Array) AddAllIs(is []interface{}) {
	this.AddAllIsByPos(is, 0, len(is))
}

func (this *Array) AddAllIsByPos(is []interface{}, offset, length int) {
	for i := 0; i < length; i++ {
		this.Items = append(this.Items, is[i+offset])
	}
}

func (this *Array) Size() int {
	return len(this.Items)
}

func (this *Array) Get(index int) (interface{}, error) {
	if this.Size() <= index {
		return nil, errors.New("IndexOutOfBounds:" + string(this.Size()) + "<=" + string(index))
	}
	return this.Items[index], nil
}

func (this *Array) Set(index int, v interface{}) error {
	if this.Size() <= index {
		return errors.New("IndexOutOfBounds:" + string(this.Size()) + "<=" + string(index))
	}
	this.Items[index] = v
	return nil
}

func (this *Array) Insert(index int, v interface{}) error {
	if this.Size() <= index {
		return errors.New("IndexOutOfBounds:" + string(this.Size()) + "<=" + string(index))
	}
	oldA := make([]interface{}, this.Size())
	copy(oldA, this.Items)
	this.Items = make([]interface{}, this.Size()+1)
	for i := 0; i < index; i++ {
		this.Items[i] = oldA[i]
	}
	this.Items[index] = v
	for i := index + 1; i < len(oldA)+1; i++ {
		this.Items[i] = oldA[i-1]
	}
	return nil
}

func (this *Array) Swap(first, second int) error {
	if first >= this.Size() {
		return errors.New("IndexOutOfBounds:" + string(this.Size()) + "<=" + string(first))
	}
	if second >= this.Size() {
		return errors.New("IndexOutOfBounds:" + string(this.Size()) + "<=" + string(second))
	}
	a := this.Items[first]
	this.Items[first] = this.Items[second]
	this.Items[second] = a
	return nil
}

func (this *Array) Contains(value interface{}) bool {
	i := this.Size() - 1
	for ; i >= 0; i-- {
		if value == this.Items[i] {
			return true
		}
	}
	return false
}

func (this *Array) IndexOf(value interface{}) int {
	for i := 0; i < this.Size(); i++ {
		if this.Items[i] == value {
			return i
		}
	}
	return -1
}

func (this *Array) LastIndexOf(value interface{}) int {

	for i := this.Size() - 1; i >= 0; i-- {
		if this.Items[i] == value {
			return i
		}
	}
	return -1
}

func (this *Array) RemoveIndex(index int) (interface{}, error) {
	if index >= this.Size() {
		return nil, errors.New("IndexOutOfBounds:" + string(this.Size()) + "<=" + string(index))
	}
	v := this.Items[index]
	oldA := make([]interface{}, this.Size())
	copy(oldA, this.Items)
	this.Items = make([]interface{}, this.Size()-1)
	for i := 0; i < index; i++ {
		this.Items[i] = oldA[i]
	}
	for i := index + 1; i < len(oldA); i++ {
		this.Items[i-1] = oldA[i]
	}
	return v, nil
}

func (this *Array) RemoveValue(value interface{}) bool {
	for i, n := 0, this.Size(); i < n; i++ {
		if this.Items[i] == value {
			this.RemoveIndex(i)
			return true
		}
	}
	return false
}

func (this *Array) RemoveAll(array *Array) {
	for i, n := 0, array.Size(); i < n; i++ {
		for j, m := 0, this.Size(); j < m; j++ {
			if this.Items[j] == array.Items[i] {
				this.RemoveIndex(j)
				m--
			}
		}
	}
}

func (this *Array) Clear() {
	this.Items = make([]interface{}, 0, 10)
}

func (this *Array) Pop() (interface{}, error) {
	if this.Size() == 0 {
		return nil, errors.New("Array is empty.")
	}
	return this.RemoveIndex(this.Size() - 1)
}

func (this *Array) Peek() (interface{}, error) {
	if this.Size() == 0 {
		return nil, errors.New("Array is empty.")
	}
	return this.Items[this.Size()-1], nil
}

func (this *Array) First() (interface{}, error) {
	if this.Size() == 0 {
		return nil, errors.New("Array is empty.")
	}
	return this.Items[0], nil
}
