package config

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	type BaseConfig struct {
		ServiceConfiguration  ServiceConfiguration
		PostgresConfiguration PostgresConfiguration
	}

	var baseConfig BaseConfig
	err := InitConfiguration("test_config", []string{
		".",
	}, &baseConfig)
	if err != nil {
		panic(err)
	}

	indent, err := json.MarshalIndent(baseConfig, "", " ")
	if err == nil {
		fmt.Println(string(indent))
	}
}

func TestReadEnvConfig(t *testing.T) {
	type BaseConfig struct {
		ServiceConfiguration  ServiceConfiguration
		PostgresConfiguration PostgresConfiguration
	}

	// 通过env 覆盖配置文件
	os.Setenv("ServiceConfiguration.Port", "8642")

	var baseConfig BaseConfig
	err := InitConfiguration("test_config", []string{
		".",
	}, &baseConfig)
	if err != nil {
		panic(err)
	}

	indent, err := json.MarshalIndent(baseConfig, "", " ")
	if err == nil {
		fmt.Println(string(indent))
	}
}
