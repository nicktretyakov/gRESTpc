package handlers

import (
	"net/http"
	"po/api/psp-api/requests"
	"po/api/psp-api/responses"
	"po/pb"

	"github.com/jinzhu/copier"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PspHandler interface {
	CreatePsp(c *gin.Context)
	ViewPsp(c *gin.Context)
	CancelPsp(c *gin.Context)
}

type pspHandler struct {
	pspClient pb.PspServiceClient
}

func (h *pspHandler) ViewPsp(c *gin.Context) {
	id := c.Param("id")
	pReq := &pb.ViewPspRequest{
		PspId: id,
	}

	pRes, err := h.pspClient.ViewPsp(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.PspResponse{}
	copier.Copy(dto, pRes)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *pspHandler) CancelPsp(c *gin.Context) {
	id := c.Param("id")
	pReq := &pb.CancelPspRequest{
		PspId: id,
	}

	pRes, err := h.pspClient.CancelPsp(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.PspResponse{}
	copier.Copy(dto, pRes)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *pspHandler) CreatePsp(c *gin.Context) {
	req := requests.CreatePspRequest{}

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

	pReq := &pb.PspRequest{}

	pRes, err := h.pspClient.CreatePsp(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.PspResponse{}
	copier.Copy(dto, pRes)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func NewPspHandler(pspServiceClient pb.PspServiceClient) PspHandler {
	return &pspHandler{
		pspClient: pspServiceClient,
	}
}
