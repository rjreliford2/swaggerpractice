package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs" // docs is generated by Swag CLI, you have to import it.
)

// UserResponse represents the response structure for the user endpoint
type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ErrorResponse represents the error response structure
type ErrorResponse struct {
	Message string `json:"message"`
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
// getUserByID returns a user by their ID
// @Summary Get user by ID
// @Description Get a user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Router /users/{id} [get]
func getUserByID(c echo.Context) error {
	userID := c.Param("id")

	// Assuming user data is stored in a map with user ID as the key
	users := map[string]UserResponse{
		"1": {ID: 1, Name: "John Doe"},
		"2": {ID: 2, Name: "Jane Smith"},
	}

	// Check if the user exists
	user, ok := users[userID]
	if !ok {
		errResponse := ErrorResponse{Message: "User not found"}
		return c.JSON(http.StatusNotFound, errResponse)
	}

	return c.JSON(http.StatusOK, user)
}
func main() {
	e := echo.New()

	// Route handler
	e.GET("/users/:id", getUserByID)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Serve Swagger JSON endpoint
	/*e.GET("/swagger.json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echoSwagger.Handler(c.Request()))
	})*/

	e.Logger.Fatal(e.Start(":1323"))
}
