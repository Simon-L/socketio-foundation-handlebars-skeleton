package socketio

import (
	"fmt"

	"github.com/googollee/go-socket.io"
)

func registerHandlers(sioServer *socketio.Server) {
	sioServer.On("disconnection", func() {
		fmt.Println("on disconnect")
	})
	sioServer.On("connection", func() {
		fmt.Println("on connect")
	})
}
