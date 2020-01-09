package config

import "sync"

var (
	once sync.Once
	Conf *Config
)

type Config struct {


}

func init()  {
	Init()
}

func Init()  {
	once.Do(func() {


	})
}