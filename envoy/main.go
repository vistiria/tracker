package main

import (
	"net/http"
	"tracker/envoy/api"
	manager_rpc "tracker/manager_grpc/rpc"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("ENVOY_ADDR", "127.0.0.1:6380")
	viper.SetDefault("COUNTER_ADDR", "127.0.0.1:6381")

	conn, err := grpc.Dial(viper.GetString("COUNTER_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create connection: %v", err)
	}
	client := manager_rpc.NewCounterServiceClient(conn)
	handler := api.NewApiHandler(client)

	router := mux.NewRouter()
	router.HandleFunc("/auth", handler.GetToken).Methods(http.MethodGet)
	router.PathPrefix("/").HandlerFunc(handler.PathCount).Methods(http.MethodGet)

	envoyAddr := viper.GetString("ENVOY_ADDR")
	log.Infof("Starting envoy service, listening on %s", envoyAddr)
	if err := http.ListenAndServe(envoyAddr, router); err != nil {
		log.Fatalf("envoy service error: %s", err)
	}
}
