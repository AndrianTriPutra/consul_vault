package app

import (
	"context"
	"errors"
	"fmt"
	"managenv/pkg/logger"
	"time"
)

func (a apps) Schedulle(ctx context.Context) error {
	counter := uint(0)
	ticker := time.NewTicker(a.setting.Env.Interval.Schedulle)
	for {
		select {
		case <-ticker.C:
			counter++
			ts := time.Now().Format(time.RFC3339)
			logger.Level("info", "Schedulle", fmt.Sprintf("[%s] counter:%d", ts, counter))

		case <-ctx.Done():
			err := errors.New("ctx is done:" + ctx.Err().Error())
			return err
		}
	}
}
