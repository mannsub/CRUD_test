package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type GetUserResponse struct {
	*APIResponse
	Users []*User `json:"result"`
}

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

func (c *CreateRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.Age,
	}
}

type CreateUserResponse struct {
	*APIResponse
}

type UpdateRequest struct {
	Name       string `json:"name" binding:"required"`
	UpdatedAge int64  `json:"updatedAge" binding:"required"`
}

type UpdateUserResponse struct {
	*APIResponse
}

type DeleteRequest struct {
	Name string `json:"name" binding:"required"`
}

func (d *DeleteRequest) ToUser() *User {
	return &User{
		Name: d.Name,
	}
}

type DeleteUserResponse struct {
	*APIResponse
}
