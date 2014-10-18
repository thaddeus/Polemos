package main

import (
    "flag"
    "log"
    "net/http"
    "text/template"
)

var (
    address = flag.String("address", ":8080", "listening address")
    clientTemplate *template.Template
    jsTemplate *template.Template
    cssTemplate *template.Template
)

func clientHandler(c http.ResponseWriter, req *http.Request) {
    clientTemplate.Execute(c, req.Host)
}

func cssHandler(c http.ResponseWriter, req *http.Request) {
    cssTemplate.Execute(c, req.Host)
}

func jsHandler(c http.ResponseWriter, req *http.Request) {
    jsTemplate.Execute(c, req.Host)
}

func main() {
    flag.Parse()
    clientTemplate = template.Must(template.ParseFiles("./Hubris/hubris.html"))
    cssTemplate = template.Must(template.ParseFiles("./Hubris/hubris.css"))
    jsTemplate = template.Must(template.ParseFiles("./Hubris/hubris.js"))
    go loginServer.run()
    http.HandleFunc("/", clientHandler)
    http.HandleFunc("/hubris.css", cssHandler)
    http.HandleFunc("/hubris.js", jsHandler)
    http.HandleFunc("/ws", wsHandler)
    if err := http.ListenAndServe(*address, nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }

}