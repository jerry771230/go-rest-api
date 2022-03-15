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

## Ref

- [Learning Go by examples: part 2 - Create an HTTP REST API Server in Go](https://dev.to/aurelievache/learning-go-by-examples-part-2-create-an-http-rest-api-server-in-go-1cdm)
- [What does this command do 'GOFLAGS=-mod=mod'?](https://stackoverflow.com/questions/71121641/what-does-this-command-do-goflags-mod-mod)
- [golang内置包管理工具go mod简明教程](https://segmentfault.com/a/1190000019314903)
