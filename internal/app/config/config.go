package config

import "github.com/BurntSushi/toml"

var global *Config

/*
	加载全局变量
*/
func LoadGlobal(path string) error {
	c, err := ParseToml(path)
	if err != nil {
		return err
	}
	global = c
	return nil
}

/*
	获取全局配置
*/
func Global() *Config {
	if global == nil {
		return &Config{}
	}
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
	Log     Log    `toml:"log"`
	HTTP    HTTP   `toml:"http"`
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

/*
	LOG log参数配置
*/
type Log struct {
	Level int `toml:"level"`
}

/*
	Redis redis配置参数
*/
type Redis struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
}

/*
	Gorm gorm配置参数
*/
type Gorm struct {
	Debug             bool   `toml:"debug"`
	DBType            string `toml:"db_type"`
	MaxLifetime       int    `toml:"max_lifetime"`
	MaxOpenConns      int    `toml:"max_open_conns"`
	MaxIdleConns      int    `toml:"max_idle_conns"`
	TablePrefix       string `toml:"table_prefix"`
	EnableAutoMigrate bool   `toml:"enable_auto_migrate"`
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

/*
	CORS 跨域请求配置参数
*/
type CORS struct {
	Enable           bool     `toml:"enable"`
	AllowOrigins     []string `toml:"allow_origins"`
	AllowMethods     []string `toml:"allow_methods"`
	AllowHeaders     []string `toml:"allow_headers"`
	AllowCredentials bool     `toml:"allow_credentials"`
	MaxAge           int      `toml:"max_age"`
}
