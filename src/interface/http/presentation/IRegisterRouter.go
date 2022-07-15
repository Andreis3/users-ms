package presentation

type IRegisterRouter interface {
	Register(router []map[string]interface{}) []map[string]interface{}
}
