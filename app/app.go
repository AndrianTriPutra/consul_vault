package app

import (
	"context"
	"managenv/pkg/env"
)

type Setting struct {
	Conf env.Setting
	Env  env.Config
}

type apps struct {
	setting Setting
}

type Application interface {
	ReadEnv(ctx context.Context) error
	Schedulle(ctx context.Context) error
}

func NewApp(setting Setting) Application {
	return &apps{
		setting: setting,
	}
}
