package backend

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// Storage отвечает за сохранение/загрузку задач
type Storage struct {
	path string
}

func NewStorage(appName string) (*Storage, error) {
	confDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	appDir := filepath.Join(confDir, appName)
	if err := os.MkdirAll(appDir, 0o755); err != nil {
		return nil, err
	}
	p := filepath.Join(appDir, "tasks.json")
	return &Storage{path: p}, nil
}

func (s *Storage) SaveTasks(tasks []*Task) error {
	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	return enc.Encode(tasks)
}

func (s *Storage) LoadTasks() ([]*Task, error) {
	b, err := os.ReadFile(s.path)
	if err != nil {
		return nil, err
	}
	var tasks []*Task
	if err := json.Unmarshal(b, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func generateID() string {
	// простая уникальная id
	return uuid.NewString()
}
