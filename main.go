package main

import (
	"portobe/pkg/router"
	"portobe/pkg/setting"
)

func main() {
	router := router.InitRouter()

	router.Run(
		setting.CONFIG.Address + ":" + setting.CONFIG.Port)
}
