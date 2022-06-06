package config

type ApiSever struct {
	Ak string `yaml:"ak"`
	Sk string `yaml:"sk"`
}

var ApiSeverConfig = new(ApiSever)
