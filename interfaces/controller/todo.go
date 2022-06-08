package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otera05/api.otera05.com/entity"
	"github.com/otera05/api.otera05.com/usecase"
)

type Todo interface {
	Get(gctx *gin.Context)
	List(gctx *gin.Context)
	Create(gctx *gin.Context)
	Update(gctx *gin.Context)
	Delete(gctx *gin.Context)
}

type todo struct {
	todoRepo usecase.TodoRepository
}

var _ Todo = (*todo)(nil)

func (t *todo) Get(gctx *gin.Context) {
	// TODO: WIP
	ctx := gctx.Request.Context()
	t.todoRepo.Find(ctx, "")
	gctx.JSON(http.StatusOK, "")
}

func (t *todo) List(gctx *gin.Context) {
	// TODO: WIP
	ctx := gctx.Request.Context()
	t.todoRepo.FindAll(ctx)
	gctx.JSON(http.StatusOK, "")
}

func (t *todo) Create(gctx *gin.Context) {
	// TODO: WIP
	ctx := gctx.Request.Context()
	t.todoRepo.Create(ctx, entity.Todo{})
	gctx.JSON(http.StatusCreated, "")
}

func (t *todo) Update(gctx *gin.Context) {
	// TODO: WIP
	ctx := gctx.Request.Context()
	t.todoRepo.Update(ctx, entity.Todo{})
	gctx.JSON(http.StatusOK, "")
}

func (t *todo) Delete(gctx *gin.Context) {
	// TODO: WIP
	ctx := gctx.Request.Context()
	t.todoRepo.Find(ctx, "")
	gctx.JSON(http.StatusNoContent, "")
}
