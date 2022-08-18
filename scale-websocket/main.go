package main

import (
	"log"
	"net/http"
	"os"

	"github.com/godbus/dbus/v5"
	"github.com/gorilla/websocket"
)

const (
	AppPath = "/medium/examples/chat"
	AppID   = "medium.examples.lock"
)

type App struct {
	Bus   *dbus.Conn
	Conns []*websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (a *App) handler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	a.Conns = append(a.Conns, conn)

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		//check message length before processing
		if len(p) > 0 {
			a.Bus.Emit(AppPath, AppID, p)
		}

	}

	// TODO : cleanup connection
}

func (a *App) listen() {

	conn, err := dbus.SystemBus()
	if err != nil {
		log.Fatal(os.Stderr, "Failed to connect to session bus:", err)
	}

	a.Bus = conn

	if err = conn.AddMatchSignal(
		dbus.WithMatchObjectPath(AppPath),
	); err != nil {
		log.Fatal("Error setting filter :", err)
	}

	c := make(chan *dbus.Signal, 10)

	conn.Signal(c)

	for v := range c {

		message := v.Body[0].([]byte)
		a.broadcast(message)

	}

}

func (a *App) broadcast(message []byte) {

	for _, conn := range a.Conns {
		// ignoring errors, with respect to UDP
		conn.WriteMessage(1, message)
	}
}

func main() {

	var app App

	http.HandleFunc("/sockettome", app.handler)

	// Listen for messages and broadcast it.
	go app.listen()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
