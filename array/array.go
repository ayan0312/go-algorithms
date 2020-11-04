package array

import (
	"errors"
)

const DEFAULT_CAPACITY = 10

type DynamicArray struct {
	size     int
	capacity int
	elements []interface{}
}

func New() *DynamicArray {
	da := new(DynamicArray)
	da.Resize()
	return da
}

func (da *DynamicArray) Change(element interface{}, index int) error {
	err := da.checkRangeFromIndex(index)

	if err != nil {
		return err
	}

	da.elements[index] = element

	return nil
}

func (da *DynamicArray) Add(element interface{}) {
	if da.size == da.capacity {
		da.Resize()
	}

	da.elements[da.size] = element
	da.size++
}

func (da *DynamicArray) Insert(element interface{}, index int) error {
	err := da.checkRangeFromIndex(index)

	if err != nil {
		return err
	}

	if da.size == da.capacity {
		da.Resize()
	}

	for i := da.size - 1; i >= index; i-- {
		da.elements[i+1] = da.elements[i]
	}

	da.elements[index] = element
	da.size++
	return nil
}

func (da *DynamicArray) Insert2(element interface{}, index int) error {
	err := da.checkRangeFromIndex(index)

	if err != nil {
		return err
	}

	if da.size == da.capacity {
		da.Resize()
	}

	copy(da.elements[index:], da.elements[index-1:da.size])
	da.elements[index] = element
	da.size++
	return nil
}

func (da *DynamicArray) Remove(index int) error {
	err := da.checkRangeFromIndex(index)

	if err != nil {
		return err
	}

	copy(da.elements[index:], da.elements[index+1:da.size])
	da.elements[da.size-1] = nil
	da.size--
	return nil
}

func (da *DynamicArray) Element(index int) (interface{}, error) {
	err := da.checkRangeFromIndex(index)

	if err != nil {
		return nil, err
	}

	return da.elements[index], nil
}

func (da *DynamicArray) IsEmpty() bool {
	return da.size == 0
}

func (da *DynamicArray) Data() []interface{} {
	return da.elements[:da.size]
}

func (da *DynamicArray) Resize() {
	if da.capacity == 0 {
		da.capacity = DEFAULT_CAPACITY
	} else {
		da.capacity = da.capacity << 1
	}

	newElements := make([]interface{}, da.capacity)

	copy(newElements, da.elements)

	da.elements = newElements
}

func (da *DynamicArray) checkRangeFromIndex(index int) error {
	if index >= da.size || index < 0 {
		return errors.New("index out of range")
	}
	return nil
}
