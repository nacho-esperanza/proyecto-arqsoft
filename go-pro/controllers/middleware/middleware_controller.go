package middlewareController

import (
	UserCliente "go-pro/clients/user"
	"go-pro/utils/errors"

	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) {
	// Get the user id(int) from the request
	userId := c.GetInt("Id")

	// Get the user from the database
	user := UserCliente.GetUserById(userId)

	/*// Get the user email from the request
	email := c.GetHeader("email")

	// Get the user from the database
	user := UserCliente.GetUserByEmail(email)
	*/
	// Check if the user is admin
	if user.Email != "admin@gmail.com" {
		c.AbortWithStatusJSON(403, errors.NewForbiddenApiError("You are not admin"))
		return
	}

	// If the user is admin, continue with the request
	c.Next()
}
