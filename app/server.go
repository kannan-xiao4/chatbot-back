package app

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func Start() error {
	app := setup()
	return endless.ListenAndServe(":3000", app)
}

func setup() *gin.Engine {
	app := gin.New()
	RouteV1(app)
	return app
}
