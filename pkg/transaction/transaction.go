package transaction

import (
	"github.com/frankgreco/fetakv/pkg/store"
)

// Transaction implements Interface
type Transaction struct {
	localStore store.Interface
}

// Interface provies an api for a transaction
type Interface interface {
	Store() store.Interface
}

// New creates a new instance of a transaction.
func New() Interface {
	return &Transaction{
		localStore: store.New(),
	}
}

// Store returns the store associated with this transaction.
func (t *Transaction) Store() store.Interface {
	return t.localStore
}
