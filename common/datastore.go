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

package common

import (
	"bytes"
	"encoding/gob"
)

type DatastoreObj map[string]string
type Datastore interface {
	Insert(table string, key string, value DatastoreObj) error
	Read(table string, key string) (DatastoreObj, error)
	Update(table string, key string, value DatastoreObj) error
	Scan(table string, startKey string, count int) ([]DatastoreObj, error)
	Delete(table string, key string) error
}

func DatastoreValuesToBytes(values map[string]string) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(values)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DatastoreBytesToValues(input []byte) (map[string]string, error) {
	buf := bytes.NewBuffer(input)
	enc := gob.NewDecoder(buf)
	rv := make(map[string]string)
	err := enc.Decode(&rv)

	if err != nil {
		return nil, err
	}

	return rv, nil
}
