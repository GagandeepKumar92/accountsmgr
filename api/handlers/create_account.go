package handlers

import (
	gserver "accountsmgr"
	"accountsmgr/domain"
	"accountsmgr/gen/restapi/operations/create_account"

	"github.com/go-openapi/runtime/middleware"
)

func NewCreateAccount(rt *gserver.Runtime) create_account.CreateAccountHandler {
	return &createAccount{rt: rt}
}

type createAccount struct {
	rt *gserver.Runtime
}

func (f *createAccount) Handle(fup create_account.CreateAccountParams) middleware.Responder {

	usr := &domain.Account{UserID: *fup.Body.UserID, UserName: "", Source: fup.Body.Source, Name: *fup.Body.Name}

	if err := f.rt.GetManager().CreateAccount(usr); err != nil {
		return create_account.NewCreateAccountBadRequest().WithPayload(asErrorResponse(err))
	}

	return create_account.NewCreateAccountCreated()

}
