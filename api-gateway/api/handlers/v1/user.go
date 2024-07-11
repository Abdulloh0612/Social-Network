package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	models "api-gateway/api/handlers/models"
	l "api-gateway/pkg/logger"
	"api-gateway/pkg/utils"
	pbc "api-gateway/protos/comment-service"
	pbu "api-gateway/protos/user-service"
)

// CreateUser ...
// @Summary CreateUser
// @Security ApiKeyAuth
// @Description Api for creating a new user
// @Tags User
// @Accept json
// @Produce json
// @Param User body models.UserCreate true "createUserModel"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users/create [post]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		body models.UserCreate
	)

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Register(ctx, &pbu.CreateUserReq{
		FirstName:      body.Name,
		LastName:       body.LastName,
		Email:          body.Email,
		Password:       body.Password,
		UserName:       body.UserName,
		PhoneNumber:    body.PhoneNumber,
		BirthDate:      body.BirthDate,
		Biography:      body.Biography,
		Gender:         body.Gender,
		ProfilePicture: body.ProfilePicture,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetUser gets user by field and value
// @Summary GetUser
// @Security ApiKeyAuth
// @Description API for retrieving user by field and value
// @Tags User
// @Accept json
// @Produce json
// @Param field query string true "Field name for filtering"
// @Param value query string true "Value for filtering"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users [get]
func (h *handlerV1) GetUser(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Get(
		ctx, &pbu.GetUserReq{
			Field: field,
			Value: value,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListUsers returns a list of users based on pagination and filtering criteria
// @Summary ListUsers
// @Security ApiKeyAuth
// @Description API for retrieving users by page and limit with optional filtering and sorting
// @Tags User
// @Accept json
// @Produce json
// @Param page query string false "Page number for pagination"
// @Param limit query string false "Number of items per page"
// @Param field query string false "Field for filtering"
// @Param value query string false "Value for filtering"
// @Param sortBy query string false "Field to sort by"
// @Param startedAt query string false "Start date for date range filtering"
// @Param endedAt query string false "End date for date range filtering"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users [get]
func (h *handlerV1) ListUsers(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetAll(
		ctx, &pbu.GetAllUsersReq{
			Limit:     params.Limit,
			Page:      params.Page,
			Field:     params.Field,
			Value:     params.Value,
			Sortby:    params.SortBy,
			StartedAt: params.StartedAt,
			EndedAt:   params.EndedAt,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser updates user by field and value
// @Summary UpdateUser
// @Security ApiKeyAuth
// @Description API for updating user by field and value
// @Tags User
// @Accept json
// @Produce json
// @Param field query string true "Field name for filtering"
// @Param value query string true "Value for filtering"
// @Param User body models.UserUpdate true "updateUserModel"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user/update [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	// Fetch existing user by field and value
	response, err := h.serviceManager.UserService().Get(
		ctx, &pbu.GetUserReq{
			Field: field,
			Value: value,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	// Bind updated user data from request body
	err = c.ShouldBindJSON(&response)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	// Update user using service method
	updated, err := h.serviceManager.UserService().Update(ctx, &pbu.UpdateUserReq{
		FirstName:      response.FirstName,
		LastName:       response.LastName,
		Password:       response.Password,
		UserName:       response.UserName,
		BirthDate:      response.BirthDate,
		Biography:      response.Biography,
		Gender:         response.Gender,
		ProfilePicture: response.ProfilePicture,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DeleteUser deletes user by field and value
// @Summary DeleteUser
// @Security ApiKeyAuth
// @Description API for deleting user by field and value
// @Tags User
// @Accept json
// @Produce json
// @Param field query string true "Field name for filtering"
// @Param value query string true "Value for filtering"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users [delete]
func (h *handlerV1) DeleteUser(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Delete(
		ctx, &pbu.DeleteUserReq{
			Field: field,
			Value: value,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllData data
// @Summary GetAllData
// @Security ApiKeyAuth
// @Description Api for get all data
// @Tags User
// @Accept json
// @Produce json
// @Param page query string true "page"
// @Param limit query string true "limit"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/all/user/data [get]
func (h *handlerV1) GetAllUserData(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetAll(
		ctx, &pbc.GetAllCommentsReq{
			Page:  cast.ToInt64(page),
			Limit: cast.ToInt64(limit),
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get all data", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
