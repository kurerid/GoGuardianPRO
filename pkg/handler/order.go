package handler

import (
	"GuardianPRO/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createPrivateOrder(c *gin.Context) {
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if input.Type != models.PRIVATE_ORDER {
		_ = c.AbortWithError(http.StatusPreconditionFailed, errors.New("unexpected order type"))
		return
	}
	if err := h.services.Order.CreatePrivate(&input); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(200)
}

func (h *Handler) getGroupOrderList(c *gin.Context) {

}
