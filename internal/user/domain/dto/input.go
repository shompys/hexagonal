package dto

type UserCreateInput struct {
	FirstName string
	LastName  string
	Email     string
	UserName  string
	Password  string
}
type UserUpdateInput struct {
	FirstName *string
	LastName  *string
	Email     *string
	UserName  *string
	Password  *string
}
