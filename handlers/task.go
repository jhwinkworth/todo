package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"todo/models"
	"todo/store"
)

var (
	TaskRe       = regexp.MustCompile(`^/api/tasks/*$`)
	TaskReWithID = regexp.MustCompile(`^/api/tasks/([0-9]+)$`)
)

type TasksHandler struct {
	store store.TaskStore
}

func NewTasksHandler(s store.TaskStore) *TasksHandler {
	return &TasksHandler{
		store: s,
	}
}

func (h *TasksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && TaskRe.MatchString(r.URL.Path):
		h.CreateTask(w, r)
		return
	case r.Method == http.MethodGet && TaskRe.MatchString(r.URL.Path):
		h.GetTasks(w, r)
		return
	case r.Method == http.MethodGet && TaskReWithID.MatchString(r.URL.Path):
		h.GetTaskByID(w, r)
		return
	case r.Method == http.MethodPut && TaskReWithID.MatchString(r.URL.Path):
		h.UpdateTask(w, r)
		return
	case r.Method == http.MethodDelete && TaskReWithID.MatchString(r.URL.Path):
		h.DeleteTask(w, r)
		return
	default:
		return
	}
}

func (h *TasksHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	// Task object that will be populated from JSON payload
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	// Call the store to add the task
	if err := h.store.Add(task); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	// Set the status code to 200
	w.WriteHeader(http.StatusOK)
}

func (h *TasksHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	resources, err := h.store.List()

	jsonBytes, err := json.Marshal(resources)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *TasksHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// Extract the resource ID/slug using a regex
	matches := TaskReWithID.FindStringSubmatch(r.URL.Path)

	// Expect matches to be length >= 2 (full string + 1 matching group)
	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	// Retrieve task from the store
	task, err := h.store.Get(matches[1])
	if err != nil {
		// Special case of NotFound Error
		if errors.Is(err, store.NotFoundErr) {
			NotFoundHandler(w, r)
			return
		}

		// Every other error
		InternalServerErrorHandler(w, r)
		return
	}

	// Convert the struct into JSON payload
	jsonBytes, err := json.Marshal(task)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	// Write the results
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *TasksHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	matches := TaskReWithID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	// Task object that will be populated from JSON payload
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	if err := h.store.Update(matches[1], task); err != nil {
		if errors.Is(err, store.NotFoundErr) {
			NotFoundHandler(w, r)
			return
		}
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TasksHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	matches := TaskReWithID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	if err := h.store.Remove(matches[1]); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}
