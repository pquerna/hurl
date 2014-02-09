/**
 *  Copyright 2014 Paul Querna
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package http

import (
	"github.com/pquerna/hurl/common"
	"github.com/pquerna/hurl/workers"
	"io/ioutil"
	nhttp "net/http"
	"net/url"
	"strconv"
)

func init() {
	workers.Register("http", NewTask)
}

type Task struct {
	conf      *common.HttpConfig
	transport *nhttp.Transport
	URL       *url.URL
}

func NewTask(c common.ConfigGetter) workers.WorkerTask {
	conf := c.GetHttpConfig()
	if conf == nil {
		panic("Invalid Configuration object for http/1 worker")
	}

	// TODO(pquerna): TLS transport configuration
	trans := &nhttp.Transport{
		/* TODO: conf.Compression */
		DisableCompression:  true,
		DisableKeepAlives:   !conf.Keepalive,
		MaxIdleConnsPerHost: conf.Concurrency,
	}
	//	client := &nhttp.Client{
	//		Transport: trans,
	//		CheckRedirect: func(req *nhttp.Request, via []*nhttp.Request) error {
	//			return errors.New("do not follow redirects")
	//		},
	//	}

	url, err := url.Parse(conf.Url)
	if err != nil {
		panic("Broken URL")
	}

	return &Task{conf: conf, transport: trans, URL: url}
}

func (t *Task) Request(requestId string) *nhttp.Request {
	req := &nhttp.Request{
		Method: t.conf.Method,
		URL:    t.URL,
		// TODO(pquerna): HTTP/2
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(nhttp.Header),
		Body:       nil,
		Host:       t.URL.Host,
	}
	// TODO: add custom headers
	req.Header.Set("User-Agent", "hurl/1 http load tester; https://github.com/pquerna/hurl")
	req.Header.Set("Request-Id", requestId)
	return req
}

func (t *Task) Work(rv *common.Result) error {
	// TOOD: capture true bytes across wire.
	req := t.Request(rv.Id)
	resp, err := t.transport.RoundTrip(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// TOOD: extract any interesting resp.Headers?
	rv.Meta["Scheme"] = t.URL.Scheme
	rv.Meta["StatusCode"] = strconv.Itoa(resp.StatusCode)
	rv.Meta["Proto"] = resp.Proto
	rv.Meta["Server"] = resp.Header.Get("Server")
	rv.Meta["Etag"] = resp.Header.Get("Etag")
	rv.Meta["Cache-Control"] = resp.Header.Get("Cache-Control")
	rv.Meta["Vary"] = resp.Header.Get("Vary")
	rv.Meta["Content-Type"] = resp.Header.Get("Content-Type")

	// TODO: make a /dev/null sink and just count the bytes.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	rv.Metrics["BodyLength"] = float64(len(body))
	return nil
}
