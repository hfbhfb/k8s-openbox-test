package logger

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	limitTypeSimpleCount = 100 //计数器法
	limitTypeWindow      = 200 //滑动窗口计数
	limitTypeLeakBucket  = 300 //漏桶算法
	limitTypeTokenBucket = 400 //令牌桶算法
)

type _LimitMgr struct {
	path      map[string]*CountType
	typeLimit int
}

var (
	gPathMgr *_LimitMgr
)

type CountType struct {
	count  int64
	time   time.Time //
	bucket int64     // 桶
}

func InitUserAndPathLimit(r *gin.Engine) {
	gPathMgr = &_LimitMgr{typeLimit: limitTypeTokenBucket}
	m := gPathMgr
	gPathMgr.path = make(map[string]*CountType)
	routers := r.Routes()
	for _, v := range routers {
		gPathMgr.path[v.Method+v.Path] = &CountType{count: 1, bucket: 600}
	}
	go func() {
		for true {
			time.Sleep(300 * time.Millisecond)
			if m.typeLimit == limitTypeTokenBucket {
				for _, v := range m.path {
					if v.bucket < 600 {
						v.bucket = v.bucket + 2 // 每秒: 1000/300 * 2
					}
				}
			}
		}
	}()
}

// GetTableName get sql table name.获取数据库名字
func (m *_LimitMgr) AddCallPath(c *gin.Context) error {
	// fmt.Println(c.Request.Method + c.Request.URL.Path)
	// fmt.Println(m.path)
	obj := m.path[c.Request.Method+c.Request.URL.Path]
	if obj != nil {
		if m.typeLimit == limitTypeTokenBucket {
			if obj.bucket > 1 {
				obj.count++
				obj.bucket--
				// fmt.Println(obj.bucket)
				return nil
			} else {
				return errors.New("erro limit")
			}
		}
	}
	return errors.New("erro limit")
}

func (m *_LimitMgr) String() string {
	for k, v := range m.path {
		if v.count > 1 {
			fmt.Println(k, "   ", v.count)
		}
	}
	return ""
}
