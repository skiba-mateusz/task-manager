package models

type UserStore interface {
	CreateUser(User) error
	GetUserByEmail(string) (User, error)
	GetUserByID(id int) (User, error) 
}


type RegisterUserPayload struct {
	Username 	string	`json:"username" validate:"required"`
	Email			string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,min=3,max=120"`	
}

type LoginUserPayload struct {
	Email 		string `json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,min=3,max=120"`	
}

type User struct {
	ID 				int64		`json:"id"`
	Username 	string	`json:"username"`
	Email			string	`json:"email"`
	Password	string	`json:"-"`
	CreatedAt string 	`json:"created_at"` 
}