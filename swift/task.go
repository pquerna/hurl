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

package swift

import (
	"fmt"
	"github.com/ncw/swift"
	"github.com/pquerna/hurl/common"
	"github.com/pquerna/hurl/workers"
	"sync"
)

func init() {
	workers.Register("swift", NewTask)
}

type Task struct {
	conf *common.SwiftConfig
	conn *swift.Connection
}

func (s *Task) Config() common.ConfigGetter {
	return s.conf
}

func (s *Task) Insert(table string, key string, values common.DatastoreObj) error {
	data, err := common.DatastoreValuesToBytes(values)
	if err != nil {
		fmt.Printf("DatastoreValuesToBytes: %v\n", err)
		return err
	}

	err = s.conn.ObjectPutBytes(table, key, data, "")
	if err != nil {
		fmt.Printf("ObjectPutBytes: %v: %v:%v->%v\n", err, table, key, data)
		return err
	}
	return nil
}

func (s *Task) Read(table string, key string) (common.DatastoreObj, error) {
	data, err := s.conn.ObjectGetBytes(table, key)

	if err != nil {
		return nil, err
	}

	v, err := common.DatastoreBytesToValues(data)

	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Task) Update(table string, key string, values common.DatastoreObj) error {
	rv, err := s.Read(table, key)

	if err != nil {
		return err
	}

	for k, v := range values {
		rv[k] = v
	}

	err = s.Insert(table, key, values)
	if err != nil {
		return err
	}

	return nil
}

func (s *Task) Scan(table string, startKey string, count int) ([]common.DatastoreObj, error) {
	p := &swift.ObjectsOpts{
		Limit:  count,
		Marker: startKey,
	}

	objs, err := s.conn.ObjectNames(table, p)
	if err != nil {
		return nil, err
	}
	// TODO(pquerna): respect concurrency.? fml.
	// TODO(pquerna): Auth Token expiration isn't go-routine safe in swift client.
	var wg sync.WaitGroup
	var m sync.Mutex
	rv := make([]common.DatastoreObj, len(objs))
	errors := make([]error, 0)
	for k, v := range objs {
		wg.Add(1)
		go func() {
			v, err := s.Read(table, v)
			m.Lock()
			defer m.Unlock()
			if err != nil {
				errors = append(errors, err)
			} else {
				rv[k] = v
			}
		}()
	}
	wg.Wait()

	if len(errors) != 0 {
		// TODO: hrm. wish this was better?
		return nil, errors[0]
	}

	return rv, nil
}

func (s *Task) Delete(table string, key string) error {
	return s.conn.ObjectDelete(table, key)
}

var authTokenCache string = ""
var storageUrlCache string = ""

func NewTask(ui common.UI) (workers.WorkerTask, error) {
	c := ui.ConfigGet()
	conf := c.GetSwiftConfig()
	if conf == nil {
		panic("Invalid Configuration object for swift worker")
	}

	conn := swift.Connection{
		UserName: conf.Username,
		ApiKey:   conf.ApiKey,
		AuthUrl:  conf.AuthUrl,
		// https://auth.api.rackspacecloud.com/v1.0
		Region:    conf.Region,
		UserAgent: fmt.Sprintf("hurl/1 http load tester; https://github.com/pquerna/hurl;  username=%s", conf.Username),
	}

	if authTokenCache != "" && storageUrlCache != "" {
		conn.AuthToken = authTokenCache
		conn.StorageUrl = storageUrlCache
	} else {
		err := conn.Authenticate()
		if err != nil {
			return nil, err
		}
		authTokenCache = conn.AuthToken
		storageUrlCache = conn.StorageUrl
	}

	return &Task{
		conf: conf,
		conn: &conn}, nil
}

func (t *Task) Work(rv *common.Result) error {
	return common.DatastoreWork(t, rv)
}
