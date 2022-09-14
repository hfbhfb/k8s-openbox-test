package cors

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	AbortOnError    bool
	AllowAllOrigins bool

	// AllowedOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// Default value is ["*"]
	AllowedOrigins []string

	// AllowOriginFunc is a custom function to validate the origin. It take the origin
	// as argument and returns true if allowed or false otherwise. If this option is
	// set, the content of AllowedOrigins is ignored.
	AllowOriginFunc func(origin string) bool

	// AllowedMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (GET and POST)
	AllowedMethods []string

	// AllowedHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
	// If the special "*" value is present in the list, all headers will be allowed.
	// Default value is [] but "Origin" is always appended to the list.
	AllowedHeaders []string

	// ExposedHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
	ExposedHeaders []string

	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
	AllowCredentials bool

	// MaxAge indicates how long (in seconds) the results of a preflight request
	// can be cached
	MaxAge time.Duration
}

func (c *Config) AddAllowedMethods(methods ...string) {
	c.AllowedMethods = append(c.AllowedMethods, methods...)
}

func (c *Config) AddAllowedHeaders(headers ...string) {
	c.AllowedHeaders = append(c.AllowedHeaders, headers...)
}

func (c *Config) AddExposedHeaders(headers ...string) {
	c.ExposedHeaders = append(c.ExposedHeaders, headers...)
}

func (c Config) Validate() error {
	if c.AllowAllOrigins && (c.AllowOriginFunc != nil || len(c.AllowedOrigins) > 0) {
		return errors.New("conflict settings: all origins are allowed. AllowOriginFunc or AllowedOrigins is not needed")
	}
	if !c.AllowAllOrigins && c.AllowOriginFunc == nil && len(c.AllowedOrigins) == 0 {
		return errors.New("conflict settings: all origins disabled")
	}
	if !c.AllowAllOrigins && c.AllowOriginFunc != nil && len(c.AllowedOrigins) > 0 {
		return errors.New("conflict settings: if a allow origin func is provided, AllowedOrigins is not needed")
	}
	for _, origin := range c.AllowedOrigins {
		if !strings.HasPrefix(origin, "http://") && !strings.HasPrefix(origin, "https://") && !strings.HasPrefix(origin, "chrome-extension://") {
			return errors.New("bad origin: origins must include http:// or https://")
		}
	}
	return nil
}

var defaultConfig = Config{
	AbortOnError:     true,
	AllowAllOrigins:  true,
	AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "HEAD", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Content-Type", "Accept", "Referer", "User-Agent", "Version", "AccessToken", "PD-Token", "X-Token", "X-AccessKeyID", "X-Claims", "X-Request-ID", "Authorization", "Content-Length", "X-CSRF-Token", "Token", "session", "X_Requested_With", "Accept", "Origin", "Host", "Connection", "Accept-Encoding", "Accept-Language", "DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Pragma"},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
}

func DefaultConfig() Config {
	cp := defaultConfig
	return cp
}

func Default(allowOrigins []string) gin.HandlerFunc {
	config := DefaultConfig()
	config.AllowedOrigins = allowOrigins
	return New(config)
}

// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//fmt.Println(method)
		ReqHeader := []string{
			"Content-Type", "Origin", "Authorization", "Accept", "token",
			"cache-control", "x-requested-with", "X-Token", "User-Agent", "Referer", "PD-Token", "access_token", "open_id", "open-id", "x-auth-token", "x-client-content-type", "x-api-language", "x-client-type", "x-api-version", "x-client-os", "x-client-version", "X-Request-ID", "lat", "lon", "origin", "accept", "access-control-allow-origin", "authorization", "content-type"}
		headers := strings.Join(ReqHeader, ", ")
		//"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type"
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", headers)
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "360000")
		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		// yclog.Infof(c, "headers : %+v \n", headers)
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
		// 处理请求
		// fmt.Printf("\n Cors Next \n")
		c.Next()

	}

	// 	return func(c *gin.Context) {
	// 		method := c.Request.Method               //请求方法
	// 		origin := c.Request.Header.Get("Origin") //请求头部
	// 		var headerKeys []string                  // 声明请求头keys
	// 		for k, _ := range c.Request.Header {
	// 			headerKeys = append(headerKeys, k)
	// 		}
	// 		headerStr := strings.Join(headerKeys, ", ")
	// 		if headerStr != "" {
	// 			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
	// 		} else {
	// 			headerStr = "access-control-allow-origin, access-control-allow-headers"
	// 		}
	// 		if origin != "" {
	// 			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
	// 			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
	// 			//  header的类型
	// 			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
	// 			//              允许跨域设置                                                                                                      可以返回其他子段
	// 			// c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
	// 			c.Header("Access-Control-Expose-Headers", "*")
	// 			c.Header("Access-Control-Max-Age", "172800")         // 缓存请求信息 单位为秒
	// 			c.Header("Access-Control-Allow-Credentials", "true") //  跨域请求是否需要带cookie信息 默认设置为true
	// 			c.Set("content-type", "application/json")            // 设置返回格式是json
	// 		}

	// 		//放行所有OPTIONS方法
	// 		if method == "OPTIONS" {
	// 			c.AbortWithStatus(http.StatusOK)
	// 			// c.JSON(http.StatusOK, "Options Request!")
	// 		}
	// 		// 处理请求
	// 		c.Next() //  处理请求
	// 	}
}

func New(config Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestOrigin := c.Request.Header.Get("Origin")
		if len(requestOrigin) == 0 {
			return
		}

		origin, valid := validateOrigin(requestOrigin, config)
		if valid {
			if c.Request.Method == "OPTIONS" {
				valid = handlePreflight(c, config)
				c.Header("Access-Control-Allow-Origin", origin)
				c.AbortWithStatus(http.StatusNoContent)
			} else {
				valid = handleNormal(c, config)
			}
		}
		c.Header("Access-Control-Allow-Origin", origin)
	}
}

func handlePreflight(c *gin.Context, cfg Config) bool {
	for key, value := range generatePreflightHeaders(cfg) {
		c.Writer.Header()[key] = value
	}
	return true
}

func handleNormal(c *gin.Context, cfg Config) bool {
	for key, value := range generateNormalHeaders(cfg) {
		c.Writer.Header()[key] = value
	}
	return true
}

func validateOrigin(origin string, c Config) (string, bool) {
	if c.AllowAllOrigins {
		return "*", true
	}
	for _, value := range c.AllowedOrigins {
		if value == origin {
			return origin, true
		}
	}
	return "", false
}

func generateNormalHeaders(c Config) http.Header {
	headers := make(http.Header)
	if c.AllowCredentials {
		headers.Set("Access-Control-Allow-Credentials", "true")
	}
	return headers
}

func generatePreflightHeaders(c Config) http.Header {
	headers := make(http.Header)
	if c.AllowCredentials {
		headers.Set("Access-Control-Allow-Credentials", "true")
	}
	if len(c.AllowedMethods) > 0 {
		headers.Set("Access-Control-Allow-Methods", strings.Join(c.AllowedMethods, ", "))
	}
	if len(c.AllowedHeaders) > 0 {
		headers.Set("Access-Control-Allow-Headers", strings.Join(c.AllowedHeaders, ", "))
	}
	if c.MaxAge > time.Duration(0) {
		headers.Set("Access-Control-Max-Age", strconv.FormatInt(int64(c.MaxAge/time.Second), 10))
	}
	return headers
}
