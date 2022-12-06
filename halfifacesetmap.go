package setmap

// HalfIfaceSetMap is map<K, set<V>>
type HalfIfaceSetMap[K, V comparable] struct {
	setmap IfaceSetMap[K, V]
}

func NewHalfIfaceSetMap[K, V comparable]() HalfIfaceSetMap[K, V] {
	return HalfIfaceSetMap[K, V]{
		setmap: NewIfaceSetMap[K, V](),
	}
}

// Add a value to the set whose ID is key
func (m HalfIfaceSetMap[K, V]) Add(key K, value IfaceSetMapItem[V]) {
	m.setmap.Add(SimpleIfaceSetMapItem[K]{T: key}, value)
}

// Remove a value from the set whose ID is key
func (m HalfIfaceSetMap[K, V]) Remove(key K, value IfaceSetMapItem[V]) {
	m.setmap.Remove(SimpleIfaceSetMapItem[K]{T: key}, value)
}

// Exists show if a value is in the set whose ID is key
func (m HalfIfaceSetMap[K, V]) Exists(key K, value IfaceSetMapItem[V]) bool {
	return m.setmap.Exists(SimpleIfaceSetMapItem[K]{T: key}, value)
}

// GetSet get all the values from the set whose ID is key
func (m HalfIfaceSetMap[K, V]) GetSet(key K) []V {
	return m.setmap.GetSet(SimpleIfaceSetMapItem[K]{T: key})
}

// HalfIfaceSetMapaMteS consists of a IfaceSetMap map<K, set<V>> and a reverse IfaceSetMap map<V, set<K>>
// so it can show K和V之间的多对多关系
type HalfIfaceSetMapaMteS[K, V comparable] struct {
	setmapamtes IfaceSetMapaMteS[K, V]
}

func NewHalfIfaceSetMapaMteS[K, V comparable]() HalfIfaceSetMapaMteS[K, V] {
	return HalfIfaceSetMapaMteS[K, V]{
		setmapamtes: NewIfaceSetMapaMteS[K, V](),
	}
}

func (m HalfIfaceSetMapaMteS[K, V]) Add(key K, value IfaceSetMapItem[V]) {
	m.setmapamtes.Add(SimpleIfaceSetMapItem[K]{T: key}, value)
}

func (m HalfIfaceSetMapaMteS[K, V]) Remove(key K, value IfaceSetMapItem[V]) {
	m.setmapamtes.Remove(SimpleIfaceSetMapItem[K]{T: key}, value)
}

func (m HalfIfaceSetMapaMteS[K, V]) RemoveValue(value IfaceSetMapItem[V]) {
	m.setmapamtes.RemoveValue(value)
}

func (m HalfIfaceSetMapaMteS[K, V]) RemoveKey(key K) {
	m.setmapamtes.RemoveKey(SimpleIfaceSetMapItem[K]{T: key})
}

func (m HalfIfaceSetMapaMteS[K, V]) GetKeys(value IfaceSetMapItem[V]) []K {
	return m.setmapamtes.GetKeys(value)
}

// GetUniqueValues only get the key's values that is not exists in other key's set
func (m HalfIfaceSetMapaMteS[K, V]) GetUniqueValues(key K) []V {
	return m.setmapamtes.GetUniqueValues(SimpleIfaceSetMapItem[K]{T: key})
}

// GetUniqueKeys only get the value's keys that is not other value's key
func (m HalfIfaceSetMapaMteS[K, V]) GetUniqueKeys(value IfaceSetMapItem[V]) []K {
	return m.setmapamtes.GetUniqueKeys(value)
}
