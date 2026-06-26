package dto

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	Senha string `json:"senha" validate:"required,min=6"`
}

type RegisterRequest struct {
	Nome   string `json:"nome" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
	Senha  string `json:"senha" validate:"required,min=6"`
	Perfil string `json:"perfil" validate:"required,oneof=administrador gerente estoquista"`
}
