package services

import (
	"errors"

	"github.com/seu-usuario/estoque-api/dto"
	"github.com/seu-usuario/estoque-api/models"
	"github.com/seu-usuario/estoque-api/repositories"
	"github.com/seu-usuario/estoque-api/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repo *repositories.UsuarioRepository
}

func NewAuthService() *AuthService {
	return &AuthService{repo: &repositories.UsuarioRepository{}}
}

type LoginResponse struct {
	Token   string         `json:"token"`
	Usuario models.Usuario `json:"usuario"`
}

func (s *AuthService) Login(req dto.LoginRequest) (*LoginResponse, error) {
	usuario, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("credenciais inválidas")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(req.Senha)); err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	token, err := utils.GenerateToken(usuario.ID, usuario.Email, usuario.Perfil)
	if err != nil {
		return nil, errors.New("erro ao gerar token")
	}

	return &LoginResponse{Token: token, Usuario: *usuario}, nil
}

func (s *AuthService) Register(req dto.RegisterRequest) (*models.Usuario, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("erro ao processar senha")
	}

	usuario := &models.Usuario{
		Nome:   req.Nome,
		Email:  req.Email,
		Senha:  string(hash),
		Perfil: req.Perfil,
	}

	if err := s.repo.Create(usuario); err != nil {
		return nil, errors.New("erro ao criar usuário, email já existe")
	}

	return usuario, nil
}
