# Socket.IO, Foundation, HandlebarsJS Revel application skeleton

This is a Revel app skeleton with some commonly used frameworks and libraries.

* WebSockets using Socket.IO for the client and the [go-socketio](https://github.com/googollee/go-socket.io) Go package for the server.
* [Foundation](https://foundation.zurb.com) CSS/JS framework for building UIs and interactions.
* Client-side templating with [Handlebars](https://handlebarsjs.com).

Install with :

    go get -d github.com/Simon-L/socketio-foundation-handlebars-skeleton

Use with :

    revel new project/path github.com/Simon-L/socketio-foundation-handlebars-skeleton

Test with :

    revel run project/path
And visit <tt>http://localhost:9000</tt> for more infos.

Integrating go-socketio to Revel taken from [here](http://www.pixeldonor.com/2014/apr/30/combining-revel-and-socketio/).
