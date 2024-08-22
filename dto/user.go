package dto

type UserRegisterRequest struct {
	Name     string `json:"name" example:"John Doe" valid:"required~Name cannot be empty"`
	Email    string `json:"email" example:"john.doe@example.com" valid:"required~Email is required,email~Invalid email format"`
	Password string `json:"password" example:"password123" valid:"required~Password is required,minstringlength(8)~Password must be at least 8 characters long"`
	Role     string `json:"role" example:"admin" valid:"required~Role cannot be empty,in(admin|user)~Invalid role. Role must be either 'admin' or 'user'"`
}

type UserRegisterResponse struct {
	Result     string `json:"result" example:"SUCCESS"`
	StatusCode int    `json:"statusCode" example:"201"`
	Message    string `json:"message" example:"User registered successfully"`
}

type UserLoginRequest struct {
	Email    string `json:"email" example:"john.doe@example.com" valid:"required~Email is required,email~Invalid email format"`
	Password string `json:"password" example:"password123" valid:"required~Password is required"`
}

type UserLoginResponse struct {
	Result     string    `json:"result" example:"SUCCESS"`
	StatusCode int       `json:"statusCode" example:"200"`
	Message    string    `json:"message" example:"User logged in successfully"`
	Data       TokenData `json:"data"`
}

type TokenData struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG4uZG9lQGV4YW1wbGUuY29tIiwiZXhwIjoxNzI0MjUxMTQyLCJpZCI6MSwicm9sZSI6ImFkbWluIn0.itsoI0KJmw9UiPL132kzzTWptIfN9K2SgimRCnQIXN8"`
}
