package main

import (
	"database/sql"
	"log"
	"net/http"
	"todoService/database"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Connect to the PostgreSQL database
	connStr := "host=localhost dbname=todo_app user=postgres password=yourpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure the database is reachable
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	/// Initialize our handler with the db connection
	queries := database.New(db)
	handle := &handler{db: db, queries: queries}

	// Define the HTTP routes
	e.GET("/", handleRoot)
	e.GET("/tasks", handle.getTasks)
	e.POST("/tasks", handle.createTask)
	e.GET("/tasks/:id", handle.getTask)
	e.PUT("/tasks/:id", handle.updateTask)
	e.DELETE("/tasks/:id", handle.deleteTask)

	// Start the server
	e.Logger.Fatal(e.Start(":1323"))
}

type handler struct {
	db      *sql.DB
	queries *database.Queries
}

func handleRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Server is up and running")
}

func (h *handler) getTasks(c echo.Context) error {
	tasks, err := h.queries.GetTasks(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *handler) createTask(c echo.Context) error {
	return c.String(http.StatusOK, "Create a task")
}

func (h *handler) getTask(c echo.Context) error {
	return c.String(http.StatusOK, "Get a single task")
}

func (h *handler) updateTask(c echo.Context) error {
	return c.String(http.StatusOK, "Update a task")
}

func (h *handler) deleteTask(c echo.Context) error {
	return c.String(http.StatusOK, "Delete a task")
}
