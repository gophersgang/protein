// Copyright © 2016 Zenly <hello@zen.ly>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protein

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gocql/gocql"
	"github.com/rainycape/memcache"
	"github.com/stretchr/testify/assert"
	"github.com/znly/protein/protoscan"
)

// -----------------------------------------------------------------------------

func TestTranscoder_Helpers_Memcached(t *testing.T) {
	c, err := memcache.New("localhost:11211")
	assert.Nil(t, err)
	assert.NotNil(t, c)
	defer c.Close()

	// clear the store, just in case
	assert.Nil(t, c.Flush(-1))

	// create Transcoder and push all local schemas via user-defined setter
	trc, err := NewTranscoder(context.Background(),
		protoscan.SHA256, "PROT-",
		TranscoderOptGetter(NewTranscoderGetterMemcached(c)),
		TranscoderOptSetter(NewTranscoderSetterMemcached(c)),
	)
	assert.Nil(t, err)
	assert.NotNil(t, trc)

	testTranscoder_Helpers_common(t, trc)
}

func TestTranscoder_Helpers_Redis(t *testing.T) {
	p := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL("redis://localhost:6379/0")
			assert.NotNil(t, c)
			assert.Nil(t, err)
			return c, nil
		},
	}
	defer p.Close()

	// clear the store, just in case
	c := p.Get()
	_, err := c.Do("FLUSHALL")
	c.Close()
	assert.Nil(t, err)

	// create Transcoder and push all local schemas via user-defined setter
	trc, err := NewTranscoder(context.Background(),
		protoscan.SHA256, "PROT-",
		TranscoderOptGetter(NewTranscoderGetterRedis(p)),
		TranscoderOptSetter(NewTranscoderSetterRedis(p)),
	)
	assert.Nil(t, err)
	assert.NotNil(t, trc)

	testTranscoder_Helpers_common(t, trc)
}

func TestTranscoder_Helpers_Cassandra(t *testing.T) {
	cluster := gocql.NewCluster("localhost:9042")
	cluster.Consistency = gocql.LocalQuorum
	cluster.NumConns = 1
	gocql.TimeoutLimit = 10
	cluster.ConnectTimeout = time.Second * 10
	cluster.Timeout = time.Second * 10

	// HACK: Connect to standard 'system' keyspace in order to create
	//       the keyspace reserved for tuyauDB's blobs.
	cluster.Keyspace = "system"
	tmp, err := cluster.CreateSession()
	assert.Nil(t, err)
	assert.NotNil(t, tmp)

	// clear the store, just in case
	keyspaceDropQuery := `DROP KEYSPACE IF EXISTS blobs;`
	assert.Nil(t, tmp.Query(keyspaceDropQuery).Exec())

	keyspaceCreateQuery := `
	CREATE KEYSPACE IF NOT EXISTS blobs
	WITH REPLICATION = {
		'class':       'NetworkTopologyStrategy',
		'datacenter1': 1
	};`
	err = tmp.Query(keyspaceCreateQuery).Exec()
	tmp.Close()
	assert.Nil(t, err)

	cluster.Keyspace = "blobs"
	s, err := cluster.CreateSession()
	assert.Nil(t, err)
	assert.NotNil(t, s)
	defer s.Close()

	tableCreateQuery := `
	CREATE TABLE IF NOT EXISTS blobs (key ascii, data blob, PRIMARY KEY (key));`
	assert.Nil(t, s.Query(tableCreateQuery).Exec())

	// create Transcoder and push all local schemas via user-defined setter
	trc, err := NewTranscoder(context.Background(),
		protoscan.SHA256, "PROT-",
		TranscoderOptGetter(NewTranscoderGetterCassandra(s, "blobs", "key", "data")),
		TranscoderOptSetter(NewTranscoderSetterCassandra(s, "blobs", "key", "data")),
	)
	assert.Nil(t, err)
	assert.NotNil(t, trc)

	testTranscoder_Helpers_common(t, trc)
}

// -----------------------------------------------------------------------------

func testTranscoder_Helpers_common(t *testing.T, trc *Transcoder) {
	payload, err := trc.Encode(_transcoderTestSchemaXXX)
	assert.Nil(t, err)
	assert.NotNil(t, payload)

	// empty local SchemaMap struct-types cache
	trc.sm = NewSchemaMap()
	trc.typeCache = map[string]reflect.Type{}

	// decoding will have to use the user-specified getter to find the schemas
	v, err := trc.Decode(payload)
	assert.Nil(t, err)
	assertFieldValues(t, reflect.ValueOf(_transcoderTestSchemaXXX), v)
}