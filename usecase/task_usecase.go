package usecase

import (
	"bank-ina-test/domain"
	"bank-ina-test/repository"
	"context"
)

type TaskUsecase struct {
	TaskRepository repository.TaskRepository
}

func (uc *TaskUsecase) CreateTask(ctx context.Context, task *domain.Task) error {
	// Call repository method to create task
	return uc.TaskRepository.Create(ctx, task)
}

func (uc *TaskUsecase) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	// Call repository method to get all tasks
	return uc.TaskRepository.GetAll(ctx)
}

func (uc *TaskUsecase) GetTaskByID(ctx context.Context, id uint) (*domain.Task, error) {
	// Call repository method to get task by ID
	return uc.TaskRepository.GetByID(ctx, id)
}

func (uc *TaskUsecase) UpdateTask(ctx context.Context, task *domain.Task) error {
	// Call repository method to update task
	return uc.TaskRepository.Update(ctx, task)
}

func (uc *TaskUsecase) DeleteTask(ctx context.Context, id uint) error {
	// Call repository method to delete task
	return uc.TaskRepository.Delete(ctx, id)
}
