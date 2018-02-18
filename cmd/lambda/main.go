package main

import (
	"context"

	"github.com/apex/gateway"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/sirupsen/logrus"
	"github.com/wolfeidau/postit"
)

type postitStore struct{}

func (ps *postitStore) SavePost(c context.Context, post *postit.Post) (*postit.Post, error) {
	logrus.WithField("post", post).Info("SavePost")
	return post, nil
}

func main() {
	xray.Configure(xray.Config{LogLevel: "trace"})
	server := postit.NewPostitServer(&postitStore{}, nil)
	logrus.Fatal(gateway.ListenAndServe(":3000", xray.Handler(xray.NewFixedSegmentNamer("Postit"), server)))
}
