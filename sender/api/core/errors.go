package core

import "net/http"

var RouteNotFound = Error{
	"error",
	http.StatusNotFound,
	ErrorInfo{
		404,
		"Route not found",
	}}

var MethodNotAllowed = Error{
	"error",
	http.StatusMethodNotAllowed,
	ErrorInfo{
		405,
		"Method not allowed",
	}}

var InternalServerError = Error{
	"error",
	http.StatusInternalServerError,
	ErrorInfo{
		500,
		"Internal server error",
	}}

var ValidationError = Error{
	"error",
	http.StatusBadRequest,
	ErrorInfo{
		2000,
		"Validation error",
	}}

var ObjectNotFound = Error{
	"error",
	http.StatusBadRequest,
	ErrorInfo{
		2010,
		"Object not found",
	}}
