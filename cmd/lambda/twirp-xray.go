package main

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/twitchtv/twirp"
)

// NewXrayServerHooks twirp server hooks for AWS Xray
func NewXrayServerHooks() *twirp.ServerHooks {

	hooks := &twirp.ServerHooks{}

	hooks.RequestRouted = func(ctx context.Context) (context.Context, error) {

		method, ok := twirp.MethodName(ctx)
		if !ok {
			return ctx, nil
		}

		logrus.WithField("method", method).Info("BeginSegment")

		if xray.GetSegment(ctx).InProgress {
			ctx, _ = xray.BeginSubsegment(ctx, method)
		}

		return ctx, nil
	}

	hooks.ResponseSent = func(ctx context.Context) {
		method, ok := twirp.MethodName(ctx)
		if !ok {
			return
		}
		logrus.WithField("method", method).Info("Close")
		seg := xray.GetSegment(ctx)
		if seg != nil {
			seg.Close(nil)
		}
	}

	return hooks
}
