package database

import (
	"database/sql"
	"sync"
)

var (
	instance *CockroachDatabase
	once     sync.Once
)

type CockroachDatabase struct {
	db *sql.DB
}

func GetInstance() IDatabase {
	once.Do(func() {
		conn, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:26257/postgres")
		if err != nil {
			panic(err)
		}
		instance = &CockroachDatabase{
			db: conn,
		}
	})
	return instance
}

func (c *CockroachDatabase) One(data any) any {
	query := "SELECT * FROM users"
	c.db.QueryRow(query).Scan(data)
	return data
}

func (c *CockroachDatabase) Many(data any) any {
	query := "SELECT * FROM users"
	rows, err := c.db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(data)
	}
	return data
}

func (c *CockroachDatabase) None(query string) {
	_, err := c.db.Query(query)
	if err != nil {
		panic(err)
	}
}
