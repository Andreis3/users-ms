package database

import "github.com/upper/db/v4"

type IDatabase interface {
	Session() *db.Session
}
