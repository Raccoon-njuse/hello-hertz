// Code generated by hertz generator.

package main

import (
	handler "github.com/Raccoon-njuse/hello-hertz/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.GET("/hello", handler.Hello)

	// your code ...
}
