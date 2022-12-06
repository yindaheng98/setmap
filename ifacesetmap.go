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

// IfaceSetMapaMteS consists of a IfaceSetMap map<K, set<V>> and a reverse IfaceSetMap map<V, set<K>>
// so it can show K和V之间的多对多关系
type IfaceSetMapaMteS[K, V comparable] struct {
	IfaceSetMap[K, V]
	reverse IfaceSetMap[V, K]
}

func NewIfaceSetMapaMteS[K, V comparable]() IfaceSetMapaMteS[K, V] {
	return IfaceSetMapaMteS[K, V]{
		IfaceSetMap: NewIfaceSetMap[K, V](),
		reverse:     NewIfaceSetMap[V, K](),
	}
}

func (m IfaceSetMapaMteS[K, V]) Add(key IfaceSetMapItem[K], value IfaceSetMapItem[V]) {
	m.IfaceSetMap.Add(key, value)
	m.reverse.Add(value, key)
}

func (m IfaceSetMapaMteS[K, V]) Remove(key IfaceSetMapItem[K], value IfaceSetMapItem[V]) {
	m.IfaceSetMap.Remove(key, value)
	m.reverse.Remove(value, key)
}

func (m IfaceSetMapaMteS[K, V]) RemoveValue(value IfaceSetMapItem[V]) {
	if set, ok := m.reverse.m[value.ID()]; ok {
		for _, key := range set {
			m.IfaceSetMap.Remove(key, value)
		}
		delete(m.reverse.m, value.ID())
	}
}

func (m IfaceSetMapaMteS[K, V]) RemoveKey(key IfaceSetMapItem[K]) {
	if set, ok := m.m[key.ID()]; ok {
		for _, value := range set {
			m.reverse.Remove(value, key)
		}
		delete(m.m, key.ID())
	}
}

func (m IfaceSetMapaMteS[K, V]) GetKeys(value IfaceSetMapItem[V]) []K {
	return m.reverse.GetSet(value)
}

// GetUniqueValues only get the key's values that is not exists in other key's set
func (m IfaceSetMapaMteS[K, V]) GetUniqueValues(key IfaceSetMapItem[K]) []V {
	var r []V
	if set, ok := m.m[key.ID()]; ok {
		for _, value := range set {
			if len(m.reverse.GetSet(value)) <= 1 { // only have this value?
				r = append(r, value.ID()) // that's what I find
			}
		}
	}
	return r
}

// GetUniqueKeys only get the value's keys that is not other value's key
func (m IfaceSetMapaMteS[K, V]) GetUniqueKeys(value IfaceSetMapItem[V]) []K {
	var r []K
	if set, ok := m.reverse.m[value.ID()]; ok {
		for _, key := range set {
			if len(m.GetSet(key)) <= 1 { // only have this key?
				r = append(r, key.ID()) // that's what I find
			}
		}
	}
	return r
}
