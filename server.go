package main

import (
    "fmt"
    "github.com/gorilla/rpc/v2"
    "github.com/gorilla/rpc/v2/json"
    "net/http"
)

func main() {
    ReadConfig()

    //RPC server for event aggregator requests
    rpcServer := rpc.NewServer()
    rpcServer.RegisterCodec(json.NewCodec(), "application/json")
    rpcServer.RegisterService(new(EventService), "")
    http.Handle("/", rpcServer)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Printf("Error ListenAndServe: %v\n", err)
    }
}