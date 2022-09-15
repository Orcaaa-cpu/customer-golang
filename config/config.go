package config

import "github.com/tkanos/gonfig"

// karna menggunakan config agak ribet saya ganti ke yg lebih simple
// yaitu menggunakan db

type Configuration struct {
	USERNAME string
	PASSWORD string
	PORT     string
	HOST     string
	NAME     string
}

func GetConfig() Configuration {

	conf := Configuration{}

	gonfig.GetConf("config/config.json", &conf)
	return conf
}

// karna menggunakan config agak ribet saya ganti ke yg lebih simple
