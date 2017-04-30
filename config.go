package main

import "os"

var (
	Conf = new(config)
)

type config struct {
	Protocol string
	Bind     string
}

func InitCfg() {
	Conf.Protocol = env("PROTOCOL", "tcp")
	Conf.Protocol = env("BIND", ":8801")
}

func env(key, defaultValue string) string {
	v := os.Getenv(key)
	if key == "" {
		return defaultValue
	}
	return v
}
