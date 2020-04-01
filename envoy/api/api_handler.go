package api

import (
	"net/http"
	manager_rpc "tracker/manager_grpc/rpc"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
)

type ApiHandler struct {
	Conn manager_rpc.CounterServiceClient
}

func NewApiHandler(conn manager_rpc.CounterServiceClient) *ApiHandler {
	return &ApiHandler{
		Conn: conn,
	}
}

func (ah *ApiHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	tokenResp, err := ah.Conn.NewToken(r.Context(), &empty.Empty{})
	if err != nil {
		log.Infof("NewToken request failed err: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	token := tokenResp.GetToken()
	RespondWithJSON(w, http.StatusOK, TokenMessage{Token: token})
	return
}

func (ah *ApiHandler) PathCount(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if len(token) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	counterResp, err := ah.Conn.UpdateCounter(r.Context(), &manager_rpc.UpdateCounterRequest{
		Token: token,
		Path:  r.URL.Path,
	})
	if err != nil {
		log.Infof("UpdateCounter request failed err: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	success := counterResp.GetSuccess()
	if !success {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	counter := counterResp.GetCounter()
	RespondWithJSON(w, http.StatusOK, CountMessage{Count: counter})
	return
}
