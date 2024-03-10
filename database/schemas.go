package database

type CreateUserSchema struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserSchema struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateBookSchema struct {
	Title  string `json:"title" binding:"required"`
}

type UpdateBookSchema struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
