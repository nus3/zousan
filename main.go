package main

import (
	"github.com/TechLoCo/Zousan-api/db"
	"github.com/TechLoCo/Zousan-api/router"
)

func main() {
	db.Init()
	defer db.Close()

	router.Init()
}
