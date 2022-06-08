package usecase

import (
	"context"

	"github.com/otera05/api.otera05.com/entity"
)

type TodoRepository interface {
	Find(ctx context.Context, id entity.TodoID) (*entity.Todo, error)
	FindAll(ctx context.Context) (entity.TodoList, error)
	Create(ctx context.Context, todo entity.Todo) (entity.TodoID, error)
	Update(ctx context.Context, todo entity.Todo) (entity.TodoID, error)
	Delete(ctx context.Context, id entity.TodoID) error
}

type todo struct{}

func NewTodoRepository() TodoRepository {
	return &todo{}
}

// Find は、指定された TodoID のデータを取得し返却する
func (t *todo) Find(ctx context.Context, id entity.TodoID) (*entity.Todo, error) {
	// TODO: WIP
	return nil, nil
}

// FindAll は、すべての Todo のデータを取得し返却する
func (t *todo) FindAll(ctx context.Context) (entity.TodoList, error) {
	// TODO: WIP
	return nil, nil
}

// Create は、新規 Todo データを作成する
func (t *todo) Create(ctx context.Context, todo entity.Todo) (entity.TodoID, error) {
	// TODO: WIP
	return "", nil
}

// Update は、指定された Todo のデータで更新する
func (t *todo) Update(ctx context.Context, todo entity.Todo) (entity.TodoID, error) {
	// TODO: WIP
	return "", nil
}

// Delete は、指定された TodoID のデータを削除する
func (t *todo) Delete(ctx context.Context, id entity.TodoID) error {
	// TODO: WIP
	return nil
}
