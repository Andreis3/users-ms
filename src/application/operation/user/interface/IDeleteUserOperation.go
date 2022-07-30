package operation_interface_user

//go:generate mockgen -package=mocks -destination ../../../../../tests/unit/mocks/DeleteUserOperationMock.go github.com/andreis3/users-ms/src/application/operation/user/interface IDeleteUserOperation
type IDeleteUserOperation interface {
	Execute(id string) error
}
