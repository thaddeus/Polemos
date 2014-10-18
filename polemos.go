package main

import (
    "flag"
    "go/build"
    "log"
    "net/http"
    "text/template"
)

var (
    address = flag.String("address", ":8080", "listening address")
    clientTemplate *template.Template
)

func clientHandler(c http.ResponseWriter, req *http.Request) {
    clientTemplate.Execute(c, req.Host)
}

func main() {
    flag.Parse()
    clientTemplate = template.Must(template.ParseFiles("./Hubris/hubris.html"))
    go h.run()
    http.HandleFunc("/", clientHandler)
    http.HandleFunc("/ws", wsHandler)
    if err := http.ListenAndServe(*address, nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }

}