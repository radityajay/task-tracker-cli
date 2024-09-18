package task

import (
	"fmt"
)

type Repository struct {
	filename string
}

func NewRepository(filename string) *Repository {
	return &Repository{
		filename: filename,
	}
}

func (r *Repository) Add(data *Task) error {
	file, err := ReadFileJSON(r.filename)
	if err != nil {
		return err
	}

	if data.ID == 0 {
		fmt.Println("masuk", file)
		data.ID = 1
		if len(file) > 0 {
			data.ID = file[len(file)-1].ID + 1
		}
	}

	file = append(file, data)
	err = WriteFileJSON(r.filename, file)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) List(status Status) ([]*Task, error) {
	file, err := ReadFileJSON(r.filename)
	if err != nil {
		return nil, err
	}

	var tasks []*Task
	for _, task := range file {
		if status != "" {
			if task.Status == status {
				tasks = append(tasks, task)
			}
		} else {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func (r *Repository) Get(id uint) (*Task, error) {
	file, err := ReadFileJSON(r.filename)
	if err != nil {
		return nil, err
	}

	for _, task := range file {
		if task.ID == id {
			return task, nil
		}
	}

	return nil, nil
}

func (r *Repository) Update(data *Task) error {
	file, err := ReadFileJSON(r.filename)
	if err != nil {
		return err
	}

	for i, task := range file {
		if task.ID == data.ID {
			file[i] = data
			break
		}
	}

	err = WriteFileJSON(r.filename, file)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(id uint) error {
	file, err := ReadFileJSON(r.filename)
	if err != nil {
		return err
	}

	indexToDelete := -1
	for i, task := range file {
		if task.ID == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete != -1 {
		file = append(file[:indexToDelete], file[indexToDelete+1:]...)
	}

	err = WriteFileJSON(r.filename, file)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetList(status Status) ([]*Task, error) {
	file, err := ReadFileJSON(r.filename)
	if err != nil {
		return nil, err
	}

	result := []*Task{}
	for _, task := range file {
		if task.Status == status {
			result = append(result, task)
		}
	}

	return result, nil
}
