package config

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	Cfg *Config

	// Used for flags.
	ConfigFile string
	Env        string
	Addr       string
	Debug      bool
	NodeID     int
)

type Config struct {
	Cors         []string
	Addr         string
	Debug        bool
	Env          string
	NodeID       int
	TokenExpired int
	DB           *Mysql
	RDB          *Redis
	Access       *Access
	Url          *Url
	Aliyun       *Aliyun
	Es           *Elasticsearch
}

type Mysql struct {
	Alias    string
	Database string
	Host     string
	Port     string
	User     string
	Password string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type Url struct {
	BlockchainUrl string
}

type Access struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
}

type Elasticsearch struct {
	URL  string
	Auth string
}

type Aliyun struct {
	AliyunurlShortmessage    string
	AliyunurlAccessKeyid     string
	AliyunurlAccessKeysecret string
	AliyunurlSignname        string
	AliyunurlEndPointName    string
	AliyunurlTemplateCode    string
}

func init() {
	pflag.StringVar(&ConfigFile, "config", "", "config file (default is ./config/config-dev.json)")
	// pflag.StringVar(&Env, "env", "dev", "environment eg:'dev','uat','pro'")
	pflag.StringVar(&Env, "env", "pro", "environment eg:'dev','uat','pro'")
	pflag.StringVar(&Addr, "addr", ":8080", "HTTP listen address")
	pflag.BoolVar(&Debug, "debug", true, "set debug logger models")
	pflag.IntVar(&NodeID, "node", 1, "service node number")
	pflag.Parse()

	viper.BindPFlag("addr", pflag.Lookup("addr"))

	viper.BindPFlag("debug", pflag.Lookup("debug"))
	viper.BindPFlag("node_id", pflag.Lookup("node"))

	initConfig()
}

func initConfig() {
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
	} else {
		fileName := getFileName(Env)
		viper.SetConfigFile(fileName)
		viper.AddConfigPath("./")
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("Using config file:", err)
			panic(2)
		}
	}

	Cfg = &Config{
		Addr: func() string {
			if Addr != ":8080" {
				return Addr
			}
			if viper.GetString("addr") != "" {
				return viper.GetString("addr")
			}
			return Addr
		}(),
		Debug:  viper.GetBool("debug"),
		DB:     getMysql(),
		RDB:    getRedis(),
		Access: getAccess(),
		Url:    getUrl(),
		Aliyun: getAliyun(),
		Env:    viper.GetString("env"),
		NodeID: NodeID,
		Cors:   viper.GetStringSlice("cors"),
		Es:     getElasticsearch(),
	}
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func getFileName(env string) string {
	filename := "config-dev"
	if Env == "uat" {
		filename = "config-uat"
	} else if Env == "pro" {
		filename = "config-pro"
	}
	filename += ".json"
	return filename
}

func getMysql() *Mysql {
	mysql := viper.GetStringMapString("mysql")
	return &Mysql{
		Alias:    mysql["alias"],
		Database: mysql["database"],
		Host:     mysql["host"],
		Port:     mysql["port"],
		User:     mysql["user"],
		Password: mysql["password"],
	}
}

func getRedis() *Redis {
	redis := viper.GetStringMapString("redis")
	return &Redis{
		Host:     redis["host"],
		Port:     redis["port"],
		Password: redis["password"],
		DB:       viper.GetInt("redis.db"),
	}
}

func getAccess() *Access {
	access := viper.GetStringMapString("access")
	return &Access{
		AccessKeyId:     access["access_key_id"],
		AccessKeySecret: access["access_key_secret"],
	}
}
func getUrl() *Url {
	url := viper.GetStringMapString("url")
	return &Url{
		BlockchainUrl: url["blockchain_url"],
	}
}

func getAliyun() *Aliyun {
	aliyun := viper.GetStringMapString("aliyun")
	return &Aliyun{
		AliyunurlShortmessage:    aliyun["aliyunurl_shortmessage"],
		AliyunurlAccessKeyid:     aliyun["aliyunurl_access_keyid"],
		AliyunurlAccessKeysecret: aliyun["aliyunurl_access_keysecret"],
		AliyunurlEndPointName:    aliyun["aliyunurl_end_point_name"],
		AliyunurlTemplateCode:    aliyun["aliyunurl_template_code"],
		AliyunurlSignname:        aliyun["aliyunurl_signname"],
	}
}

func getElasticsearch() *Elasticsearch {
	es := viper.GetStringMapString("elasticsearch")
	return &Elasticsearch{
		URL:  es["url"],
		Auth: es["auth"],
	}
}
