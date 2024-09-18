package task

import (
	"encoding/json"
	"io"
	"os"
)

func getFile(filename string) (*os.File, error) {

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if info.Size() == 0 {
		_, err = file.Write([]byte("[]"))
		if err != nil {
			return nil, err
		}
	}

	return file, nil
}

func ReadFileJSON(filename string) ([]*Task, error) {
	file, err := getFile(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var task []*Task
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	err = json.Unmarshal(data, &task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func WriteFileJSON(filename string, data []*Task) error {
	file, err := getFile(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
