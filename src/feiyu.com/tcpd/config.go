package main

type serverConfig struct {
	host string //服务地址，e.g: 192.168.93.129:10003
}

func NewServerConfig() *serverConfig {
	return &serverConfig{}
}

var defServerConfig = NewServerConfig()
