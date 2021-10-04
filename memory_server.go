//go:build local
// +build local

//go:generate hugo --gc --minify
//go:generate go build memory_server.go

// memory_server is an in-memory static site server.
//
// The server is intended to be run as a service. To do this place the
// following at /lib/systemd/system/memory-server.service
//  [Unit]
//  Description=In-memory tatic site server for local website.
//
//  [Service]
//  Type=simple
//  ExecStart=/usr/local/bin/memory_server
//
//  [Install]
//  WantedBy=multi-user.target
// and install the binary at /usr/local/bin/memory_server. Then the usual
// systemd controls can be used, for example, systemctl start memory-server.
package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed public
var content embed.FS

func main() {
	content, err := fs.Sub(content, "public")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(content)))
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}
