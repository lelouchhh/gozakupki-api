package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
	domain "gozakupki-api/domain"
	"gozakupki-api/pkg/response"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type AuthHandler struct {
	AUsecase domain.AuthUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewAuthHandler(e *gin.Engine, us domain.AuthUsecase) {
	handler := &AuthHandler{
		AUsecase: us,
	}
	auth := e.Group("/auth")
	{
		auth.POST("/sign_in", handler.SignIn)
		auth.POST("/sign_up", handler.SignUp)
		auth.GET("/check", handler.Check)
		auth.POST("/confirm", handler.Confirm)
	}
}

// SignIn godoc
// @Summary Sign in
// @Description Sign in with user credentials
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body domain.Auth true "User object"
// @Success 200 {object} response.Success
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /auth/sign_in [post]
// @Security BearerAuth
func (a *AuthHandler) SignIn(c *gin.Context) {
	var auth domain.Auth
	err := c.BindJSON(&auth)
	if err != nil {
		c.JSON(
			getStatusCode(err),
			response.SendErrorResponse(response.Error{Message: domain.ErrBadParamInput.Error(), Code: getStatusCode(err)}),
		)
		return
	}
	//err = validateUserFields(auth, "Password")
	if err != nil {
		c.JSON(getStatusCode(err), response.SendErrorResponse(response.Error{
			Message: domain.ErrBadParamInput.Error(),
			Details: err,
			Code:    getStatusCode(err),
		}))
		return
	}
	ctx := c.Request.Context()
	token, err := a.AUsecase.SignIn(ctx, auth)
	if err != nil {
		c.JSON(getStatusCode(err), response.SendErrorResponse(response.Error{
			Message: err.Error(),
			Code:    getStatusCode(err),
		}))
		return
	}
	c.JSON(getStatusCode(err), response.SendSuccessResponse(response.Success{
		Data: map[string]interface{}{
			"token": token,
		},
	}))
}

// SignUp godoc
// @Summary Sign up
// @Description Sign up with user credentials. Send hash to user email
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body domain.Auth true "User object"
// @Success 201 "Created"
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /auth/sign_up [post]
func (a *AuthHandler) SignUp(c *gin.Context) {
	var auth domain.Auth
	err := c.BindJSON(&auth)
	if err != nil {
		c.JSON(
			getStatusCode(err),
			response.SendErrorResponse(response.Error{Message: domain.ErrBadParamInput.Error(), Code: getStatusCode(err)}),
		)
		return
	}
	//err = validateFields(auth, "Email", "Password", "Login")
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
	//}
	ctx := c.Request.Context()
	err = a.AUsecase.SignUp(ctx, auth)
	if err != nil {
		c.JSON(
			getStatusCode(err),
			response.SendErrorResponse(response.Error{Message: err.Error(), Code: getStatusCode(err)}),
		)
		return
	}
	c.JSON(http.StatusCreated, "")
	return
}

// Check godoc
// @Summary Check authentication token
// @Description Check if the authentication token is valid
// @Tags Authentication
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authentication token" default(Bearer )
// @Success 200 {object} response.Success
// @Failure 401 {object} response.Error
// @Router /auth/check [get]
func (a *AuthHandler) Check(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, response.SendErrorResponse(
			response.Error{
				domain.ErrUnauthorized.Error(),
				nil,
				http.StatusUnauthorized,
			}))
		return
	}
	ctx := c.Request.Context()
	err := a.AUsecase.CheckToken(ctx, authHeader[7:])
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.SendErrorResponse(response.Error{
			domain.ErrUnauthorized.Error(),
			nil,
			http.StatusUnauthorized,
		}))
		return
	}
	c.JSON(http.StatusOK, response.SendSuccessResponse(response.Success{""}))
}

func (a *AuthHandler) Confirm(c *gin.Context) {
	var auth domain.Auth
	err := c.BindJSON(&auth)
	if err != nil {
		c.JSON(
			getStatusCode(err),
			response.SendErrorResponse(response.Error{Message: domain.ErrBadParamInput.Error(), Code: getStatusCode(err)}),
		)
		return
	}
	//err = validateFields(auth, "Email", "Password", "Login")
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
	//}
	ctx := c.Request.Context()
	err = a.AUsecase.ConfirmUser(ctx, auth)
	if err != nil {
		c.JSON(
			getStatusCode(err),
			response.SendErrorResponse(response.Error{Message: err.Error(), Code: getStatusCode(err)}),
		)
		return
	}
	c.JSON(
		getStatusCode(err),
		response.SendSuccessResponse(response.Success{}),
	)
	return
}
func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func validateUserFields(user domain.Auth, fields ...string) error {
	v := validator.New()

	// Perform validation on specific fields
	for _, field := range fields {
		switch field {
		case "Password":
			if err := v.Var(user.ID, "required,min=6"); err != nil {
				return err
			}
		case "Login":
			if err := v.Var(user.Login, "required,min=4"); err != nil {
				return err
			}
		case "Email":
			if err := v.Var(user.Email, "required,email"); err != nil {
				return err
			}
		}
	}
	return v.StructPartial(user, fields...)
}
