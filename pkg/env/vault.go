package env

import (
	"encoding/json"
	"errors"
	"fmt"
	"managenv/pkg/logger"
	"time"

	"github.com/hashicorp/vault/api"
)

func ReadEnv_Vault(c Setting) (Config, error) {
	var conf Config

	// Config Vault Client
	config := api.DefaultConfig()
	config.Address = c.Host

	// create Client
	client, err := api.NewClient(config)
	if err != nil {
		errN := errors.New("[NewClient]" + fmt.Sprintf("create client vault:%s", err.Error()))
		return conf, errN
	}

	// Set Token Authentication
	client.SetToken(c.Token)

	// get Secret from Vault
	secret, err := client.Logical().Read(c.Path)
	if err != nil {
		errN := errors.New("[Logical.Read]" + fmt.Sprintf("read secret:%s", err.Error()))
		return conf, errN
	}

	// check secret
	if secret == nil || secret.Data["data"] == nil {
		errN := errors.New("[check]" + "secret not found or nil")
		return conf, errN
	}

	// get data from response with format JSON
	js, err := json.Marshal(secret.Data["data"])
	if err != nil {
		errN := errors.New("[Marshal]" + fmt.Sprintf("change secret to format json:%s", err.Error()))
		return conf, errN
	}

	// Unmarshal JSON to struct
	err = json.Unmarshal(js, &conf)
	if err != nil {
		errN := errors.New("[unmarshal]" + fmt.Sprintf("change json to struct:%s", err.Error()))
		return conf, errN
	}

	//logger access
	logger.Load(conf.Internal.Logger)

	logger.Trace(" ============= [Secret from Vault] ============= ", "")
	logger.Trace(" ======= [internal] ======= ", "")
	logger.Trace("Logger       :", conf.Internal.Logger)
	logger.Trace("Core         :", conf.Internal.Core)

	logger.Trace(" ======= [interval] ======= ", "")
	conf.Interval.Environment *= time.Second
	conf.Interval.Schedulle *= time.Second
	logger.Trace("Environment  :", conf.Interval.Environment)
	logger.Trace("Schedulle    :", conf.Interval.Schedulle)

	logger.Trace(" =======  [database] ======= ", "")
	logger.Trace("Host         :", conf.Database.Host)
	logger.Trace("Port         :", conf.Database.Port)
	logger.Trace("Name         :", conf.Database.Name)
	logger.Trace("User         :", conf.Database.User)
	logger.Trace("Pass         :", conf.Database.Pass)

	logger.Trace(" =======   [rabbit]  ======= ", "")
	logger.Trace("Host         :", conf.Rabbit.Host)
	logger.Trace("Tag          :", conf.Rabbit.Tag)
	logger.Trace("Queue        :", conf.Rabbit.Que)
	logger.Trace("Key          :", conf.Rabbit.Key)

	return conf, nil

}
