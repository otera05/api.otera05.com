package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/otera05/api.otera05.com/entity"
	"github.com/otera05/api.otera05.com/usecase"
)

type TodoAccess struct {
	db *sqlx.DB
}

var _ usecase.TodoRepository = (*TodoAccess)(nil)

func (t TodoAccess) Find(ctx context.Context, id entity.TodoID) (*entity.Todo, error) {
	// TODO: WIP
	return nil, nil
}
func (t TodoAccess) FindAll(ctx context.Context) (entity.TodoList, error) {
	// TODO: WIP
	return nil, nil
}
func (t TodoAccess) Create(ctx context.Context, todo entity.Todo) (entity.TodoID, error) {
	// TODO: WIP
	return "", nil
}
func (t TodoAccess) Update(ctx context.Context, todo entity.Todo) (entity.TodoID, error) {
	// TODO: WIP
	return "", nil
}
func (t TodoAccess) Delete(ctx context.Context, id entity.TodoID) error {
	// TODO: WIP
	return nil
}
