package conf

import (
	_ "embed"
	"evildoer/pkg/cryptic"
	"evildoer/utils"

	"gopkg.in/yaml.v3"
)

var confGetter utils.OnceGetter[config] = utils.MakeGetter(initConfig)

//go:embed config.yaml
var buffer []byte

func initConfig() (c config) {
	utils.AssertErr(yaml.Unmarshal(buffer, &c))
	if len(c.Secret) != 32 {
		panic("todo")
	}
	return
}

type config struct {
	Secret   cryptic.Secret `yaml:"secret"`
	BindAddr string         `yaml:"bind-addr"`
	Metadata struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Mongo struct {
		Address  string `yaml:"address"`
		Database string `yaml:"db"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"mongodb"`
}

func GetConf() *config {
	conf := confGetter.Getter()
	return &conf
}
