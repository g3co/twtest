package rest

import (
	"net/http"
)

func NewHTTPROuter(svc Parser) http.Handler {
	rtr := http.NewServeMux()

	h := handler{svc: svc}

	rtr.HandleFunc("/block", h.CurrentBlockHandler)
	rtr.HandleFunc("/address/", h.AddressHandler)
	return rtr
}
