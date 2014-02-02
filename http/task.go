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
	nhttp "net/http"
	"net/url"
)

func init() {
	workers.Register("http", NewTask)
}

type Task struct {
	conf   *common.HttpConfig
	client *nhttp.Client
	URL    *url.URL
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
	client := &nhttp.Client{Transport: trans}

	url, err := url.Parse(conf.Url)
	if err != nil {
		panic("Broken URL")
	}

	return &Task{conf: conf, client: client, URL: url}
}

func (t *Task) Request() *nhttp.Request {
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
	return req
}

func (t *Task) Work(rv *common.Result) error {
	t.client.Do(t.Request())
	return nil
}
