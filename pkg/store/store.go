package store

import (
	"fmt"
	"sync"
)

// KVStore represents a thread safe key value store.
type KVStore struct {
	data  map[string]string
	mutex sync.RWMutex
}

// Interface provies an api for a store.
type Interface interface {
	Read(key string) (val string, err error)
	Write(key, val string)
	Delete(key string) error
	AddAll(src Interface)
	Iterator() Iterator
}

// Iterator provides an api for an iterator.
type Iterator interface {
	HasNext() bool
	Next() (key, value string)
}

// KVIterator implements the Iterator interface.
type KVIterator struct {
	data []string
	curr int
}

// New instantiates a new store.
func New() Interface {
	return &KVStore{
		data:  map[string]string{},
		mutex: sync.RWMutex{},
	}
}

// Read the value associated with the given key.
// If the key does not exist, an error is returned.
func (s *KVStore) Read(key string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.read(key)
}

func (s *KVStore) read(key string) (string, error) {
	err := fmt.Errorf("Key not found: %s", key)
	if s.data == nil {
		return "", err
	}
	val, ok := s.data[key]
	if !ok {
		return "", err
	}
	return val, nil
}

// Write will store the given value in the given key.
func (s *KVStore) Write(key, val string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.write(key, val)
}

func (s *KVStore) write(key, val string) {
	if s.data == nil {
		s.data = map[string]string{}
	}
	if len(key) < 1 || len(val) < 1 {
		return
	}
	s.data[key] = val
}

// Delete removes the given key from the store.
// If the key does not exist, an error is returned.
func (s *KVStore) Delete(key string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.delete(key)
}

func (s *KVStore) delete(key string) error {
	if _, err := s.read(key); err != nil {
		return err
	}
	delete(s.data, key)
	return nil
}

// AddAll will add every element of a given
// store into the receiver's store.
func (s *KVStore) AddAll(src Interface) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.addAll(src)
}

func (s *KVStore) addAll(src Interface) {
	if src == nil {
		return
	}

	it := src.Iterator()
	for it.HasNext() {
		s.write(it.Next())
	}
}

// Iterator returns an iterator over the store.
// Multiple calls of Iterator() are not guarenteed
// to return in the same order.
func (s *KVStore) Iterator() Iterator {
	if s.data == nil {
		s.data = map[string]string{}
	}

	data := make([]string, len(s.data)*2)
	index := 0
	for k, v := range s.data {
		// more efficient than append
		data[index] = k
		index++
		data[index] = v
		index++
	}

	return &KVIterator{
		data: data,
		curr: -2,
	}
}

// HasNext will relay whether there is another element
// that can be retreived. It HasNext() returns true, it
// is safe to call Next().
func (it *KVIterator) HasNext() bool {
	return it.curr < len(it.data)-3
}

// Next returns the next element.
func (it *KVIterator) Next() (key, value string) {
	it.curr += 2
	return it.data[it.curr], it.data[it.curr+1]
}
