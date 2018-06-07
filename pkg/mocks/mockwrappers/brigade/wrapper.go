package mocks

import (
	"github.com/Azure/brigade/pkg/storage"
)

// Store is the wrapper for brigade store interface for mocks.
type Store interface {
	storage.Store
}
