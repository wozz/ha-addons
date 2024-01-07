package main

import (
    "flag"
    "log"
    "net/http"
    "net/url"

    "github.com/gorilla/websocket"
)

var (
  addr = flag.String("addr", ":8099", "http service address")
  staticDir = flag.String("static-dir", "/static", "Directory for static files")
)

var upgrader = websocket.Upgrader{} // use default options

func proxyWs(w http.ResponseWriter, r *http.Request) {
    // Upgrade initial HTTP connection to a WebSocket connection
    clientConn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade:", err)
        return
    }
    defer clientConn.Close()

    params := r.URL.Query()
    targetHost := params.Get("host")
    targetPort := params.Get("port")
    if targetPort == "" {
      targetPort = "5580"
    }
    targetPath := params.Get("path")
    if targetPath == "" {
      targetPath = "/ws"
    }

    targetURL := url.URL{Scheme: "ws", Host: targetHost + ":" + targetPort, Path: targetPath}
    targetConn, _, err := websocket.DefaultDialer.Dial(targetURL.String(), nil)
    if err != nil {
        log.Print("dial:", err)
        return
    }
    defer targetConn.Close()

    // Proxying data between client and target
    go transferMessages(clientConn, targetConn)
    transferMessages(targetConn, clientConn)
}

func transferMessages(src, dst *websocket.Conn) {
    for {
        mt, message, err := src.ReadMessage()
        if err != nil {
            log.Println("read:", err)
            break
        }
        err = dst.WriteMessage(mt, message)
        if err != nil {
            log.Println("write:", err)
            break
        }
    }
}

func main() {
    flag.Parse()
    log.SetFlags(0)
    http.HandleFunc("/wsproxy", proxyWs)
    fileServer := http.FileServer(http.Dir(*staticDir))
    http.Handle("/", fileServer)
    log.Printf("Starting server on %s\n", *addr)
    log.Fatal(http.ListenAndServe(*addr, nil))
}
