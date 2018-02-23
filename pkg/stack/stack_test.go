package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/frankgreco/fetakv/pkg/transaction"
)

func TestNew(t *testing.T) {
	s := New()
	assert.Zero(t, len(s.data))
}

func TestPop(t *testing.T) {
	s := New()
	trans := transaction.New()
	s.data = append(s.data, trans)
	val := s.Pop()
	assert.Equal(t, val, trans)
	assert.Nil(t, s.Pop())
}

func TestPeek(t *testing.T) {
	s := New()
	assert.Nil(t, s.Peek())
	trans := transaction.New()
	s.data = append(s.data, trans)
	val := s.Peek()
	assert.Equal(t, val, trans)
	assert.Equal(t, val, trans)
}

func TestPush(t *testing.T) {
	s := New()
	trans := transaction.New()
	s.Push(trans)
	assert.Equal(t, len(s.data), 1)
	assert.Equal(t, s.data[0], trans)
}

func TestSize(t *testing.T) {
	s := New()
	assert.Zero(t, s.Size())
	trans := transaction.New()
	s.Push(trans)
	assert.Equal(t, s.Size(), 1)
	s.Peek()
	assert.Equal(t, s.Size(), 1)
	s.Pop()
	assert.Zero(t, s.Size())
	s.Pop()
	assert.Zero(t, s.Size())
}
