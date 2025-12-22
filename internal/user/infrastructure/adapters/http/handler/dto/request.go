package dto

type CreateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
}

type UpdateRequest struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email"`
	UserName  *string `json:"userName"`
	Password  *string `json:"password"`
}
