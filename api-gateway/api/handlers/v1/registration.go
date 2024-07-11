package v1

import (
	"context"
	"log"
	"net/http"

	"time"

	"api-gateway/api/handlers/models"
	token "api-gateway/api/handlers/tokens"
	l "api-gateway/pkg/logger"
	pbu "api-gateway/protos/user-service"

	"github.com/gin-gonic/gin"
)

// Register ...
// @Summary Register
// @Description API for user registration
// @Tags Authorizations
// @Accept json
// @Produce json
// @Param name query string true "First name"
// @Param lastName query string false "Last name"
// @Param email query string true "Email"
// @Param password query string true "Password"
// @Param userName query string true "Username"
// @Param phoneNumber query string false "Phone number"
// @Param birthDate query string false "Birth date (YYYY-MM-DD)"
// @Param biography query string false "Biography"
// @Param gender query string false "Gender"
// @Param profilePicture query string false "Profile picture URL"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/register/ [post]
func (h *handlerV1) Register(c *gin.Context) {
	var body models.RegisterUser

	// Bind query parameters
	body.Name = c.Query("name")
	body.LastName = c.Query("lastName")
	body.Email = c.Query("email")
	body.Password = c.Query("password")
	body.UserName = c.Query("userName")
	body.PhoneNumber = c.Query("phoneNumber")
	body.BirthDate = c.Query("birthDate")
	body.Biography = c.Query("biography")
	body.Gender = c.Query("gender")
	body.ProfilePicture = c.Query("profilePicture")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This password is already in use or there is an email error. Please choose another.",
		})
		h.log.Error("failed to check email uniqueness", l.Error(err))
		return
	}

	res, err := h.serviceManager.UserService().Register(ctx, &pbu.CreateUserReq{
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
		h.log.Error("failed to register user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "email": res.User.Email})
}

// LogIn
// @Summary LogIn User
// @Description LogIn - Api for login users
// @Tags Authorizations
// @Accept json
// @Produce json
// @Param email query string false "Email"
// @Param username query string false "Username"
// @Param password query string true "Password"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/login [post]
func (h *handlerV1) LogIn(c *gin.Context) {
	email := c.Query("email")
	username := c.Query("username")
	password := c.Query("password")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	// Validate if either email or username is provided
	if email == "" && username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "either email or username must be provided",
		})
		return
	}

	response, err := h.serviceManager.UserService().Login(ctx, &pbu.LoginUserReq{
		Email:    email,
		UserName: username,
		Password: password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to log in user",
		})
		return
	}

	h.jwthandler = token.JWTHandler{
		Sub:       response.User.Id,
		Iss:       time.Now().String(),
		Exp:       time.Now().Add(time.Hour * 6).String(),
		Role:      "user",
		SigninKey: h.cfg.SigningKey,
		Timeot:    h.cfg.AccessTokenTimout,
	}

	access, refresh_token, err := h.jwthandler.GenerateAuthJWT()
	if err != nil {
		log.Fatal("error while generating auth token")
	}

	var respModel models.UserBYtokens
	respModel.Id = response.User.Id
	respModel.Name = response.User.FirstName
	respModel.LastName = response.User.LastName
	respModel.UserName = response.User.UserName
	respModel.Email = response.User.Email
	respModel.RefreshToken = refresh_token
	respModel.Password = response.User.Password
	respModel.AccessToken = access

	c.JSON(http.StatusOK, respModel)
}

// Verification
// @Summary Verification User
// @Description LogIn - Api for verification users
// @Tags Authorizations
// @Accept json
// @Produce json
// @Param email query string true "Email"
// @Param code query string true "Code"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/verification [post]
func (h *handlerV1) Verification(c *gin.Context) {

	email := c.Query("email")
	code := c.Query("code")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	createdUser, err := h.serviceManager.UserService().Authorization(ctx, &pbu.AuthUser{
		Email: email,
		Code:  code,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to verify user",
		})
		return
	}

	// Create access and refresh tokens JWT
	h.jwthandler = token.JWTHandler{
		Sub:       createdUser.User.Id,
		Iss:       time.Now().String(),
		Exp:       time.Now().Add(time.Hour * 6).String(),
		Role:      "user",
		SigninKey: h.cfg.SigningKey,
		Timeot:    h.cfg.AccessTokenTimout,
	}
	// aksestoken bn refreshtokeni generatsa qiliah
	access, refresh, err := h.jwthandler.GenerateAuthJWT()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "error generating token")
		return
	}

	response := &models.UserBYtokens{
		Id:           createdUser.User.Id,
		Name:         createdUser.User.FirstName,
		LastName:     createdUser.User.LastName,
		UserName:     createdUser.User.UserName,
		Email:        createdUser.User.Email,
		AccessToken:  access,
		RefreshToken: refresh,
	}

	c.JSON(http.StatusOK, response)
}
