package models

type Usuario struct {
	Base
	Nome   string `gorm:"not null" json:"nome"`
	Email  string `gorm:"unique;not null" json:"email"`
	Senha  string `gorm:"not null" json:"-"`
	Perfil string `gorm:"not null;default:'estoquista'" json:"perfil"`
}
