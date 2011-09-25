package web

import (
	"http"
	"net"
)

var (
	welcomeText = `
<html>
  <head>
    <title>electric dragon</title>
  </head>
  <body>
    <h1>welcome to the electric dragon</h1>
  </body>
</html>`
)

func Serve(listener net.Listener) {
	http.HandleFunc("/", welcomeHandler)
	http.Serve(listener, nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Write([]byte(welcomeText))
}