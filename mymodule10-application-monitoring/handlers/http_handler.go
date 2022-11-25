package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type HttpHandler = func(http.ResponseWriter, *http.Request, httprouter.Params) // alias
