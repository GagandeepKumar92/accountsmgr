package inmemory

import (
	"fmt"

	"accountsmgr/db"
	"accountsmgr/domain"
)

func init() {
	db.RegisterDataStore("inmemory", NewClient)
	fmt.Println("Call is in Init")
}

func NewClient() (db.DataStore, error) {
	return &client{ds: make(map[string]*domain.Account)}, nil
}

type client struct {
	ds map[string]*domain.Account
}

func (c *client) CreateAccount(account *domain.Account) (string, error) {
	c.ds[account.ID] = account
	return account.ID, nil
}
