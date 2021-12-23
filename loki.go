//go:build bd_all || bd_loki || loki

package blackdatura

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type lokiRequest struct {
	Streams []lokiRequestStream `json:"streams"`
}

type lokiRequestStream struct {
	Stream map[string]interface{}   `json:"stream"`
	Values []lokiRequestStreamValue `json:"values"`
}

type lokiRequestStreamValue [2]interface{}

func newLokiRequest() *lokiRequest {
	var arg lokiRequest
	arg.Streams = make([]lokiRequestStream, 1)
	arg.Streams[0].Stream = make(map[string]interface{})
	return &arg
}

type LokiSink struct {
	Client  *redis.Client
	Key     []string
	apiAddr string

	httpClient *http.Client
}

func NewLoki() zap.Sink {
	var i LokiSink
	i.httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxConnsPerHost:       100,
			MaxIdleConnsPerHost:   100,
		},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       30 * time.Second,
	}

	return i
}

func (r LokiSink) Sink(*url.URL) (zap.Sink, error) { return r, nil }

// Close implement zap.Sink func Close
func (r LokiSink) Close() error { return r.Client.Close() }

// Write implement zap.Sink func Write
func (r LokiSink) Write(b []byte) (n int, err error) {
	var (
		req = newLokiRequest()
		arg map[string]interface{}
	)
	if err = json.Unmarshal(b, &arg); err != nil {
		return
	}

	if str, ok := arg["ts"].(string); ok {
		var t time.Time
		if t, err = time.Parse("2006-01-02T15:04:05.999-0700", str); err != nil {
			return
		}
		req.Streams[0].Values = []lokiRequestStreamValue{{t.UnixNano(), string(b)}}
	}

	for _, v := range r.Key {
		if data, has := arg[v]; has {
			req.Streams[0].Stream[v] = data
		}
	}

	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(req); err != nil {
		return
	}

	var hr *http.Request
	if hr, err = http.NewRequest("POST", r.apiAddr, &buf); err != nil {
		return
	}

	var resp *http.Response
	if resp, err = r.httpClient.Do(hr); err == nil {
		_ = resp.Body.Close()
	}

	return len(b), err
}

// Sync implement zap.Sink func Sync
func (r LokiSink) Sync() error { return nil }

func (r LokiSink) String() string { return "loki" }
