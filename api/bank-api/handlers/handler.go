package handlers

import (
	"net/http"
	"po/api/bank-api/requests"
	"po/api/bank-api/responses"
	"po/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BankHandler interface {
	CreateBank(c *gin.Context)
	UpdateBank(c *gin.Context)
	ListBank(c *gin.Context)
}

type bankHandler struct {
	bankClient pb.BankServiceClient
}

func NewBankHandler(bankClient pb.BankServiceClient) BankHandler {
	return &bankHandler{
		bankClient: bankClient,
	}
}

func (h *bankHandler) UpdateBank(c *gin.Context) {
	req := requests.UpdateBankRequest{}

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
	pReq := &pb.Bank{
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		AvailableSlot: req.AvailableSlot,
	}
	pRes, err := h.bankClient.Update(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.BankResponse{
		ID:            pRes.Id,
		Name:          pRes.Name,
		From:          pRes.From,
		To:            pRes.To,
		AvailableSlot: pRes.AvailableSlot,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *bankHandler) ListBank(c *gin.Context) {
	req := requests.ListBankRequest{}

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
	pReq := &pb.ListBankRequest{
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		AvailableSlot: req.AvailableSlot,
	}
	pRes, err := h.bankClient.List(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dtos := make([]*responses.BankResponse, 0)

	for _, v := range pRes.Banks {
		dto := &responses.BankResponse{
			ID:            v.Id,
			Name:          v.Name,
			From:          v.From,
			To:            v.To,
			AvailableSlot: v.AvailableSlot,
		}

		dtos = append(dtos, dto)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dtos,
	})
}

func (h *bankHandler) CreateBank(c *gin.Context) {
	req := requests.CreateBankRequest{}

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
	pReq := &pb.Bank{
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		AvailableSlot: req.AvailableSlot,
	}
	pRes, err := h.bankClient.Create(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.BankResponse{
		ID:            pRes.Id,
		Name:          pRes.Name,
		From:          pRes.From,
		To:            pRes.To,
		AvailableSlot: pRes.AvailableSlot,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
