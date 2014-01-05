// Package main is the core package of the Compute Engine instance.
package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/rpc/v2"
    "github.com/gorilla/rpc/v2/json"

    "socialvibes/config"
    svrpc "socialvibes/rpc"
)

func main() {
    config.ReadConfig()

    //RPC server for event aggregator requests
    rpcServer := rpc.NewServer()
    rpcServer.RegisterCodec(json.NewCodec(), "application/json")
    rpcServer.RegisterService(new(svrpc.EventService), "")
    http.Handle("/", rpcServer)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Printf("Error ListenAndServe: %v\n", err)
    }
}