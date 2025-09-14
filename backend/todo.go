package backend

import (
	"errors"
	"time"
)

// Task модель задачи
type Task struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TodoService экспортируемый сервис для фронтенда
type TodoService struct {
	storage *Storage
	tasks   []*Task
}

func NewTodoService() *TodoService {
	s, err := NewStorage("wails-todo-app")
	if err != nil {
		panic(err)
	}

	svc := &TodoService{storage: s}
	svc.load()
	return svc
}

// GetTasks возвращает список задач
func (t *TodoService) GetTasks() ([]*Task, error) {
	return t.tasks, nil
}

// AddTask добавляет задачу и сохраняет
func (t *TodoService) AddTask(text string) (*Task, error) {
	if text == "" {
		return nil, errors.New("text is empty")
	}
	n := &Task{
		ID:        generateID(),
		Text:      text,
		Done:      false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	t.tasks = append([]*Task{n}, t.tasks...)
	err := t.storage.SaveTasks(t.tasks)
	if err != nil {
		return nil, err
	}
	return n, nil
}

// ToggleDone переключает статус задачи
func (t *TodoService) ToggleDone(id string) (*Task, error) {
	for _, task := range t.tasks {
		if task.ID == id {
			task.Done = !task.Done
			task.UpdatedAt = time.Now()
			err := t.storage.SaveTasks(t.tasks)
			return task, err
		}
	}
	return nil, errors.New("task not found")
}

// UpdateTask обновляет текст задачи
func (t *TodoService) UpdateTask(id, text string) (*Task, error) {
	for _, task := range t.tasks {
		if task.ID == id {
			task.Text = text
			task.UpdatedAt = time.Now()
			err := t.storage.SaveTasks(t.tasks)
			return task, err
		}
	}
	return nil, errors.New("task not found")
}

// DeleteTask удаляет задачу
func (t *TodoService) DeleteTask(id string) error {
	for i, task := range t.tasks {
		if task.ID == id {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			return t.storage.SaveTasks(t.tasks)
		}
	}
	return errors.New("task not found")
}

func (t *TodoService) load() {
	tasks, err := t.storage.LoadTasks()
	if err != nil {
		// если файла ещё нет — инициализируем пустым слайсом
		t.tasks = []*Task{}
		return
	}
	t.tasks = tasks
}
