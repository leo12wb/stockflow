package repositories

import (
	"github.com/seu-usuario/estoque-api/config"
	"github.com/seu-usuario/estoque-api/models"
)

type UsuarioRepository struct{}

func (r *UsuarioRepository) FindByEmail(email string) (*models.Usuario, error) {
	var usuario models.Usuario
	result := config.DB.Where("email = ?", email).First(&usuario)
	return &usuario, result.Error
}

func (r *UsuarioRepository) Create(usuario *models.Usuario) error {
	return config.DB.Create(usuario).Error
}

func (r *UsuarioRepository) FindByID(id uint) (*models.Usuario, error) {
	var usuario models.Usuario
	result := config.DB.First(&usuario, id)
	return &usuario, result.Error
}

func (r *UsuarioRepository) FindAll() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	result := config.DB.Find(&usuarios)
	return usuarios, result.Error
}
