package app

import (
	"net/http"
	"strings"

	"github.com/googollee/go-socket.io"
	"github.com/revel/revel"
)

var (
	sioServer   *socketio.Server
	revelHandle http.Handler
)

func handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	// route socketio requests to the socketio handler
	// and send everything else to the revel handler
	if strings.HasPrefix(path, "/socket.io/") {
		sioServer.ServeHTTP(w, r)
	} else {
		revelHandle.ServeHTTP(w, r)
	}
}

func PatchServer() {
	// create socketio server config
	// config := &socketio.Config{}
	// config.HeartbeatTimeout = 2
	// config.ClosingTimeout = 4

	// create socketio server
	sioServer, _ = socketio.NewServer(nil)
	// if err != nil {
	// 	fmt.Println("Error")
	// }
	// sioServer.SetPingInterval(2)
	// sioServer.SetPingTimeout(4)

	// register global and namespace handlers
	registerHandlers(sioServer)

	// store a reference to revel's old http.Handler
	revelHandle = revel.Server.Handler

	// replace revel.Server.Handler with our new handler
	revel.Server.Handler = http.HandlerFunc(handle)
}
