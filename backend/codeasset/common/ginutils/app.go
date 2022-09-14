package ginutils

import (
	"fmt"
	"time"

	"backend/codeasset/common/middleware/cors"
	"backend/codeasset/common/middleware/logger"
	"backend/codeasset/models"
	"backend/codeasset/utils/lib"
	"backend/codeasset/utils/tlog"

	// "common/middleware/cors"

	// "common/middleware/logger"
	// . "common/share"
	// "fclient"
	// "utils/lib"
	// "utils/tlog"

	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type App struct {
	Cors    []string
	Addr    string
	Name    string
	Env     string
	Version string
	Debug   bool
	NodeID  int
}

func NewApp(name, version string) *App {
	log.SetFlags(log.Flags() | log.Lshortfile)

	p := &App{
		Name:    name,
		Version: version,
	}
	tlog.Logger(version, name, strconv.Itoa(models.Config.NodeID))
	return p
}

func (a *App) Init() *gin.Engine {
	// set gin model
	if a.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	// set logger level
	tlog.SetLevel(a.Debug)

	//
	// fabric-sdk
	// client := fclient.NewDefaultClient()
	// version, err := client.QueryCC(FABRIC_VERSION, [][]byte{})
	// if err != nil {
	// 	panic("please set fabric")
	// }

	localIPAddr, _ := lib.GetLocalIPAddr()
	fmt.Printf("\n")
	fmt.Printf("//////////////////////////////////////////////////\n")
	fmt.Printf("App Name: %s\n", a.Name)
	fmt.Printf("App Env: %s\n", a.Env)
	fmt.Printf("App Version: %s\n", a.Version)
	fmt.Printf("Start-up time: %s", time.Now().Format("2006-01-02 15:04:05\n"))
	fmt.Printf("Local Address: %s\n", localIPAddr)
	fmt.Printf("Listening and serving HTTP on %s\n", a.Addr)
	fmt.Printf("Snowflake Node: %d\n", a.NodeID)
	// fmt.Printf("Chaincode Version: %s\n", string(version))
	fmt.Printf("//////////////////////////////////////////////////\n\n")

	// return gin router
	// r := gin.Default()
	r := gin.New()
	r.Use(cors.Cors())

	// if len(a.Cors) > 0 {
	// 	r.Use(cors.Default(a.Cors))
	// }
	// logger
	r.Use(logger.Logger())
	// r.Use(logger.Qps())
	// r.Use(logger.UserAndPathLimit(r))
	return r
}
