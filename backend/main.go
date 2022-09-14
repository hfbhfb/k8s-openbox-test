package main

import (
	"log"

	// "backend/codeasset/app"
	"backend/codeasset/models"
	"backend/codeasset/routers"

	"backend/codeasset/common/ginutils"
	. "backend/codeasset/common/share"
	"backend/codeasset/utils/db"
)

func main() {
	app := ginutils.NewApp(NAME, VERSION)
	app.Env = models.Config.Env
	app.Debug = models.Config.Debug
	app.NodeID = models.Config.NodeID
	app.Addr = models.Config.Addr
	router := app.Init()
	routers.Register(router)
	// log.Println(app)
	// log.Println(router)
	// ID := models.IDNode.Generate().String()
	// fmt.Println("snowid:", ID)
	// fmt.Println("snowid:", ID)
	// fmt.Println("snowid:", ID)
	// tlog.Debug("testlog")

	router.Run(app.Addr)

	defer log.Printf("main end")
	// defer controls.GrpcConnClose()
	defer db.CloseRedis()

}
