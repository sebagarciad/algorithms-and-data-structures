package mymap

import (
	"fmt"
	"hash/crc32"
)

const (
	_PANIC_HASH      = "The key does not belong to the map"
	_PANIC_ITERATOR  = "The iterator has finished iterating"
	_INITIAL_SIZE    = 17
	_LOAD_FACTOR_INC = 0.7
	_LOAD_FACTOR_DEC = _LOAD_FACTOR_INC / 4
	_RESIZE_FACTOR   = 2
)

// ===================== Types ======================

type state int

const (
	EMPTY state = iota
	OCCUPIED
	DELETED
)

type hashCell[K comparable, V any] struct {
	key   K
	value V
	state state
}

type closedHash[K comparable, V any] struct {
	table   []hashCell[K, V]
	count   int
	size    int
	deleted int
}

type closedHashIterator[K comparable, V any] struct {
	hash  *closedHash[K, V]
	index int
}

// =================== Hash Function ===================

func toBytes[K comparable](key K) []byte {
	return []byte(fmt.Sprintf("%v", key))
}

// Hash function: CRC-32 ChecksumIEEE.
// URL: https://cs.opensource.google/go/go/+/refs/tags/go1.23.2:src/hash/crc32/crc32.go;l=236

func fnvHashing64[K comparable, V any](bytes []byte, hash *closedHash[K, V]) int32 {
	h := crc32.ChecksumIEEE(bytes)
	return int32(h % uint32(hash.size))
}

// =============== Hash Auxiliaries ==================

func (hash *closedHash[K, V]) getKeyHash(key K) int {
	return int(fnvHashing64(toBytes(key), hash))
}

func (hash *closedHash[K, V]) resize(newSize int) {
	if newSize < _INITIAL_SIZE {
		newSize = _INITIAL_SIZE
	}
	newHash := new(closedHash[K, V])
	newHash.createTable(newSize)

	for i := 0; i < hash.size; i++ {
		if hash.table[i].state == OCCUPIED {
			newHash.Save(hash.table[i].key, hash.table[i].value)
		}
	}
	*hash = *newHash
}

// getPosition searches for the position of a key in the hash table.
// If the key is found, returns its index.
// If the key is not found and searchEmpty is true, returns the index of an available empty position.
// If the key is not found and searchEmpty is false, returns -1.
func (hash *closedHash[K, V]) getPosition(key K, searchEmpty bool) int {
	keyHash := hash.getKeyHash(key)
	for hash.table[keyHash].state != EMPTY {
		if hash.table[keyHash].state == OCCUPIED && hash.table[keyHash].key == key {
			return keyHash
		}
		keyHash = (keyHash + 1) % hash.size
	}
	if searchEmpty {
		return keyHash
	}
	return -1
}

func (hash *closedHash[K, V]) createTable(size int) {
	hash.table = make([]hashCell[K, V], size)
	hash.size = size
}

// ================= Hash Primitives ==================

func NewHash[K comparable, V any]() Map[K, V] {
	hash := new(closedHash[K, V])
	hash.createTable(_INITIAL_SIZE)
	return hash
}

func (hash *closedHash[K, V]) Save(key K, value V) {
	cell := hashCell[K, V]{
		key:   key,
		value: value,
		state: OCCUPIED,
	}
	keyHash := hash.getPosition(key, true)

	if hash.table[keyHash].state == EMPTY {
		hash.count++
	}
	hash.table[keyHash] = cell

	if loadFactor := float64(hash.count+hash.deleted) / float64(hash.size); loadFactor >= _LOAD_FACTOR_INC {
		hash.resize(hash.size * _RESIZE_FACTOR)
	}
}

func (hash *closedHash[K, V]) Contains(key K) bool {
	return hash.getPosition(key, false) != -1
}

func (hash *closedHash[K, V]) Get(key K) V {
	if pos := hash.getPosition(key, false); pos != -1 {
		return hash.table[pos].value
	}
	panic(_PANIC_HASH)
}

func (hash *closedHash[K, V]) Remove(key K) V {
	if pos := hash.getPosition(key, false); pos != -1 {
		value := hash.table[pos].value
		hash.table[pos].state = DELETED
		hash.count--
		hash.deleted++
		if loadFactor := float64(hash.count+hash.deleted) / float64(hash.size); loadFactor <= _LOAD_FACTOR_DEC {
			hash.resize(hash.size / _RESIZE_FACTOR)
		}
		return value
	}
	panic(_PANIC_HASH)
}

func (hash closedHash[K, V]) Count() int {
	return hash.count
}

// =================== Internal Iterator ===================

func (hash closedHash[K, V]) Iterate(visit func(key K, value V) bool) {
	for _, cell := range hash.table {
		if cell.state == OCCUPIED && !visit(cell.key, cell.value) {
			break
		}
	}
}

// =========== External Iterator Auxiliaries ============

func (it *closedHashIterator[K, V]) nextOccupied() {
	for it.index < it.hash.size && it.hash.table[it.index].state != OCCUPIED {
		it.index++
	}
}

// =================== External Iterator ===================

func (hash *closedHash[K, V]) Iterator() MapIterator[K, V] {
	it := new(closedHashIterator[K, V])
	it.hash = hash
	it.nextOccupied()
	return it
}

func (it *closedHashIterator[K, V]) HasNext() bool {
	for it.index < it.hash.size {
		if it.hash.table[it.index].state == OCCUPIED {
			return true
		}
		it.index++
	}
	return false
}

func (it *closedHashIterator[K, V]) Current() (K, V) {
	if !it.HasNext() {
		panic(_PANIC_ITERATOR)
	}
	return it.hash.table[it.index].key, it.hash.table[it.index].value
}

func (it *closedHashIterator[K, V]) Next() {
	if !it.HasNext() {
		panic(_PANIC_ITERATOR)
	}
	it.index++
	it.nextOccupied()
}
