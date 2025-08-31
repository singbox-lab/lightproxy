package main

import (
    "log"

    "github.com/yourname/light-proxy/config"
    "github.com/yourname/light-proxy/inbound"
)

func main() {
    cfg, err := config.Load("examples/config.json")
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }

    for _, inb := range cfg.Inbounds {
        switch inb.Type {
        case "socks5":
            go inbound.StartSOCKS5(inb.Listen)
        case "http":
            go inbound.StartHTTPProxy(inb.Listen)
        case "shadowsocks":
            go inbound.StartShadowsocks(inb.Listen, inb.Method, inb.Password)
        default:
            log.Println("Unknown inbound type:", inb.Type)
        }
    }

    select {} // Block forever
}
