package main

import (
        "context"
        "fmt"
        "log"
        "net/http"
        "sync"
        "time"
)

func startHttpServer(wg *sync.WaitGroup) *http.Server {
        srv := &http.Server{Addr: ":1100"}
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
        })

        go func() {
                defer wg.Done()
                if err := srv.ListenAndServe(); err != http.ErrServerClosed {
                        log.Fatalf("ListenAndServe(): %v", err)
                }
        }()

        return srv
}
