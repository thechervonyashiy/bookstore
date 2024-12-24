package config

import "time"

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage"`

	// Анонимное поле HTTPServer(так как необходимо отношение композиция - строгая зависимость между объектами. 
	// При этом подходе объекты существуют только в контексте друг друга, один объект является неотъемлемой частью другого.
	HTTPServer  `yaml:"http_server"` 
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func MustLoad() *Config {

}