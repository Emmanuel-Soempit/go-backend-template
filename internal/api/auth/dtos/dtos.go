package dtos

type LoginUserPayload struct {
	Email    string `json:"email" xml:"email" form:"email"`
	Password string `json:"password" xml:"password" form:"password"`
}

type UserDTO struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type LoginResponse struct {
	User  UserDTO `json:"user" xml:"user" form:"user"`
	Token string  `json:"token" xml:"token" form:"token"`
}

type RegisterUserPayload struct {
	Firstname string `json:"firstname" xml:"firstname" form:"firstname"`
	Lastname  string `json:"lastname" xml:"lastname" form:"lastname"`
	Email     string `json:"email" xml:"email" form:"email"`
	Password  string `json:"password" xml:"password" form:"password"`
}
