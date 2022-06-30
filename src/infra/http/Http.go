package http

type Http interface {
	On(method string, path string, fn func(c any))
	Filter(fn func(c any) bool)
	Listen(port string)
}
