package main

import (
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jerry771230/go-restful-api/pkg/swagger/server/restapi"
	"github.com/jerry771230/go-restful-api/pkg/swagger/server/restapi/operations"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func main() {
	// Initialize zap
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar = logger.Sugar()

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if nil != err {
		sugar.Fatal(err)
	}

	api := operations.NewHelloAPIAPI((swaggerSpec))
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); nil != err {
			sugar.Fatal(err)
		}
	}()

	server.Port = 8080

	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)
	api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(GetHelloUser)
	api.GetGopherNameHandler = operations.GetGopherNameHandlerFunc(GetGopherByName)

	if err := server.Serve(); nil != err {
		sugar.Fatal(err)
	}
}

func Health(operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("OK")
}

func GetHelloUser(user operations.GetHelloUserParams) middleware.Responder {
	return operations.NewGetHelloUserOK().WithPayload("Hello " + user.User + "!")
}

func GetGopherByName(gopher operations.GetGopherNameParams) middleware.Responder {
	// url := "https://raw.githubusercontent.com/scraly/gophers/main/"
	url := "https://github.com/scraly/gophers/raw/main/"
	if "" != gopher.Name {
		url = url + gopher.Name + ".png"
	} else {
		url = url + "dr-who.png"
	}
	sugar.Info(url)

	response, err := http.Get(url)
	sugar.Info(response)
	if nil != err {
		sugar.Info(err)
	}

	return operations.NewGetGopherNameOK().WithPayload(response.Body)
}
