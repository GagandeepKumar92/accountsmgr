package db

import (
	"fmt"

	"accountsmgr/domain"
)

type DataStore interface {
	CreateAccount(account *domain.Account) (string, error)
}
type DatastoreFactory func() (DataStore, error)

var factories map[string]DatastoreFactory

func RegisterDataStore(key string, value DatastoreFactory) {
	if factories == nil {
		factories = make(map[string]DatastoreFactory)
	}
	factories[key] = value
}

func NewDataStore(dbType string) (DataStore, error) {
	fmt.Println("The length is", len(factories))
	return factories[dbType]()
}
