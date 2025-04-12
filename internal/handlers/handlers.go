package handlers

import (
	"net/http"

)

type AuthHandler interface{
	Greeting(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}