package app

import (
	"github.com/DataPsycho/demoapp/src/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
