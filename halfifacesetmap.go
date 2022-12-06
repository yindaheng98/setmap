package setmap

// HalfIfaceSetMap is map<K, set<V>>
type HalfIfaceSetMap[K, V comparable] struct {
	m map[K]map[V]IfaceSetMapItem[V]
}

func NewHalfIfaceSetMap[K, V comparable]() HalfIfaceSetMap[K, V] {
	return HalfIfaceSetMap[K, V]{
		m: make(map[K]map[V]IfaceSetMapItem[V]),
	}
}

// Add a value to the set whose ID is key
func (m HalfIfaceSetMap[K, V]) Add(key K, value IfaceSetMapItem[V]) {
	if set, ok := m.m[key]; ok {
		set[value.ID()] = value
	} else {
		m.m[key] = map[V]IfaceSetMapItem[V]{value.ID(): value}
	}
}

// Remove a value from the set whose ID is key
func (m HalfIfaceSetMap[K, V]) Remove(key K, value IfaceSetMapItem[V]) {
	if set, ok := m.m[key]; ok {
		delete(set, value.ID())
		if len(set) <= 0 {
			delete(m.m, key)
		}
	}
}

// Exists show if a value is in the set whose ID is key
func (m HalfIfaceSetMap[K, V]) Exists(key K, value IfaceSetMapItem[V]) bool {
	var set map[V]IfaceSetMapItem[V]
	var ok bool
	if set, ok = m.m[key]; ok {
		_, ok = set[value.ID()]
	}
	return ok
}

// GetSet get all the values from the set whose ID is key
func (m HalfIfaceSetMap[K, V]) GetSet(key K) []V {
	var r []V
	if set, ok := m.m[key]; ok {
		r = make([]V, len(set))
		i := 0
		for value := range set {
			r[i] = value
			i++
		}
	}
	return r
}
