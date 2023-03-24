package handler

import "github.com/gin-gonic/gin"

func (h *Handler) getAccountList(c *gin.Context) {
	output, err := h.services.Account.GetList()
	if err != nil {
		_ = c.AbortWithError(500, err)
	}
	c.JSON(200, &output)
}
