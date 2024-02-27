package repository

import (
	"bank-ina-test/domain"
	"context"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) Create(ctx context.Context, task *domain.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]domain.Task, error) {
	var tasks []domain.Task
	if err := r.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id uint) (*domain.Task, error) {
	var task domain.Task
	if err := r.DB.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Update(ctx context.Context, task *domain.Task) error {
	result := r.DB.Model(&task).Updates(domain.Task{UserID: task.UserID, Title: task.Title, Description: task.Description, Status: task.Status})
	return result.Error
}

func (r *TaskRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.Delete(&domain.Task{}, id).Error
}
