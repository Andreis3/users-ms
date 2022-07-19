package database

type IDatabase interface {
	One(data any) any
	Many(data any) any
	None(query string)
}
