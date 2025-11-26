package db

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Taras-Rm/technical-practice/task-tracker/domain"
)

const FileModePermit = 0644

type TaskStore interface {
	GetAll() ([]domain.Task, error)
	Add(description string) (*domain.Task, error)
	Delete(id int) error
	Update(id int, description string) error
	SetStatus(id int, status string) error
}

type store struct {
	fileName string
}

func InitTaskStore(fileName string) TaskStore {
	return &store{
		fileName,
	}
}

func (s *store) GetAll() ([]domain.Task, error) {
	return s.getTasksList()
}

func (s *store) Add(description string) (*domain.Task, error) {
	createAt := time.Now()

	newTask := domain.Task{
		Id:          323, // TODO: randon id
		Description: description,
		Status:      "todo",
		CreatedAt:   createAt,
		UpdatedAt:   createAt,
	}

	db, err := s.getDb()
	if err != nil {
		return nil, err
	}

	db.Tasks = append(db.Tasks, newTask)

	byteJson, err := json.Marshal(db)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(s.fileName, byteJson, FileModePermit)
	if err != nil {
		return nil, err
	}

	return &newTask, nil
}

func (s *store) Delete(id int) error {
	db, err := s.getDb()
	if err != nil {
		return err
	}

	var newTasks []domain.Task

	for _, task := range db.Tasks {
		if task.Id == id {
			continue
		}
		newTasks = append(newTasks, task)
	}

	if len(db.Tasks) == len(newTasks) {
		return fmt.Errorf("task with id %d not found", id)
	}

	db.Tasks = newTasks

	byteJson, err := json.Marshal(db)
	if err != nil {
		return err
	}

	err = os.WriteFile(s.fileName, byteJson, FileModePermit)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) Update(id int, description string) error {
	db, err := s.getDb()
	if err != nil {
		return err
	}

	for idx, task := range db.Tasks {
		if task.Id == id {
			db.Tasks[idx].Description = description
			db.Tasks[idx].UpdatedAt = time.Now()
		}
	}

	byteJson, err := json.Marshal(db)
	if err != nil {
		return err
	}

	err = os.WriteFile(s.fileName, byteJson, FileModePermit)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) SetStatus(id int, status string) error {
	db, err := s.getDb()
	if err != nil {
		return err
	}

	for idx, task := range db.Tasks {
		if task.Id == id {
			db.Tasks[idx].Description = status
			db.Tasks[idx].UpdatedAt = time.Now()
		}
	}

	byteJson, err := json.Marshal(db)
	if err != nil {
		return err
	}

	err = os.WriteFile(s.fileName, byteJson, FileModePermit)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) getDb() (*domain.Db, error) {
	jsonFile, err := os.Open(s.fileName)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteJson, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var db domain.Db

	err = json.Unmarshal(byteJson, &db)
	if err != nil {
		return nil, err
	}

	return &db, nil
}

func (s *store) getTasksList() ([]domain.Task, error) {
	db, err := s.getDb()
	if err != nil {
		return nil, err
	}

	return db.Tasks, nil
}
