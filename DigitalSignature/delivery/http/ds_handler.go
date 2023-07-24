package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	domain "gozakupki-api/domain"
	"gozakupki-api/pkg/response"
	"net/http"
)

// ArticleHandler  represent the httphandler for article
type DSHandler struct {
	SUsecase domain.DigitalSignatureUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewDSHandler(e *gin.Engine, us domain.DigitalSignatureUsecase) {
	handler := &DSHandler{
		SUsecase: us,
	}
	auth := e.Group("/auth")
	{
		auth.POST("/sign_in", handler.SignIn)
		auth.POST("/sign_up", handler.SignUp)
	}
}
func (a *DSHandler) SignIn(c *gin.Context) {
	var ds domain.DigitalSignature
	err := c.BindJSON(&ds)
	if err != nil {
		c.JSON(
			getStatusCode(err),
			response.SendErrorResponse(response.Error{Message: domain.ErrBadParamInput.Error(), Code: getStatusCode(err)}),
		)
		return
	}
	ctx := c.Request.Context()
	token, err := a.SUsecase.SignIn(ctx, ds)
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
func (a *DSHandler) SignUp(c *gin.Context) {

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
