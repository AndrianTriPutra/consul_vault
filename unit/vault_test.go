package unit_test

import (
	"encoding/json"
	"managenv/pkg/env"
	"managenv/pkg/logger"
	"testing"
)

func Test_Vault(t *testing.T) {
	setting := env.Setting{
		Host:  "http://localhost:8200",
		Token: "hvs.tZxPAm",
		Path:  "smartcity/data/development/api",
	}

	conf, err := env.ReadEnv_Vault(setting)
	if err != nil {
		logger.Level("fatal", "ReadEnv_Vault", err.Error())
	}
	js, _ := json.MarshalIndent(conf, " ", " ")
	logger.Level("debug", "conf", string(js))
}
