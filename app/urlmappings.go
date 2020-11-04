package app

import (
	"github.com/DataPsycho/demoapp/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
