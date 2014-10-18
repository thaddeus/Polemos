package main

type lobby struct {
    // Registered connections.
    connections map[*connection]bool

    // Inbound messages from the connections.
    broadcast chan []byte

    // Register requests from the connections.
    register chan *connection

    // Unregister requests from connections.
    unregister chan *connection
}

var lobby = lobby{
    broadcast:   make(chan []byte),
    register:    make(chan *connection),
    unregister:  make(chan *connection),
    connections: make(map[*connection]bool),
}

func (l *lobby) run() {
    for {
        select {
        case c := <-l.register:
            l.connections[c] = true
        case c := <-l.unregister:
            if _, ok := l.connections[c]; ok {
                delete(l.connections, c)
                close(c.send)
            }
        case m := <-l.broadcast:
            for c := range l.connections {
                select {
                case c.send <- m:
                default:
                    delete(l.connections, c)
                    close(c.send)
                }
            }
        }
    }
}
