package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"todoapp/database/generated"
	"todoapp/i18n"

	"github.com/gorilla/mux"
)

type TaskAPI struct {
	Q *generated.Queries
}

// middleware
func (a *TaskAPI) ListTasks(w http.ResponseWriter, r *http.Request) {
	userID, ok := MustUserID(r)
	if !ok {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}
	items, err := a.Q.ListTasks(r.Context(), userID)
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(items)
}

func (a *TaskAPI) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := MustUserID(r)
	if !ok {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var in struct {
		Title    string     `json:"title"`
		Deadline *time.Time `json:"deadline"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Title == "" {
		i18n.ErrorResponse(w, r, "bad json", http.StatusBadRequest)
		return
	}
	var deadline sql.NullTime
	if in.Deadline != nil {
		deadline = sql.NullTime{Time: *in.Deadline, Valid: true}
	}

	created, err := a.Q.CreateTask(r.Context(), generated.CreateTaskParams{
		Title:    in.Title,
		UserID:   userID,
		Deadline: deadline,
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(created)
}

func (a *TaskAPI) UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := MustUserID(r)
	if !ok {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		i18n.ErrorResponse(w, r, "invalid id", http.StatusBadRequest)
		return
	}

	var body struct {
		Title    *string    `json:"title"`
		Done     *bool      `json:"done"`
		Deadline *time.Time `json:"deadline"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		i18n.ErrorResponse(w, r, "invalid json", http.StatusBadRequest)
		return
	}

	var updated generated.Task
	if body.Title != nil || body.Deadline != nil {
		title := ""
		if body.Title != nil {
			title = *body.Title
		}

		var deadline sql.NullTime
		if body.Deadline != nil {
			deadline = sql.NullTime{Time: *body.Deadline, Valid: true}
		}

		updated, err = a.Q.UpdateTask(r.Context(), generated.UpdateTaskParams{
			ID:       int32(id),
			Title:    title,
			UserID:   userID,
			Deadline: deadline,
		})
		if err != nil {
			i18n.ErrorResponse(w, r, "database error", 500)
			return
		}
	}
	if body.Done != nil {
		setDoneResult, err := a.Q.SetDone(r.Context(), generated.SetDoneParams{
			ID:     int32(id),
			Done:   *body.Done,
			UserID: userID,
		})
		if err != nil {
			i18n.ErrorResponse(w, r, "database error", 500)
			return
		}
		// Преобразуем SetDoneRow в Task (без deadline, так как SetDoneRow его не содержит)
		updated = generated.Task{
			ID:     setDoneResult.ID,
			Title:  setDoneResult.Title,
			Done:   setDoneResult.Done,
			UserID: setDoneResult.UserID,
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(updated)
}

func (a *TaskAPI) DeleteTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := MustUserID(r)
	if !ok {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		i18n.ErrorResponse(w, r, "invalid id", http.StatusBadRequest)
		return
	}
	if err := a.Q.DeleteTask(r.Context(), generated.DeleteTaskParams{
		ID:     int32(id),
		UserID: userID,
	}); err != nil {
		i18n.ErrorResponse(w, r, "database error", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
