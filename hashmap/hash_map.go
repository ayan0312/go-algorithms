package hashmap

import (
	"fmt"
	"hash/fnv"
)

const DEFAULT_CAPACITY uint64 = 1024

type Node struct {
	key   interface{}
	value interface{}
	next  *Node
}

type HashMap struct {
	capacity uint64
	size     uint64
	table    []*Node
}

func New() *HashMap {
	return &HashMap{
		capacity: DEFAULT_CAPACITY,
		table:    make([]*Node, DEFAULT_CAPACITY),
	}
}

func (hm *HashMap) Get(key interface{}) interface{} {
	node := hm.getNodeByHash(hm.hash(key))

	if node != nil {
		return node.value
	}

	return nil
}

func (hm *HashMap) Put(key interface{}, value interface{}) interface{} {
	return hm.putValue(hm.hash(key), key, value)
}

func (hm *HashMap) Contains(key interface{}) bool {
	node := hm.getNodeByHash(hm.hash(key))
	return node != nil
}

func (hm *HashMap) putValue(hash uint64, key interface{}, value interface{}) interface{} {
	if hm.capacity == 0 {
		hm.capacity = DEFAULT_CAPACITY
		hm.table = make([]*Node, DEFAULT_CAPACITY)
	}

	node := hm.getNodeByHash(hash)

	if node == nil {
		hm.table[hash] = newNode(key, value)
	} else if node.key == key {
		hm.table[hash] = newNodeWithNext(key, value, node)
		return value
	} else {
		hm.resize()
		return hm.putValue(hash, key, value)
	}

	hm.size++
	return value
}

func (hm *HashMap) getNodeByHash(hash uint64) *Node {
	return hm.table[hash]
}

func (hm *HashMap) resize() {
	hm.capacity <<= 1

	tempTable := hm.table

	hm.table = make([]*Node, hm.capacity)

	for i := 0; i < len(tempTable); i++ {
		node := tempTable[i]
		if node == nil {
			continue
		}

		hm.table[hm.hash(node.key)] = node
	}
}

func newNode(key interface{}, value interface{}) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

func newNodeWithNext(key interface{}, value interface{}, next *Node) *Node {
	return &Node{
		key:   key,
		value: value,
		next:  next,
	}
}

func (hm *HashMap) hash(key interface{}) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))

	hashValue := h.Sum64()

	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}
