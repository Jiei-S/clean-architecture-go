// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package openapi

// Defines values for HealthStatus.
const (
	Healthy   HealthStatus = "healthy"
	Unhealthy HealthStatus = "unhealthy"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Health defines model for Health.
type Health struct {
	Status HealthStatus `json:"status"`
}

// HealthStatus defines model for Health.Status.
type HealthStatus string

// User defines model for User.
type User struct {
	Age       int    `json:"age"`
	FirstName string `json:"firstName"`
	Id        string `json:"id"`
	LastName  string `json:"lastName"`
}

// Conflict defines model for Conflict.
type Conflict = Error

// NotFound defines model for NotFound.
type NotFound = Error

// Unauthorized defines model for Unauthorized.
type Unauthorized = Error

// AddUser defines model for AddUser.
type AddUser = interface{}

// AddUserJSONBody defines parameters for AddUser.
type AddUserJSONBody = interface{}

// AddUserJSONRequestBody defines body for AddUser for application/json ContentType.
type AddUserJSONRequestBody = AddUserJSONBody