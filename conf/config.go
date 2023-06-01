package conf

import (
	"evildoer/utils"
	"os"

	"gopkg.in/yaml.v3"
)

var confGetter utils.OnceGetter[config] = utils.MakeGetter(initConfig)

func initConfig() (c config) {
	buffer, err := os.ReadFile("config.yaml")
	utils.AssertErr(err)
	utils.AssertErr(yaml.Unmarshal(buffer, &c))
	return
}

type config struct {
	BindAddr string `yaml:"bind-addr"`
	Metadata struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
}

func GetConf() *config {
	conf := confGetter.Getter()
	return &conf
}
