package main

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	// server.OnConnect("/", func(s socketio.Conn) error {
	// 	s.SetContext("")
	// 	log.Println("New Connection")
	// 	return nil
	// })

	server.OnConnect("connection", func(so socketio.Socket) {

		log.Println("on connection")
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.Handle("/socket.io/", server)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
