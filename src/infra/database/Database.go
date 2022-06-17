package database

type Database interface {
	One(data any) any
}
