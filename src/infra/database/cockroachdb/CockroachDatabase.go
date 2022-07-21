package cockroachdb

import (
	"sync"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"

	"github.com/andreis3/users-ms/src/infra/database"
)

var (
	instance *CockroachDatabase
	once     sync.Once
)

type CockroachDatabase struct {
	session *db.Session
}

func GetInstance() database.IDatabase {
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
