// Package maps provides some functions for working with Go maps.
package maps

// Clone clones a map.
func Clone[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// Update updates a map with items from another map.
func Update[K comparable, V any](m1, m2 map[K]V) {
	if m1 == nil {
		panic("cannot update nil map")
	}
	for k, v := range m2 {
		m1[k] = v
	}
}

// Clear removes all items from a map.
func Clear[K comparable, V any](m map[K]V) {
	for k := range m {
		delete(m, k)
	}
}

// Contains returns true if key is in map m.
func Contains[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// Get returns the value for key from map m or a default value.
func Get[K comparable, V any](m map[K]V, key K, dflt V) V {
	v, ok := m[key]
	if ok {
		return v
	}
	return dflt
}

// Keys returns a slice with all keys from map m.
func Keys[K comparable, V any](m map[K]V) []K {
	if m == nil {
		return nil
	}
	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// Values returns a slice with all values from map m.
func Values[K comparable, V any](m map[K]V) []V {
	if m == nil {
		return nil
	}
	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// Equal returns true if the two maps are equal (containing the same keys with the same values).
func Equal[K, V comparable](m1, m2 map[K]V) bool {
	return EqualFunc(m1, m2, func(v1, v2 V) bool {
		return v1 == v2
	})
}

// EqualFunc returns true if the two maps are equal using a function to compare values.
func EqualFunc[K comparable, V any](m1, m2 map[K]V, equal func(v1, v2 V) bool) bool {
	if len(m1) != len(m2) {
		return false
	}
	if m1 == nil && m2 != nil || m1 != nil && m2 == nil {
		return false
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok || !equal(v1, v2) {
			return false
		}
	}
	return true
}

// Type Item represents a key-value pair.
type Item[K comparable, V any] struct {
	Key   K
	Value V
}

// Items returns a slice of [Item] objects for the given map.
func Items[K comparable, V any](m map[K]V) []Item[K, V] {
	if m == nil {
		return nil
	}
	result := make([]Item[K, V], 0, len(m))
	for k, v := range m {
		result = append(result, Item[K, V]{k, v})
	}
	return result
}

// FromItems makes a map from a slice of [Item] objects.
func FromItems[K comparable, V any](items []Item[K, V]) map[K]V {
	if items == nil {
		return nil
	}
	m := make(map[K]V, len(items))
	for _, item := range items {
		m[item.Key] = item.Value
	}
	return m
}

// FromSlices makes a map from two slices.
// If the keys slice is longer then the values slice, the surplus keys will have
// the zero value of type V as their value. Surplus values will be ignored.
func FromSlices[K comparable, V any](keys []K, values []V) map[K]V {
	if keys == nil {
		return nil
	}
	m := make(map[K]V, len(keys))
	for i, k := range keys {
		if i < len(values) {
			m[k] = values[i]
		} else {
			var v V
			m[k] = v
		}
	}
	return m
}

// FromFuncs makes a map from the return values of two functions (e.g. from math.random).
// Panics if the keys functions is nil or size is negative.
// If the values function is nil, the zero value of type V will be used for all values.
func FromFuncs[K comparable, V any](size int, keys func() K, values func() V) map[K]V {
	if size < 0 {
		panic("size must be >= 0")
	}
	if keys == nil {
		panic("nil functions")
	}
	m := make(map[K]V, size)
	for i := 0; i < size; i++ {
		var v V
		if values != nil {
			v = values()
		}
		m[keys()] = v
	}
	return m
}

// KeysForValue returns a slice with keys which have the given value.
func KeysForValue[K, V comparable](m map[K]V, value V) []K {
	return KeysForValueFunc(m, value, func(v1 V, v2 V) bool {
		return v1 == v2
	})
}

// KeysForValueFunc returns a slice with keys which have the given value using
// a function to compare values.
func KeysForValueFunc[K comparable, V any](m map[K]V, value V, equal func(v1, v2 V) bool) []K {
	if m == nil {
		return nil
	}
	keys := []K{}
	for k, v := range m {
		if equal(v, value) {
			keys = append(keys, k)
		}
	}
	return keys
}

// Delete deletes all items from m for which fn returns true and
// returns the number of deleted items.
func Delete[K comparable, V any](m map[K]V, fn func(k K, v V) bool) int {
	cnt := 0
	for k, v := range m {
		if fn(k, v) {
			delete(m, k)
			cnt++
		}
	}
	return cnt
}
