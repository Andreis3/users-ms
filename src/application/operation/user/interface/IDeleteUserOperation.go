package operation_interface_user

type IDeleteUserOperation interface {
	Execute(id string) error
}
