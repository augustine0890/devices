# Device Service

## Project Structure
## Introduction
- The REST API to get the devices
- The list of HTTP endpoints
 - GET `/devices`: returns a list of devices
 - POST `/devices`: creates a new devices
 - GET `/devices/{id}`: returns an existing device by ID
 - GET `/devices/{type}: returns a list of devices by type
 - GET `/devices/{status}: returns a list of devices by status
 
## Run the Server

## Dependencies
- Install the dependencies used by the project
- Gin Web Framework
  - `go get -u github.com/gin-gonic/gin`