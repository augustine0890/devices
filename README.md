# Device Service

## Project Structure
```
.
├── README.md
├── devices.json
├── go.mod
├── go.sum
└── main.go
```

## Introduction
- The REST API to get the devices
- The list of HTTP endpoints
 - GET `/devices`: returns a list of devices
 - POST `/devices`: creates a new devices
 - GET `/devices/{id}`: returns an existing device by ID
 - GET `/devices/type?type=X`: returns a list of devices by type
 - GET `/devices/status?status=X`: returns a list of devices by status
 
## Run the Server
- Just run `go mod download` and then `go run main.go` to start the project.
- The default port for listening is `8080`.
- Use a `cURL` command to issue an HTTP request or `Postman`:
  - GET `/devices`: returns a list of devices
    - `curl -X GET http://localhost:8080/devices`

## Dependencies
- Install the dependencies used by the project
- Gin Web Framework
  - `go get -u github.com/gin-gonic/gin`


## TODO
- Reorganize the project structure by following:
```
.
├── README.md
├── devices.json
├── go.mod
├── go.sum
├── handlers
│   └── handler.go
├── main.go
└── models
    └── device.go
```
- Writing unit-test for HTTP endpoint.