package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-bookstore/crud"
	"gin-bookstore/database"
)

// @Summary 	List users
// @Description Retrieve a list of users
// @Tags 		users
// @Produce 	json
// @Success 	200 {object} UsersResponse
// @Router 		/users [get]
func FindUsers(c *gin.Context) {
	users, err := crud.ReadUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users, "count": len(users)})
}

// @Summary      Get Current User
// @Description  Retrieve the currently logged-in user's details
// @Tags         users
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  database.User "Current user details"
// @Failure      400  {object}  ErrorResponse "Bad Request"
// @Failure      401  {object}  ErrorResponse "Unauthorized"
// @Router       /users/current [get]
func FindCurrentUser(c *gin.Context) {
	email := c.MustGet("currentUserEmail")
	user, err := crud.ReadUserBy("email", email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}


// @Summary 	Get user by ID
// @Description Retrieve a user's details by their ID
// @Tags 		users
// @Produce 	json
// @Param 		id path string true "User ID"
// @Success 	200 {object} database.User
// @Failure 	400 {object} ErrorResponse
// @Router 		/users/{id} [get]
func FindUser(c *gin.Context) {
	user, err := crud.ReadUserBy("id", c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// @Summary 	Create a new user
// @Description Add a new user to the system
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Param 		user body database.CreateUserSchema true "User data"
// @Success 	200 {object} database.User
// @Failure 	400 {object} ErrorResponse
// @Router 		/users [post]
func CreateUser(c *gin.Context) {
	// Validate input
	var input database.CreateUserSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create User
	user, err := crud.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// @Summary 	Update user by ID
// @Description Modify a user's details by their ID
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Security 	BearerAuth
// @Param 		id path string true "User ID"
// @Param 		user body database.UpdateUserSchema true "Updated user data"
// @Success 	200 {object} database.User
// @Failure 	400 {object} ErrorResponse
// @Router 		/users/{id} [patch]
func UpdateUser(c *gin.Context) {
	// Validate input
	var input database.UpdateUserSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update User
	user, err := crud.UpdateUserById(c.Param("id"), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// @Summary Delete user by ID
// @Description Remove a user from the system by their ID
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} database.User
// @Failure 400 {object} ErrorResponse
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	user, err := crud.DeleteUserByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// UsersResponse represents the structure of the response for the list of users
// @Description Response structure for list of users
// @Param data query []database.User true "List of users"
// @Param count query int true "Total number of users"
// @Success 200 {object} UsersResponse
type UsersResponse struct {
	Data  []database.User `json:"data"`
	Count int             `json:"count"`
}
