package handlers

import (
	"net/http"
	"po/api/user-api/requests"
	"po/api/user-api/responses"
	"po/pb"

	"github.com/jinzhu/copier"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	ChangePassword(c *gin.Context)
	PspHistory(c *gin.Context)
}

type userHandler struct {
	userServiceClient pb.UserServiceClient
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	req := requests.UpdateUserRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.User{}
	err := copier.Copy(&pReq, req)
	if err != nil {
	}
	pRes, err := h.userServiceClient.Update(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.UserResponse{}
	err = copier.Copy(dto, pRes)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *userHandler) ChangePassword(c *gin.Context) {
	req := requests.UpdateUserRequest{}

	pReq := &pb.ChangePasswordRequest{}
	err := copier.Copy(&pReq, req)
	if err != nil {
	}
	pRes, err := h.userServiceClient.ChangePassword(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.UserResponse{}
	err = copier.Copy(dto, pRes)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *userHandler) PspHistory(c *gin.Context) {
	id := c.Param("id")
	pReq := &pb.PspHistoryRequest{
		UserId: id,
	}
	pRes, err := h.userServiceClient.PspHistory(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.UserResponse{}
	err = copier.Copy(dto, pRes)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *userHandler) CreateUser(c *gin.Context) {
	req := requests.CreateUserRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.User{}
	err := copier.Copy(&pReq, req)
	if err != nil {
	}
	pRes, err := h.userServiceClient.Create(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.UserResponse{}
	err = copier.Copy(dto, pRes)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func NewPeopleHandler(userServiceClient pb.UserServiceClient) UserHandler {
	return &userHandler{
		userServiceClient: userServiceClient,
	}
}
