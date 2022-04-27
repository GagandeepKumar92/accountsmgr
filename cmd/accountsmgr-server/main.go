package main

import (
	"accountsmgr"
	"accountsmgr/api/handlers"
	"accountsmgr/gen/restapi"
	"accountsmgr/gen/restapi/operations"
	"flag"
	"log"

	"github.com/go-openapi/loads"
	//_ "accountsmgr/db/inmemory"
	_ "accountsmgr/db/mongo"
)

func main() {

	var portFlag = flag.Int("port", 5000, "Port to run this service on")

	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewAccountsmgrAPI(swaggerSpec)

	v := accountsmgr.NewRunTime("accountsmgr")
	api.CreateAccountCreateAccountHandler = handlers.NewCreateAccount(v)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// TODO: Set Handle

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
