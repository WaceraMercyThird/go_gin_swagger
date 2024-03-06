package handlers

import (
	"go_api/model"
	"net/http"

	
	"github.com/gin-gonic/gin"
)


type userResponse struct {
	Data []model.User `json: "data"`
}

type ErrorResponse struct {
    Message string `json:"message"`
}

// GetUsers returns a list of all users from the database.
// @Summary Return a list of all users from the database
// @Description Return a list of all users from the database
// @Tags Users
// @Success 200 {object} userResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/users/ [get]
func GetUsers(c *gin.Context) {
	users := model.ListUsers()
    if len(users) == 0 {
        // Assume no users found as an example of an error response
        c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found"})
        return
    }
    c.JSON(http.StatusOK, userResponse{Data: users})
}

