package server

import (
	"github.com/Backend-GoAtreugo-server/server/middleware"
	"github.com/Backend-GoAtreugo-server/server/router"
	"github.com/Backend-GoAtreugo-server/utils"
	"github.com/savsgio/atreugo/v11"
)

func Start() {
	config := atreugo.Config{
		Addr: "0.0.0.0:4030",
	}
	server := atreugo.New(config)
	server.UseBefore(middleware.BeforeGlobal)

	router.Start_v1(server)

	server.GET("/", func(rc *atreugo.RequestCtx) error {
		rc.HTTPResponse("root dir api test", 200)
		return nil
	})
	err := server.ListenAndServe()
	utils.HandleErr(err)
}
