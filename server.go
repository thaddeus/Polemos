package main

type server struct {
    // Registered connections.
    connections map[*connection]bool

    // Inbound messages from the connections.
    broadcast chan []byte

    // Register requests from the connections.
    register chan *connection

    // Unregister requests from connections.
    unregister chan *connection
}

var loginServer = server{
    broadcast:   make(chan []byte),
    register:    make(chan *connection),
    unregister:  make(chan *connection),
    connections: make(map[*connection]bool),
}

func (server *server) run() {
    for {
        select {
        case c := <-server.register:
            server.connections[c] = true
        case c := <-server.unregister:
            if _, ok := server.connections[c]; ok {
                delete(server.connections, c)
                close(c.send)
            }
        case m := <-server.broadcast:
            for c := range server.connections {
                select {
                case c.send <- m:
                default:
                    delete(server.connections, c)
                    close(c.send)
                }
            }
        }
    }
}
