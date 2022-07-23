package operation

type IDeleteUserOperation interface {
	Execute(id string) error
}
