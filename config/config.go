package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/ranxx/observer"
	"github.com/ranxx/proxy/internal/model"
	proto "github.com/ranxx/proxy/proto/msg/v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gopkg.in/yaml.v3"
)

// Config ...
type Config struct {
	App     string   `json:"app" yaml:"app"`
	Server  Server   `json:"server" yaml:"server"`
	Auth    Auth     `json:"auth" yaml:"auth"`
	Client  Client   `json:"client" yaml:"client"`
	Tunnels []Tunnel `json:"tunnels" yaml:"tunnels"`
	Mysql   Mysql    `json:"mysql" yaml:"mysql"`
	HTTP    HTTP     `json:"http"`
	GRPC    GRPC     `json:"grpc"`
}

// Server 服务配置
type Server struct {
	IP      string `json:"ip" yaml:"ip"`
	Port    int    `json:"port" yaml:"port"`
	APIPort int    `json:"apiPort" yaml:"apiPort"`
}

// Auth auth
type Auth struct {
	Key            string `json:"key" yaml:"key"`
	ExpireInterval int64  `json:"expireInterval" yaml:"expireInterval"`
}

// Client 客户端配置
type Client struct {
	UserID int64 `json:"userId" yaml:"userId"`
	Port   int   `json:"port" yaml:"port"`
}

// Tunnel 隧道
type Tunnel struct {
	Laddr   proto.Addr        `json:"laddr" yaml:"laddr"`
	Raddr   proto.Addr        `json:"raddr" yaml:"raddr"`
	Match   model.TunnelMatch `json:"match" yaml:"match"`
	Network proto.NetworkType `json:"network" yaml:"network"`
	Account string            `json:"account" yaml:"account"`
}

// Mysql mysql
type Mysql struct {
	Dsn string `json:"dsn" yaml:"dsn"`
}

// GRPC ...
type GRPC struct {
	Port int `json:"port"`
}

// HTTP ...
type HTTP struct {
	Port int `json:"port"`
}

// var
var (
	// Cfg 配置文件
	Cfg *Config

	Obs = observer.NewObserver()

	_db *gorm.DB

	// 解析函数
	parser map[string]func(file string) *Config = map[string]func(file string) *Config{
		"yaml": ParseYamlFile,
		"json": ParseJSONFile,
		"":     ParseYamlFile,
	}
)

// ParseYamlFile 解析 yaml 配置文件
func ParseYamlFile(file string) *Config {
	body, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	Cfg = new(Config)
	if err = yaml.Unmarshal(body, Cfg); err != nil {
		panic(err)
	}
	return Cfg
}

// ParseJSONFile 解析 json 配置文件
func ParseJSONFile(file string) *Config {
	body, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	Cfg = new(Config)
	if err = json.Unmarshal(body, Cfg); err != nil {
		panic(err)
	}
	return Cfg
}

// ParseFile 解析配置文件，会根据 file 的后缀名匹配不同的解析函数
// 目前支持 yaml，json；其余默认使用 yaml 方式进行解析
func ParseFile(file string) *Config {
	// 截取后缀名
	lindex := strings.LastIndex(file, ".")
	if lindex <= 0 || lindex == len(file)-1 {
		return ParseYamlFile(file)
	}

	suffix := file[lindex+1:]
	v, ok := parser[suffix]
	if !ok {
		return ParseYamlFile(file)
	}
	return v(file)
}

// GetConfig  config
func GetConfig() *Config {
	return Cfg
}

// GetDB db
func GetDB() *gorm.DB {
	return _db
}

// GetObs obs
func GetObs() observer.Observer {
	return Obs
}

// Init init
func Init() {
	// 初始化 db
	db, err := gorm.Open(mysql.Open(Cfg.Mysql.Dsn))
	if err != nil {
		panic(err)
	}
	db.CreateBatchSize = 1024
	db = db.Debug()
	_db = db
}
