package inbound

import (
    "log"

    "github.com/armon/go-socks5"
)

func StartSOCKS5(listen string) {
    conf := &socks5.Config{}
    server, err := socks5.New(conf)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("[SOCKS5] Listening on", listen)
    if err := server.ListenAndServe("tcp", listen); err != nil {
        log.Fatal(err)
    }
}
