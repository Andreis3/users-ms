package http

type Http interface {
	On(method string, url string, fn func(c any))
	Filter(filter func(c any) bool)
	Listen(port string)
}
