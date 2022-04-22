package mongo

import (
	"context"
	"fmt"
	"time"

	"accountsmgr/db"
	"accountsmgr/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	db.RegisterDataStore("mongo", NewClient)
}

func NewClient() (db.DataStore, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientCurrent, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println("The error is", err)
		return nil, err
	}

	err = clientCurrent.Ping(ctx, readpref.Primary())

	if err != nil {
		fmt.Println("The error is", err)
		return nil, err
	}

	return &client{dbc: clientCurrent.Database("accounts_mgr")}, nil
}

type client struct {
	dbc *mongo.Database
}

func (c *client) CreateAccount(account *domain.Account) (string, error) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := c.dbc.Collection("accountsmgr").InsertOne(ctx, account)

	if err != nil {
		fmt.Println("The error is", err)
		return "", err
	}

	return "", nil
}
