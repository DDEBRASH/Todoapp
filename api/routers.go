package api

import (
	"todoapp/database/generated"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, q *generated.Queries) {
	taskAPI := &TaskAPI{Q: q}
	apiRouter := r.PathPrefix("/api").Subrouter()

	tasks := apiRouter.PathPrefix("/tasks").Subrouter()
	tasks.Use(WithAuth)
	tasks.HandleFunc("", taskAPI.ListTasks).Methods("GET")
	tasks.HandleFunc("", taskAPI.CreateTask).Methods("POST")
	tasks.HandleFunc("/{id}", taskAPI.UpdateTask).Methods("PATCH")
	tasks.HandleFunc("/{id}", taskAPI.DeleteTask).Methods("DELETE")

	userAPI := &UserAPI{Q: q}
	auth := apiRouter.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", userAPI.Register).Methods("POST")
	auth.HandleFunc("/login", userAPI.Login).Methods("POST")
	auth.HandleFunc("/logout", userAPI.Logout).Methods("POST")
	auth.HandleFunc("/request-password-reset", userAPI.RequestPasswordReset).Methods("POST")
	auth.HandleFunc("/reset-password", userAPI.ResetPassword).Methods("POST")

	// Team Projects API
	teamProjectAPI := &TeamProjectAPI{Q: q}
	teamProjects := apiRouter.PathPrefix("/team-projects").Subrouter()
	teamProjects.Use(WithAuth)
	teamProjects.HandleFunc("", teamProjectAPI.CreateTeamProject).Methods("POST")
	teamProjects.HandleFunc("/join", teamProjectAPI.JoinTeamProject).Methods("POST")
	teamProjects.HandleFunc("/my", teamProjectAPI.GetUserTeamProjects).Methods("GET")
	teamProjects.HandleFunc("/{id}/members", teamProjectAPI.GetTeamProjectMembers).Methods("GET")
	teamProjects.HandleFunc("/{id}/tasks", teamProjectAPI.GetTeamTasks).Methods("GET")
	teamProjects.HandleFunc("/{id}/tasks", teamProjectAPI.CreateTeamTask).Methods("POST")
	teamProjects.HandleFunc("/{id}/tasks/{taskId}", teamProjectAPI.UpdateTeamTask).Methods("PATCH")
	teamProjects.HandleFunc("/{id}/tasks/{taskId}", teamProjectAPI.DeleteTeamTask).Methods("DELETE")
}
