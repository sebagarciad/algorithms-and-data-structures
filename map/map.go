package mymap

type Map[K comparable, V any] interface {
	// Save stores the key-value pair in the Map. If the key is already present in the Map,
	// the associated value is updated
	Save(key K, value V)

	// Contains returns true if a key is already present in the Map. Otherwise, returns false
	Contains(key K) bool

	// Get returns the value associated with a key. If the key does not belong to the Map, it panics
	// with the message 'The key does not belong to the map'
	Get(key K) V

	// Remove removes the given key from the Map and returns the associated value. If the key does not
	// belong to the Map, it panics with the message 'The key does not belong to the map'
	Remove(key K) V

	// Count returns the number of elements in the Map
	Count() int

	// Iterate iterates over the Map internally, applying the function passed as a parameter to each element
	Iterate(func(key K, value V) bool)

	// Iterator returns an IterMap to iterate over the Map
	Iterator() MapIterator[K, V]
}

type MapIterator[K comparable, V any] interface {
	// HasNext returns true if there are more elements to see, that is, if the iterator is at an element.
	// Otherwise, returns false
	HasNext() bool

	// Current returns the key and value of the element where the iterator is. If not HasNext,
	// it panics with the message 'The iterator has finished iterating'
	Current() (K, V)

	// Next advances to the next element of the Map. If not HasNext, it panics with the message
	// 'The iterator has finished iterating'
	Next()
}
