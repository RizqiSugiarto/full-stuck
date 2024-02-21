package repository

import (
	"fmt"
	"list/model"

	"gorm.io/gorm"
)

type NoteRepository interface {
	Insert(userData model.UserInput) error
	GetByHeader(header string) (model.UserInput, error)
	Delete(header string) error
}

type noteRepository struct {
	gorm *gorm.DB
}

func (n *noteRepository) Insert(userData model.UserInput) error {
	result := n.gorm.Create(&userData)

	if result.Error != nil {
		return fmt.Errorf("error when insert data postgres: %s", result.Error)
	}
	return nil
}

func (n *noteRepository) GetByHeader(header string) (model.UserInput, error) {
	var DataUser model.UserInput

	if err := n.gorm.Where("header = ?", header).First(&DataUser).Error; err != nil {
		return model.UserInput{}, fmt.Errorf("error when getByHeader data postgres: %s", err)
	}
	return DataUser, nil
}

func (n *noteRepository) Delete(header string) error {
	var DataUser model.UserInput

	if err := n.gorm.Where("header = ?", header).Delete(&DataUser).Error; err != nil {
		return fmt.Errorf("error when delete data postgres: %s", err)
	}
	return nil
}

func NewNoteRepository(gorm *gorm.DB) NoteRepository {
	return &noteRepository{
		gorm: gorm,
	}
}
