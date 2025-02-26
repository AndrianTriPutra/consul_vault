package app

import (
	"context"
	"encoding/json"
	"errors"
	"managenv/pkg/env"
	"managenv/pkg/logger"
	"os"
	"time"
)

func (a apps) ReadEnv(ctx context.Context) error {
	ticker := time.NewTicker(a.setting.Env.Interval.Environment)
	for {
		select {
		case <-ticker.C:
			environment, err := env.ReadEnv_Consul(a.setting.Conf)
			if err != nil {
				logger.Level("error", "ReadEnv_Consul", err.Error())
			} else {
				jsNow, _ := json.MarshalIndent(environment, " ", " ")
				logger.Level("debug", "config Now", string(jsNow))

				jsExist, _ := json.MarshalIndent(a.setting.Env, " ", " ")
				if string(jsNow) != string(jsExist) {
					logger.Level("info", "compare", "config now and exist its different, RESTART service")
					os.Exit(1)
				}
			}

		case <-ctx.Done():
			err := errors.New("ctx is done:" + ctx.Err().Error())
			return err
		}
	}
}
