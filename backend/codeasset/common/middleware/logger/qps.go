package logger

import (
	"fmt"
	"net/http"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	flagPrintQps bool = true
)

func OptionQps(flag bool) {
	flagPrintQps = flag
}

func RLimit() {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println("Error Getting Rlimit ", err)
	}
	//fmt.Println(rLimit)
	rLimit.Max = 999999
	rLimit.Cur = 999999
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		//fmt.Println("Error Setting Rlimit ", err)
	}
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		//fmt.Println("Error Getting Rlimit ", err)
	}
	fmt.Println("Rlimit Info:", rLimit)
}

func UserAndPathLimit(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := gPathMgr.AddCallPath(c)
		if err != nil {
			c.AbortWithStatus(http.StatusTooManyRequests)
		}
	}
}

func Qps() gin.HandlerFunc {
	RLimit() // init Rlimit
	timeStart := time.Now()
	arrayQps := []int64{}
	var countQps int64 = 0
	fmt.Println("enable qps count")
	go func() {
		for true {
			time.Sleep(time.Second)
			if flagPrintQps {

				if len(arrayQps) > 400 {
					arrayQps = []int64{}
				}

				if len(arrayQps) > 0 {
					if arrayQps[len(arrayQps)-1] != countQps && countQps != 0 {
						arrayQps = append(arrayQps, countQps)
					}
				} else {
					if countQps != 0 {
						arrayQps = append(arrayQps, countQps)
					}
				}
				// log.Println(arrayQps)
				// log.Println(gPathMgr)
			}
		}
	}()
	return func(c *gin.Context) {
		t := time.Now()
		//requestID := t.UnixNano() % 1000000000
		intervalTime := t.UnixNano() - timeStart.UnixNano()
		if intervalTime > 1000000000 {
			//fmt.Println(t.UnixNano())
			//fmt.Println(timeStart.UnixNano())
			timeStart = time.Now()
			countQps = 1
		} else {
			countQps++
		}
	}
}
