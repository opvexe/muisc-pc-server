package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"sync"
)

var global *Config
var once sync.Once

/*
	加载全局变量 单利模式只初始化一次
*/
func LoadGlobal(path string) error {
	once.Do(func() {
		c, _ := ParseToml(path)
		global = c
	})
	if global == nil {
		return errors.New("get global equal nil")
	}
	return nil
}

/*
	获取全局配置
*/
func Global() *Config {
	return global
}

/*
	解析配置文件
*/
func ParseToml(path string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(path, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

type Config struct {
	RunMode string `toml:"run_mode"`
	Version string `toml:"version"`
	AppName string `toml:"app_name"`
	Email   Email  `toml:"email"`
	HTTP    HTTP   `toml:"http"`
	Log     Log    `toml:"log"`
	CORS    CORS   `toml:"cors"`
}

/*
	HTTP http参数配置
*/
type HTTP struct {
	Host            string `toml:"host"`
	Port            int    `toml:"port"`
	ShutdownTimeout int    `toml:"shutdown_timeout"`
}

// Email 邮箱配置
type Email struct {
	Host   string `toml:"host"`
	Port   int    `toml:"port"`
	Send   string `toml:"send"`
	Pass   string `toml:"pass"`
	Recive string `toml:"recive"`
}

/*
	LOG log参数配置
*/
type Log struct {
	Level   string `toml:"level"`
	OutFile string `toml:"out_file"`
}

/*
	CORS 跨域请求配置参数
*/
type CORS struct {
	Enable bool `toml:"enable"`
	//AllowOrigins     []string `toml:"allow_origins"`
	//AllowMethods     []string `toml:"allow_methods"`
	//AllowHeaders     []string `toml:"allow_headers"`
	//AllowCredentials bool     `toml:"allow_credentials"`
	//MaxAge           int      `toml:"max_age"`
}

/*
	Redis redis配置参数
*/
type Redis struct {
	Host     string `toml:"host"`
	Password string `toml:"password"`
}

/*
	Gorm gorm配置参数
*/
type Gorm struct {
	Debug       bool   `toml:"debug"`
	TablePrefix string `toml:"table_prefix"`
	//MaxLifetime       int    `toml:"max_lifetime"`
	//MaxOpenConns      int    `toml:"max_open_conns"`
	//MaxIdleConns      int    `toml:"max_idle_conns"`
	//TablePrefix       string `toml:"table_prefix"`
	//EnableAutoMigrate bool   `toml:"enable_auto_migrate"`
}

/*
 MySQL mysql配置参数
*/
type MySQL struct {
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	User       string `toml:"user"`
	Password   string `toml:"password"`
	DBName     string `toml:"db_name"`
	Parameters string `toml:"parameters"`
}
