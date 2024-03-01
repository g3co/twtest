package rest

import (
	"encoding/json"
	"github.com/g3co/twtest/pkg/ethrpc"
	"io"
	"net/http"
	"strings"
)

type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []ethrpc.Transaction
}

type handler struct {
	svc Parser
}

// CurrentBlockHandler returns current block
func (h *handler) CurrentBlockHandler(w http.ResponseWriter, req *http.Request) {
	resp := struct {
		CurrentBlock int `json:"current_block"`
	}{
		CurrentBlock: h.svc.GetCurrentBlock(),
	}

	jResp, _ := json.Marshal(resp)
	w.WriteHeader(200)
	_, _ = w.Write(jResp)
}

// AddressHandler get transactions by address and subscribe handler
func (h *handler) AddressHandler(w http.ResponseWriter, r *http.Request) {
	address := strings.TrimPrefix(r.URL.Path, "/address/")

	// As per requirement, no external libraries. Had to use stdlib to handle it.
	switch r.Method {
	case http.MethodGet:
		resp := h.svc.GetTransactions(address)

		jResp, _ := json.Marshal(resp)
		w.WriteHeader(200)
		_, _ = w.Write(jResp)
	case http.MethodPost:
		addrReq := struct {
			Address string `json:"address"`
			Status  bool   `json:"status"`
		}{}

		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Wrong request", http.StatusBadRequest)
			return
		}

		if err = json.Unmarshal(reqBody, &addrReq); err != nil {
			http.Error(w, "Wrong request format", http.StatusBadRequest)
			return
		}

		if addrReq.Address == "" {
			http.Error(w, "The address was not specified", http.StatusBadRequest)
			return
		}

		if ok := h.svc.Subscribe(addrReq.Address); !ok {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		addrReq.Status = true
		jResp, _ := json.Marshal(addrReq)
		w.WriteHeader(200)
		_, _ = w.Write(jResp)
	default:
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}
