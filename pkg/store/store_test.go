package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	i := New()
	assert.Implements(t, (*Interface)(nil), i)
	kv, ok := i.(*KVStore)
	assert.True(t, ok)
	assert.Zero(t, len(kv.data))
}

func TestRead(t *testing.T) {
	kv := &KVStore{}
	_, err := kv.Read("foo")
	assert.NotNil(t, err)

	kv.data = map[string]string{}
	_, err = kv.Read("foo")
	assert.NotNil(t, err)

	kv.data["foo"] = "bar"
	val, err := kv.Read("foo")
	assert.Nil(t, err)
	assert.Equal(t, val, "bar")
}

func TestWrite(t *testing.T) {
	kv := &KVStore{}
	kv.Write("foo", "bar")
	val, err := kv.Read("foo")
	assert.Nil(t, err)
	assert.Equal(t, val, "bar")
	kv.Write("", "bar")
	assert.Equal(t, len(kv.data), 1)
}

func TestDelete(t *testing.T) {
	kv := &KVStore{}
	assert.NotNil(t, kv.Delete("foo"))
	assert.NotNil(t, kv.Delete(""))
	kv.Write("foo", "bar")
	assert.Nil(t, kv.Delete("foo"))
}

func TestAddAll(t *testing.T) {
	kv1 := &KVStore{
		data: map[string]string{
			"one": "two",
		},
	}
	kv2 := &KVStore{
		data: map[string]string{
			"three": "four",
		},
	}
	kv1.AddAll(kv2)
	kv1.AddAll(nil)
	assert.Equal(t, len(kv1.data), 2)
	val, err := kv1.Read("one")
	assert.Nil(t, err)
	assert.Equal(t, val, "two")

	val, err = kv1.Read("three")
	assert.Nil(t, err)
	assert.Equal(t, val, "four")
}

func TestIterator(t *testing.T) {
	kv := &KVStore{}
	i := kv.Iterator()
	kvIt := i.(*KVIterator)
	assert.Equal(t, len(kvIt.data), 0)
	assert.Implements(t, (*Iterator)(nil), i)
	kv.Write("foo", "bar")
	i = kv.Iterator()
	kvIt = i.(*KVIterator)
	assert.Equal(t, len(kvIt.data), 2)
}

func TestHasNext(t *testing.T) {
	kv := &KVStore{}
	kv.Write("foo", "bar")
	i := kv.Iterator()
	assert.True(t, i.HasNext())
	assert.True(t, i.HasNext())
	i = kv.Iterator()
	assert.True(t, i.HasNext())

	kvIt := &KVIterator{
		data: []string{"foo"},
		curr: -2,
	}
	assert.False(t, kvIt.HasNext())
}

func TestNext(t *testing.T) {
	kv := &KVStore{}
	i := kv.Iterator()
	assert.False(t, i.HasNext())
	assert.Panics(t, assert.PanicTestFunc(func() {
		i.Next()
	}))
	kv.Write("foo", "bar")
	i = kv.Iterator()
	assert.True(t, i.HasNext())
	k, v := i.Next()
	assert.Equal(t, k, "foo")
	assert.Equal(t, v, "bar")
	assert.Panics(t, assert.PanicTestFunc(func() {
		i.Next()
	}))
}
