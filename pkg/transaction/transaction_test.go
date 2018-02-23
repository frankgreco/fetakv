package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	i := New()
	assert.Implements(t, (*Interface)(nil), i)
}

func TestStore(t *testing.T) {
	i := New()
	s := i.(*Transaction)
	assert.Equal(t, s.localStore, i.Store())
}
