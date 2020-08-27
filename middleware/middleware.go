package middleware

// https://github.com/micro/go-plugins/tree/master/wrapper

import (
	"context"

	"github.com/micro/go-micro/v3/server"
	"github.com/sirupsen/logrus"
)

type Wrapper struct {
}

func (wp *Wrapper) NewWrapHandler() server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			logrus.Infof("Before serving request method: %+v", req)
			err := fn(ctx, req, rsp)
			logrus.Infof("After serving request")
			return err
		}
	}
}

func (wp *Wrapper) NewWrapSubscriber() server.SubscriberWrapper {
	return func(fn server.SubscriberFunc) server.SubscriberFunc {
		return func(ctx context.Context, msg server.Message) error {
			logrus.Infof("Before sub request method: %+v", msg)
			err := fn(ctx, msg)
			logrus.Infof("After sub request")
			return err
		}
	}
}
