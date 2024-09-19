package task

import (
	"fmt"
	"time"
)

type UseCase struct {
	Repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		Repo: repo,
	}
}

func (uc *UseCase) Add(data *Task) error {
	timeNow := time.Now()

	if data.Status == "" {
		data.Status = STATUS_TODO
	}

	if data.CreatedAt.IsZero() {
		data.CreatedAt = timeNow
	}

	if data.UpdatedAt.IsZero() {
		data.UpdatedAt = timeNow
	}

	if data.Description == "" {
		return fmt.Errorf("description cannot be empty")
	}

	return uc.Repo.Add(data)
}

func (uc *UseCase) List(status Status)([]*Task, error) {
	validStatus := status == STATUS_TODO || status == STATUS_DONE || status == STATUS_IN_PROGRESS || status == ""
	if !validStatus {
		return nil, fmt.Errorf("invalid status")
	}

	return uc.Repo.List(status)
}

func (uc *UseCase) Update(id uint, desc string) error {
	if id == 0 {
		return fmt.Errorf("id cannot be empty")
	}

	data, err := uc.Repo.Get(id)
	if err != nil {
		return err
	}

	data.Description = desc
	data.UpdatedAt = time.Now()

	return uc.Repo.Update(data)
}

func (uc *UseCase) Delete(id uint) error {
	if id == 0 {
		return fmt.Errorf("id cannot be empty")
	}

	return uc.Repo.Delete(id)
}

func (uc *UseCase) UpdateStatus(id uint, status Status) error {
	if id == 0 {
		return fmt.Errorf("id cannot be empty")
	}

	data, err := uc.Repo.Get(id)
	if err != nil {
		return err
	}

	data.Status = status
	data.UpdatedAt = time.Now()

	return uc.Repo.Update(data)
}

func (uc *UseCase) GetListByStatus(status Status) ([]*Task, error) {
	if status != STATUS_TODO && status != STATUS_DONE && status != STATUS_IN_PROGRESS && status != "" {
		return nil, fmt.Errorf("invalid status")
	}
	return uc.Repo.List(status)
}