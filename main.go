package main

import (
	"adame/repository"
	"adame/router"
)

func main() {
	repository.InitDB()
	router.InitRouter()
}
