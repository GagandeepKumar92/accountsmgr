package clients

import (
	apiclient "github.com/GagandeepKumar92/go-usersclient/client"
	"github.com/GagandeepKumar92/go-usersclient/client/users"
	apimodel "github.com/GagandeepKumar92/go-usersclient/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type ClientManager struct {
	cs *apiclient.GoUserclient
}

func NewClientManager() *ClientManager {
	rt := client.New("localhost:4000", "", apiclient.DefaultSchemes)

	rt.Consumers["application/json"] = runtime.JSONConsumer()
	rt.Producers["application/json"] = runtime.JSONProducer()

	return &ClientManager{cs: apiclient.New(rt, nil)}
}

func (c *ClientManager) FindUser() ([]*apimodel.User, error) {

	req := users.NewFindUsersParams()
	res, err := c.cs.Users.FindUsers(req)

	if err != nil {
		return nil, err
	}

	return res.Payload, nil
}
