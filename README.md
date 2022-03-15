# go-restful-api

Create first HTTP RESTful API server in Go.

## Commands

### Run app

```bash
$ go run ./internal/main.go
```

Test

```bash
$ curl localhost:8080

Hello, "/"%
```

### Generate an executable binary

```bash
$ go build -p ./bin/go-restful-api ./internal/main.go
```

## Taskfile

A Makefile alternative.

Install `task`

```bash
$ brew install go-task/tap/go-task
```

Display the list of available tasks:

```bash
$ task --list

task: Available tasks for this project:
* build: 		Build the app
* build.linux: 		Build the app for linux
* run: 			Run the app
* swagger.doc: 		Doc for swagger
* swagger.gen: 		Generate Go code
* swagger.validate: 	Validate swagger

```

## Swagger

Install swagger

```bash
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
```

Validate the swagger file

```bash
$ task swagger.validate

task: [swagger.validate] swagger validate pkg/swagger/swagger.yml
2022/03/15 14:40:19
The swagger spec at "pkg/swagger/swagger.yml" is valid against swagger specification 2.0
```

Create the swagger definitions in an HTML doc.

```bash
$ task swagger.doc

task: [swagger.doc] docker run -i yousan/swagger-yaml-to-html < pkg/swagger/swagger.yml > doc/index.html
```

Generate Go code

```bash
$ task swagger.gen

task: [swagger.gen] GOFLAGS=-mod=mod go generate github.com/jerry771230/go-restful-api/internal github.com/jerry771230/go-restful-api/pkg/swagger
```

The command will generate several useful files in `pkg/swagger/server`

```pre
pkg/swagger
├── gen.go
├── server
│   └── restapi
│       ├── configure_hello_api.go
│       ├── doc.go
│       ├── embedded_spec.go
│       ├── operations
│       │   ├── check_health.go
│       │   ├── check_health_parameters.go
│       │   ├── check_health_responses.go
│       │   ├── check_health_urlbuilder.go
│       │   ├── get_gopher_name.go
│       │   ├── get_gopher_name_parameters.go
│       │   ├── get_gopher_name_responses.go
│       │   ├── get_gopher_name_urlbuilder.go
│       │   ├── get_hello_user.go
│       │   ├── get_hello_user_parameters.go
│       │   ├── get_hello_user_responses.go
│       │   ├── get_hello_user_urlbuilder.go
│       │   └── hello_api_api.go
│       └── server.go
└── swagger.yml
```

## Ref

- [Learning Go by examples: part 2 - Create an HTTP REST API Server in Go](https://dev.to/aurelievache/learning-go-by-examples-part-2-create-an-http-rest-api-server-in-go-1cdm)
- [What does this command do 'GOFLAGS=-mod=mod'?](https://stackoverflow.com/questions/71121641/what-does-this-command-do-goflags-mod-mod)
- [golang内置包管理工具go mod简明教程](https://segmentfault.com/a/1190000019314903)
- [What does brew tap mean?](https://stackoverflow.com/questions/34408147/what-does-brew-tap-mean)
