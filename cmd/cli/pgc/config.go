package pgc

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Postgresql struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"postgresql"`
	Databases []struct {
		Name  string `yaml:"name"`
		Owner string `yaml:"owner"`
	} `yaml:"databases"`
	Users []struct {
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
	} `yaml:"users"`
	Privileges []struct {
		Database   string `yaml:"database"`
		User       string `yaml:"user"`
		Privileges string `yaml:"privileges"`
	} `yaml:"privileges"`
	PgHba []struct {
		Type     string `yaml:"type"`
		Database string `yaml:"database"`
		User     string `yaml:"user"`
		Address  string `yaml:"address"`
		Method   string `yaml:"method"`
	} `yaml:"pg_hba"`
	PgHbaConfPath  string `yaml:"pg_hba_conf_path"`
	RestartCommand string `yaml:"restart_cmd"`
}

// Load the config from a file
func LoadConfig(filename string) (Config, error) {
	config := Config{}
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	return config, err
}
