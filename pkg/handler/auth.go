package handler

import (
	"GuardianPRO/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.Account
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	isRegistered, err := h.services.Account.IsRegistered(input.Email)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if isRegistered {
		c.AbortWithStatus(http.StatusPreconditionFailed)
		return
	}
	session, err := h.services.Auth.SignUp(&input)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Header("Authorization", *session)
	c.Status(201)
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.Account
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isRegistered, err := h.services.Account.IsRegistered(input.Email)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if !isRegistered {
		c.AbortWithStatus(http.StatusPreconditionFailed)
		return
	}
	session, err := h.services.SignIn(&input)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Header("Authorization", *session)
	c.Status(200)
}
