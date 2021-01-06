package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("Hello Sam")

	server := socketio.NewServer(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("New Connection")
		return nil
	})

	http.Handle("/socket.io/", server)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
