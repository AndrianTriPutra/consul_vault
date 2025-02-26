package unit_test

import (
	"encoding/json"
	"managenv/pkg/env"
	"managenv/pkg/logger"
	"testing"
)

func Test_Consul(t *testing.T) {
	setting := env.Setting{
		Host:  "http://localhost:8500/v1/kv/",
		Token: "",
		Path:  "smartcity/development/api",
	}

	conf, err := env.ReadEnv_Consul(setting)
	if err != nil {
		logger.Level("fatal", "ReadEnv_Consul", err.Error())
	}
	js, _ := json.MarshalIndent(conf, " ", " ")
	logger.Level("debug", "conf", string(js))
}
