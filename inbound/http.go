package inbound

import (
    "log"
    "net/http"
)

func StartHTTPProxy(listen string) {
    proxy := &http.Server{
        Addr:    listen,
        Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            http.Error(w, "HTTP Proxy Not Implemented", http.StatusNotImplemented)
        }),
    }
    log.Println("[HTTP] Listening on", listen)
    if err := proxy.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}
