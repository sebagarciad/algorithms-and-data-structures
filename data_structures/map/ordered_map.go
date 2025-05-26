package mymap

type BSTMap[K comparable, V any] interface {
	Map[K, V]

	// IterateRange iterates only including the elements that are within the indicated range,
	// including them if they are found
	IterateRange(from *K, to *K, visit func(key K, value V) bool)

	// IteratorRange creates an IterMap that only iterates over the keys that are within the indicated range
	IteratorRange(from *K, to *K) MapIterator[K, V]
}
