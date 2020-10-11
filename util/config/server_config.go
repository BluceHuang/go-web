package config

import (
	"goweb/log"
	"goweb/model"
	"goweb/util"
	"path/filepath"
)

const (
	defaultConfigPath = "./config/config.yaml"
)

var serverConfig model.ServerConfig

// ServerConfig return global server config
func ServerConfig() model.ServerConfig {
	return serverConfig
}

func init() {
	var err error
	var path, dir string
	dir, err = filepath.Abs(filepath.Dir("."))
	if err != nil {
		log.Fatal(err)
	}
	path = filepath.Join(dir, defaultConfigPath)
	if serverConfig, err = util.GetHelper().ParseServerConfig(path); err != nil {
		log.Fatalf("parse config fail: %v", err)
	}
}
