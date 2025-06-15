package maps

import "iter"

// Iter returns an iterator over key-value pairs from m.
func Iter[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func IterKeys[K comparable, V any](m map[K]V) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range m {
			if !yield(k) {
				return
			}
		}
	}
}

// IterValues returns an iterator over the values from m.
func IterValues[K comparable, V any](m map[K]V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range m {
			if !yield(v) {
				return
			}
		}
	}
}

// Collect returns a new map with the key-value pairs from seq.
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	result := make(map[K]V)
	for k, v := range seq {
		result[k] = v
	}
	return result
}
