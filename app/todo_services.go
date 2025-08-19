package app

import (
	"github.com/IkhsanDS/golang-api/models"
	"gorm.io/gorm"
)

type TodoService struct{ db *gorm.DB }

func NewTodoService(db *gorm.DB) *TodoService { return &TodoService{db: db} }

func (s *TodoService) Create(title string) (models.Todo, error) {
	t := models.Todo{Title: title}
	return t, s.db.Create(&t).Error
}

func (s *TodoService) List() ([]models.Todo, error) {
	var items []models.Todo
	return items, s.db.Order("created_at DESC").Find(&items).Error
}

func (s *TodoService) Get(id string) (models.Todo, error) {
	var t models.Todo
	return t, s.db.First(&t, "id = ?", id).Error
}

func (s *TodoService) Update(id string, title *string, completed *bool) (models.Todo, error) {
	var t models.Todo
	if err := s.db.First(&t, "id = ?", id).Error; err != nil {
		return t, err
	}
	if title != nil {
		t.Title = *title
	}
	if completed != nil {
		t.Completed = *completed
	}
	return t, s.db.Save(&t).Error
}

func (s *TodoService) Delete(id string) error {
	return s.db.Where("id = ?", id).Delete(&models.Todo{}).Error
}
