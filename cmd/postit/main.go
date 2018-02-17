package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"

	"github.com/alecthomas/kingpin"
	"github.com/sirupsen/logrus"

	"github.com/wolfeidau/postit"
)

var (
	version = "unknown"

	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	url     = kingpin.Arg("url", "URL of service.").Required().String()
)

// transport is an http.RoundTripper that keeps track of the in-flight
// request and implements hooks to report HTTP tracing events.
type transport struct {
	current *http.Request
}

// RoundTrip wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {

	// req.Header.Set("Accept", "application/protobuf")

	t.current = req
	data, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, err
	}

	logrus.WithField("req", string(data)).Debug("RoundTrip")

	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	data, err = httputil.DumpResponse(res, true)
	if err != nil {
		return nil, err
	}

	logrus.WithField("res", string(data)).Debug("RoundTrip")

	return res, nil
}

// GotConn prints whether the connection has been used previously
// for the current request.
func (t *transport) GotConn(info httptrace.GotConnInfo) {
	fmt.Printf("Connection reused for %v? %v\n", t.current.URL, info.Reused)
}

func main() {
	kingpin.Version(version)
	kingpin.Parse()

	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	t := &transport{}

	client := postit.NewPostitJSONClient(*url, &http.Client{Transport: t})

	post := &postit.Post{
		Title:   "build a new thing",
		Slug:    "build_a_new_thing_2018-02-14",
		Date:    "2018-02-14T10:19:31.194Z",
		Content: "# test\n\nddd\n\ndd",
	}

	post, err := client.SavePost(context.TODO(), post)
	if err != nil {
		logrus.WithError(err).Fatal("failed to save post")
	}

	logrus.WithField("post", post).Info("SavePost")
}
