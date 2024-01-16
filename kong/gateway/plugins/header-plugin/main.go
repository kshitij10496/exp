package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

const Version = "1.0.0"
const Priority = 1

type CustomHeader struct{}

func New() interface{} {
	return &CustomHeader{}
}

func (h *CustomHeader) Access(kong *pdk.PDK) {
	kong.Response.SetHeader("X-Custom-Header", "Hello, World!")
}

func main() {
	server.StartServer(New, Version, Priority)
}
