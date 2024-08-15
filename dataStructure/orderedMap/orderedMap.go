package orderedMap

import (
	"container/list"
)

// OrderedMap represents a map with keys that are ordered by insertion order.
type OrderedMap[K comparable, V any] struct {
	m    map[K]*list.Element
	list *list.List
}

// entry represents a key-value pair stored in the OrderedMap.
type entry[K comparable, V any] struct {
	key   K
	value V
}

// NewOrderedMap creates a new OrderedMap.
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		m:    make(map[K]*list.Element),
		list: list.New(),
	}
}

// Set sets the value for the given key in the OrderedMap.
func (om *OrderedMap[K, V]) Set(key K, value V) {
	if el, exists := om.m[key]; exists {
		el.Value.(*entry[K, V]).value = value
	} else {
		el := om.list.PushBack(&entry[K, V]{key, value})
		om.m[key] = el
	}
}

// Get retrieves the value for the given key from the OrderedMap.
func (om *OrderedMap[K, V]) Get(key K) (V, bool) {
	if el, exists := om.m[key]; exists {
		return el.Value.(*entry[K, V]).value, true
	}
	var zero V
	return zero, false
}

// Delete removes the key-value pair from the OrderedMap.
func (om *OrderedMap[K, V]) Delete(key K) {
	if el, exists := om.m[key]; exists {
		om.list.Remove(el)
		delete(om.m, key)
	}
}

// Keys returns all the keys in insertion order.
func (om *OrderedMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(om.m))
	for el := om.list.Front(); el != nil; el = el.Next() {
		keys = append(keys, el.Value.(*entry[K, V]).key)
	}
	return keys
}

// Values returns all the values in insertion order.
func (om *OrderedMap[K, V]) Values() []V {
	values := make([]V, 0, len(om.m))
	for el := om.list.Front(); el != nil; el = el.Next() {
		values = append(values, el.Value.(*entry[K, V]).value)
	}
	return values
}

// At returns the key-value pair at the given index.
func (om *OrderedMap[K, V]) At(index int) (K, V, bool) {
	if index < 0 || index >= len(om.m) {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	el := om.list.Front()
	for i := 0; i < index; i++ {
		el = el.Next()
	}
	if el != nil {
		entry := el.Value.(*entry[K, V])
		return entry.key, entry.value, true
	}
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}
