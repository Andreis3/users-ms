package operationUser

type IDeleteUserOperation interface {
	Execute(id string) error
}
