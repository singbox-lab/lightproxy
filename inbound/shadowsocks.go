package inbound

import (
    "log"

    ss "github.com/shadowsocks/go-shadowsocks2/core"
    "github.com/shadowsocks/go-shadowsocks2/socks"
    "net"
)

func StartShadowsocks(listen, method, password string) {
    cipher, err := ss.PickCipher(method, nil, password)
    if err != nil {
        log.Fatal(err)
    }

    l, err := net.Listen("tcp", listen)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("[Shadowsocks] Listening on %s using %s\n", listen, method)
    for {
        conn, err := l.Accept()
        if err != nil {
            log.Println("Accept error:", err)
            continue
        }
        go socks.HandleConn(conn, cipher.StreamConn)
    }
}
