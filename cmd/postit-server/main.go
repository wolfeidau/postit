package main

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/wolfeidau/postit"
)

type postitStore struct{}

func (ps *postitStore) SavePost(c context.Context, post *postit.Post) (*postit.Post, error) {
	logrus.WithField("post", post).Info("SavePost")
	return post, nil
}

func main() {
	server := postit.NewPostitServer(&postitStore{}, nil)
	logrus.Fatal(http.ListenAndServe(":3000", server))
}
