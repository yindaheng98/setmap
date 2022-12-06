package setmap

type IfaceSetMapItem[T comparable] interface {
	ID() T
}

type SimpleIfaceSetMapItem[T comparable] struct {
	T T
}

func (s SimpleIfaceSetMapItem[T]) ID() T {
	return s.T
}

// IfaceSetMap is map<K, set<V>>
type IfaceSetMap[K, V comparable] struct {
	m map[K]map[V]IfaceSetMapItem[V]
}

func NewIfaceSetMap[K, V comparable]() IfaceSetMap[K, V] {
	return IfaceSetMap[K, V]{
		m: make(map[K]map[V]IfaceSetMapItem[V]),
	}
}

// Add a value to the set whose ID is key
func (m IfaceSetMap[K, V]) Add(key IfaceSetMapItem[K], value IfaceSetMapItem[V]) {
	if set, ok := m.m[key.ID()]; ok {
		set[value.ID()] = value
	} else {
		m.m[key.ID()] = map[V]IfaceSetMapItem[V]{value.ID(): value}
	}
}

// Remove a value from the set whose ID is key
func (m IfaceSetMap[K, V]) Remove(key IfaceSetMapItem[K], value IfaceSetMapItem[V]) {
	if set, ok := m.m[key.ID()]; ok {
		delete(set, value.ID())
		if len(set) <= 0 {
			delete(m.m, key.ID())
		}
	}
}

// Exists show if a value is in the set whose ID is key
func (m IfaceSetMap[K, V]) Exists(key IfaceSetMapItem[K], value IfaceSetMapItem[V]) bool {
	var set map[V]IfaceSetMapItem[V]
	var ok bool
	if set, ok = m.m[key.ID()]; ok {
		_, ok = set[value.ID()]
	}
	return ok
}

// GetSet get all the values from the set whose ID is key
func (m IfaceSetMap[K, V]) GetSet(key IfaceSetMapItem[K]) []V {
	var r []V
	if set, ok := m.m[key.ID()]; ok {
		r = make([]V, len(set))
		i := 0
		for value := range set {
			r[i] = value
			i++
		}
	}
	return r
}
