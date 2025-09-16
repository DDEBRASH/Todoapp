package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"todoapp/database/generated"
	"todoapp/i18n"

	"github.com/gorilla/mux"
)

type TeamProjectAPI struct {
	Q *generated.Queries
}

// Генерация 6-значного кода
func generateProjectCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// CREATE TEAM PROJECT
func (a *TeamProjectAPI) CreateTeamProject(w http.ResponseWriter, r *http.Request) {
	userID, ok := MustUserID(r)
	if !ok {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var in struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Name == "" {
		i18n.ErrorResponse(w, r, "bad json", http.StatusBadRequest)
		return
	}

	// Генерируем уникальный код
	var code string
	for {
		code = generateProjectCode()
		_, err := a.Q.GetTeamProjectByCode(r.Context(), code)
		if err != nil {
			// Код не найден, значит он уникален
			break
		}
	}

	// Создаем проект
	project, err := a.Q.CreateTeamProject(r.Context(), generated.CreateTeamProjectParams{
		Name:      in.Name,
		Code:      code,
		CreatedBy: userID,
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	// Добавляем создателя как участника
	_, err = a.Q.AddTeamProjectMember(r.Context(), generated.AddTeamProjectMemberParams{
		ProjectID: project.ID,
		UserID:    userID,
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(project)
}

// JOIN TEAM PROJECT
func (a *TeamProjectAPI) JoinTeamProject(w http.ResponseWriter, r *http.Request) {
	userID, ok := MustUserID(r)
	if !ok {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var in struct {
		Code string `json:"code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Code == "" {
		i18n.ErrorResponse(w, r, "bad json", http.StatusBadRequest)
		return
	}

	// Проверяем существование проекта
	project, err := a.Q.GetTeamProjectByCode(r.Context(), in.Code)
	if err != nil {
		i18n.ErrorResponse(w, r, "user not found", http.StatusNotFound)
		return
	}

	// Проверяем, не является ли пользователь уже участником
	isMember, err := a.Q.CheckTeamProjectMember(r.Context(), generated.CheckTeamProjectMemberParams{
		ProjectID: project.ID,
		UserID:    userID,
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	if isMember {
		i18n.ErrorResponse(w, r, "user not found", http.StatusBadRequest)
		return
	}

	// Добавляем пользователя в проект
	_, err = a.Q.AddTeamProjectMember(r.Context(), generated.AddTeamProjectMemberParams{
		ProjectID: project.ID,
		UserID:    userID,
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(project)
}

// GET USER TEAM PROJECTS
func (a *TeamProjectAPI) GetUserTeamProjects(w http.ResponseWriter, r *http.Request) {
	userID, ok := MustUserID(r)
	if !ok {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}

	projects, err := a.Q.GetUserTeamProjects(r.Context(), userID)
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(projects)
}

// GET TEAM PROJECT MEMBERS
func (a *TeamProjectAPI) GetTeamProjectMembers(w http.ResponseWriter, r *http.Request) {
	userID, ok := MustUserID(r)
	if !ok {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}

	projectID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		i18n.ErrorResponse(w, r, "invalid id", http.StatusBadRequest)
		return
	}

	// Проверяем, является ли пользователь участником проекта
	isMember, err := a.Q.CheckTeamProjectMember(r.Context(), generated.CheckTeamProjectMemberParams{
		ProjectID: int32(projectID),
		UserID:    userID,
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	if !isMember {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}

	members, err := a.Q.GetTeamProjectMembers(r.Context(), int32(projectID))
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(members)
}

// CREATE TEAM TASK
func (a *TeamProjectAPI) CreateTeamTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := MustUserID(r)
	if !ok {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}

	projectID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		i18n.ErrorResponse(w, r, "invalid id", http.StatusBadRequest)
		return
	}

	// Проверяем, является ли пользователь участником проекта
	isMember, err := a.Q.CheckTeamProjectMember(r.Context(), generated.CheckTeamProjectMemberParams{
		ProjectID: int32(projectID),
		UserID:    userID,
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	if !isMember {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var in struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Deadline    *time.Time `json:"deadline"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Title == "" {
		i18n.ErrorResponse(w, r, "bad json", http.StatusBadRequest)
		return
	}

	var deadline sql.NullTime
	if in.Deadline != nil {
		deadline = sql.NullTime{Time: *in.Deadline, Valid: true}
	}

	task, err := a.Q.CreateTeamTask(r.Context(), generated.CreateTeamTaskParams{
		ProjectID:   int32(projectID),
		Title:       in.Title,
		Description: sql.NullString{String: in.Description, Valid: in.Description != ""},
		CreatedBy:   userID,
		Deadline:    deadline,
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	// Преобразуем sql.NullBool в обычный bool для JSON
	taskResponse := map[string]interface{}{
		"id":          task.ID,
		"project_id":  task.ProjectID,
		"title":       task.Title,
		"description": task.Description.String,
		"done":        task.Done.Bool,
		"created_by":  task.CreatedBy,
		"created_at":  task.CreatedAt,
		"updated_at":  task.UpdatedAt,
		"deadline":    task.Deadline,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(taskResponse)
}

// GET TEAM TASKS
func (a *TeamProjectAPI) GetTeamTasks(w http.ResponseWriter, r *http.Request) {
	userID, ok := MustUserID(r)
	if !ok {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}

	projectID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		i18n.ErrorResponse(w, r, "invalid id", http.StatusBadRequest)
		return
	}

	// Проверяем, является ли пользователь участником проекта
	isMember, err := a.Q.CheckTeamProjectMember(r.Context(), generated.CheckTeamProjectMemberParams{
		ProjectID: int32(projectID),
		UserID:    userID,
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	if !isMember {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tasks, err := a.Q.GetTeamTasks(r.Context(), int32(projectID))
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	// Преобразуем задачи в правильный формат
	var taskResponses []map[string]interface{}
	for _, task := range tasks {
		taskResponse := map[string]interface{}{
			"id":                  task.ID,
			"project_id":          task.ProjectID,
			"title":               task.Title,
			"description":         task.Description.String,
			"done":                task.Done.Bool,
			"created_by":          task.CreatedBy,
			"created_by_username": task.CreatedByUsername,
			"created_at":          task.CreatedAt,
			"updated_at":          task.UpdatedAt,
			"deadline":            task.Deadline,
		}
		taskResponses = append(taskResponses, taskResponse)
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(taskResponses)
}

// UPDATE TEAM TASK
func (a *TeamProjectAPI) UpdateTeamTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := MustUserID(r)
	if !ok {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}

	projectID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		i18n.ErrorResponse(w, r, "invalid id", http.StatusBadRequest)
		return
	}

	taskID, err := strconv.Atoi(mux.Vars(r)["taskId"])
	if err != nil {
		i18n.ErrorResponse(w, r, "invalid id", http.StatusBadRequest)
		return
	}

	// Проверяем, является ли пользователь участником проекта
	isMember, err := a.Q.CheckTeamProjectMember(r.Context(), generated.CheckTeamProjectMemberParams{
		ProjectID: int32(projectID),
		UserID:    userID,
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	if !isMember {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var in struct {
		Title       *string    `json:"title"`
		Description *string    `json:"description"`
		Done        *bool      `json:"done"`
		Deadline    *time.Time `json:"deadline"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		i18n.ErrorResponse(w, r, "invalid json", http.StatusBadRequest)
		return
	}

	var updated generated.TeamTask
	if in.Title != nil || in.Description != nil || in.Deadline != nil {
		title := ""
		description := ""
		if in.Title != nil {
			title = *in.Title
		}
		if in.Description != nil {
			description = *in.Description
		}

		var deadline sql.NullTime
		if in.Deadline != nil {
			deadline = sql.NullTime{Time: *in.Deadline, Valid: true}
		}

		updated, err = a.Q.UpdateTeamTask(r.Context(), generated.UpdateTeamTaskParams{
			ID:          int32(taskID),
			Title:       title,
			Description: sql.NullString{String: description, Valid: description != ""},
			ProjectID:   int32(projectID),
			Deadline:    deadline,
		})
		if err != nil {
			i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
			return
		}
	}

	if in.Done != nil {
		updated, err = a.Q.SetTeamTaskDone(r.Context(), generated.SetTeamTaskDoneParams{
			ID:        int32(taskID),
			Done:      sql.NullBool{Bool: *in.Done, Valid: true},
			ProjectID: int32(projectID),
		})
		if err != nil {
			i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
			return
		}
	}

	// Преобразуем обновленную задачу в правильный формат
	updatedResponse := map[string]interface{}{
		"id":          updated.ID,
		"project_id":  updated.ProjectID,
		"title":       updated.Title,
		"description": updated.Description.String,
		"done":        updated.Done.Bool,
		"created_by":  updated.CreatedBy,
		"created_at":  updated.CreatedAt,
		"updated_at":  updated.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(updatedResponse)
}

// DELETE TEAM TASK
func (a *TeamProjectAPI) DeleteTeamTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := MustUserID(r)
	if !ok {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}

	projectID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		i18n.ErrorResponse(w, r, "invalid id", http.StatusBadRequest)
		return
	}

	taskID, err := strconv.Atoi(mux.Vars(r)["taskId"])
	if err != nil {
		i18n.ErrorResponse(w, r, "invalid id", http.StatusBadRequest)
		return
	}

	// Проверяем, является ли пользователь участником проекта
	isMember, err := a.Q.CheckTeamProjectMember(r.Context(), generated.CheckTeamProjectMemberParams{
		ProjectID: int32(projectID),
		UserID:    userID,
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	if !isMember {
		i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = a.Q.DeleteTeamTask(r.Context(), generated.DeleteTeamTaskParams{
		ID:        int32(taskID),
		ProjectID: int32(projectID),
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
