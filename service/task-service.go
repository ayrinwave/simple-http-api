package service

import (
	"Simple_http_api/models"
	"context"
	"github.com/google/uuid"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type TaskService interface {
	CreateTask() *models.Task
	GetTask(id string) (*models.Task, bool)
	DeleteTask(id string) bool
	GetAllTasks() []*models.Task
}

type taskService struct {
	mutex sync.RWMutex
	tasks map[string]*models.Task
}

func NewTaskService() TaskService {
	return &taskService{
		tasks: make(map[string]*models.Task),
	}
}

func (t *taskService) CreateTask() *models.Task {
	id := uuid.New().String()
	task := &models.Task{
		ID:        id,
		Status:    models.Waiting,
		CreatedAt: time.Now(),
	}
	ctx, cancel := context.WithCancel(context.Background())
	task.CancelFunc = cancel
	t.mutex.Lock()
	t.tasks[id] = task
	t.mutex.Unlock()

	go t.runTask(ctx, task)
	return task
}

func (t *taskService) runTask(ctx context.Context, task *models.Task) {
	start := time.Now()

	t.mutex.Lock()
	task.Status = models.Running
	task.StartedAt = &start
	t.mutex.Unlock()

	usefulWork := 3*time.Minute + time.Duration(rand.Intn(2*60))*time.Second

	select {
	case <-ctx.Done():
		t.mutex.Lock()
		task.Status = models.Deleted
		now := time.Now()
		task.EndedAt = &now
		errMsg := "task cancelled/deleted"
		task.Error = &errMsg
		t.mutex.Unlock()
		return
	case <-time.After(usefulWork):
		t.mutex.Lock()
		now := time.Now()
		task.EndedAt = &now
		task.Duration = now.Sub(*task.StartedAt).String()
		task.Status = models.Success
		resMsg := "task success"
		task.Result = &resMsg
		task.Error = nil
		t.mutex.Unlock()
	}
}

func (t *taskService) GetAllTasks() []*models.Task {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	tasks := make([]*models.Task, 0, len(t.tasks))
	for _, task := range t.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}
func (t *taskService) GetTask(id string) (*models.Task, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	task, ok := t.tasks[id]
	return task, ok
}

func (t *taskService) DeleteTask(id string) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	task, ok := t.tasks[id]
	if !ok {
		return false
	}

	if task.Status == models.Success || task.Status == models.Running {
		task.CancelFunc()
	}
	delete(t.tasks, id)
	return true
}
