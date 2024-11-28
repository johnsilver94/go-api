package todo

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/johnsilver94/go-api/services/auth"
	"github.com/johnsilver94/go-api/types"
	"github.com/johnsilver94/go-api/utils"
)

type Handler struct {
	store     types.TodoStore
	userStore types.UserStore
}

func NewHandler(store types.TodoStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/todos", h.handleGetProducts).Methods(http.MethodGet)
	// router.HandleFunc("/todos/{todoID}", h.handleGetProduct).Methods(http.MethodGet)

	// admin routes
	router.HandleFunc("/todos", auth.WithJWTAuth(h.handleCreateTodo, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo types.CreateTodoPayload
	if err := utils.ParseJSON(r, &todo); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(todo); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.store.CreateTodo(types.Todo{
		ID:          uuid.New(),
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		UserID:      todo.UserID,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, nil)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	todos, err := h.store.GetTodos()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, todos)
}
