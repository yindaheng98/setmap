package setmap

// HalfIfaceSetMap is map<K, set<V>>
type HalfIfaceSetMap[K, V comparable] struct {
	IfaceSetMap[K, V]
}

func NewHalfIfaceSetMap[K, V comparable]() HalfIfaceSetMap[K, V] {
	return HalfIfaceSetMap[K, V]{
		IfaceSetMap: NewIfaceSetMap[K, V](),
	}
}

// Add a value to the set whose ID is key
func (m HalfIfaceSetMap[K, V]) Add(key K, value IfaceSetMapItem[V]) {
	m.IfaceSetMap.Add(SimpleIfaceSetMapItem[K]{T: key}, value)
}

// Remove a value from the set whose ID is key
func (m HalfIfaceSetMap[K, V]) Remove(key K, value IfaceSetMapItem[V]) {
	m.IfaceSetMap.Remove(SimpleIfaceSetMapItem[K]{T: key}, value)
}

// Exists show if a value is in the set whose ID is key
func (m HalfIfaceSetMap[K, V]) Exists(key K, value IfaceSetMapItem[V]) bool {
	return m.IfaceSetMap.Exists(SimpleIfaceSetMapItem[K]{T: key}, value)
}

// GetSet get all the values from the set whose ID is key
func (m HalfIfaceSetMap[K, V]) GetSet(key K) []V {
	return m.IfaceSetMap.GetSet(SimpleIfaceSetMapItem[K]{T: key})
}

// HalfIfaceSetMapaMteS consists of a IfaceSetMap map<K, set<V>> and a reverse IfaceSetMap map<V, set<K>>
// so it can show K和V之间的多对多关系
type HalfIfaceSetMapaMteS[K, V comparable] struct {
	IfaceSetMapaMteS[K, V]
}

func NewHalfIfaceSetMapaMteS[K, V comparable]() HalfIfaceSetMapaMteS[K, V] {
	return HalfIfaceSetMapaMteS[K, V]{
		IfaceSetMapaMteS: NewIfaceSetMapaMteS[K, V](),
	}
}

func (m HalfIfaceSetMapaMteS[K, V]) Add(key K, value IfaceSetMapItem[V]) {
	m.IfaceSetMapaMteS.Add(SimpleIfaceSetMapItem[K]{T: key}, value)
}

func (m HalfIfaceSetMapaMteS[K, V]) Remove(key K, value IfaceSetMapItem[V]) {
	m.IfaceSetMapaMteS.Remove(SimpleIfaceSetMapItem[K]{T: key}, value)
}

func (m HalfIfaceSetMapaMteS[K, V]) RemoveValue(value IfaceSetMapItem[V]) {
	m.IfaceSetMapaMteS.RemoveValue(value)
}

func (m HalfIfaceSetMapaMteS[K, V]) RemoveKey(key K) {
	m.IfaceSetMapaMteS.RemoveKey(SimpleIfaceSetMapItem[K]{T: key})
}

func (m HalfIfaceSetMapaMteS[K, V]) GetKeys(value IfaceSetMapItem[V]) []K {
	return m.IfaceSetMapaMteS.GetKeys(value)
}

// GetUniqueValues only get the key's values that is not exists in other key's set
func (m HalfIfaceSetMapaMteS[K, V]) GetUniqueValues(key K) []V {
	return m.IfaceSetMapaMteS.GetUniqueValues(SimpleIfaceSetMapItem[K]{T: key})
}

// GetUniqueKeys only get the value's keys that is not other value's key
func (m HalfIfaceSetMapaMteS[K, V]) GetUniqueKeys(value IfaceSetMapItem[V]) []K {
	return m.IfaceSetMapaMteS.GetUniqueKeys(value)
}

// Exists show if a value is in the set whose ID is key
func (m HalfIfaceSetMapaMteS[K, V]) Exists(key K, value IfaceSetMapItem[V]) bool {
	return m.IfaceSetMapaMteS.Exists(SimpleIfaceSetMapItem[K]{T: key}, value)
}

// GetSet get all the values from the set whose ID is key
func (m HalfIfaceSetMapaMteS[K, V]) GetSet(key K) []V {
	return m.IfaceSetMapaMteS.GetSet(SimpleIfaceSetMapItem[K]{T: key})
}
