package service

import (
	"fmt"
	"time"

	"accountsmgr/db"

	"accountsmgr/clients"

	"accountsmgr/domain"

	"github.com/segmentio/ksuid"
)

type mgr struct {
	ds db.DataStore
}

func (m *mgr) CreateAccount(acc *domain.Account) *domain.Error {

	res, err := clients.NewClientManager().FindUser()

	/*if err != nil {
		derr, ok := err.(domain.Err)
		//fmt.Println("The typecasted Delete status from handler", derr.StatusCode())
		if ok {
			switch derr.StatusCode() {
			//case 404:
			//	return users.NewFindUsersParams().WithPayload(asErrorResponse(err.(*domain.Error)))
			}
		} else {
			return users.NewFindUsersDefault(500).FindUsersDefault.Payload
		}
	}*/

	if err != nil {
		return &domain.Error{Code: 500, Message: "Internal Server Error"}
	}

	userName := ""
	if len(res) > 0 {
		for i := 0; i < len(res); i++ {
			if res[i].ID == acc.UserID {
				userName = *res[i].Name
				fmt.Println("User name is ", userName)
			}
		}
	}

	if userName == "" {
		return &domain.Error{Code: 400, Message: "User doesn't exist"}
	}

	acc.CreatedAt = time.Now().UTC()
	acc.UpdatedAt = time.Now().UTC()
	acc.ID = ksuid.New().String()
	acc.UserName = userName

	m.ds.CreateAccount(acc)
	return nil

}
