package cache

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/dgraph-io/badger/v3"
	"github.com/gomodule/redigo/redis"
)

var trc RedisCache  //TestRedisCache
var tbc BadgerCache //TestBadgerCache

func TestMain(m *testing.M) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	pool := redis.Pool{
		MaxIdle:     50,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", s.Addr())
		},
	}

	trc.Conn = &pool
	trc.Prefix = "test-celeritas"

	defer trc.Conn.Close()

	//BADGER
	badgerFolder := "./testdata/tmp/badger"
	//remove any existing test data
	_ = os.RemoveAll(badgerFolder)
	//create the badger database
	if _, err := os.Stat("./testdata/tmp"); os.IsNotExist(err) {
		err := os.Mkdir("./testdata/tmp", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = os.Mkdir(badgerFolder, 0755)
	if err != nil {
		log.Fatal(err)
	}

	db, _ := badger.Open(badger.DefaultOptions(badgerFolder))
	tbc.Conn = db

	os.Exit(m.Run())
}
