package service

import (
	"time"

	"accountsmgr/db"

	"accountsmgr/domain"

	"github.com/segmentio/ksuid"
)

type mgr struct {
	ds db.DataStore
}

func (m *mgr) CreateAccount(usr *domain.Account) *domain.Error {

	usr.CreatedAt = time.Now().UTC()
	usr.UpdatedAt = time.Now().UTC()
	usr.ID = ksuid.New().String()

	/*if len(usr.Name) < 3 {
		return &domain.Error{Code: 400, Message: "Name should be at least 3 characters long"}
	}*/

	m.ds.CreateAccount(usr)
	return nil

}
