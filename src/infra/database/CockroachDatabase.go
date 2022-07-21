package database

import (
	"sync"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
)

var (
	instance *CockroachDatabase
	once     sync.Once
)

type CockroachDatabase struct {
	session *db.Session
}

func GetInstance() IDatabase {
	once.Do(func() {
		settings := cockroachdb.ConnectionURL{
			Host:     "localhost",
			Database: "users_db",
			User:     "root",
			Password: "",
			Options: map[string]string{
				"sslmode": "disable",
			},
		}
		session, err := cockroachdb.Open(settings)
		if session == nil {
			session, err = cockroachdb.Open(settings)
		}

		if err != nil {
			panic(err)
		}
		session.SetConnMaxIdleTime(10)
		session.SetMaxOpenConns(10)
		instance = &CockroachDatabase{
			session: &session,
		}
	})
	return instance
}

func (c *CockroachDatabase) Session() *db.Session {
	return c.session
}
