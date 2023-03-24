package handler

import (
	"GuardianPRO/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.Account
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.AbortWithError(400, err)
		return
	}
	isRegistered, err := h.services.Account.IsRegistered(input.Email)
	if err != nil {
		_ = c.AbortWithError(500, err)
		return
	}
	if isRegistered {
		_ = c.AbortWithError(412, err)
		return
	}
	session, err := h.services.Auth.SignUp(&input)
	if err != nil {
		_ = c.AbortWithError(500, err)
		return
	}
	c.Header("Authorization", *session)
	c.Status(201)
}

func (h *Handler) signIn(c *gin.Context) {

}
