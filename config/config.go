package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	Port              int32
	MidtransServerKey string
	MidtransClientKey string
}

var Conf Configuration

func LoadConfiguration(path string) {
	Conf = Configuration{}
	err := gonfig.GetConf(path, &Conf)
	if err != nil {
		panic(err)
	}
}
