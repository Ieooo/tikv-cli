package config

import (
	"errors"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

var (
	configFilePath string = ".tikvcli/config"
)

type Config struct {
	CurrentTikv string       `yaml:"currentTikv"`
	Tikvs       []TikvConfig `yaml:"tikvs"`
}

type TikvConfig struct {
	Name     string   `yaml:"name"`
	Address  []string `yaml:"address"`
	Security Security `yaml:"security"`
}

type Security struct {
	ClusterSSLCA    string   `yaml:"cluster-ssl-ca"`
	ClusterSSLCert  string   `yaml:"cluster-ssl-cert"`
	ClusterSSLKey   string   `yaml:"cluster-ssl-key"`
	ClusterVerifyCN []string `yaml:"cluster-verify-cn"`
}

func (c *Config) Load() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	file, err := os.ReadFile(path.Join(homeDir, configFilePath))
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(file, c); err != nil {
		return err
	}
	return nil
}

func (c *Config) Save() error {
	configData, err := yaml.Marshal(*c)
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dirPath := path.Join(homeDir, path.Dir(configFilePath))

	if _, err := os.Stat(dirPath); err != nil && errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(dirPath, 0755); err != nil {
			return err
		}
	}
	if err := os.WriteFile(path.Join(homeDir, configFilePath), configData, 0755); err != nil {
		return err
	}
	return nil
}
