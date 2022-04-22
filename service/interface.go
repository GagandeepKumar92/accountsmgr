package service

import (
	"fmt"

	"accountsmgr/db"

	"accountsmgr/domain"
)

type Manager interface {
	CreateAccount(usr *domain.Account) *domain.Error
}

func NewManager(dbType string) Manager {

	ds, err := db.NewDataStore(dbType)
	if err != nil {
		fmt.Println("The err is", err)
		return nil
	}

	return &mgr{ds: ds}

}
