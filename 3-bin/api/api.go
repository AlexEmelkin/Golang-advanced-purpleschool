package api

import "HomeTask/3-struct/config"

func GetConfig() *config.Config {

	conf := config.NewConfig()
	return conf
}
